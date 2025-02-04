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
	Class   string
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
		if n == 0 {
			e.Class = "java/lang/Throwable"
		} else {
			e.Class = consts[n-1].(*ConstantClass).Name
		}
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

type AttrExceptions struct {
	Exceptions []string
}

func (*AttrExceptions) Name() string { return "Exceptions" }
func (a *AttrExceptions) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	n, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Exceptions = make([]string, n)
	for i := range n {
		if n, err = readUint16(r); err != nil {
			return err
		}
		a.Exceptions[i] = consts[n-1].(*ConstantClass).Name
	}
	return nil
}
func (a *AttrExceptions) String() string {
	return fmt.Sprint(a.Exceptions)
}

type AttrInnerClasses struct {
	Classes []*InnerClassRecord
}

type InnerClassRecord struct {
	Class      *ConstantClass
	OuterClass *ConstantClass
	Name       string
	Access     AccessFlag
}

func (*AttrInnerClasses) Name() string { return "InnerClasses" }
func (a *AttrInnerClasses) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	n, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Classes = make([]*InnerClassRecord, n)
	for i := range n {
		c := new(InnerClassRecord)
		if n, err = readUint16(r); err != nil {
			return err
		}
		c.Class = consts[n-1].(*ConstantClass)
		if n, err = readUint16(r); err != nil {
			return err
		}
		c.OuterClass = consts[n-1].(*ConstantClass)
		if n, err = readUint16(r); err != nil {
			return err
		}
		if n != 0 {
			c.Name = consts[n-1].(*ConstantUtf8).Value
		}
		if n, err = readUint16(r); err != nil {
			return err
		}
		c.Access = (AccessFlag)(n)
		a.Classes[i] = c
	}
	return nil
}
func (a *AttrInnerClasses) String() string {
	return fmt.Sprintf("%#v", a.Classes)
}

type AttrEnclosingMethod struct {
	Class  *ConstantClass
	Method *ConstantNameAndType
}

func (*AttrEnclosingMethod) Name() string { return "EnclosingMethod" }
func (a *AttrEnclosingMethod) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	n, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Class = consts[n-1].(*ConstantClass)
	if n, err = readUint16(r); err != nil {
		return err
	}
	if n != 0 {
		a.Method = consts[n-1].(*ConstantNameAndType)
	}
	return nil
}
func (a *AttrEnclosingMethod) String() string {
	return a.Class.Name + "." + a.Method.String()
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
