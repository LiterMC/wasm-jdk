package ir

import (
	"unsafe"
)

type ArrayKind byte

const (
	KindNone ArrayKind = iota
	KindRef
	KindBoolean
	KindByte
	KindChar
	KindShort
	KindInt
	KindLong
	KindFloat
	KindDouble
)

func (k ArrayKind) Size() uintptr {
	switch k {
	case KindRef:
		return unsafe.Sizeof((Ref)(nil))
	case KindBoolean, KindByte:
		return unsafe.Sizeof((int8)(0))
	case KindChar, KindShort:
		return unsafe.Sizeof((int16)(0))
	case KindInt, KindFloat:
		return unsafe.Sizeof((int32)(0))
	case KindLong, KindDouble:
		return unsafe.Sizeof((int64)(0))
	default:
		panic("unknown kind")
	}
}

type Ref interface {
	IsArray() bool
	ArrayKind() ArrayKind
	Len() int32
	Data() unsafe.Pointer
	GetArrRef() []Ref
	GetArrInt8() []int8
	GetArrInt16() []int16
	GetArrInt32() []int32
	GetArrInt64() []int64
}
