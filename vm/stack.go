package vm

import (
	"math"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type Stack struct {
	class     *jcls.Class
	method    *jcls.Method
	pc        *ir.ICNode
	vars      []uint32
	varRefs   []*Ref
	stack     []uint32
	stackRefs []*Ref
}

var _ ir.Stack = (*Stack)(nil)

func (s *Stack) GetVarInt8(i uint16) int8 {
	return (int8)(s.vars[i])
}

func (s *Stack) GetVarInt16(i uint16) int16 {
	return (int16)(s.vars[i])
}

func (s *Stack) GetVarInt32(i uint16) int32 {
	return (int32)(s.vars[i])
}

func (s *Stack) GetVarInt64(i uint16) int64 {
	return ((int64)(s.vars[i]) << 32) | (int64)(s.vars[i+1])
}

func (s *Stack) GetVarFloat32(i uint16) float32 {
	return math.Float32frombits((uint32)(s.GetVarInt32(i)))
}

func (s *Stack) GetVarFloat64(i uint16) float64 {
	return math.Float64frombits((uint64)(s.GetVarInt64(i)))
}

func (s *Stack) GetVarRef(i uint16) ir.Ref {
	return s.varRefs[i]
}

func (s *Stack) SetVarInt8(i uint16, v int8) {
	s.vars[i] = (uint32)(v)
	s.varRefs[i] = nil
}

func (s *Stack) SetVarInt16(i uint16, v int16) {
	s.vars[i] = (uint32)(v)
	s.varRefs[i] = nil
}

func (s *Stack) SetVarInt32(i uint16, v int32) {
	s.vars[i] = (uint32)(v)
	s.varRefs[i] = nil
}

func (s *Stack) SetVarInt64(i uint16, v int64) {
	s.vars[i] = (uint32)(v >> 32)
	s.vars[i+1] = (uint32)(v)
	s.varRefs[i] = nil
	s.varRefs[i+1] = nil
}

func (s *Stack) SetVarFloat32(i uint16, v float32) {
	s.SetVarInt32(i, (int32)(math.Float32bits(v)))
}

func (s *Stack) SetVarFloat64(i uint16, v float64) {
	s.SetVarInt64(i, (int64)(math.Float64bits(v)))
}

func (s *Stack) SetVarRef(i uint16, v ir.Ref) {
	s.varRefs[i] = v.(*Ref)
}

func (s *Stack) PeekInt8() int8 {
	return (int8)(s.stack[len(s.stack)-1])
}

func (s *Stack) PeekInt16() int16 {
	return (int16)(s.stack[len(s.stack)-1])
}

func (s *Stack) PeekInt32() int32 {
	return (int32)(s.stack[len(s.stack)-1])
}

func (s *Stack) PeekInt64() int64 {
	return ((int64)(s.stack[len(s.stack)-2]) << 32) | (int64)(s.stack[len(s.stack)-1])
}

func (s *Stack) PeekFloat32() float32 {
	return math.Float32frombits((uint32)(s.PeekInt32()))
}

func (s *Stack) PeekFloat64() float64 {
	return math.Float64frombits((uint64)(s.PeekInt64()))
}

func (s *Stack) PeekRef() ir.Ref {
	return s.stackRefs[len(s.stack)-1]
}

func (s *Stack) PopInt8() int8 {
	i := len(s.stack) - 1
	v := (int8)(s.stack[i])
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	return v
}

func (s *Stack) PopInt16() int16 {
	i := len(s.stack) - 1
	v := (int16)(s.stack[i])
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	return v
}

func (s *Stack) PopInt32() int32 {
	i := len(s.stack) - 1
	v := (int32)(s.stack[i])
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	return v
}

func (s *Stack) PopInt64() int64 {
	i := len(s.stack) - 2
	v := ((int64)(s.stack[i]) << 32) | (int64)(s.stack[i+1])
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	s.stackRefs[i+1] = nil
	return v
}

func (s *Stack) PopFloat32() float32 {
	return math.Float32frombits((uint32)(s.PopInt32()))
}

func (s *Stack) PopFloat64() float64 {
	return math.Float64frombits((uint64)(s.PopInt64()))
}

func (s *Stack) PopRef() ir.Ref {
	i := len(s.stack) - 1
	v := s.stackRefs[i]
	s.stack = s.stack[:i]
	s.stackRefs[i] = nil
	return v
}

func (s *Stack) PushInt8(v int8) {
	s.stack = append(s.stack, (uint32)(v))
}

func (s *Stack) PushInt16(v int16) {
	s.stack = append(s.stack, (uint32)(v))
}

func (s *Stack) PushInt32(v int32) {
	s.stack = append(s.stack, (uint32)(v))
}

func (s *Stack) PushInt64(v int64) {
	s.stack = append(s.stack, (uint32)(v>>32), (uint32)(v))
}

func (s *Stack) PushFloat32(v float32) {
	s.PushInt32((int32)(math.Float32bits(v)))
}

func (s *Stack) PushFloat64(v float64) {
	s.PushInt64((int64)(math.Float64bits(v)))
}

func (s *Stack) PushRef(v ir.Ref) {
	i := len(s.stack)
	s.PushInt32(0)
	if n := i - len(s.stackRefs); n >= 0 {
		s.stackRefs = append(s.stackRefs, make([]*Ref, n+1)...)
	}
	s.stackRefs[i] = v.(*Ref)
}

// returns whether the top element is a reference or not
func (s *Stack) IsRef() bool {
	i := len(s.stack) - 1
	return len(s.stackRefs) > i && s.stackRefs[i] != nil
}
