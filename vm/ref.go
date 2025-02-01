package vm

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/jcls"
)

const (
	kindNone uint8 = iota
	kindRef
	kindBoolean
	kindByte
	kindChar
	kindShort
	kindInt
	kindLong
	kindFloat
	kindDouble
)

type Ref = unsafe.Pointer

type refHeader struct {
	class     *jcls.Class
	arrayKind uint8
	len       int
	data      [0]byte
}

func getKindSize(kind uint8) uintptr {
	switch kind {
	case kindRef:
		return unsafe.Sizeof((Ref)(nil))
	case kindBoolean, kindByte:
		return unsafe.Sizeof((int8)(0))
	case kindChar, kindShort:
		return unsafe.Sizeof((int16)(0))
	case kindInt, kindFloat:
		return unsafe.Sizeof((int32)(0))
	case kindLong, kindDouble:
		return unsafe.Sizeof((int64)(0))
	default:
		panic("unknown kind")
	}
}

func newRefArray(class *jcls.Class, kind uint8, length int) Ref {
	bytes := make([]byte, unsafe.Offsetof(refHeader{}.data)+getKindSize(kind)*(uintptr)(length))
	r := (*refHeader)((unsafe.Pointer)(unsafe.SliceData(bytes)))
	r.class = class
	r.arrayKind = kind
	r.len = length
	return r.AsRef()
}

func newObjectRef(class *jcls.Class, size int) Ref {
	bytes := make([]byte, unsafe.Offsetof(refHeader{}.data)+(uintptr)(size))
	r := (*refHeader)((unsafe.Pointer)(unsafe.SliceData(bytes)))
	r.class = class
	r.arrayKind = kindRef
	return r.AsRef()
}

func getRefHeader(ref Ref) *refHeader {
	return (*refHeader)((unsafe.Pointer)((uintptr)(ref) - unsafe.Offsetof(refHeader{}.data)))
}

func (r *refHeader) AsRef() Ref {
	return (Ref)(&r.data)
}

func (r *refHeader) GetRefArray() []Ref {
	if r.arrayKind != kindRef {
		panic("Underlying array is not reference")
	}
	return unsafe.Slice((*Ref)((Ref)(&r.data)), r.len)
}

func (r *refHeader) GetArrInt8() []int8 {
	if r.arrayKind != kindBoolean && r.arrayKind != kindByte {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*int8)((Ref)(&r.data)), r.len)
}

func (r *refHeader) GetArrInt16() []int16 {
	if r.arrayKind != kindChar && r.arrayKind != kindShort {
		panic("Underlying array is not int16")
	}
	return unsafe.Slice((*int16)((Ref)(&r.data)), r.len)
}

func (r *refHeader) GetArrInt32() []int32 {
	if r.arrayKind != kindInt && r.arrayKind != kindFloat {
		panic("Underlying array is not int32")
	}
	return unsafe.Slice((*int32)((Ref)(&r.data)), r.len)
}

func (r *refHeader) GetArrInt64() []int64 {
	if r.arrayKind != kindLong && r.arrayKind != kindDouble {
		panic("Underlying array is not int64")
	}
	return unsafe.Slice((*int64)((Ref)(&r.data)), r.len)
}
