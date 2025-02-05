package ir

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ops"
)

// Intermediate Code Representation
type IC interface {
	// The operation code
	Op() ops.Op
	// Execute
	Execute(vm VM) error
}

type ICJumpable interface {
	Offsets() []int32
	SetNode(i int, n *ICNode)
}

type ICNode struct {
	IC
	Next   *ICNode
	Offset int32
}

type VM interface {
	GetStack() Stack

	New(Class) Ref
	NewString(string) Ref
	NewArray(*desc.Desc, int32) Ref
	NewArrayByClass(Class, int32) Ref
	NewArrayMultiDim(*desc.Desc, []int32) Ref

	GetObjectClass() Class
	GetThrowableClass() Class
	GetStringClass() Class
	GetString(Ref) string
	GetStringIntern(Ref) Ref
	GetClassRef(Class) Ref

	GetDesc(uint16) *desc.Desc
	GetClassByIndex(uint16) (Class, error)
	GetClass(Ref) Class

	GetClassLoader() ClassLoader
	GetCurrentClass() Class
	GetCurrentMethod() Method
	LoadNativeMethod(Method, func(VM) error)
	Invoke(Method, Ref)
	InvokeStatic(Method)
	InvokeDynamic(uint16) error
	Return()
	Throw(Ref)
	Goto(*ICNode)

	MonitorLock(Ref) error
	MonitorUnlock(Ref) error
}

type Stack interface {
	GetVar(uint16) uint32
	GetVar64(uint16) uint64
	GetVarInt8(uint16) int8
	GetVarInt16(uint16) int16
	GetVarInt32(uint16) int32
	GetVarInt64(uint16) int64
	GetVarFloat32(uint16) float32
	GetVarFloat64(uint16) float64
	GetVarRef(uint16) Ref
	SetVar(uint16, uint32)
	SetVar64(uint16, uint64)
	SetVarInt8(uint16, int8)
	SetVarInt16(uint16, int16)
	SetVarInt32(uint16, int32)
	SetVarInt64(uint16, int64)
	SetVarFloat32(uint16, float32)
	SetVarFloat64(uint16, float64)
	SetVarRef(uint16, Ref)

	Peek() uint32
	Peek64() uint64
	PeekInt8() int8
	PeekInt16() int16
	PeekInt32() int32
	PeekInt64() int64
	PeekFloat32() float32
	PeekFloat64() float64
	PeekRef() Ref
	Pop() uint32
	Pop64() uint64
	PopInt8() int8
	PopInt16() int16
	PopInt32() int32
	PopInt64() int64
	PopFloat32() float32
	PopFloat64() float64
	PopRef() Ref
	Push(uint32)
	Push64(uint64)
	PushInt8(int8)
	PushInt16(int16)
	PushInt32(int32)
	PushInt64(int64)
	PushFloat32(float32)
	PushFloat64(float64)
	PushRef(Ref)

	// returns whether the top element is a reference or not
	IsRef() bool
}

type ClassLoader interface {
	LoadClass(name string) (Class, error)
}
