package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICbaload struct{}

func (*ICbaload) Op() ops.Op { return ops.Baload }
func (*ICbaload) Execute(vm VM) error {
	stack := vm.GetStack()
	index := stack.PopInt32()
	arr := stack.PopRef().GetArrInt8()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt8(arr[index])
	return nil
}

type ICbastore struct{}

func (*ICbastore) Op() ops.Op { return ops.Bastore }
func (*ICbastore) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt8()
	index := stack.PopInt32()
	arr := stack.PopRef().GetArrInt8()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type ICbipush struct {
	Value int8
}

func (*ICbipush) Op() ops.Op { return ops.Bipush }
func (ic *ICbipush) Execute(vm VM) error {
	vm.GetStack().PushInt8(ic.Value)
	return nil
}
