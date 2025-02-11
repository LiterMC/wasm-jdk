package ir

import (
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

	Modifiers() int32
	Super() Class
	Interfaces() []Class
	IsInterface() bool
	IsAssignableFrom(Class) bool
	IsInstance(Ref) bool

	GetAndPushConst(uint16, Stack) error
	GetField(uint16) Field
	GetFieldByName(string) Field
	GetMethod(uint16) Method
	GetMethodByName(string) Method
	GetMethodByNameAndType(name, typ string) Method
}

type Field interface {
	Name() string
	Offset() int64

	GetDeclaringClass() Class
	IsStatic() bool

	GetPointer(Ref) unsafe.Pointer
	GetAndPush(Stack) error
	PopAndSet(Stack) error
}

type Method interface {
	Name() string
	Desc() *desc.MethodDesc

	GetDeclaringClass() Class
	IsStatic() bool
}
