package jcls

import (
	"io"
)

type Method struct {
	AccessFlags AccessFlag
	Name        string
	Desc        string
	Attrs       []Attribute
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
	m.Name = consts[n-1].(*ConstantUtf8).Value
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	m.Desc = consts[n-1].(*ConstantUtf8).Value
	if n, err = readUint16(r); err != nil {
		return nil, err
	}
	m.Attrs = make([]Attribute, n)
	for i := range n {
		if m.Attrs[i], err = ParseAttr(r, consts); err != nil {
			return nil, err
		}
	}
	return m, nil
}
