package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRi2b struct{}

func (*IRi2b) Op() ops.Op            { return ops.I2b }
func (*IRi2b) Operands() int         { return 0 }
func (*IRi2b) Parse(operands []byte) {}
func (*IRi2b) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32((int32)((int8)(value)))
	return nil
}

type IRi2c struct{}

func (*IRi2c) Op() ops.Op            { return ops.I2c }
func (*IRi2c) Operands() int         { return 0 }
func (*IRi2c) Parse(operands []byte) {}
func (*IRi2c) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32(value & 0xffff)
	return nil
}

type IRi2d struct{}

func (*IRi2d) Op() ops.Op            { return ops.I2d }
func (*IRi2d) Operands() int         { return 0 }
func (*IRi2d) Parse(operands []byte) {}
func (*IRi2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushFloat64((float64)(value))
	return nil
}

type IRi2f struct{}

func (*IRi2f) Op() ops.Op            { return ops.I2f }
func (*IRi2f) Operands() int         { return 0 }
func (*IRi2f) Parse(operands []byte) {}
func (*IRi2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushFloat32((float32)(value))
	return nil
}

type IRi2l struct{}

func (*IRi2l) Op() ops.Op            { return ops.I2l }
func (*IRi2l) Operands() int         { return 0 }
func (*IRi2l) Parse(operands []byte) {}
func (*IRi2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt64((int64)(value))
	return nil
}

type IRi2s struct{}

func (*IRi2s) Op() ops.Op            { return ops.I2s }
func (*IRi2s) Operands() int         { return 0 }
func (*IRi2s) Parse(operands []byte) {}
func (*IRi2s) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt32()
	stack.PushInt32((int32)((int16)(value)))
	return nil
}

type IRiadd struct{}

func (*IRiadd) Op() ops.Op            { return ops.Iadd }
func (*IRiadd) Operands() int         { return 0 }
func (*IRiadd) Parse(operands []byte) {}
func (*IRiadd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a + b)
	return nil
}

type IRiaload struct{}

func (*IRiaload) Op() ops.Op            { return ops.Iaload }
func (*IRiaload) Operands() int         { return 0 }
func (*IRiaload) Parse(operands []byte) {}
func (*IRiaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt32()
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

type IRiand struct{}

func (*IRiand) Op() ops.Op            { return ops.Iand }
func (*IRiand) Operands() int         { return 0 }
func (*IRiand) Parse(operands []byte) {}
func (*IRiand) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a & b)
	return nil
}

type IRiastore struct{}

func (*IRiastore) Op() ops.Op            { return ops.Iastore }
func (*IRiastore) Operands() int         { return 0 }
func (*IRiastore) Parse(operands []byte) {}
func (*IRiastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt32()
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

type IRiconst_m1 struct{}

func (*IRiconst_m1) Op() ops.Op            { return ops.Iconst_m1 }
func (*IRiconst_m1) Operands() int         { return 0 }
func (*IRiconst_m1) Parse(operands []byte) {}
func (*IRiconst_m1) Execute(vm VM) error {
	vm.GetStack().PushInt32(-1)
	return nil
}

type IRiconst_0 struct{}

func (*IRiconst_0) Op() ops.Op            { return ops.Iconst_0 }
func (*IRiconst_0) Operands() int         { return 0 }
func (*IRiconst_0) Parse(operands []byte) {}
func (*IRiconst_0) Execute(vm VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}

type IRiconst_1 struct{}

func (*IRiconst_1) Op() ops.Op            { return ops.Iconst_1 }
func (*IRiconst_1) Operands() int         { return 0 }
func (*IRiconst_1) Parse(operands []byte) {}
func (*IRiconst_1) Execute(vm VM) error {
	vm.GetStack().PushInt32(1)
	return nil
}

type IRiconst_2 struct{}

func (*IRiconst_2) Op() ops.Op            { return ops.Iconst_2 }
func (*IRiconst_2) Operands() int         { return 0 }
func (*IRiconst_2) Parse(operands []byte) {}
func (*IRiconst_2) Execute(vm VM) error {
	vm.GetStack().PushInt32(2)
	return nil
}

type IRiconst_3 struct{}

func (*IRiconst_3) Op() ops.Op            { return ops.Iconst_3 }
func (*IRiconst_3) Operands() int         { return 0 }
func (*IRiconst_3) Parse(operands []byte) {}
func (*IRiconst_3) Execute(vm VM) error {
	vm.GetStack().PushInt32(3)
	return nil
}

type IRiconst_4 struct{}

func (*IRiconst_4) Op() ops.Op            { return ops.Iconst_4 }
func (*IRiconst_4) Operands() int         { return 0 }
func (*IRiconst_4) Parse(operands []byte) {}
func (*IRiconst_4) Execute(vm VM) error {
	vm.GetStack().PushInt32(4)
	return nil
}

type IRiconst_5 struct{}

func (*IRiconst_5) Op() ops.Op            { return ops.Iconst_5 }
func (*IRiconst_5) Operands() int         { return 0 }
func (*IRiconst_5) Parse(operands []byte) {}
func (*IRiconst_5) Execute(vm VM) error {
	vm.GetStack().PushInt32(5)
	return nil
}

type IRidiv struct{}

func (*IRidiv) Op() ops.Op            { return ops.Idiv }
func (*IRidiv) Operands() int         { return 0 }
func (*IRidiv) Parse(operands []byte) {}
func (*IRidiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a / b)
	return nil
}

