package ir

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ops"
)

type Ref = unsafe.Pointer

// Intermediate Cide Representation
type IC interface {
	// The operation code
	Op() ops.Op
	// Execute
	Execute(vm VM) error
}

type ICNode struct {
	IC
	Offset int32
	Next   *ICNode
	Nexts  []*ICNode
}

type VM interface {
	GetStack() Stack

	New(Class) Ref
	NewArrInt8(int32) []int8
	NewArrInt16(int32) []int16
	NewArrInt32(int32) []int32
	NewArrInt64(int32) []int64
	NewArrRef(Class, int32) []Ref
	NewArrRefMultiDim(Class, []int32) []Ref

	GetObjectClass() Class
	GetThrowableClass() Class
	GetClassByIndex(uint16) (Class, error)
	GetClass(Ref) Class
	GetArrClass([]Ref) Class

	GetCurrentClass() Class
	GetCurrentMethod() Method
	Invoke(Method, Ref)
	InvokeStatic(Method)
	Return()
	Throw(Ref)
	Goto(*ICNode)

	MonitorLock(Ref) error
	MonitorUnlock(Ref) error
}

type Stack interface {
	GetVarInt8(uint16) int8
	GetVarInt16(uint16) int16
	GetVarInt32(uint16) int32
	GetVarInt64(uint16) int64
	GetVarFloat32(uint16) float32
	GetVarFloat64(uint16) float64
	GetVarRef(uint16) Ref
	SetVarInt8(uint16, int8)
	SetVarInt16(uint16, int16)
	SetVarInt32(uint16, int32)
	SetVarInt64(uint16, int64)
	SetVarFloat32(uint16, float32)
	SetVarFloat64(uint16, float64)
	SetVarRef(uint16, Ref)

	PeekInt8() int8
	PeekInt16() int16
	PeekInt32() int32
	PeekInt64() int64
	PeekFloat32() float32
	PeekFloat64() float64
	PeekRef() Ref
	PopInt8() int8
	PopInt16() int16
	PopInt32() int32
	PopInt64() int64
	PopFloat32() float32
	PopFloat64() float64
	PopRef() Ref
	PushInt8(int8)
	PushInt16(int16)
	PushInt32(int32)
	PushInt64(int64)
	PushFloat32(float32)
	PushFloat64(float64)
	PushRef(Ref)

	PopArrInt8() []int8
	PopArrInt16() []int16
	PopArrInt32() []int32
	PopArrInt64() []int64
	PopArrFloat32() []float32
	PopArrFloat64() []float64
	PopArrRef() []Ref
	PushArrInt8([]int8)
	PushArrInt16([]int16)
	PushArrInt32([]int32)
	PushArrInt64([]int64)
	PushArrFloat32([]float32)
	PushArrFloat64([]float64)
	PushArrRef([]Ref)

	// returns whether the top element is a reference or not
	IsRef() bool
}

type Class interface {
	IsAssignableFrom(Class) bool
	IsInstance(Ref) bool
	Name() string
	GetAndPushConst(uint16, Stack) error
	GetField(uint16) Field
	GetMethod(uint16) Method
}

type Field interface {
	GetDeclaringClass() Class
	IsStatic() bool
	Modifiers() int

	Name() string
	Type() Class
	GetAndPush(Ref, Stack)
	PopAndSet(Ref, Stack)
}

type Method interface {
	GetDeclaringClass() Class
	IsStatic() bool
	Modifiers() int

	Name() string
	ParameterTypes() []Class
	ReturnType() Class
}
