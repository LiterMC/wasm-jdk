package vm

import (
	"fmt"
	"iter"
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

	loadedFieldAccesors map[uint16]func(ir.VM) *Field
	loadedMethods       map[uint16]func(ir.VM) *Method
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
	elem := c.elem
	if elem == nil {
		elem = c
	} else {
		elem = c.elem
		dim += c.arrayDim
	}
	return &Class{
		arrayDim: dim,
		elem:     elem,
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

	c.loadedFieldAccesors = make(map[uint16]func(ir.VM) *Field)
	c.loadedMethods = make(map[uint16]func(ir.VM) *Method)
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
		return c.Desc().String()
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
	if c == k {
		return true
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
	if r == nil || r == (*Ref)(nil) {
		return true
	}
	return c.IsAssignableFrom(r.Class())
}

func (c *Class) GetAndPushConst(vm ir.VM, i uint16, s ir.Stack) error {
	v := c.ConstPool[i-1]
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
		s.PushRef(vm.GetStringInternOrNew(v.Utf8))
	case *jcls.ConstantClass:
		class, err := vm.(*VM).getClassFromDescString(v.Name)
		if err != nil {
			return err
		}
		s.PushRef(class.AsRef(vm))
	default:
		return fmt.Errorf("Unexpected constant type %T", v)
	}
	return nil
}

func (c *Class) ForEachField(yield func(ir.Field) bool) {
	for i := range len(c.Fields) {
		if !yield(&c.Fields[i]) {
			return
		}
	}
	if c.super != nil {
		c.super.(*Class).ForEachField(yield)
	}
}

func (c *Class) GetFields() iter.Seq[ir.Field] {
	return c.ForEachField
}

func (c *Class) GetField(vm ir.VM, i uint16) ir.Field {
	return c.loadedFieldAccesors[i](vm)
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

func (c *Class) ForEachMethod(yield func(ir.Method) bool) {
	for i := range len(c.Methods) {
		if !yield(&c.Methods[i]) {
			return
		}
	}
	if c.super != nil {
		c.super.(*Class).ForEachMethod(yield)
	}
}

func (c *Class) GetMethods() iter.Seq[ir.Method] {
	return c.ForEachMethod
}

func (c *Class) GetMethod(vm ir.VM, i uint16) ir.Method {
	return c.loadedMethods[i](vm)
}

func (c *Class) GetMethodByName(vm ir.VM, location string) ir.Method {
	i := strings.IndexByte(location, desc.Method)
	if i < 0 {
		panic("method name missing descriptor")
	}
	return c.GetMethodByNameAndType(vm, location[:i], location[i:])
}

func (c *Class) GetMethodByNameAndType(vm ir.VM, name, typ string) ir.Method {
	dc, err := desc.ParseMethodDesc(typ)
	if err != nil {
		panic(err)
	}
	return c.GetMethodByDesc(vm, name, dc)
}

func (c *Class) GetMethodByDesc(vm ir.VM, name string, dc *desc.MethodDesc) ir.Method {
	if c.arrayDim > 0 {
		return vm.(*VM).getArrayMethod(c, name)
	}
	x := c
	for {
		for i := range len(x.Methods) {
			m := &x.Methods[i]
			if m.Name() == name && m.Desc().EqInputs(dc) {
				// if !m.Desc().Output.Eq(dc.Output) {
				// 	panic("Methods " + name + dc.String() + " have same inputs but different output " + m.Desc().String())
				// }
				return m
			}
		}
		if x.super == nil {
			return nil
		}
		x = x.super.(*Class)
	}
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
	f := c.loadField(c.initVM.Load(), ref, false)
	if f != nil {
		if f.IsStatic() != static {
			panic(fmt.Errorf("field status %v not match getfield/getstatic command", f.IsStatic()))
		}
		c.loadedFieldAccesors[ind] = func(ir.VM) *Field { return f }
	} else {
		c.loadedFieldAccesors[ind] = OnceApply(func(vm ir.VM) *Field {
			return c.loadField(vm, ref, true)
		})
	}
}

func (c *Class) loadField(vm ir.VM, ref *jcls.ConstantRef, canLoadClass bool) *Field {
	var x *Class
	if ref.Class.Name == c.Name() {
		x = c
	} else if canLoadClass {
		k, err := c.loader.LoadClass(ref.Class.Name)
		if err != nil {
			panic(err)
		}
		x = k.(*Class)
		x.InitBeforeUse(vm.(*VM))
	} else {
		return nil
	}
	for {
		for i, f := range x.Class.Fields {
			if f.Name() == ref.NameAndType.Name {
				if f.Desc.String() != ref.NameAndType.Desc {
					panic(fmt.Errorf("cannot load class: field %s is %s, but one operation requires %s", ref.NameAndType.Name, f.Desc.String(), ref.NameAndType.Desc))
				}
				return &x.Fields[i]
			}
		}
		if x.super == nil {
			return nil
		}
		x = x.super.(*Class)
	}
}

func (c *Class) loadMethodGetter(ind uint16) {
	if _, ok := c.loadedMethods[ind]; ok {
		return
	}
	ref, ok := c.ConstPool[ind-1].(*jcls.ConstantRef)
	if !ok || (ref.ConstTag != jcls.TagMethodref && ref.ConstTag != jcls.TagInterfaceMethodref) {
		panic(fmt.Errorf("cannot load class: constant at %d is not a method reference", ind-1))
	}
	if ref.Class.Name[0] == '[' {
		if ref.Class.Name[len(ref.Class.Name)-1] == ';' {
			// lazy load class
			c.loadedMethods[ind] = OnceApply(func(vm ir.VM) *Method {
				return vm.(*VM).getArrayMethodByName(ref)
			})
		} else {
			arrayMethod := ((*VM)(nil)).getArrayMethodByName(ref)
			c.loadedMethods[ind] = func(ir.VM) *Method {
				return arrayMethod
			}
		}
	} else {
		c.loadedMethods[ind] = OnceApply(func(vm ir.VM) *Method {
			return c.loadMethod(vm, ref)
		})
	}
}

func (c *Class) loadMethod(vm ir.VM, ref *jcls.ConstantRef) *Method {
	k, err := c.loader.LoadClass(ref.Class.Name)
	if err != nil {
		panic(err)
	}
	x := k.(*Class)
	x.InitBeforeUse(vm.(*VM))
	method := x.loadMethod0(ref)
	if method == nil {
		panic(fmt.Errorf("cannot load class: missing method %s", ref))
	}
	return method
}

func (c *Class) loadMethod0(ref *jcls.ConstantRef) *Method {
	for x := c; x != nil; x = x.super.(*Class) {
		for i, m := range x.Class.Methods {
			if m.Name() == ref.NameAndType.Name && m.Desc().String() == ref.NameAndType.Desc {
				return &x.Methods[i]
			}
		}
		if x.super == nil {
			break
		}
	}
	for _, in := range c.interfaces {
		m := in.(*Class).loadMethod0(ref)
		if m != nil {
			return m
		}
	}
	return nil
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