type IRiinc struct {
	index uint16
	value int16
}

func (*IRiinc) Op() ops.Op    { return ops.Iinc }
func (*IRiinc) Operands() int { return 2 }
func (ir *IRiinc) Parse(operands []byte) {
	ir.index = (uint16)(operands[0])
	ir.value = (int16)((int8)(operands[1]))
}
func (ir *IRiinc) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(ir.index)
	stack.SetVarInt32(ir.index, val+(int32)(ir.value))
	return nil
}

type IRiload struct {
	index byte
}

func (*IRiload) Op() ops.Op    { return ops.Iload }
func (*IRiload) Operands() int { return 1 }
func (ir *IRiload) Parse(operands []byte) {
	ir.index = operands[0]
}
func (ir *IRiload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32((uint16)(ir.index))
	stack.PushInt32(val)
	return nil
}

type IRiload_0 struct{}

func (*IRiload_0) Op() ops.Op            { return ops.Iload_0 }
func (*IRiload_0) Operands() int         { return 0 }
func (*IRiload_0) Parse(operands []byte) {}
func (*IRiload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(0)
	stack.PushInt32(val)
	return nil
}

type IRiload_1 struct{}

func (*IRiload_1) Op() ops.Op            { return ops.Iload_1 }
func (*IRiload_1) Operands() int         { return 0 }
func (*IRiload_1) Parse(operands []byte) {}
func (*IRiload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(1)
	stack.PushInt32(val)
	return nil
}

type IRiload_2 struct{}

func (*IRiload_2) Op() ops.Op            { return ops.Iload_2 }
func (*IRiload_2) Operands() int         { return 0 }
func (*IRiload_2) Parse(operands []byte) {}
func (*IRiload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(2)
	stack.PushInt32(val)
	return nil
}

type IRiload_3 struct{}

func (*IRiload_3) Op() ops.Op            { return ops.Iload_3 }
func (*IRiload_3) Operands() int         { return 0 }
func (*IRiload_3) Parse(operands []byte) {}
func (*IRiload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(3)
	stack.PushInt32(val)
	return nil
}

type IRimul struct{}

func (*IRimul) Op() ops.Op            { return ops.Imul }
func (*IRimul) Operands() int         { return 0 }
func (*IRimul) Parse(operands []byte) {}
func (*IRimul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a * b)
	return nil
}

type IRineg struct{}

func (*IRineg) Op() ops.Op            { return ops.Ineg }
func (*IRineg) Operands() int         { return 0 }
func (*IRineg) Parse(operands []byte) {}
func (*IRineg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	stack.PushInt32(-a)
	return nil
}

type IRior struct{}

