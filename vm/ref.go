package vm

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type Ref struct {
	class     uint32 // index class in classloader
	arrayKind ir.arrayKind
	arrayLen  int32
	data      unsafe.Pointer
}

var _ ir.Ref = (*Ref)(nil)

func newRefArray(class *jcls.Class, kind uint8, length int) *Ref {
	r := &Ref{
		class:     class,
		arrayKind: kind,
		arrayLen:  (int32)(length),
	}
	if kind == kindRef {
		data := make([]unsafe.Pointer, length)
	} else {
		bytes := make([]byte, kind.Size()*(uintptr)(length))
		r.data = (unsafe.Pointer)(unsafe.SliceData(bytes))
	}
	return r
}

func newObjectRef(class *jcls.Class, size int) *Ref {
	return &Ref{
		class:     class,
		arrayKind: kindNone,
	}
}

func (r *Ref) IsArray() bool {
	return r.arrayKind != kindNone
}

func (r *Ref) ArrayKind() ArrayKind {
	return r.arrayKind
}

func (r *Ref) Len() int32 {
	return r.arrayLen
}

func (r *Ref) Data() unsafe.Pointer {
	return r.data
}

func (r *Ref) GetRefArray() []Ref {
	if r.arrayKind != kindRef {
		panic("Underlying array is not reference")
	}
	return unsafe.Slice((*Ref)(r.data), r.len)
}

func (r *Ref) GetArrInt8() []int8 {
	if r.arrayKind != kindBoolean && r.arrayKind != kindByte {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*int8)(r.data), r.len)
}

func (r *Ref) GetArrInt16() []int16 {
	if r.arrayKind != kindChar && r.arrayKind != kindShort {
		panic("Underlying array is not int16")
	}
	return unsafe.Slice((*int16)(r.data), r.len)
}

func (r *Ref) GetArrInt32() []int32 {
	if r.arrayKind != kindInt && r.arrayKind != kindFloat {
		panic("Underlying array is not int32")
	}
	return unsafe.Slice((*int32)(r.data), r.len)
}

func (r *Ref) GetArrInt64() []int64 {
	if r.arrayKind != kindLong && r.arrayKind != kindDouble {
		panic("Underlying array is not int64")
	}
	return unsafe.Slice((*int64)(r.data), r.len)
}
