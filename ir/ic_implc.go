package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICcaload struct{}

func (*ICcaload) Op() ops.Op { return ops.Caload }
func (*ICcaload) Execute(vm VM) error {
	stack := vm.GetStack()
	index := stack.PopInt32()
	arr := stack.PopRef().GetArrInt16()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt16(arr[index])
	return nil
}

type ICcastore struct{}

func (*ICcastore) Op() ops.Op { return ops.Castore }
func (*ICcastore) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt16()
	index := stack.PopInt32()
	arr := stack.PopRef().GetArrInt16()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type ICsaload struct{}

func (*ICsaload) Op() ops.Op { return ops.Saload }
func (*ICsaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopRef().GetArrInt16()
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

type ICsastore struct{}

func (*ICsastore) Op() ops.Op { return ops.Sastore }
func (*ICsastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopRef().GetArrInt16()
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

type ICsipush struct {
	Value int16
}

func (*ICsipush) Op() ops.Op { return ops.Sipush }
func (ic *ICsipush) Execute(vm VM) error {
	vm.GetStack().PushInt16(ic.Value)
	return nil
}
