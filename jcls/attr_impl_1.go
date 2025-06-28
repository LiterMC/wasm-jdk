package jcls

import (
	"bytes"
	"fmt"
	"slices"
	"strings"

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
	MaxStack    uint16
	MaxLocals   uint16
	Code        *ir.ICNode
	Exceptions  []ExceptionHandlers
	Attrs       []Attribute
	LineNumbers []*LineNumberEntry
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
	a.Attrs = make([]Attribute, 0, n)
	for range n {
		var at Attribute
		if at, err = ParseAttr(r, consts); err != nil {
			return err
		}
		if table, ok := at.(*AttrLineNumberTable); ok {
			a.LineNumbers = append(a.LineNumbers, table.Items...)
		} else {
			a.Attrs = append(a.Attrs, at)
		}
	}
	slices.SortFunc(a.LineNumbers, func(a, b *LineNumberEntry) int {
		if a.StartPc > b.StartPc {
			return 1
		}
		return -1
	})
	return nil
}

func (a *AttrCode) String() string {
	return fmt.Sprintf("%#v", a.Code)
}

func (a *AttrCode) GetLine(pc uint16) int {
	if len(a.LineNumbers) == 0 {
		return -1
	}
	ind, _ := slices.BinarySearchFunc(a.LineNumbers, pc, func(e *LineNumberEntry, pc uint16) int {
		if e.StartPc < pc {
			return -1
		}
		if e.StartPc == pc {
			return 0
		}
		return 1
	})
	if ind >= len(a.LineNumbers) {
		ind = len(a.LineNumbers) - 1
	}
	e := a.LineNumbers[ind]
	if e.StartPc > pc {
		e = a.LineNumbers[ind-1]
	}
	return (int)(e.LineNum)
}

type AttrLineNumberTable struct {
	Items []*LineNumberEntry
}

type LineNumberEntry struct {
	StartPc uint16
	LineNum uint16
}

func (*AttrLineNumberTable) Name() string { return "LineNumberTable" }
func (a *AttrLineNumberTable) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	n, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Items = make([]*LineNumberEntry, n)
	for i := range n {
		entry := new(LineNumberEntry)
		if entry.StartPc, err = readUint16(r); err != nil {
			return err
		}
		if entry.LineNum, err = readUint16(r); err != nil {
			return err
		}
		a.Items[i] = entry
	}
	return nil
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
		if n != 0 {
			c.OuterClass = consts[n-1].(*ConstantClass)
		}
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

type AttrBootstrapMethods struct {
	Methods []*BootstrapMethod
}

type BootstrapMethod struct {
	Method *ConstantMethodHandle
	Args   []ConstantInfo
}

func (*AttrBootstrapMethods) Name() string { return "BootstrapMethods" }
func (a *AttrBootstrapMethods) Parse(r *bytes.Buffer, consts []ConstantInfo) error {
	n, err := readUint16(r)
	if err != nil {
		return err
	}
	a.Methods = make([]*BootstrapMethod, n)
	for i := range len(a.Methods) {
		m := new(BootstrapMethod)
		if n, err = readUint16(r); err != nil {
			return err
		}
		m.Method = consts[n-1].(*ConstantMethodHandle)
		if n, err = readUint16(r); err != nil {
			return err
		}
		m.Args = make([]ConstantInfo, n)
		for i := range n {
			if n, err = readUint16(r); err != nil {
				return err
			}
			m.Args[i] = consts[n-1]
		}
		a.Methods[i] = m
	}
	return nil
}

func (m *BootstrapMethod) String() string {
	var sb strings.Builder
	sb.WriteString(m.Method.String())
	sb.WriteString(" -> [")
	for i, arg := range m.Args {
		if i != 0 {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", arg)
	}
	sb.WriteByte(']')
	return sb.String()
}

func init() {
	RegisterAttr(func() ParsableAttribute { return new(AttrConstantValue) })
	RegisterAttr(func() ParsableAttribute { return new(AttrCode) })
	RegisterAttr(func() ParsableAttribute { return new(AttrLineNumberTable) })
	RegisterAttr(func() ParsableAttribute { return new(AttrExceptions) })
	RegisterAttr(func() ParsableAttribute { return new(AttrInnerClasses) })
	RegisterAttr(func() ParsableAttribute { return new(AttrEnclosingMethod) })
	RegisterAttr(func() ParsableAttribute { return new(AttrSourceFile) })
	RegisterAttr(func() ParsableAttribute { return new(AttrBootstrapMethods) })
}
