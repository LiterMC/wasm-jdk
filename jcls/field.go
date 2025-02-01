package jcls

import (
	"fmt"
	"io"
	"strings"
)

type Field struct {
	AccessFlags AccessFlag
	Name        string
	Desc        string
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
	f.Desc = consts[n-1].(*ConstantUtf8).Value
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
	sb.WriteString(f.Desc)
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
