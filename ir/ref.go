package ir

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
)

type Ref interface {
	Desc() *desc.Desc
	Len() int32
	Data() unsafe.Pointer
	GetArrRef() []Ref
	GetArrInt8() []int8
	GetArrInt16() []int16
	GetArrInt32() []int32
	GetArrInt64() []int64
}
