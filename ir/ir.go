package ir

import (
	"unsafe"

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
	Running() bool
	Step() error

	New(Class) Ref
	NewString(string) Ref
	NewArray(*desc.Desc, int32) Ref
	NewArrayMultiDim(*desc.Desc, []int32) Ref
	NewObjectArray(Class, int32) Ref

	RefToPtr(Ref) unsafe.Pointer
	PtrToRef(unsafe.Pointer) Ref

	GetObjectClass() Class
	GetThrowableClass() Class
	GetStringClass() Class
	GetString(Ref) string
	GetStringIntern(Ref) Ref
	GetStringInternOrNew(string) Ref
	GetClassRef(Class) Ref

	GetDesc(uint16) *desc.Desc
	GetClassByIndex(uint16) (Class, error)
	GetClass(Ref) Class

	GetBootLoader() ClassLoader
	GetClassLoader() ClassLoader
	GetCurrentClass() Class
	GetCurrentMethod() Method

	LoadNativeMethod(Method, func(VM) error)
	Invoke(Method)
	InvokeStatic(Method)
	InvokeVirtual(Method)
	InvokeDynamic(uint16) error

	Return()
	Throw(Ref)
	Throwing() Ref
	Goto(*ICNode)

	GetCarrierThread() Ref
	GetCurrentThread() Ref
	SetCurrentThread(Ref)
	Interrupt(Ref)
	ClearInterrupt()

	FillThrowableStackTrace(Ref)

	NewSubVM(Ref) VM
}

type Stack interface {
	Prev() Stack
	Method() Method
	PC() *ICNode

	GetVar(uint16) uint32
	GetVar64(uint16) uint64
	GetVarInt8(uint16) int8
	GetVarInt16(uint16) int16
	GetVarInt32(uint16) int32
	GetVarInt64(uint16) int64
	GetVarFloat32(uint16) float32
	GetVarFloat64(uint16) float64
	GetVarRef(uint16) Ref
	GetVarPointer(uint16) unsafe.Pointer
	SetVar(uint16, uint32)
	SetVar64(uint16, uint64)
	SetVarInt8(uint16, int8)
	SetVarInt16(uint16, int16)
	SetVarInt32(uint16, int32)
	SetVarInt64(uint16, int64)
	SetVarFloat32(uint16, float32)
	SetVarFloat64(uint16, float64)
	SetVarRef(uint16, Ref)
	SetVarPointer(uint16, unsafe.Pointer)

	Peek() uint32
	Peek64() uint64
	PeekInt8() int8
	PeekInt16() int16
	PeekInt32() int32
	PeekInt64() int64
	PeekFloat32() float32
	PeekFloat64() float64
	PeekRef() Ref
	PeekPointer() unsafe.Pointer
	Pop() uint32
	Pop64() uint64
	PopInt8() int8
	PopInt16() int16
	PopInt32() int32
	PopInt64() int64
	PopFloat32() float32
	PopFloat64() float64
	PopRef() Ref
	PopPointer() unsafe.Pointer
	Push(uint32)
	Push64(uint64)
	PushInt8(int8)
	PushInt16(int16)
	PushInt32(int32)
	PushInt64(int64)
	PushFloat32(float32)
	PushFloat64(float64)
	PushRef(Ref)
	PushPointer(unsafe.Pointer)

	// returns whether the top element is a reference or not
	IsRef() bool
}
