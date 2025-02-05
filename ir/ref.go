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
	GetArrRef() []Ref
	GetArrInt8() []int8
	GetArrInt16() []int16
	GetArrInt32() []int32
	GetArrInt64() []int64

	Lock(vm VM) int
	Unlock(vm VM) (int, error)
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

	GetDeclaringClass() Class
	IsStatic() bool

	GetPointer(Ref) unsafe.Pointer
	GetAndPush(Ref, Stack)
	PopAndSet(Ref, Stack)
}

type Method interface {
	Name() string
	Desc() *desc.MethodDesc

	GetDeclaringClass() Class
	IsStatic() bool
}
