package vm

import (
	"fmt"
	"reflect"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

// vm.Class represents a loaded class
type Class struct {
	*jcls.Class
	vm *VM

	super      ir.Class
	interfaces []ir.Class

	refType reflect.Type
	Fields  []Field
	Methods []Method

	loadedFields  map[uint16]*Field
	loadedMethods map[uint16]*Method
}

var _ ir.Class = (*Class)(nil)

func loadClass(cls *jcls.Class, vm *VM) *Class {
	c := &Class{
		Class: cls,
		vm:    vm,
	}
	var err error
	if c.super, err = vm.loader.LoadClass(c.SuperSym.Name); err != nil {
		panic(err)
	}
	c.interfaces = make([]ir.Class, len(c.InterfacesSym))
	for i, in := range c.InterfacesSym {
		if c.interfaces[i], err = vm.loader.LoadClass(in.Name); err != nil {
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
	c.Fields = make([]Field, len(c.Class.Fields))
	for i, f := range c.Class.Fields {
		cf := &c.Fields[i]
		cf.Field = f
		cf.class = c
		cf.offset = c.refType.Field(i + 1).Offset
		if f.Desc.EndType == desc.Class {
			if cf.typ, err = vm.loader.LoadClass(f.Desc.Class); err != nil {
				panic(err)
			}
		}
	}

	// parseFields
	c.loadedFields = make(map[uint16]*Field)
	c.loadedMethods = make(map[uint16]*Method)
	c.scanCodes()
	return c
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
	return c.loadedFields[i]
}

func (c *Class) GetMethod(i uint16) ir.Method {
	return c.loadedMethods[i]
}

func (c *Class) GetMethodByName(name string, desc *desc.MethodDesc) ir.Method {
	for i := range len(c.Methods) {
		m := &c.Methods[i]
		if !m.AccessFlags.Has(jcls.AccPrivate) && m.Name() == name && m.Desc.EqInputs(desc) {
			if !m.Desc.Output.Eq(desc.Output) {
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
				c.loadField(ic.Field)
			case *ir.ICgetstatic:
				c.loadField(ic.Field)
			case *ir.ICinvokedynamic:
				c.loadMethod(ic.Method)
			case *ir.ICinvokeinterface:
				c.loadMethod(ic.Method)
			case *ir.ICinvokespecial:
				c.loadMethod(ic.Method)
			case *ir.ICinvokestatic:
				c.loadMethod(ic.Method)
			}
		}
	}
}

func (c *Class) loadField(ind uint16) {
	nat := c.ConstPool[ind-1].(*jcls.ConstantNameAndType)
	x := c
	for i, f := range x.Class.Fields {
		if !f.AccessFlags.Has(jcls.AccPrivate) && f.Name() == nat.Name {
			if f.Desc.String() != nat.Desc {
				panic(fmt.Errorf("cannot load class: field %s is %s, but one operation requires %s", nat.Name, f.Desc.String(), nat.Desc))
			}
			c.loadedFields[ind] = &x.Fields[i]
			return
		}
	}
}

func (c *Class) loadMethod(ind uint16) {
}
