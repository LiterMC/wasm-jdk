package jcls

import (
	"bytes"
	"fmt"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ir/parser"
)

type AttrConstantValue struct {
	Value ConstantInfo
}

func (*AttrConstantValue) Name() string { return "ConstantValue" }
func (a *AttrConstantValue) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	ind, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Value = consts[ind-1]
	return nil
}
func (a *AttrConstantValue) String() string {
	return fmt.Sprint(a.Value)
}

type AttrCode struct {
	MaxStack   uint16
	MaxLocals  uint16
	Code       *ir.ICNode
	Exceptions []ExceptionHandlers
	Attrs      []Attribute
}

type ExceptionHandlers struct {
	Start   uint16
	End     uint16
	Handler uint16
	Type    *ConstantClass
}

func (*AttrCode) Name() string { return "Code" }
func (a *AttrCode) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	var (
		n    uint16
		size uint32
		err  error
	)
	if a.MaxStack, err = readUint16(r); err != nil {
		return err
	}
	if a.MaxLocals, err = readUint16(r); err != nil {
		return err
	}

	if size, err = readUint32(r); err != nil {
		return err
	}
	if a.Code, err = parser.ParseCode(r.Next((int)(size))); err != nil {
		return err
	}

	if n, err = readUint16(r); err != nil {
		return err
	}
	a.Exceptions = make([]ExceptionHandlers, n)
	for i := range n {
		var e ExceptionHandlers
		if e.Start, err = readUint16(r); err != nil {
			return err
		}
		if e.End, err = readUint16(r); err != nil {
			return err
		}
		if e.Handler, err = readUint16(r); err != nil {
			return err
		}
		if n, err = readUint16(r); err != nil {
			return err
		}
		e.Type = consts[n-1].(*ConstantClass)
		a.Exceptions[i] = e
	}

	if n, err = readUint16(r); err != nil {
		return err
	}
	a.Attrs = make([]Attribute, n)
	for i := range n {
		if a.Attrs[i], err = ParseAttr(r, consts); err != nil {
			return err
		}
	}
	return nil
}
func (a *AttrCode) String() string {
	return fmt.Sprintf("%#v", a.Code)
}

type AttrSourceFile struct {
	Value string
}

func (*AttrSourceFile) Name() string { return "SourceFile" }
func (a *AttrSourceFile) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	ind, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Value = consts[ind-1].(*ConstantUtf8).Value
	return nil
}
func (a *AttrSourceFile) String() string {
	return a.Value
}

func init() {
	RegisterAttr(func() ParsableAttribute { return new(AttrConstantValue) })
	RegisterAttr(func() ParsableAttribute { return new(AttrCode) })
	RegisterAttr(func() ParsableAttribute { return new(AttrSourceFile) })
}
