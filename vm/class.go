package vm

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

// vm.Class represents a loaded class
type Class struct {
	*jcls.Class
	loader ir.ClassLoader
	initVM atomic.Pointer[VM]

	arrayDim   int // -1: primary type; 0: normal class; 1+: array class
	elem       *Class
	super      ir.Class
	interfaces []ir.Class
	refType    reflect.Type
	classRef   atomic.Pointer[Ref]

	initFMOnce sync.Once
	Fields     []Field
	Methods    []Method
	staticInit *Method
	staticData unsafe.Pointer

	loadedFieldAccesors map[uint16]func() *Field
	loadedMethods       map[uint16]func() *Method
	loadedDynamics      map[uint16]*dynamicInfo
}

var _ ir.Class = (*Class)(nil)

func LoadClass(cls *jcls.Class, loader ir.ClassLoader) *Class {
	c := &Class{
		Class:  cls,
		loader: loader,
	}
	var err error
	if c.SuperSym != nil {
		if c.super, err = loader.LoadClass(c.SuperSym.Name); err != nil {
			panic(err)
		}
	}
	c.interfaces = make([]ir.Class, len(c.InterfacesSym))
	for i, in := range c.InterfacesSym {
		if c.interfaces[i], err = loader.LoadClass(in.Name); err != nil {
			panic(err)
		}
	}

	statics := make([]reflect.StructField, 0, len(c.Class.Fields))
	fields := make([]reflect.StructField, 1, len(c.Class.Fields)+1)
	if c.super == nil {
		fields[0] = reflect.StructField{
			Name: "S",
			Type: reflect.TypeOf(struct{}{}),
		}
	} else {
		fields[0] = reflect.StructField{
			Name: "S",
			Type: c.super.Reflect(),
		}
	}
	for _, f := range c.Class.Fields {
		fName := "F" + j2goName(f.Name())
		if f.IsStatic() {
			statics = append(statics, reflect.StructField{
				Name: fName,
				Type: f.Desc.AsReflect(),
			})
		} else {
			fields = append(fields, reflect.StructField{
				Name: fName,
				Type: f.Desc.AsReflect(),
			})
		}
	}
	staticType := reflect.StructOf(statics)
	c.refType = reflect.StructOf(fields)
	c.Fields = make([]Field, len(c.Class.Fields))
	for i, f := range c.Class.Fields {
		cf := &c.Fields[i]
		cf.Field = f
		cf.class = c
		fName := "F" + j2goName(cf.Name())
		if cf.IsStatic() {
			rf, _ := staticType.FieldByName(fName)
			cf.offset = rf.Offset
		} else {
			rf, _ := c.refType.FieldByName(fName)
			cf.offset = rf.Offset
		}
	}
	c.staticData = reflect.New(staticType).UnsafePointer()

	c.Methods = make([]Method, len(c.Class.Methods))
	for i, m := range c.Class.Methods {
		cm := &c.Methods[i]
		cm.Method = m
		cm.class = c
		if cm.Name() == "<clinit>" {
			if !cm.IsStatic() {
				panic("class initalize method is not static")
			}
			c.staticInit = cm
		}
	}
	return c
}

func (c *Class) NewArrayClass(dim int) *Class {
	return &Class{
		arrayDim: dim,
		elem:     c,
	}
}

func (c *Class) ShouldInit() bool {
	if c.arrayDim > 0 {
		return c.elem.ShouldInit()
	}
	return c.initVM.Load() == nil
}

func (c *Class) InitBeforeUse(vm *VM) {
	if c.arrayDim > 0 {
		c.elem.InitBeforeUse(vm)
		return
	}
	if c.arrayDim < 0 {
		return
	}
	if c.initVM.Load() == vm {
		return
	}
	c.initVM.CompareAndSwap(nil, vm)
	c.initFMOnce.Do(c.initFM)
}

func (c *Class) initFM() {
	fmt.Println("initing", c.Name())
	defer fmt.Println("post init", c.Name())

	ivm := c.initVM.Load()
	if super, ok := c.super.(*Class); ok {
		// fmt.Println("waiting", super.Name())
		super.InitBeforeUse(ivm)
	}

	var err error
	for i, f := range c.Class.Fields {
		cf := &c.Fields[i]
		if f.Desc.EndType == desc.Class {
			if cf.typ, err = c.loader.LoadClass(f.Desc.Class); err != nil {
				panic(err)
			}
		}
	}

	c.loadedFieldAccesors = make(map[uint16]func() *Field)
	c.loadedMethods = make(map[uint16]func() *Method)
	c.loadedDynamics = make(map[uint16]*dynamicInfo)
	c.scanCodes()

	if c.staticInit != nil {
		fmt.Println("==> invoking " + c.Name() + ".<clinit>")
		vm := c.initVM.Load()
		prev := vm.stack
		prev.pc = vm.nextPc
		vm.stack = &Stack{
			prev:   prev,
			class:  c,
			method: c.staticInit,
		}
		vm.nextPc = c.staticInit.Code.Code
		err := vm.RunStack()
		if err != nil {
			panic(err)
		}
	}
}

