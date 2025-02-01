package jcls

import (
	"fmt"
	"io"
	"strings"
)

const ClassMagic uint32 = 0xCAFEBABE

type Class struct {
	Magic       uint32
	Minor       uint16
	Major       uint16
	ConstPool   []ConstantInfo
	AccessFlags AccessFlag
	This        *ConstantClass
	Super       *ConstantClass
	Interfaces  []*ConstantClass
	Fields      []*Field
	Methods     []*Method
	Attrs       []Attribute
}

func ParseClass(r io.Reader) (*Class, error) {
	var (
		c   = new(Class)
		n   uint16
		err error
	)
	if c.Magic, err = readUint32(r); err != nil {
		return nil, err
	}
	if c.Magic != ClassMagic {
		return nil, fmt.Errorf("Unexpected class header 0x%08x", c.Magic)
	}
	if c.Minor, err = readUint16(r); err != nil {
		return nil, err
	}
	if c.Major, err = readUint16(r); err != nil {
		return nil, err
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	n--
	c.ConstPool = make([]ConstantInfo, n)
	for i := (uint16)(0); i < n; i++ {
		v, err := ParseConstant(r)
		if err != nil {
			return nil, err
		}
		c.ConstPool[i] = v
		if v.IsWide() {
			i++
		}
	}
	for _, v := range c.ConstPool {
		if v != nil {
			v.Resolve(c.ConstPool)
		}
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.AccessFlags = (AccessFlag)(n)

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.This = c.ConstPool[n-1].(*ConstantClass)

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	if n != 0 {
		c.Super = c.ConstPool[n-1].(*ConstantClass)
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.Interfaces = make([]*ConstantClass, n)
	for i := range n {
		if n, err = readUint16(r); err != nil {
			return nil, err
		}
		c.Interfaces[i] = c.ConstPool[n-1].(*ConstantClass)
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.Fields = make([]*Field, n)
	for i := range n {
		if c.Fields[i], err = ParseField(r, c.ConstPool); err != nil {
			return nil, err
		}
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.Methods = make([]*Method, n)
	for i := range n {
		if c.Methods[i], err = ParseMethod(r, c.ConstPool); err != nil {
			return nil, err
		}
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.Attrs = make([]Attribute, n)
	for i := range n {
		if c.Attrs[i], err = ParseAttr(r, c.ConstPool); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Class) String() string {
	var sb strings.Builder
	sb.WriteString("Class ")
	sb.WriteString(c.AccessFlags.String())
	sb.WriteString(c.This.Name)
	if c.Super != nil {
		sb.WriteString(" extends ")
		sb.WriteString(c.Super.Name)
	}
	if len(c.Interfaces) > 0 {
		sb.WriteString(" implements ")
		for i, it := range c.Interfaces {
			if i != 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(it.Name)
		}
	}
	sb.WriteString(" {\n")
	fmt.Fprintf(&sb, "  Fields: (%d)\n", len(c.Fields))
	for _, f := range c.Fields {
		sb.WriteString("    ")
		sb.WriteString(f.String())
		sb.WriteByte('\n')
	}
	fmt.Fprintf(&sb, "  Methods: (%d)\n", len(c.Methods))
	for _, m := range c.Methods {
		sb.WriteString("    ")
		sb.WriteString(m.String())
		sb.WriteByte('\n')
	}
	fmt.Fprintf(&sb, "  Attributes: (%d)\n", len(c.Attrs))
	for _, a := range c.Attrs {
		sb.WriteString("    ")
		sb.WriteString(a.Name())
		sb.WriteByte('\n')
	}
	sb.WriteByte('}')
	return sb.String()
}
