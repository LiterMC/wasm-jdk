package vm

import (
	"math"

	"github.com/LiterMC/wasm-jdk/ir"
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
	return s.stackRefs[len(s.stack)-1]
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

// returns whether the top element is a reference or not
func (s *Stack) IsRef() bool {
	i := len(s.stack) - 1
	return len(s.stackRefs) > i && s.stackRefs[i] != nil
}
