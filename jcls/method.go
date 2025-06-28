package jcls

import (
	"fmt"
	"io"
	"strings"

	"github.com/LiterMC/wasm-jdk/desc"
)

type Method struct {
	AccessFlags AccessFlag
	name        string
	desc        *desc.MethodDesc
	Attrs       []Attribute
	Code        *AttrCode
	Exceptions  []string
}

func ParseMethod(r io.Reader, consts []ConstantInfo) (*Method, error) {
	m := new(Method)
	n, err := readUint16(r)
	if err != nil {
		return nil, err
	}
	m.AccessFlags = (AccessFlag)(n)
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	m.name = consts[n-1].(*ConstantUtf8).Value
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	if m.desc, err = consts[n-1].(*ConstantUtf8).AsMethodDesc(); err != nil {
		return nil, err
	}
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	m.Attrs = make([]Attribute, n)
	for i := range n {
		var a Attribute
		if a, err = ParseAttr(r, consts); err != nil {
			return nil, err
		}
		m.Attrs[i] = a
		switch a := a.(type) {
		case *AttrCode:
			m.Code = a
		case *AttrExceptions:
			m.Exceptions = a.Exceptions
		}
	}
	return m, nil
}

func NewMethod(flags AccessFlag, name string, descriptor *desc.MethodDesc, attrs []Attribute) *Method {
	m := new(Method)
	m.AccessFlags = flags
	m.name = name
	m.desc = descriptor
	m.Attrs = attrs
	return m
}

func (m *Method) Name() string {
	return m.name
}

func (m *Method) Desc() *desc.MethodDesc {
	return m.desc
}

func (m *Method) Modifiers() int32 {
	return (int32)(m.AccessFlags)
}

func (m *Method) IsPublic() bool {
	return m.AccessFlags.Has(AccPublic)
}

func (m *Method) IsStatic() bool {
	return m.AccessFlags.Has(AccStatic)
}

func (m *Method) IsConstructor() bool {
	return !m.IsStatic() && m.name == "<init>"
}

func (m *Method) String() string {
	var sb strings.Builder
	sb.WriteString(m.AccessFlags.String())
	sb.WriteString(m.name)
	sb.WriteString(m.desc.String())
	fmt.Fprintf(&sb, " (%d attrs);", len(m.Attrs))
	for _, a := range m.Attrs {
		sb.WriteByte(' ')
		sb.WriteString(a.Name())
	}
	sb.WriteByte(';')
	return sb.String()
}

func (m *Method) GetAttr(name string) Attribute {
	for _, a := range m.Attrs {
		if a.Name() == name {
			return a
		}
	}
	return nil
}
