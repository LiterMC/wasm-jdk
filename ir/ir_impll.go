package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRl2d struct{}

func (*IRl2d) Op() ops.Op            { return ops.L2d }
func (*IRl2d) Operands() int         { return 0 }
func (*IRl2d) Parse(operands []byte) {}
func (*IRl2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushFloat64((float64)(value))
	return nil
}

type IRl2f struct{}

func (*IRl2f) Op() ops.Op            { return ops.L2f }
func (*IRl2f) Operands() int         { return 0 }
func (*IRl2f) Parse(operands []byte) {}
func (*IRl2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushFloat32((float32)(value))
	return nil
}

type IRl2i struct{}

func (*IRl2i) Op() ops.Op            { return ops.L2i }
func (*IRl2i) Operands() int         { return 0 }
func (*IRl2i) Parse(operands []byte) {}
func (*IRl2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopInt64()
	stack.PushInt32((int32)(value))
	return nil
}

type IRladd struct{}

func (*IRladd) Op() ops.Op            { return ops.Ladd }
func (*IRladd) Operands() int         { return 0 }
func (*IRladd) Parse(operands []byte) {}
func (*IRladd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a + b)
	return nil
}

type IRlaload struct{}

func (*IRlaload) Op() ops.Op            { return ops.Laload }
func (*IRlaload) Operands() int         { return 0 }
func (*IRlaload) Parse(operands []byte) {}
func (*IRlaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt64()
	index := stack.PopInt64()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushInt64(arr[index])
	return nil
}

type IRland struct{}

func (*IRland) Op() ops.Op            { return ops.Land }
func (*IRland) Operands() int         { return 0 }
func (*IRland) Parse(operands []byte) {}
func (*IRland) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a & b)
	return nil
}

type IRlastore struct{}

func (*IRlastore) Op() ops.Op            { return ops.Lastore }
func (*IRlastore) Operands() int         { return 0 }
func (*IRlastore) Parse(operands []byte) {}
func (*IRlastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrInt64()
	index := stack.PopInt32()
	value := stack.PopInt64()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IRlcmp struct{}

func (*IRlcmp) Op() ops.Op            { return ops.Lcmp }
func (*IRlcmp) Operands() int         { return 0 }
func (*IRlcmp) Parse(operands []byte) {}
func (*IRlcmp) Execute(vm VM) error {
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

type IRlconst_0 struct{}

func (*IRlconst_0) Op() ops.Op            { return ops.Lconst_0 }
func (*IRlconst_0) Operands() int         { return 0 }
func (*IRlconst_0) Parse(operands []byte) {}
func (*IRlconst_0) Execute(vm VM) error {
	vm.GetStack().PushInt64(0)
	return nil
}

type IRlconst_1 struct{}

func (*IRlconst_1) Op() ops.Op            { return ops.Lconst_1 }
func (*IRlconst_1) Operands() int         { return 0 }
func (*IRlconst_1) Parse(operands []byte) {}
func (*IRlconst_1) Execute(vm VM) error {
	vm.GetStack().PushInt64(1)
	return nil
}

type IRldc2_w struct {
	index uint16
}

func (*IRldc2_w) Op() ops.Op    { return ops.Ldc2_w }
func (*IRldc2_w) Operands() int { return 2 }
func (ir *IRldc2_w) Parse(operands []byte) {
	ir.index = bytesToUint16(operands)
}
func (ir *IRldc2_w) Execute(vm VM) error {
	return vm.GetCurrentClass().GetAndPushConst(ir.index, vm.GetStack())
}

type IRldiv struct{}

func (*IRldiv) Op() ops.Op            { return ops.Ldiv }
func (*IRldiv) Operands() int         { return 0 }
func (*IRldiv) Parse(operands []byte) {}
func (*IRldiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a / b)
	return nil
}

type IRlload struct {
	index byte
}

func (*IRlload) Op() ops.Op    { return ops.Lload }
func (*IRlload) Operands() int { return 1 }
func (ir *IRlload) Parse(operands []byte) {
	ir.index = operands[0]
}
func (ir *IRlload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64((uint16)(ir.index))
	stack.PushInt64(val)
	return nil
}

type IRlload_0 struct{}

func (*IRlload_0) Op() ops.Op            { return ops.Lload_0 }
func (*IRlload_0) Operands() int         { return 0 }
func (*IRlload_0) Parse(operands []byte) {}
func (*IRlload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(0)
	stack.PushInt64(val)
	return nil
}

type IRlload_1 struct{}

func (*IRlload_1) Op() ops.Op            { return ops.Lload_1 }
func (*IRlload_1) Operands() int         { return 0 }
func (*IRlload_1) Parse(operands []byte) {}
func (*IRlload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(1)
	stack.PushInt64(val)
	return nil
}

type IRlload_2 struct{}

func (*IRlload_2) Op() ops.Op            { return ops.Lload_2 }
func (*IRlload_2) Operands() int         { return 0 }
func (*IRlload_2) Parse(operands []byte) {}
func (*IRlload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(2)
	stack.PushInt64(val)
	return nil
}

type IRlload_3 struct{}

func (*IRlload_3) Op() ops.Op            { return ops.Lload_3 }
func (*IRlload_3) Operands() int         { return 0 }
func (*IRlload_3) Parse(operands []byte) {}
func (*IRlload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(3)
	stack.PushInt64(val)
	return nil
}

type IRlmul struct{}

func (*IRlmul) Op() ops.Op            { return ops.Lmul }
func (*IRlmul) Operands() int         { return 0 }
func (*IRlmul) Parse(operands []byte) {}
func (*IRlmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a * b)
	return nil
}

