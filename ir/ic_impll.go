package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICl2d struct{}

func (*ICl2d) Op() ops.Op { return ops.L2d }
func (*ICl2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushFloat64((float64)(value))
	return nil
}

type ICl2f struct{}

func (*ICl2f) Op() ops.Op { return ops.L2f }
func (*ICl2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushFloat32((float32)(value))
	return nil
}

type ICl2i struct{}

func (*ICl2i) Op() ops.Op { return ops.L2i }
func (*ICl2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushInt32((int32)(value))
	return nil
}

type ICladd struct{}

func (*ICladd) Op() ops.Op { return ops.Ladd }
func (*ICladd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a + b)
	return nil
}

type IClaload struct{}

func (*IClaload) Op() ops.Op { return ops.Laload }
func (*IClaload) Execute(vm VM) error {
	stack := vm.GetStack()
	index := stack.PopInt64()
	arr := stack.PopRef().GetInt64Arr()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt64(arr[index])
	return nil
}

type ICland struct{}

func (*ICland) Op() ops.Op { return ops.Land }
func (*ICland) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a & b)
	return nil
}

type IClastore struct{}

func (*IClastore) Op() ops.Op { return ops.Lastore }
func (*IClastore) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	index := stack.PopInt32()
	arr := stack.PopRef().GetInt64Arr()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IClcmp struct{}

func (*IClcmp) Op() ops.Op { return ops.Lcmp }
func (*IClcmp) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	var res int32 = 0
	if a > b {
		res = 1
	} else if a < b {
		res = -1
	}
	stack.PushInt32(res)
	return nil
}

type IClconst_0 struct{}

func (*IClconst_0) Op() ops.Op { return ops.Lconst_0 }
func (*IClconst_0) Execute(vm VM) error {
	vm.GetStack().PushInt64(0)
	return nil
}

type IClconst_1 struct{}

func (*IClconst_1) Op() ops.Op { return ops.Lconst_1 }
func (*IClconst_1) Execute(vm VM) error {
	vm.GetStack().PushInt64(1)
	return nil
}

type ICldc2_w struct {
	Index uint16
}

func (*ICldc2_w) Op() ops.Op { return ops.Ldc2_w }
func (ic *ICldc2_w) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst(ic.Index, vm.GetStack())
}

type ICldiv struct{}

func (*ICldiv) Op() ops.Op { return ops.Ldiv }
func (*ICldiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a / b)
	return nil
}

type IClload struct {
	Index uint16
}

func (*IClload) Op() ops.Op { return ops.Lload }
func (ic *IClload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64((uint16)(ic.Index))
	stack.PushInt64(val)
	return nil
}

type IClload_0 struct{}

func (*IClload_0) Op() ops.Op { return ops.Lload_0 }
func (*IClload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(0)
	stack.PushInt64(val)
	return nil
}

type IClload_1 struct{}

func (*IClload_1) Op() ops.Op { return ops.Lload_1 }
func (*IClload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(1)
	stack.PushInt64(val)
	return nil
}

type IClload_2 struct{}

func (*IClload_2) Op() ops.Op { return ops.Lload_2 }
func (*IClload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(2)
	stack.PushInt64(val)
	return nil
}

type IClload_3 struct{}

func (*IClload_3) Op() ops.Op { return ops.Lload_3 }
func (*IClload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(3)
	stack.PushInt64(val)
	return nil
}

type IClmul struct{}

func (*IClmul) Op() ops.Op { return ops.Lmul }
func (*IClmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a * b)
	return nil
}

type IClneg struct{}

func (*IClneg) Op() ops.Op { return ops.Lneg }
func (*IClneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt64()
	stack.PushInt64(-a)
	return nil
}

type IClor struct{}

func (*IClor) Op() ops.Op { return ops.Lor }
func (*IClor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a | b)
	return nil
}

type IClrem struct{}

func (*IClrem) Op() ops.Op { return ops.Lrem }
func (*IClrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a % b)
	return nil
}

type IClshl struct{}

func (*IClshl) Op() ops.Op { return ops.Lshl }
func (*IClshl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt64()
	stack.PushInt64(a << (b & 0x3f))
	return nil
}

type IClshr struct{}

func (*IClshr) Op() ops.Op { return ops.Lshr }
func (*IClshr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt64()
	stack.PushInt64(a >> (b & 0x3f))
	return nil
}

type IClstore struct {
	Index uint16
}

func (*IClstore) Op() ops.Op { return ops.Lstore }
func (ic *IClstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64((uint16)(ic.Index), val)
	return nil
}

type IClstore_0 struct{}

func (*IClstore_0) Op() ops.Op { return ops.Lstore_0 }
func (*IClstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(0, val)
	return nil
}

type IClstore_1 struct{}

func (*IClstore_1) Op() ops.Op { return ops.Lstore_1 }
func (*IClstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(1, val)
	return nil
}

type IClstore_2 struct{}

func (*IClstore_2) Op() ops.Op { return ops.Lstore_2 }
func (*IClstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(2, val)
	return nil
}

type IClstore_3 struct{}

func (*IClstore_3) Op() ops.Op { return ops.Lstore_3 }
func (*IClstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(3, val)
	return nil
}

type IClsub struct{}

func (*IClsub) Op() ops.Op { return ops.Lsub }
func (*IClsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a - b)
	return nil
}

type IClushr struct{}

func (*IClushr) Op() ops.Op { return ops.Lushr }
func (*IClushr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.Pop64()
	stack.Push64(a >> (b & 0x1f))
	return nil
}

type IClxor struct{}

func (*IClxor) Op() ops.Op { return ops.Lxor }
func (*IClxor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a ^ b)
	return nil
}
