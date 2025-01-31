package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRcaload struct{}

func (*IRcaload) Op() ops.Op            { return ops.Caload }
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
	Value int16
}

func (*IRsipush) Op() ops.Op    { return ops.Sipush }
func (ir *IRsipush) Execute(vm VM) error {
	vm.GetStack().PushInt32((int32)(ir.Value))
	return nil
}
