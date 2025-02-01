package vm

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type Ref struct {
	class     *jcls.Class
	arrayKind ir.ArrayKind
	arrayLen  int32
	data      unsafe.Pointer
}

var _ ir.Ref = (*Ref)(nil)

func newRefArray(class *jcls.Class, kind ir.ArrayKind, length int) *Ref {
	r := &Ref{
		class:     class,
		arrayKind: kind,
		arrayLen:  (int32)(length),
	}
	if kind == ir.KindRef {
		data := make([]ir.Ref, length)
		r.data = (unsafe.Pointer)(unsafe.SliceData(data))
	} else {
		bytes := make([]byte, kind.Size()*(uintptr)(length))
		r.data = (unsafe.Pointer)(unsafe.SliceData(bytes))
	}
	return r
}

func newObjectRef(class *jcls.Class, size int) *Ref {
	return &Ref{
		class:     class,
		arrayKind: ir.KindNone,
	}
}

func (r *Ref) IsArray() bool {
	return r.arrayKind != ir.KindNone
}

func (r *Ref) ArrayKind() ir.ArrayKind {
	return r.arrayKind
}

func (r *Ref) Len() int32 {
	return r.arrayLen
}

func (r *Ref) Data() unsafe.Pointer {
	return r.data
}

func (r *Ref) GetArrRef() []ir.Ref {
	if r.arrayKind != ir.KindRef {
		panic("Underlying array is not reference")
	}
	return unsafe.Slice((*ir.Ref)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt8() []int8 {
	if r.arrayKind != ir.KindBoolean && r.arrayKind != ir.KindByte {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*int8)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt16() []int16 {
	if r.arrayKind != ir.KindChar && r.arrayKind != ir.KindShort {
		panic("Underlying array is not int16")
	}
	return unsafe.Slice((*int16)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt32() []int32 {
	if r.arrayKind != ir.KindInt && r.arrayKind != ir.KindFloat {
		panic("Underlying array is not int32")
	}
	return unsafe.Slice((*int32)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt64() []int64 {
	if r.arrayKind != ir.KindLong && r.arrayKind != ir.KindDouble {
		panic("Underlying array is not int64")
	}
	return unsafe.Slice((*int64)(r.data), r.arrayLen)
}
