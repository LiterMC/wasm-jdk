package ir

import (
	"iter"
	"reflect"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
)

type Ref interface {
	Desc() *desc.Desc
	Class() Class
	Id() int32
	Len() int32
	UserData() *any
	Data() unsafe.Pointer
	GetRefArr() []unsafe.Pointer
	GetByteArr() []byte
	GetInt8Arr() []int8
	GetInt16Arr() []int16
	GetInt32Arr() []int32
	GetInt64Arr() []int64

	IsLocked(VM) int
	Lock(VM) int
	Unlock(VM) (int, error)
	Notify(VM) error
	NotifyAll(VM) error
	Wait(VM, int64) error
}

type Class interface {
	Name() string
	Desc() *desc.Desc
	ArrayDim() int
	Elem() Class
	Reflect() reflect.Type
	AsRef(VM) Ref

	Modifiers() int32
	Super() Class
	Interfaces() []Class
	IsInterface() bool
	IsAssignableFrom(Class) bool
	IsInstance(Ref) bool

	GetAndPushConst(VM, uint16, Stack) error
	GetFields() iter.Seq[Field]
	GetField(VM, uint16) Field
	GetFieldByName(string) Field
	GetMethods() iter.Seq[Method]
	GetMethod(VM, uint16) Method
	GetMethodByName(VM, string) Method
	GetMethodByNameAndType(vm VM, name, typ string) Method
}

type Field interface {
	Name() string
	Offset() int64
	GetDeclaringClass() Class
	Modifiers() int32
	IsStatic() bool

	AsRef(VM) Ref
	GetPointer(Ref) unsafe.Pointer
	GetAndPush(Stack) error
	PopAndSet(Stack) error
}

type Method interface {
	Name() string
	Desc() *desc.MethodDesc
	GetDeclaringClass() Class
	Modifiers() int32
	IsStatic() bool

	AsRef(VM) Ref
}
