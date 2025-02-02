package vm

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

// vm.Class represents a loaded class
type Class struct {
	*jcls.Class
	loader ClassLoader

	super      ir.Class
	interfaces []ir.Class
	refType    reflect.Type

	initFMOnce sync.Once
	Fields     []Field
	Methods    []Method

	loadedFieldAccesors map[uint16]func() *Field
	loadedMethods       map[uint16]func() *Method
}

var _ ir.Class = (*Class)(nil)

func LoadClass(cls *jcls.Class, loader ClassLoader) *Class {
	c := &Class{
		Class:  cls,
		loader: loader,
	}
	var err error
	if c.super, err = c.loader.LoadClass(c.SuperSym.Name); err != nil {
		panic(err)
	}
	c.interfaces = make([]ir.Class, len(c.InterfacesSym))
	for i, in := range c.InterfacesSym {
		if c.interfaces[i], err = c.loader.LoadClass(in.Name); err != nil {
			panic(err)
		}
	}

	fields := make([]reflect.StructField, len(c.Class.Fields)+1)
	fields[0] = reflect.StructField{
		Name: "S",
		Type: c.super.Reflect(),
	}
	for i, f := range c.Class.Fields {
		fields[i+1] = reflect.StructField{
			Name: "F" + f.Name(),
			Type: f.Desc.AsReflect(),
		}
	}
	c.refType = reflect.StructOf(fields)
	return c
}

func (c *Class) InitBeforeUse() {
	c.initFMOnce.Do(c.initFM)
}

func (c *Class) initFM() {
	var err error
	c.Fields = make([]Field, len(c.Class.Fields))
	for i, f := range c.Class.Fields {
		cf := &c.Fields[i]
		cf.Field = f
		cf.class = c
		cf.offset = c.refType.Field(i + 1).Offset
		if f.Desc.EndType == desc.Class {
			if cf.typ, err = c.loader.LoadClass(f.Desc.Class); err != nil {
				panic(err)
			}
		}
	}
	c.Methods = make([]Method, len(c.Class.Methods))
	for i, m := range c.Class.Methods {
		cf := &c.Methods[i]
		cf.Method = m
		cf.class = c
	}

	c.loadedFieldAccesors = make(map[uint16]func() *Field)
	c.loadedMethods = make(map[uint16]func() *Method)
	c.scanCodes()
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

func (c *Class) IsAssignableFrom(k ir.Class) bool {
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
	switch v := v.(type) {
	case *jcls.ConstantString:
		s.PushRef(nil) // TODO: create string ref
	case *jcls.ConstantInteger:
		s.Push(v.Value)
	case *jcls.ConstantFloat:
		s.Push(v.Value)
	case *jcls.ConstantLong:
		s.Push64(v.Value)
	case *jcls.ConstantDouble:
		s.Push64(v.Value)
	default:
		return fmt.Errorf("Unexpected constant type %T", v)
	}
	return nil
}

func (c *Class) GetField(i uint16) ir.Field {
	return c.loadedFieldAccesors[i]()
}

func (c *Class) GetMethod(i uint16) ir.Method {
	return c.loadedMethods[i]()
}

func (c *Class) GetMethodByName(name string, desc *desc.MethodDesc) ir.Method {
	for i := range len(c.Methods) {
		m := &c.Methods[i]
		if !m.AccessFlags.Has(jcls.AccPrivate) && m.Name() == name && m.Desc().EqInputs(desc) {
			if !m.Desc().Output.Eq(desc.Output) {
				panic("Methods have same inputs but different output")
			}
			return m
		}
	}
	return nil
}

func (c *Class) scanCodes() {
	for _, m := range c.Class.Methods {
		for node := m.Code.Code; node != nil; node = node.Next {
			switch ic := node.IC.(type) {
			case *ir.ICgetfield:
				c.loadFieldGetter(ic.Field, false)
			case *ir.ICgetstatic:
				c.loadFieldGetter(ic.Field, true)
			case *ir.ICinvokedynamic:
				c.loadMethodDynamicGetter(ic.Method)
			case *ir.ICinvokeinterface:
				c.loadMethodGetter(ic.Method)
			case *ir.ICinvokespecial:
				c.loadMethodGetter(ic.Method)
			case *ir.ICinvokestatic:
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
	} else {
		return nil
	}
	for i, f := range x.Class.Fields {
		if (x == c || !f.AccessFlags.Has(jcls.AccPrivate)) && f.Name() == ref.NameAndType.Name {
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
	k, err := c.loader.LoadClass(c.Class.Name())
	if err != nil {
		panic(err)
	}
	x := k.(*Class)
	for i, m := range x.Class.Methods {
		if (x == c || !m.AccessFlags.Has(jcls.AccPrivate)) && m.Name() == ref.NameAndType.Name && m.Desc().String() == ref.NameAndType.Desc {
			return &x.Methods[i]
		}
	}
	panic(fmt.Errorf("cannot load class: cannot find method %s%s", ref.NameAndType.Name, ref.NameAndType.Desc))
}

func (c *Class) loadMethodDynamicGetter(ind uint16) {
	ref, ok := c.ConstPool[ind-1].(*jcls.ConstantDynamics)
	if !ok || ref.ConstTag != jcls.TagInvokeDynamic {
		panic(fmt.Errorf("cannot load class: constant at %d is not a invoke dynamic", ind-1))
	}
	c.loadedMethods[ind] = sync.OnceValue(func() *Method {
		return c.loadMethodDynamic(ref)
	})
}

func (c *Class) loadMethodDynamic(ref *jcls.ConstantDynamics) *Method {
	panic(fmt.Errorf("TODO: Trying to invoke dynamic %#v %#v", ref, ref.NameAndType))
}
