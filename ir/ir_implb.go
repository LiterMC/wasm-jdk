package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRbaload struct{}

func (*IRbaload) Op() ops.Op            { return ops.Baload }
func (*IRbaload) Operands() int         { return 0 }
func (*IRbaload) Parse(operands []byte) {}
func (*IRbaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt8()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt8(arr[index])
	return nil
}

type IRbastore struct{}

func (*IRbastore) Op() ops.Op            { return ops.Bastore }
func (*IRbastore) Operands() int         { return 0 }
func (*IRbastore) Parse(operands []byte) {}
func (*IRbastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt8()
	index := stack.PopInt32()
	value := stack.PopInt8()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IRbipush struct {
	value int8
}

func (*IRbipush) Op() ops.Op    { return ops.Bipush }
func (*IRbipush) Operands() int { return 1 }
func (ir *IRbipush) Parse(operands []byte) {
	ir.value = (int8)(operands[0])
}
func (ir *IRbipush) Execute(vm VM) error {
	vm.GetStack().PushInt32((int32)(ir.value))
	return nil
}
