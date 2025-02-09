package vm

import (
	"sync/atomic"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
)

var (
	threadId        int64 = 2
	threadIdAddress int64 = (int64)((uintptr)((unsafe.Pointer)(&threadId)))
)

func GetNextThreadId() int64 {
	return atomic.AddInt64(&threadId, 1) - 1
}

func GetNextThreadIdAddress() int64 {
	return threadIdAddress
}

type ThreadUserData struct {
	VM       *VM
	Name     string
	Priority int32

	ScopedValueCache ir.Ref
}

func (vm *VM) GetCarrierThread() ir.Ref {
	return vm.carrierThread
}

func (vm *VM) GetCurrentThread() ir.Ref {
	return vm.currentThread
}

func (vm *VM) SetCurrentThread(thread ir.Ref) {
	vm.currentThread = thread.(*Ref)
}

func (vm *VM) MarkInterrupt() {
	select {
	case vm.interruptNotifier <- struct{}{}:
	default:
	}
}

func (vm *VM) Interrupt(thread ir.Ref) {
	ref := thread.(*Ref)
	target := ref.userData.(*ThreadUserData).VM
	target.MarkInterrupt()
}

func (vm *VM) ClearInterrupt() {
	select {
	case <-vm.interruptNotifier:
	default:
	}
}

func (vm *VM) GetAndClearInterrupt() bool {
	ptr := (*int32)(vm.javaLangThread_interrupted.GetPointer(vm.currentThread))
	return atomic.CompareAndSwapInt32(ptr, 1, 0)
}
