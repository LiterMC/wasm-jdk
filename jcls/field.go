package jcls

import (
	"fmt"
	"io"
	"strings"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
)

type Field struct {
	AccessFlags AccessFlag
	name        string
	Desc        *desc.Desc
	Attrs       []ir.Attribute
}

func ParseField(r io.Reader, consts []ConstantInfo) (*Field, error) {
	f := new(Field)
	n, err := readUint16(r)
	if err != nil {
		return nil, err
	}
	f.AccessFlags = (AccessFlag)(n)
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	f.name = consts[n-1].(*ConstantUtf8).Value
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	if f.Desc, err = consts[n-1].(*ConstantUtf8).AsDesc(); err != nil {
		return nil, err
	}
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	f.Attrs = make([]ir.Attribute, n)
	for i := range n {
		if f.Attrs[i], err = ParseAttr(r, consts); err != nil {
			return nil, err
		}
	}
	return f, nil
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Modifiers() int32 {
	return (int32)(f.AccessFlags)
}

func (f *Field) IsPublic() bool {
	return f.AccessFlags.Has(AccPublic)
}

func (f *Field) IsStatic() bool {
	return f.AccessFlags.Has(AccStatic)
}

func (f *Field) String() string {
	var sb strings.Builder
	sb.WriteString(f.AccessFlags.String())
	sb.WriteString(f.Desc.String())
	sb.WriteByte(' ')
	sb.WriteString(f.name)
	fmt.Fprintf(&sb, " (%d attrs);", len(f.Attrs))
	for _, a := range f.Attrs {
		sb.WriteByte(' ')
		sb.WriteString(a.Name())
	}
	sb.WriteByte(';')
	return sb.String()
}

func (f *Field) GetAttr(name string) ir.Attribute {
	for _, a := range f.Attrs {
		if a.Name() == name {
			return a
		}
	}
	return nil
}
