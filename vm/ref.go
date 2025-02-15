package vm

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ir"
)

type Ref struct {
	lock      sync.Mutex
	locked    atomic.Pointer[VM]
	lockCount int32
	notify    chan struct{}
	notifyAll atomic.Pointer[chan struct{}]

	desc  *desc.Desc
	class *Class

	identity int32

	// userData stored native datas for specific objects, such as java.lang.Class and java.lang.Thread
	userData any

	arrayLen int32
	data     unsafe.Pointer
}

var _ ir.Ref = (*Ref)(nil)

func newRefBase(cls *Class, data unsafe.Pointer) *Ref {
	r := &Ref{
		notify:   make(chan struct{}, 0),
		desc:     cls.Desc(),
		class:    cls,
		identity: (int32)(rand.Int63()),
		data:     data,
	}
	nch := make(chan struct{}, 0)
	r.notifyAll.Store(&nch)
	fmt.Printf("New Ref: %#v\n", r)
	runtime.SetFinalizer(r, func(r *Ref) {
		fmt.Printf("Finalizing: %p %#v\n", r, r)
	})
	return r
}

func newObjectRef(cls ir.Class) *Ref {
	class := cls.(*Class)
	data := reflect.New(class.refType).UnsafePointer()
	return newRefBase(class, data)
}

func newRefArray(cls ir.Class, length int32) *Ref {
	var ptr unsafe.Pointer
	et := cls.Desc().ElemType()
	if et.IsRef() {
		data := make([]unsafe.Pointer, length)
		ptr = (unsafe.Pointer)(unsafe.SliceData(data))
	} else {
		bytes := make([]byte, et.Size()*(uintptr)(length))
		ptr = (unsafe.Pointer)(unsafe.SliceData(bytes))
	}
	return newRefArrayWithData(cls, length, ptr)
}

func newRefArrayWithData(cls ir.Class, length int32, data unsafe.Pointer) *Ref {
	ref := newRefBase(cls.(*Class), data)
	ref.arrayLen = length
	return ref
}

func newMultiDimArray(cls ir.Class, lengths []int32) *Ref {
	l := lengths[0]
	arr := newRefArray(cls, l)
	elem := cls.Elem()
	lengths = lengths[1:]
	if len(lengths) > 0 && elem.ArrayDim() > 0 {
		refs := arr.GetRefArr()
		for i := (int32)(0); i < l; i++ {
			refs[i] = (unsafe.Pointer)(newMultiDimArray(elem, lengths))
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

func (r *Ref) GoString() string {
	return fmt.Sprintf("<Ref 0x%08x type=%s data=%p>", (uint32)(r.identity), r.desc, r.data)
}

func (r *Ref) GetRefArr() []unsafe.Pointer {
	if !r.desc.ElemType().IsRef() {
		panic("Underlying array is not reference")
	}
	return unsafe.Slice((*unsafe.Pointer)(r.data), r.arrayLen)
}

func (r *Ref) GetByteArr() []byte {
	if r.desc.ElemType().Size() != 1 {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*byte)(r.data), r.arrayLen)
}

func (r *Ref) GetInt8Arr() []int8 {
	if r.desc.ElemType().Size() != 1 {
		panic("Underlying array is not int8")
	}
	return unsafe.Slice((*int8)(r.data), r.arrayLen)
}

func (r *Ref) GetInt16Arr() []int16 {
	if r.desc.ElemType().Size() != 2 {
		panic("Underlying array is not int16")
	}
	return unsafe.Slice((*int16)(r.data), r.arrayLen)
}

func (r *Ref) GetInt32Arr() []int32 {
	if r.desc.ElemType().Size() != 4 {
		panic("Underlying array is not int32")
	}
	return unsafe.Slice((*int32)(r.data), r.arrayLen)
}

func (r *Ref) GetInt64Arr() []int64 {
	if r.desc.ElemType().Size() != 8 {
		panic("Underlying array is not int64")
	}
	return unsafe.Slice((*int64)(r.data), r.arrayLen)
}

func (r *Ref) IsLocked(vm ir.VM) int {
	if r.locked.Load() == vm.(*VM) {
		return (int)(r.lockCount)
	}
	return 0
}

func (r *Ref) Lock(vm ir.VM) int {
	return r.Lock0(vm.(*VM))
}

func (r *Ref) Lock0(vm *VM) int {
	if r.locked.Load() != vm {
		r.lock.Lock()
		r.locked.Store(vm)
	}
	// if r.locked == vm, it is impossible to unlock concurrently
	r.lockCount++
	return (int)(r.lockCount)
}

func (r *Ref) Unlock(vm ir.VM) (int, error) {
	return r.Unlock0(vm.(*VM))
}

func (r *Ref) Unlock0(vm *VM) (int, error) {
	if r.locked.Load() != vm {
		return 0, errs.IllegalMonitorStateException
	}
	r.lockCount--
	c := (int)(r.lockCount)
	if c == 0 {
		r.locked.Store(nil)
		r.lock.Unlock()
	}
	return c, nil
}

func (r *Ref) Notify(vm ir.VM) error {
	if r.locked.Load() != vm {
		return errs.IllegalMonitorStateException
	}
	select {
	case r.notify <- struct{}{}:
	default:
	}
	return nil
}

func (r *Ref) NotifyAll(vm ir.VM) error {
	if r.locked.Load() != vm {
		return errs.IllegalMonitorStateException
	}
	nch := make(chan struct{}, 0)
	old := r.notifyAll.Swap(&nch)
	close(*old)
	return nil
}

func (r *Ref) Wait(vm ir.VM, millis int64) error {
	return r.Wait0(vm.(*VM), time.Millisecond*(time.Duration)(millis))
}

func (r *Ref) Wait0(vm *VM, dur time.Duration) error {
	if r.locked.Load() != vm {
		return errs.IllegalMonitorStateException
	}
	r.lock.Unlock()
	select {
	case <-vm.interruptNotifier:
		if vm.GetAndClearInterrupt() {
			r.lock.Lock()
			return errs.InterruptedException
		}
	default:
	}
	if dur == 0 {
	SELECT_NO_TIMER:
		select {
		case <-vm.interruptNotifier:
			if vm.GetAndClearInterrupt() {
				r.lock.Lock()
				return errs.InterruptedException
			}
			goto SELECT_NO_TIMER
		case <-r.notify:
		case <-*r.notifyAll.Load():
		}
	} else {
	SELECT_WITH_TIMER:
		select {
		case <-vm.interruptNotifier:
			if vm.GetAndClearInterrupt() {
				r.lock.Lock()
				return errs.InterruptedException
			}
			goto SELECT_WITH_TIMER
		case <-r.notify:
		case <-*r.notifyAll.Load():
		case <-time.After(dur):
		}
	}
	r.lock.Lock()
	return nil
}