func (c *Class) ArrayDim() int {
	return c.arrayDim
}

func (c *Class) Elem() ir.Class {
	if c.arrayDim == 0 {
		panic("not an array class")
	}
	if c.arrayDim == 1 {
		return c.elem
	}
	return c.elem.NewArrayClass(c.arrayDim - 1)
}

func (c *Class) Name() string {
	if c.arrayDim > 0 {
		return c.elem.Name() + "[]"
	}
	return c.Class.Name()
}

func (c *Class) Desc() *desc.Desc {
	if c.arrayDim > 0 {
		d := c.elem.Class.Desc().Clone()
		d.ArrDim = c.arrayDim
		return d
	}
	return c.Class.Desc()
}

func (c *Class) Reflect() reflect.Type {
	return c.refType
}

func (c *Class) Modifiers() int32 {
	return (int32)(c.Class.AccessFlags)
}

func (c *Class) Super() ir.Class {
	return c.super
}

func (c *Class) Interfaces() []ir.Class {
	return c.interfaces
}

func (c *Class) IsInterface() bool {
	if c.arrayDim != 0 {
		return false
	}
	return c.Class.IsInterface()
}

func (c *Class) IsAssignableFrom(k ir.Class) bool {
	if k == nil {
		return false
	}
	kk := k.(*Class)
	if (c == nil || kk == nil) && c != kk {
		return false
	}
	if c.arrayDim != 0 {
		if c.arrayDim != kk.arrayDim {
			return false
		}
		if c.arrayDim == -1 {
			return c == kk
		}
		return c.elem.IsAssignableFrom(kk.elem)
	}
	if !c.IsInterface() {
		if c == k {
			return true
		}
		if c.AccessFlags.Has(jcls.AccFinal) {
			return false
		}
		for {
			k = k.Super()
			if k == nil {
				return false
			}
			if c == k {
				return true
			}
		}
	}
	for k != nil {
		for _, in := range k.Interfaces() {
			if c == in {
				return true
			}
		}
		k = k.Super()
	}
	return false
}

func (c *Class) IsInstance(r ir.Ref) bool {
	if r == nil {
		return true
	}
	return c.IsAssignableFrom(r.Class())
}

func (c *Class) GetAndPushConst(i uint16, s ir.Stack) error {
	v := c.ConstPool[i-1]
	fmt.Printf("pushing const %#v\n", v)
	switch v := v.(type) {
	case *jcls.ConstantInteger:
		s.Push(v.Value)
	case *jcls.ConstantFloat:
		s.Push(v.Value)
	case *jcls.ConstantLong:
		s.Push64(v.Value)
	case *jcls.ConstantDouble:
		s.Push64(v.Value)
	case *jcls.ConstantString:
		s.PushRef(c.initVM.Load().GetStringInternOrNew(v.Utf8))
	case *jcls.ConstantClass:
		vm := c.initVM.Load()
		class, err := vm.getClassFromDescString(v.Name)
		if err != nil {
			return err
		}
		s.PushRef(vm.GetClassRef(class))
	default:
		return fmt.Errorf("Unexpected constant type %T", v)
	}
	return nil
}

func (c *Class) GetField(i uint16) ir.Field {
	return c.loadedFieldAccesors[i]()
}

func (c *Class) GetFieldByName(name string) ir.Field {
	for i := range len(c.Fields) {
		f := &c.Fields[i]
		if f.Name() == name {
			return f
		}
	}
	return nil
}

func (c *Class) GetMethod(i uint16) ir.Method {
	return c.loadedMethods[i]()
}

func (c *Class) GetMethodByName(location string) ir.Method {
	i := strings.IndexByte(location, desc.Method)
	if i < 0 {
		panic("method name missing descriptor")
	}
	return c.GetMethodByNameAndType(location[:i], location[i:])
}

func (c *Class) GetMethodByNameAndType(name, typ string) ir.Method {
	dc, err := desc.ParseMethodDesc(typ)
	if err != nil {
		panic(err)
	}
	return c.GetMethodByDesc(name, dc)
}

func (c *Class) GetMethodByDesc(name string, dc *desc.MethodDesc) ir.Method {
	for x := c; x != nil; x = x.super.(*Class) {
		for i := range len(x.Methods) {
			m := &x.Methods[i]
			if m.Name() == name && m.Desc().EqInputs(dc) {
				// if !m.Desc().Output.Eq(dc.Output) {
				// 	panic("Methods " + name + dc.String() + " have same inputs but different output " + m.Desc().String())
				// }
				return m
			}
		}
	}
	return nil
}

