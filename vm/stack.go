package vm

import (
	"fmt"
	"math"
	"strings"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type Stack struct {
	prev      *Stack
	class     *Class
	method    ir.Method
	pc        *ir.ICNode
	vars      []uint32
	varRefs   []*Ref
	stack     []uint32
	stackRefs []*Ref
}

var _ ir.Stack = (*Stack)(nil)

func (s *Stack) Prev() ir.Stack {
	return s.prev
}

func (s *Stack) Method() ir.Method {
	return s.method
}

func (s *Stack) PC() *ir.ICNode {
	return s.pc
}

func (s *Stack) GetVar(i uint16) uint32 {
	return s.vars[i]
}

func (s *Stack) GetVar64(i uint16) uint64 {
	return ((uint64)(s.vars[i]) << 32) | (uint64)(s.vars[i+1])
}

func (s *Stack) GetVarInt8(i uint16) int8 {
	return (int8)(s.GetVar(i))
}

func (s *Stack) GetVarInt16(i uint16) int16 {
	return (int16)(s.GetVar(i))
}

func (s *Stack) GetVarInt32(i uint16) int32 {
	return (int32)(s.GetVar(i))
}

func (s *Stack) GetVarInt64(i uint16) int64 {
	return (int64)(s.GetVar64(i))
}

func (s *Stack) GetVarFloat32(i uint16) float32 {
	return math.Float32frombits(s.GetVar(i))
}

func (s *Stack) GetVarFloat64(i uint16) float64 {
	return math.Float64frombits(s.GetVar64(i))
}

func (s *Stack) GetVarRef(i uint16) ir.Ref {
	v := s.varRefs[i]
	if v == nil {
		return nil
	}
	return v
}

func (s *Stack) GetVarPointer(i uint16) unsafe.Pointer {
	return (unsafe.Pointer)(s.varRefs[i])
}

func (s *Stack) SetVar(i uint16, v uint32) {
	if n := (int)(i) - len(s.vars) + 1; n > 0 {
		s.vars = append(s.vars, make([]uint32, n)...)
		s.varRefs = append(s.varRefs, make([]*Ref, n)...)
	}
	s.vars[i] = v
	s.varRefs[i] = nil
}

func (s *Stack) SetVar64(i uint16, v uint64) {
	if n := (int)(i) - len(s.vars) + 2; n > 0 {
		s.vars = append(s.vars, make([]uint32, n)...)
		s.varRefs = append(s.varRefs, make([]*Ref, n)...)
	}
	s.vars[i] = (uint32)(v >> 32)
	s.vars[i+1] = (uint32)(v)
	s.varRefs[i] = nil
	s.varRefs[i+1] = nil
}

func (s *Stack) SetVarInt8(i uint16, v int8) {
	s.SetVarInt32(i, (int32)(v))
}

func (s *Stack) SetVarInt16(i uint16, v int16) {
	s.SetVarInt32(i, (int32)(v))
}

func (s *Stack) SetVarInt32(i uint16, v int32) {
	s.SetVar(i, (uint32)(v))
}

func (s *Stack) SetVarInt64(i uint16, v int64) {
	s.SetVar64(i, (uint64)(v))
}

func (s *Stack) SetVarFloat32(i uint16, v float32) {
	s.SetVar(i, math.Float32bits(v))
}

func (s *Stack) SetVarFloat64(i uint16, v float64) {
	s.SetVar64(i, math.Float64bits(v))
}

func (s *Stack) SetVarRef(i uint16, v ir.Ref) {
	s.SetVar(i, 0)
	if v == nil {
		s.varRefs[i] = nil
	} else {
		s.varRefs[i] = v.(*Ref)
	}
}

func (s *Stack) SetVarPointer(i uint16, v unsafe.Pointer) {
	s.SetVar(i, 0)
	s.varRefs[i] = (*Ref)(v)
}

func (s *Stack) Peek() uint32 {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Peek64() uint64 {
	return ((uint64)(s.stack[len(s.stack)-2]) << 32) | (uint64)(s.stack[len(s.stack)-1])
}

func (s *Stack) PeekInt8() int8 {
	return (int8)(s.Peek())
}

func (s *Stack) PeekInt16() int16 {
	return (int16)(s.Peek())
}

func (s *Stack) PeekInt32() int32 {
	return (int32)(s.Peek())
}

func (s *Stack) PeekInt64() int64 {
	return (int64)(s.Peek64())
}

func (s *Stack) PeekFloat32() float32 {
	return math.Float32frombits(s.Peek())
}

func (s *Stack) PeekFloat64() float64 {
	return math.Float64frombits(s.Peek64())
}

func (s *Stack) PeekRef() ir.Ref {
	v := s.stackRefs[len(s.stack)-1]
	if v == nil {
		return nil
	}
	return v
}

func (s *Stack) PeekPointer() unsafe.Pointer {
	return (unsafe.Pointer)(s.stackRefs[len(s.stack)-1])
}

func (s *Stack) Pop() uint32 {
	i := len(s.stack) - 1
	v := s.stack[i]
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	s.stackRefs = s.stackRefs[:i]
	return v
}

func (s *Stack) Pop64() uint64 {
	i := len(s.stack) - 2
	v := ((uint64)(s.stack[i]) << 32) | (uint64)(s.stack[i+1])
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	s.stackRefs[i+1] = nil
	s.stackRefs = s.stackRefs[:i]
	return v
}

func (s *Stack) PopInt8() int8 {
	return (int8)(s.Pop())
}

func (s *Stack) PopInt16() int16 {
	return (int16)(s.Pop())
}

func (s *Stack) PopInt32() int32 {
	return (int32)(s.Pop())
}

func (s *Stack) PopInt64() int64 {
	return (int64)(s.Pop64())
}

func (s *Stack) PopFloat32() float32 {
	return math.Float32frombits(s.Pop())
}

func (s *Stack) PopFloat64() float64 {
	return math.Float64frombits(s.Pop64())
}

func (s *Stack) PopRef() ir.Ref {
	i := len(s.stack) - 1
	v := s.stackRefs[i]
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	s.stackRefs = s.stackRefs[:i]
	if v == nil {
		return nil
	}
	return v
}

func (s *Stack) PopPointer() unsafe.Pointer {
	i := len(s.stack) - 1
	v := s.stackRefs[i]
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	s.stackRefs = s.stackRefs[:i]
	return (unsafe.Pointer)(v)
}

func (s *Stack) Push(v uint32) {
	s.stack = append(s.stack, v)
	s.stackRefs = append(s.stackRefs, nil)
}

func (s *Stack) Push64(v uint64) {
	s.stack = append(s.stack, (uint32)(v>>32), (uint32)(v))
	s.stackRefs = append(s.stackRefs, nil, nil)
}

func (s *Stack) PushInt8(v int8) {
	s.PushInt32((int32)(v))
}

func (s *Stack) PushInt16(v int16) {
	s.PushInt32((int32)(v))
}

func (s *Stack) PushInt32(v int32) {
	s.Push((uint32)(v))
}

func (s *Stack) PushInt64(v int64) {
	s.Push64((uint64)(v))
}

func (s *Stack) PushFloat32(v float32) {
	s.Push(math.Float32bits(v))
}

func (s *Stack) PushFloat64(v float64) {
	s.Push64(math.Float64bits(v))
}

func (s *Stack) PushRef(v ir.Ref) {
	i := len(s.stack)
	s.Push(0)
	if v != nil {
		s.stackRefs[i] = v.(*Ref)
	}
}

func (s *Stack) PushPointer(v unsafe.Pointer) {
	i := len(s.stack)
	s.Push(0)
	s.stackRefs[i] = (*Ref)(v)
}

// returns true only if the top element is a non-null reference
func (s *Stack) IsRef() bool {
	i := len(s.stack) - 1
	return len(s.stackRefs) > i && s.stackRefs[i] != nil
}

func (s *Stack) GoString() string {
	var sb strings.Builder
	sb.WriteString("Stack {\n")
	sb.WriteString("  Vars:\n")
	for i, v := range s.vars {
		fmt.Fprintf(&sb, "    0x%02x: ", i)
		if r := s.varRefs[i]; r != nil {
			sb.WriteString(r.GoString())
		} else {
			fmt.Fprintf(&sb, "%d", v)
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("  Stack:\n")
	for i, v := range s.stack {
		sb.WriteString("    - ")
		if r := s.stackRefs[i]; r != nil {
			sb.WriteString(r.GoString())
		} else {
			fmt.Fprintf(&sb, "%d", v)
		}
		sb.WriteByte('\n')
	}
	sb.WriteByte('}')
	return sb.String()
}

type StackInfo struct {
	Frames     []StackFrameInfo
	TotalDepth int
}

type StackFrameInfo struct {
	Method *Method
	PC     *ir.ICNode
}

func NewStackInfo(vm *VM, stack ir.Stack, depth int) *StackInfo {
	var frames []StackFrameInfo
	if depth == -1 {
		frames = make([]StackFrameInfo, 0, 8)
	} else {
		frames = make([]StackFrameInfo, 0, depth)
	}
	si := new(StackInfo)

	throwable := vm.javaLangThrowable
	s := stack
	for throwable.IsAssignableFrom(s.Method().GetDeclaringClass()) && s.Method().Name() == "<init>" {
		s = s.Prev()
	}
	for ; s != nil; s = s.Prev() {
		if s.Method() == nil { // JVM initialization stack
			break
		}
		if si.TotalDepth++; si.TotalDepth <= depth {
			continue
		}
		frames = append(frames, StackFrameInfo{
			Method: s.Method().(*Method),
			PC:     s.PC(),
		})
	}
	si.Frames = frames
	return si
}

func (s *StackInfo) String() string {
	var sb strings.Builder
	more := s.TotalDepth - len(s.Frames)
	for _, f := range s.Frames {
		sb.WriteString("  at ")
		sb.WriteString(f.String())
		sb.WriteByte('\n')
	}
	if more > 0 {
		fmt.Fprintf(&sb, " ... %d more", more)
	}
	return sb.String()
}

func (fi *StackFrameInfo) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s.%s%s (",
		fi.Method.class.Name(),
		fi.Method.Name(),
		fi.Method.Desc().String())
	if sourceFile, ok := fi.Method.class.GetAttr("SourceFile").(*jcls.AttrSourceFile); ok {
		sb.WriteString(sourceFile.String())
		sb.WriteByte(':')
	}
	if fi.Method.AccessFlags.Has(jcls.AccNative) {
		sb.WriteString("native")
	} else {
		line := fi.Method.Code.GetLine((uint16)(fi.PC.Offset))
		if line >= 0 {
			fmt.Fprintf(&sb, "%d", line)
		} else {
			fmt.Fprintf(&sb, "0x%04x", fi.PC.Offset)
		}
	}
	sb.WriteByte(')')
	return sb.String()
}
