package jcls

import (
	"fmt"
	"io"
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
	c.ConstPool = make([]ConstantInfo, n)
	for i := range n {
		if c.ConstPool[i], err = ParseConstant(r); err != nil {
			return nil, err
		}
	}
	for i := range n {
		c.ConstPool[i].Resolve(c.ConstPool)
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
	c.Attrs = make([]Attribute, n)
	for i := range n {
		if c.Attrs[i], err = ParseAttr(r, c.ConstPool); err != nil {
			return nil, err
		}
	}
	return c, nil
}