func (*IRior) Op() ops.Op            { return ops.Ior }
func (*IRior) Operands() int         { return 0 }
func (*IRior) Parse(operands []byte) {}
func (*IRior) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a | b)
	return nil
}

type IRirem struct{}

func (*IRirem) Op() ops.Op            { return ops.Irem }
func (*IRirem) Operands() int         { return 0 }
func (*IRirem) Parse(operands []byte) {}
func (*IRirem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a % b)
	return nil
}

type IRishl struct{}

func (*IRishl) Op() ops.Op            { return ops.Ishl }
func (*IRishl) Operands() int         { return 0 }
func (*IRishl) Parse(operands []byte) {}
func (*IRishl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a << (b & 0x1f))
	return nil
}

type IRishr struct{}

func (*IRishr) Op() ops.Op            { return ops.Ishr }
func (*IRishr) Operands() int         { return 0 }
func (*IRishr) Parse(operands []byte) {}
func (*IRishr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a >> (b & 0x1f))
	return nil
}

type IRistore struct {
	index byte
}

func (*IRistore) Op() ops.Op    { return ops.Istore }
func (*IRistore) Operanfs() int { return 1 }
func (ir *IRistore) Parse(operanfs []byte) {
	ir.index = operanfs[0]
}
func (ir *IRistore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32((uint16)(ir.index), val)
	return nil
}

type IRistore_0 struct{}

func (*IRistore_0) Op() ops.Op            { return ops.Istore_0 }
func (*IRistore_0) Operands() int         { return 0 }
func (*IRistore_0) Parse(operands []byte) {}
func (*IRistore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(0, val)
	return nil
}

type IRistore_1 struct{}

func (*IRistore_1) Op() ops.Op            { return ops.Istore_1 }
func (*IRistore_1) Operands() int         { return 0 }
func (*IRistore_1) Parse(operands []byte) {}
func (*IRistore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(1, val)
	return nil
}

type IRistore_2 struct{}

func (*IRistore_2) Op() ops.Op            { return ops.Istore_2 }
func (*IRistore_2) Operands() int         { return 0 }
func (*IRistore_2) Parse(operands []byte) {}
func (*IRistore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(2, val)
	return nil
}

type IRistore_3 struct{}

func (*IRistore_3) Op() ops.Op            { return ops.Istore_3 }
func (*IRistore_3) Operands() int         { return 0 }
func (*IRistore_3) Parse(operands []byte) {}
func (*IRistore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(3, val)
	return nil
}

type IRisub struct{}

func (*IRisub) Op() ops.Op            { return ops.Isub }
func (*IRisub) Operands() int         { return 0 }
func (*IRisub) Parse(operands []byte) {}
func (*IRisub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a - b)
	return nil
}

type IRiushr struct{}

func (*IRiushr) Op() ops.Op            { return ops.Iushr }
func (*IRiushr) Operands() int         { return 0 }
func (*IRiushr) Parse(operands []byte) {}
func (*IRiushr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32((int32)((uint32)(a) >> (b & 0x1f)))
	return nil
}

type IRixor struct{}

func (*IRixor) Op() ops.Op            { return ops.Ixor }
func (*IRixor) Operands() int         { return 0 }
func (*IRixor) Parse(operands []byte) {}
func (*IRixor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	stack.PushInt32(a ^ b)
	return nil
}

type IRldc struct {
	index byte
}

func (*IRldc) Op() ops.Op    { return ops.Ldc }
func (*IRldc) Operanfs() int { return 1 }
func (ir *IRldc) Parse(operanfs []byte) {
	ir.index = operanfs[0]
}
func (ir *IRldc) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst((uint16)(ir.index), vm.GetStack())
}

type IRldc_w struct {
	index uint16
}

func (*IRldc_w) Op() ops.Op    { return ops.Ldc_w }
func (*IRldc_w) Operanfs() int { return 2 }
func (ir *IRldc_w) Parse(operanfs []byte) {
	ir.index = ((uint16)(operanfs[0]) << 8) | (uint16)(operanfs[1])
}
func (ir *IRldc_w) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst(ir.index, vm.GetStack())
}
