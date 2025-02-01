package jcls

import (
	"fmt"
	"io"
	"strings"

	"github.com/LiterMC/wasm-jdk/desc"
)

type Field struct {
	AccessFlags AccessFlag
	Name        string
	Desc        *desc.Desc
	Attrs       []Attribute
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
	f.Name = consts[n-1].(*ConstantUtf8).Value
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	if f.Desc, err = consts[n-1].(*ConstantUtf8).AsDesc(); err != nil {
		return nil, err
	}
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	f.Attrs = make([]Attribute, n)
	for i := range n {
		if f.Attrs[i], err = ParseAttr(r, consts); err != nil {
			return nil, err
		}
	}
	return f, nil
}

func (f *Field) String() string {
	var sb strings.Builder
	sb.WriteString(f.AccessFlags.String())
	sb.WriteString(f.Desc.String())
	sb.WriteByte(' ')
	sb.WriteString(f.Name)
	fmt.Fprintf(&sb, " (%d attrs);", len(f.Attrs))
	for _, a := range f.Attrs {
		sb.WriteByte(' ')
		sb.WriteString(a.Name())
	}
	sb.WriteByte(';')
	return sb.String()
}

func (f *Field) GetAttr(name string) Attribute {
	for _, a := range f.Attrs {
		if a.Name() == name {
			return a
		}
	}
	return nil
}
