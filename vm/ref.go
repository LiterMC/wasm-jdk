package vm

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ir"
)

type Ref struct {
	lock      sync.Cond
	locked    ir.VM
	lockCount int32

	desc  *desc.Desc
	class ir.Class

	identity int32

	userData any
	arrayLen int32
	data     unsafe.Pointer
}

var _ ir.Ref = (*Ref)(nil)

func newObjectRef(cls ir.Class) *Ref {
	class := cls.(*Class)
	data := reflect.New(class.refType).UnsafePointer()
	return &Ref{
		lock:     sync.Cond{L: new(sync.Mutex)},
		desc:     cls.Desc(),
		class:    class,
		identity: (int32)(rand.Int63n(32)),
		data:     data,
	}
}

func newRefArray(cls ir.Class, length int32) *Ref {
	var ptr unsafe.Pointer
	et := cls.Desc().ElemType()
	if et.IsRef() {
		data := make([]ir.Ref, length)
		ptr = (unsafe.Pointer)(unsafe.SliceData(data))
	} else {
		bytes := make([]byte, et.Size()*(uintptr)(length))
		ptr = (unsafe.Pointer)(unsafe.SliceData(bytes))
	}
	return newRefArrayWithData(cls, length, ptr)
}

func newRefArrayWithData(cls ir.Class, length int32, data unsafe.Pointer) *Ref {
	return &Ref{
		lock:     sync.Cond{L: new(sync.Mutex)},
		desc:     cls.Desc(),
		class:    cls.(*Class),
		identity: (int32)(rand.Int63n(32)),
		arrayLen: length,
		data:     data,
	}
}

func newMultiDimArray(cls ir.Class, lengths []int32) *Ref {
	l := lengths[0]
	arr := newRefArray(cls, l)
	elem := cls.Elem()
	lengths = lengths[1:]
	if len(lengths) > 0 {
		refs := arr.GetArrRef()
		for i := (int32)(0); i < l; i++ {
			refs[i] = newMultiDimArray(elem, lengths)
		}
	}
	return arr
}

func (r *Ref) Desc() *desc.Desc {
	return r.desc
}

func (r *Ref) Class() ir.Class {
	return r.class
}

func (r *Ref) Id() int32 {
	return r.identity
}

func (r *Ref) Len() int32 {
	return r.arrayLen
}

func (r *Ref) UserData() *any {
	return &r.userData
}

func (r *Ref) Data() unsafe.Pointer {
	return r.data
}

func (r *Ref) GetArrRef() []ir.Ref {
	if !r.desc.ElemType().IsRef() {
		panic("Underlying array is not reference")
	}
	return unsafe.Slice((*ir.Ref)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt8() []int8 {
	if r.desc.ElemType().Size() != 1 {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*int8)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt16() []int16 {
	if r.desc.ElemType().Size() != 2 {
		panic("Underlying array is not int16")
	}
	return unsafe.Slice((*int16)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt32() []int32 {
	if r.desc.ElemType().Size() != 4 {
		panic("Underlying array is not int32")
	}
	return unsafe.Slice((*int32)(r.data), r.arrayLen)
}

func (r *Ref) GetArrInt64() []int64 {
	if r.desc.ElemType().Size() != 8 {
		panic("Underlying array is not int64")
	}
	return unsafe.Slice((*int64)(r.data), r.arrayLen)
}

func (r *Ref) Lock(vm ir.VM) int {
	if r.locked != vm {
		r.lock.L.Lock()
	}
	r.lockCount++
	return (int)(r.lockCount)
}

func (r *Ref) Unlock(vm ir.VM) (int, error) {
	if r.locked != vm {
		return 0, errs.IllegalMonitorStateException
	}
	r.lockCount--
	c := (int)(r.lockCount)
	if c == 0 {
		r.lock.L.Unlock()
	}
	return c, nil
}

func (r *Ref) GoString() string {
	return fmt.Sprintf("<Ref %x type=%s data=%p>", (uint32)(r.identity), r.desc, r.data)
}
