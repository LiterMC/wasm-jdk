package jcls

import (
	"bytes"
	"io"

	"github.com/LiterMC/wasm-jdk/ir"
)

type ParsableAttribute interface {
	ir.Attribute
	Parse(r *bytes.Buffer, consts []ConstantInfo) error
}

var attributeFactories = make(map[string]func() ParsableAttribute)

func RegisterAttr(newer func() ParsableAttribute) {
	attr := newer()
	name := attr.Name()
	if _, ok := attributeFactories[name]; ok {
		panic("Attribute " + name + " is already registered")
	}
	attributeFactories[name] = newer
}

func ParseAttr(r io.Reader, consts []ConstantInfo) (ir.Attribute, error) {
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
	var a ParsableAttribute
	newer := attributeFactories[name]
	if newer == nil {
		a = &AttributeRaw{
			AName: name,
		}
	} else {
		a = newer()
	}
	if err = a.Parse(bytes.NewBuffer(data), consts); err != nil {
		return nil, err
	}
	return a, nil
}

type AttributeRaw struct {
	AName string
	Data  []byte
}

func (a *AttributeRaw) Name() string { return a.AName }
func (a *AttributeRaw) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	if r.Len() > 0 {
		a.Data = r.Bytes()
	}
	return nil
}
