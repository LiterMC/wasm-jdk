package ir

import (
	"math"

	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRd2f struct{}

func (*IRd2f) Op() ops.Op            { return ops.D2f }
func (*IRd2f) Operands() int         { return 0 }
func (*IRd2f) Parse(operands []byte) {}
func (*IRd2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	stack.PushFloat32((float32)(value))
	return nil
}

type IRd2i struct{}

func (*IRd2i) Op() ops.Op            { return ops.D2i }
func (*IRd2i) Operands() int         { return 0 }
func (*IRd2i) Parse(operands []byte) {}
func (*IRd2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	if value != value {
		stack.PushInt32(0)
	} else {
		stack.PushInt32((int32)(value))
	}
	return nil
}

type IRd2l struct{}

func (*IRd2l) Op() ops.Op            { return ops.D2l }
func (*IRd2l) Operands() int         { return 0 }
func (*IRd2l) Parse(operands []byte) {}
func (*IRd2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	if value != value {
		stack.PushInt64(0)
	} else {
		stack.PushInt64((int64)(value))
	}
	return nil
}

type IRdadd struct{}

func (*IRdadd) Op() ops.Op            { return ops.Dadd }
func (*IRdadd) Operands() int         { return 0 }
func (*IRdadd) Parse(operands []byte) {}
func (*IRdadd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a + b)
	return nil
}

type IRdaload struct{}

func (*IRdaload) Op() ops.Op            { return ops.Daload }
func (*IRdaload) Operands() int         { return 0 }
func (*IRdaload) Parse(operands []byte) {}
func (*IRdaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrFloat64()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushFloat64(arr[index])
	return nil
}

type IRdastore struct{}

func (*IRdastore) Op() ops.Op            { return ops.Dastore }
func (*IRdastore) Operands() int         { return 0 }
func (*IRdastore) Parse(operands []byte) {}
func (*IRdastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrFloat64()
	index := stack.PopInt32()
	value := stack.PopFloat64()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	arr[index] = value
	return nil
}

type IRdcmpg struct{}

func (*IRdcmpg) Op() ops.Op            { return ops.Dcmpg }
func (*IRdcmpg) Operands() int         { return 0 }
func (*IRdcmpg) Parse(operands []byte) {}
func (*IRdcmpg) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	var res int32 = 0
	if a > b {
		res = 1
	} else if a < b {
		res = -1
	} else if a != a || b != b {
		res = 1
	}
	stack.PushInt32(res)
	return nil
}

type IRdcmpl struct{}

func (*IRdcmpl) Op() ops.Op            { return ops.Dcmpl }
func (*IRdcmpl) Operands() int         { return 0 }
func (*IRdcmpl) Parse(operands []byte) {}
func (*IRdcmpl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	var res int32 = 0
	if a > b {
		res = 1
	} else if a < b {
		res = -1
	} else if a != a || b != b {
		res = -1
	}
	stack.PushInt32(res)
	return nil
}

type IRdconst_0 struct{}

func (*IRdconst_0) Op() ops.Op            { return ops.Dconst_0 }
func (*IRdconst_0) Operands() int         { return 0 }
func (*IRdconst_0) Parse(operands []byte) {}
func (*IRdconst_0) Execute(vm VM) error {
	vm.GetStack().PushFloat64(0)
	return nil
}

type IRdconst_1 struct{}

func (*IRdconst_1) Op() ops.Op            { return ops.Dconst_1 }
func (*IRdconst_1) Operands() int         { return 0 }
func (*IRdconst_1) Parse(operands []byte) {}
func (*IRdconst_1) Execute(vm VM) error {
	vm.GetStack().PushFloat64(1)
	return nil
}

type IRddiv struct{}

func (*IRddiv) Op() ops.Op            { return ops.Ddiv }
func (*IRddiv) Operands() int         { return 0 }
func (*IRddiv) Parse(operands []byte) {}
func (*IRddiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a / b)
	return nil
}

type IRdload struct {
	Index uint16
}