func (c *Class) scanCodes() {
	for _, m := range c.Class.Methods {
		if m.AccessFlags.Has(jcls.AccNative | jcls.AccAbstract) {
			continue
		}
		for node := m.Code.Code; node != nil; node = node.Next {
			switch ic := node.IC.(type) {
			case *ir.ICgetfield:
				c.loadFieldGetter(ic.Field, false)
			case *ir.ICgetstatic:
				c.loadFieldGetter(ic.Field, true)
			case *ir.ICputfield:
				c.loadFieldGetter(ic.Field, false)
			case *ir.ICputstatic:
				c.loadFieldGetter(ic.Field, true)
			case *ir.ICinvokedynamic:
				c.loadMethodDynamic(ic.Method)
			case *ir.ICinvokeinterface:
				c.loadMethodGetter(ic.Method)
			case *ir.ICinvokespecial:
				c.loadMethodGetter(ic.Method)
			case *ir.ICinvokestatic:
				c.loadMethodGetter(ic.Method)
			case *ir.ICinvokevirtual:
				c.loadMethodGetter(ic.Method)
			}
		}
	}
}

func (c *Class) loadFieldGetter(ind uint16, static bool) {
	if _, ok := c.loadedFieldAccesors[ind]; ok {
		return
	}
	ref, ok := c.ConstPool[ind-1].(*jcls.ConstantRef)
	if !ok || ref.ConstTag != jcls.TagFieldref {
		panic(fmt.Errorf("cannot load class: constant at %d is not a field ref", ind-1))
	}
	f := c.loadField(ref, false)
	if f != nil {
		if f.IsStatic() != static {
			panic(fmt.Errorf("field status %v not match getfield/getstatic command", f.IsStatic()))
		}
		c.loadedFieldAccesors[ind] = func() *Field { return f }
	} else {
		c.loadedFieldAccesors[ind] = sync.OnceValue(func() *Field {
			return c.loadField(ref, true)
		})
	}
}

func (c *Class) loadField(ref *jcls.ConstantRef, canLoadClass bool) *Field {
	var x *Class
	if ref.Class.Name == c.Name() {
		x = c
	} else if canLoadClass {
		k, err := c.loader.LoadClass(ref.Class.Name)
		if err != nil {
			panic(err)
		}
		x = k.(*Class)
		x.InitBeforeUse(c.initVM.Load())
	} else {
		return nil
	}
	for i, f := range x.Class.Fields {
		if f.Name() == ref.NameAndType.Name {
			if f.Desc.String() != ref.NameAndType.Desc {
				panic(fmt.Errorf("cannot load class: field %s is %s, but one operation requires %s", ref.NameAndType.Name, f.Desc.String(), ref.NameAndType.Desc))
			}
			return &x.Fields[i]
		}
	}
	return nil
}

func (c *Class) loadMethodGetter(ind uint16) {
	ref, ok := c.ConstPool[ind-1].(*jcls.ConstantRef)
	if !ok || (ref.ConstTag != jcls.TagMethodref && ref.ConstTag != jcls.TagInterfaceMethodref) {
		panic(fmt.Errorf("cannot load class: constant at %d is not a method reference", ind-1))
	}
	c.loadedMethods[ind] = sync.OnceValue(func() *Method {
		return c.loadMethod(ref)
	})
}

func (c *Class) loadMethod(ref *jcls.ConstantRef) *Method {
	k, err := c.loader.LoadClass(ref.Class.Name)
	if err != nil {
		panic(err)
	}
	x := k.(*Class)
	x.InitBeforeUse(c.initVM.Load())
	for x != nil {
		for i, m := range x.Class.Methods {
			if m.Name() == ref.NameAndType.Name && m.Desc().String() == ref.NameAndType.Desc {
				return &x.Methods[i]
			}
		}
		if x.super == nil {
			break
		}
		x = x.super.(*Class)
	}
	panic(fmt.Errorf("cannot load class: cannot find method %s", ref))
}

type dynamicInfo struct {
	info      *jcls.ConstantDynamics
	bootstrap *jcls.BootstrapMethod
	callSite  *Ref
}

func (c *Class) loadMethodDynamic(ind uint16) {
	info, ok := c.ConstPool[ind-1].(*jcls.ConstantDynamics)
	if !ok || info.ConstTag != jcls.TagInvokeDynamic {
		panic(fmt.Errorf("cannot load class: constant at %d is not a invokedynamic", ind-1))
	}
	bootstrap := c.GetAttr("BootstrapMethods").(*jcls.AttrBootstrapMethods).Methods[info.BootstrapMethod]
	c.loadedDynamics[ind] = &dynamicInfo{
		info:      info,
		bootstrap: bootstrap,
	}
}

func (c *Class) invokeMethodDynamic(ref *jcls.ConstantDynamics) *Method {
	panic(fmt.Errorf("TODO: Trying to invoke dynamic %#v %#v", ref, ref.NameAndType))
}

func j2goName(name string) string {
	return strings.ReplaceAll(name, "$", "__")
}
