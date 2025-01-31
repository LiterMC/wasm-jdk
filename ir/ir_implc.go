package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRcaload struct{}

func (*IRcaload) Op() ops.Op            { return ops.Caload }
func (*IRcaload) Operands() int         { return 0 }
func (*IRcaload) Parse(operands []byte) {}
func (*IRcaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt16()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt16(arr[index])
	return nil
}

type IRcastore struct{}

func (*IRcastore) Op() ops.Op            { return ops.Castore }
func (*IRcastore) Operands() int         { return 0 }
func (*IRcastore) Parse(operands []byte) {}
func (*IRcastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt16()
	index := stack.PopInt32()
	value := stack.PopInt16()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IRsaload struct{}

func (*IRsaload) Op() ops.Op            { return ops.Saload }
func (*IRsaload) Operands() int         { return 0 }
func (*IRsaload) Parse(operands []byte) {}
func (*IRsaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt16()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt16(arr[index])
	return nil
}

type IRsastore struct{}

func (*IRsastore) Op() ops.Op            { return ops.Sastore }
func (*IRsastore) Operands() int         { return 0 }
func (*IRsastore) Parse(operands []byte) {}
func (*IRsastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt16()
	index := stack.PopInt32()
	value := stack.PopInt16()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IRsipush struct {
	value int16
}

func (*IRsipush) Op() ops.Op    { return ops.Sipush }
func (*IRsipush) Operands() int { return 2 }
func (ir *IRsipush) Parse(operands []byte) {
	ir.value = bytesToInt16(operands)
}
func (ir *IRsipush) Execute(vm VM) error {
	vm.GetStack().PushInt32((int32)(ir.value))
	return nil
}