func (*IRdload) Op() ops.Op    { return ops.Dload }
func (*IRdload) Operands() int { return 1 }
func (ir *IRdload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}
func (ir *IRdload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64((uint16)(ir.Index))
	stack.PushInt64(val)
	return nil
}

type IRdload_0 struct{}

func (*IRdload_0) Op() ops.Op            { return ops.Dload_0 }
func (*IRdload_0) Operands() int         { return 0 }
func (*IRdload_0) Parse(operands []byte) {}
func (*IRdload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(0)
	stack.PushInt64(val)
	return nil
}

type IRdload_1 struct{}

func (*IRdload_1) Op() ops.Op            { return ops.Dload_1 }
func (*IRdload_1) Operands() int         { return 0 }
func (*IRdload_1) Parse(operands []byte) {}
func (*IRdload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(1)
	stack.PushInt64(val)
	return nil
}

type IRdload_2 struct{}

func (*IRdload_2) Op() ops.Op            { return ops.Dload_2 }
func (*IRdload_2) Operands() int         { return 0 }
func (*IRdload_2) Parse(operands []byte) {}
func (*IRdload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(2)
	stack.PushInt64(val)
	return nil
}

type IRdload_3 struct{}

func (*IRdload_3) Op() ops.Op            { return ops.Dload_3 }
func (*IRdload_3) Operands() int         { return 0 }
func (*IRdload_3) Parse(operands []byte) {}
func (*IRdload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(3)
	stack.PushInt64(val)
	return nil
}

type IRdmul struct{}

func (*IRdmul) Op() ops.Op            { return ops.Dmul }
func (*IRdmul) Operands() int         { return 0 }
func (*IRdmul) Parse(operands []byte) {}
func (*IRdmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a * b)
	return nil
}

type IRdneg struct{}

func (*IRdneg) Op() ops.Op            { return ops.Dneg }
func (*IRdneg) Operands() int         { return 0 }
func (*IRdneg) Parse(operands []byte) {}
func (*IRdneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat64()
	stack.PushFloat64(-a)
	return nil
}

type IRdrem struct{}

func (*IRdrem) Op() ops.Op            { return ops.Drem }
func (*IRdrem) Operands() int         { return 0 }
func (*IRdrem) Parse(operands []byte) {}
func (*IRdrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(math.Mod(a, b))
	return nil
}

type IRdstore struct {
	Index uint16
}

func (*IRdstore) Op() ops.Op    { return ops.Dstore }
func (*IRdstore) Operands() int { return 1 }
func (ir *IRdstore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}
func (ir *IRdstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64((uint16)(ir.Index), val)
	return nil
}

type IRdstore_0 struct{}

func (*IRdstore_0) Op() ops.Op            { return ops.Dstore_0 }
func (*IRdstore_0) Operands() int         { return 0 }
func (*IRdstore_0) Parse(operands []byte) {}
func (*IRdstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(0, val)
	return nil
}

type IRdstore_1 struct{}

func (*IRdstore_1) Op() ops.Op            { return ops.Dstore_1 }
func (*IRdstore_1) Operands() int         { return 0 }
func (*IRdstore_1) Parse(operands []byte) {}
func (*IRdstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(1, val)
	return nil
}

type IRdstore_2 struct{}

func (*IRdstore_2) Op() ops.Op            { return ops.Dstore_2 }
func (*IRdstore_2) Operands() int         { return 0 }
func (*IRdstore_2) Parse(operands []byte) {}
func (*IRdstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(2, val)
	return nil
}

type IRdstore_3 struct{}

func (*IRdstore_3) Op() ops.Op            { return ops.Dstore_3 }
func (*IRdstore_3) Operands() int         { return 0 }
func (*IRdstore_3) Parse(operands []byte) {}
func (*IRdstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(3, val)
	return nil
}

type IRdsub struct{}

func (*IRdsub) Op() ops.Op            { return ops.Dsub }
func (*IRdsub) Operands() int         { return 0 }
func (*IRdsub) Parse(operands []byte) {}
func (*IRdsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a - b)
	return nil
}
