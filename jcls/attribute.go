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
func (a *AttributeRaw) Parse(r io.Reader, consts []ConstantInfo) error {
	name, err := readUint16(r)
	if err != nil {
		return err
	}
	a.AName = consts[name].(*ConstantUtf8).Value
	size, err := readUint32(r)
	if err != nil {
		return err
	}
	a.Data = make([]byte, size)
	if _, err = io.ReadFull(r, a.Data); err != nil {
		return err
	}
	return nil
}
