package jcls

import (
	"io"
)

type Attribute interface {
	Name() string
}

type AttributeRaw struct {
	AName string
	Data  []byte
}

func (a *AttributeRaw) Name() string { return a.AName }

func ParseAttr(r io.Reader, consts []ConstantInfo) (Attribute, error) {
	nameInd, err := readUint16(r)
	if err != nil {
		return nil, err
	}
	name := consts[nameInd-1].(*ConstantUtf8).Value
	size, err := readUint32(r)
	if err != nil {
		return nil, err
	}
	data := make([]byte, size)
	if _, err = io.ReadFull(r, data); err != nil {
		return nil, err
	}
	switch name {
	// TODO:
	}
	return &AttributeRaw{
		AName: name,
		Data:  data,
	}, nil
}
