package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICi2b struct{}

func (*ICi2b) Op() ops.Op { return ops.I2b }
func (*ICi2b) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32((int32)((int8)(value)))
	return nil
}

type ICi2c struct{}

func (*ICi2c) Op() ops.Op { return ops.I2c }
func (*ICi2c) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32(value & 0xffff)
	return nil
}

type ICi2d struct{}

func (*ICi2d) Op() ops.Op { return ops.I2d }
func (*ICi2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushFloat64((float64)(value))
	return nil
}

type ICi2f struct{}

func (*ICi2f) Op() ops.Op { return ops.I2f }
func (*ICi2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushFloat32((float32)(value))
	return nil
}

type ICi2l struct{}

func (*ICi2l) Op() ops.Op { return ops.I2l }
func (*ICi2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt64((int64)(value))
	return nil
}

type ICi2s struct{}

func (*ICi2s) Op() ops.Op { return ops.I2s }
func (*ICi2s) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32((int32)((int16)(value)))
	return nil
}

type ICiadd struct{}

func (*ICiadd) Op() ops.Op { return ops.Iadd }
func (*ICiadd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a + b)
	return nil
}

type ICiaload struct{}

func (*ICiaload) Op() ops.Op { return ops.Iaload }
func (*ICiaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopRef().GetArrInt32()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt32(arr[index])
	return nil
}

type ICiand struct{}

func (*ICiand) Op() ops.Op { return ops.Iand }
func (*ICiand) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a & b)
	return nil
}

type ICiastore struct{}

func (*ICiastore) Op() ops.Op { return ops.Iastore }
func (*ICiastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopRef().GetArrInt32()
	index := stack.PopInt32()
	value := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type ICiconst_m1 struct{}

func (*ICiconst_m1) Op() ops.Op { return ops.Iconst_m1 }
func (*ICiconst_m1) Execute(vm VM) error {
	vm.GetStack().PushInt32(-1)
	return nil
}

type ICiconst_0 struct{}

func (*ICiconst_0) Op() ops.Op { return ops.Iconst_0 }
func (*ICiconst_0) Execute(vm VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}

type ICiconst_1 struct{}

func (*ICiconst_1) Op() ops.Op { return ops.Iconst_1 }
func (*ICiconst_1) Execute(vm VM) error {
	vm.GetStack().PushInt32(1)
	return nil
}

type ICiconst_2 struct{}

func (*ICiconst_2) Op() ops.Op { return ops.Iconst_2 }
func (*ICiconst_2) Execute(vm VM) error {
	vm.GetStack().PushInt32(2)
	return nil
}

type ICiconst_3 struct{}

func (*ICiconst_3) Op() ops.Op { return ops.Iconst_3 }
func (*ICiconst_3) Execute(vm VM) error {
	vm.GetStack().PushInt32(3)
	return nil
}

type ICiconst_4 struct{}

func (*ICiconst_4) Op() ops.Op { return ops.Iconst_4 }
func (*ICiconst_4) Execute(vm VM) error {
	vm.GetStack().PushInt32(4)
	return nil
}

type ICiconst_5 struct{}

func (*ICiconst_5) Op() ops.Op { return ops.Iconst_5 }
func (*ICiconst_5) Execute(vm VM) error {
	vm.GetStack().PushInt32(5)
	return nil
}

type ICidiv struct{}

func (*ICidiv) Op() ops.Op { return ops.Idiv }
func (*ICidiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a / b)
	return nil
}

type ICiinc struct {
	Index uint16
	Const int16
}

func (*ICiinc) Op() ops.Op { return ops.Iinc }
func (ir *ICiinc) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(ir.Index)
	stack.SetVarInt32(ir.Index, val+(int32)(ir.Const))
	return nil
}

type ICiload struct {
	Index uint16
}

func (*ICiload) Op() ops.Op { return ops.Iload }
func (ir *ICiload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32((uint16)(ir.Index))
	stack.PushInt32(val)
	return nil
}

type ICiload_0 struct{}

func (*ICiload_0) Op() ops.Op { return ops.Iload_0 }
func (*ICiload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(0)
	stack.PushInt32(val)
	return nil
}

type ICiload_1 struct{}

func (*ICiload_1) Op() ops.Op { return ops.Iload_1 }
func (*ICiload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(1)
	stack.PushInt32(val)
	return nil
}

type ICiload_2 struct{}

func (*ICiload_2) Op() ops.Op { return ops.Iload_2 }
func (*ICiload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(2)
	stack.PushInt32(val)
	return nil
}

type ICiload_3 struct{}

func (*ICiload_3) Op() ops.Op { return ops.Iload_3 }
func (*ICiload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(3)
	stack.PushInt32(val)
	return nil
}

type ICimul struct{}

func (*ICimul) Op() ops.Op { return ops.Imul }
func (*ICimul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a * b)
	return nil
}

type ICineg struct{}

func (*ICineg) Op() ops.Op { return ops.Ineg }
func (*ICineg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	stack.PushInt32(-a)
	return nil
}

type ICior struct{}

func (*ICior) Op() ops.Op { return ops.Ior }
func (*ICior) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a | b)
	return nil
}

type ICirem struct{}

func (*ICirem) Op() ops.Op { return ops.Irem }
func (*ICirem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a % b)
	return nil
}

type ICishl struct{}

func (*ICishl) Op() ops.Op { return ops.Ishl }
func (*ICishl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a << (b & 0x1f))
	return nil
}

type ICishr struct{}

func (*ICishr) Op() ops.Op { return ops.Ishr }
func (*ICishr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a >> (b & 0x1f))
	return nil
}

type ICistore struct {
	Index uint16
}

func (*ICistore) Op() ops.Op { return ops.Istore }
func (ir *ICistore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32((uint16)(ir.Index), val)
	return nil
}

type ICistore_0 struct{}

func (*ICistore_0) Op() ops.Op { return ops.Istore_0 }
func (*ICistore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(0, val)
	return nil
}

type ICistore_1 struct{}

func (*ICistore_1) Op() ops.Op { return ops.Istore_1 }
func (*ICistore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(1, val)
	return nil
}

type ICistore_2 struct{}

func (*ICistore_2) Op() ops.Op { return ops.Istore_2 }
func (*ICistore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(2, val)
	return nil
}

type ICistore_3 struct{}

func (*ICistore_3) Op() ops.Op { return ops.Istore_3 }
func (*ICistore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(3, val)
	return nil
}

type ICisub struct{}

func (*ICisub) Op() ops.Op { return ops.Isub }
func (*ICisub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a - b)
	return nil
}

type ICiushr struct{}

func (*ICiushr) Op() ops.Op { return ops.Iushr }
func (*ICiushr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32((int32)((uint32)(a) >> (b & 0x1f)))
	return nil
}

type ICixor struct{}

func (*ICixor) Op() ops.Op { return ops.Ixor }
func (*ICixor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a ^ b)
	return nil
}

type ICldc struct {
	Index byte
}

func (*ICldc) Op() ops.Op { return ops.Ldc }
func (ir *ICldc) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst((uint16)(ir.Index), vm.GetStack())
}

type ICldc_w struct {
	Index uint16
}

func (*ICldc_w) Op() ops.Op { return ops.Ldc_w }
func (ir *ICldc_w) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst(ir.Index, vm.GetStack())
}