type IRlneg struct{}

func (*IRlneg) Op() ops.Op            { return ops.Lneg }
func (*IRlneg) Operands() int         { return 0 }
func (*IRlneg) Parse(operands []byte) {}
func (*IRlneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt64()
	stack.PushInt64(-a)
	return nil
}

type IRlor struct{}

func (*IRlor) Op() ops.Op            { return ops.Lor }
func (*IRlor) Operands() int         { return 0 }
func (*IRlor) Parse(operands []byte) {}
func (*IRlor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a | b)
	return nil
}

type IRlrem struct{}

func (*IRlrem) Op() ops.Op            { return ops.Lrem }
func (*IRlrem) Operands() int         { return 0 }
func (*IRlrem) Parse(operands []byte) {}
func (*IRlrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a % b)
	return nil
}

type IRlshl struct{}

func (*IRlshl) Op() ops.Op            { return ops.Lshl }
func (*IRlshl) Operands() int         { return 0 }
func (*IRlshl) Parse(operands []byte) {}
func (*IRlshl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt64()
	stack.PushInt64(a << (b & 0x3f))
	return nil
}

type IRlshr struct{}

func (*IRlshr) Op() ops.Op            { return ops.Lshr }
func (*IRlshr) Operands() int         { return 0 }
func (*IRlshr) Parse(operands []byte) {}
func (*IRlshr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt64()
	stack.PushInt64(a >> (b & 0x3f))
	return nil
}

type IRlstore struct {
	index byte
}

func (*IRlstore) Op() ops.Op    { return ops.Lstore }
func (*IRlstore) Operands() int { return 1 }
func (ir *IRlstore) Parse(operands []byte) {
	ir.index = operands[0]
}
func (ir *IRlstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64((uint16)(ir.index), val)
	return nil
}

type IRlstore_0 struct{}

func (*IRlstore_0) Op() ops.Op            { return ops.Lstore_0 }
func (*IRlstore_0) Operands() int         { return 0 }
func (*IRlstore_0) Parse(operands []byte) {}
func (*IRlstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(0, val)
	return nil
}

type IRlstore_1 struct{}

func (*IRlstore_1) Op() ops.Op            { return ops.Lstore_1 }
func (*IRlstore_1) Operands() int         { return 0 }
func (*IRlstore_1) Parse(operands []byte) {}
func (*IRlstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(1, val)
	return nil
}

type IRlstore_2 struct{}

func (*IRlstore_2) Op() ops.Op            { return ops.Lstore_2 }
func (*IRlstore_2) Operands() int         { return 0 }
func (*IRlstore_2) Parse(operands []byte) {}
func (*IRlstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(2, val)
	return nil
}

type IRlstore_3 struct{}

func (*IRlstore_3) Op() ops.Op            { return ops.Lstore_3 }
func (*IRlstore_3) Operands() int         { return 0 }
func (*IRlstore_3) Parse(operands []byte) {}
func (*IRlstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(3, val)
	return nil
}

type IRlsub struct{}

func (*IRlsub) Op() ops.Op            { return ops.Lsub }
func (*IRlsub) Operands() int         { return 0 }
func (*IRlsub) Parse(operands []byte) {}
func (*IRlsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a - b)
	return nil
}

type IRlushr struct{}

func (*IRlushr) Op() ops.Op            { return ops.Lushr }
func (*IRlushr) Operands() int         { return 0 }
func (*IRlushr) Parse(operands []byte) {}
func (*IRlushr) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt64()
	stack.PushInt64((int64)((uint64)(a) >> (b & 0x1f)))
	return nil
}

type IRlxor struct{}

func (*IRlxor) Op() ops.Op            { return ops.Lxor }
func (*IRlxor) Operands() int         { return 0 }
func (*IRlxor) Parse(operands []byte) {}
func (*IRlxor) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt64()
	a := stack.PopInt64()
	stack.PushInt64(a ^ b)
	return nil
}
