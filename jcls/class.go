package jcls

import (
	"fmt"
	"io"
	"strings"

	"github.com/LiterMC/wasm-jdk/desc"
)

const ClassMagic uint32 = 0xCAFEBABE

// jcls.Class represents a parsed class
type Class struct {
	Minor         uint16
	Major         uint16
	ConstPool     []ConstantInfo
	AccessFlags   AccessFlag
	ThisSym       *ConstantClass
	SuperSym      *ConstantClass
	InterfacesSym []*ConstantClass
	Fields        []*Field
	Methods       []*Method
	Attrs         []Attribute

	ThisDesc *desc.Desc
}

func ParseClass(r io.Reader) (*Class, error) {
	var (
		c     = new(Class)
		n     uint16
		magic uint32
		err   error
	)
	if magic, err = readUint32(r); err != nil {
		return nil, err
	}
	if magic != ClassMagic {
		return nil, fmt.Errorf("Unexpected class header 0x%08x", magic)
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
	c.ThisSym = c.ConstPool[n-1].(*ConstantClass)

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	if n != 0 {
		c.SuperSym = c.ConstPool[n-1].(*ConstantClass)
	}

	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	c.InterfacesSym = make([]*ConstantClass, n)
	for i := range n {
		if n, err = readUint16(r); err != nil {
			return nil, err
		}
		c.InterfacesSym[i] = c.ConstPool[n-1].(*ConstantClass)
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

	c.ThisDesc = &desc.Desc{
		EndType: desc.Class,
		Class:   c.ThisSym.Name,
	}
	return c, nil
}

func NewClass(flags AccessFlag, name string, super string, interfaces []string, fields []*Field, methods []*Method, attrs []Attribute) *Class {
	c := new(Class)
	c.AccessFlags = flags
	c.ThisSym = &ConstantClass{
		Name: name,
	}
	c.SuperSym = &ConstantClass{
		Name: super,
	}
	c.InterfacesSym = make([]*ConstantClass, len(interfaces))
	for i, name := range interfaces {
		c.InterfacesSym[i] = &ConstantClass{
			Name: name,
		}
	}
	c.Fields = fields
	c.Methods = methods
	c.Attrs = attrs
	c.ThisDesc = &desc.Desc{
		EndType: desc.Class,
		Class:   c.ThisSym.Name,
	}
	return c
}

func (c *Class) String() string {
	var sb strings.Builder
	sb.WriteString("Class ")
	sb.WriteString(c.AccessFlags.String())
	sb.WriteString(c.ThisDesc.Class)
	if c.SuperSym != nil {
		sb.WriteString(" extends ")
		sb.WriteString(c.SuperSym.Name)
	}
	if len(c.InterfacesSym) > 0 {
		sb.WriteString(" implements ")
		for i, it := range c.InterfacesSym {
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
		fmt.Fprintf(&sb, ": %v", a)
		sb.WriteByte('\n')
	}
	sb.WriteByte('}')
	return sb.String()
}

func (c *Class) Name() string {
	return c.ThisDesc.Class
}

func (c *Class) Desc() *desc.Desc {
	return c.ThisDesc
}

func (c *Class) Modifiers() int32 {
	return (int32)(c.AccessFlags)
}

func (c *Class) IsInterface() bool {
	return c.AccessFlags.Has(AccInterface)
}

func (c *Class) GetAttr(name string) Attribute {
	for _, a := range c.Attrs {
		if a.Name() == name {
			return a
		}
	}
	return nil
}
