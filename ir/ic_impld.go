package ir

import (
	"math"

	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICd2f struct{}

func (*ICd2f) Op() ops.Op { return ops.D2f }
func (*ICd2f) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	stack.PushFloat32((float32)(value))
	return nil
}

type ICd2i struct{}

func (*ICd2i) Op() ops.Op { return ops.D2i }
func (*ICd2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	if value != value {
		stack.PushInt32(0)
	} else {
		stack.PushInt32((int32)(value))
	}
	return nil
}

type ICd2l struct{}

func (*ICd2l) Op() ops.Op { return ops.D2l }
func (*ICd2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat64()
	if value != value {
		stack.PushInt64(0)
	} else {
		stack.PushInt64((int64)(value))
	}
	return nil
}

type ICdadd struct{}

func (*ICdadd) Op() ops.Op { return ops.Dadd }
func (*ICdadd) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a + b)
	return nil
}

type ICdaload struct{}

func (*ICdaload) Op() ops.Op { return ops.Daload }
func (*ICdaload) Execute(vm VM) error {
	stack := vm.GetStack()
	index := stack.PopInt32()
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

type ICdastore struct{}

func (*ICdastore) Op() ops.Op { return ops.Dastore }
func (*ICdastore) Execute(vm VM) error {
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

type ICdcmpg struct{}

func (*ICdcmpg) Op() ops.Op { return ops.Dcmpg }
func (*ICdcmpg) Execute(vm VM) error {
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

type ICdcmpl struct{}

func (*ICdcmpl) Op() ops.Op { return ops.Dcmpl }
func (*ICdcmpl) Execute(vm VM) error {
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

type ICdconst_0 struct{}

func (*ICdconst_0) Op() ops.Op { return ops.Dconst_0 }
func (*ICdconst_0) Execute(vm VM) error {
	vm.GetStack().PushFloat64(0)
	return nil
}

type ICdconst_1 struct{}

func (*ICdconst_1) Op() ops.Op { return ops.Dconst_1 }
func (*ICdconst_1) Execute(vm VM) error {
	vm.GetStack().PushFloat64(1)
	return nil
}

type ICddiv struct{}

func (*ICddiv) Op() ops.Op { return ops.Ddiv }
func (*ICddiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a / b)
	return nil
}

type ICdload struct {
	Index uint16
}

func (*ICdload) Op() ops.Op { return ops.Dload }
func (ic *ICdload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64((uint16)(ic.Index))
	stack.PushInt64(val)
	return nil
}

type ICdload_0 struct{}

func (*ICdload_0) Op() ops.Op { return ops.Dload_0 }
func (*ICdload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(0)
	stack.PushInt64(val)
	return nil
}

type ICdload_1 struct{}

func (*ICdload_1) Op() ops.Op { return ops.Dload_1 }
func (*ICdload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(1)
	stack.PushInt64(val)
	return nil
}

type ICdload_2 struct{}

func (*ICdload_2) Op() ops.Op { return ops.Dload_2 }
func (*ICdload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(2)
	stack.PushInt64(val)
	return nil
}

type ICdload_3 struct{}

func (*ICdload_3) Op() ops.Op { return ops.Dload_3 }
func (*ICdload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt64(3)
	stack.PushInt64(val)
	return nil
}

type ICdmul struct{}

func (*ICdmul) Op() ops.Op { return ops.Dmul }
func (*ICdmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a * b)
	return nil
}

type ICdneg struct{}

func (*ICdneg) Op() ops.Op { return ops.Dneg }
func (*ICdneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat64()
	stack.PushFloat64(-a)
	return nil
}

type ICdrem struct{}

func (*ICdrem) Op() ops.Op { return ops.Drem }
func (*ICdrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(math.Mod(a, b))
	return nil
}

type ICdstore struct {
	Index uint16
}

func (*ICdstore) Op() ops.Op { return ops.Dstore }
func (ic *ICdstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64((uint16)(ic.Index), val)
	return nil
}

type ICdstore_0 struct{}

func (*ICdstore_0) Op() ops.Op { return ops.Dstore_0 }
func (*ICdstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(0, val)
	return nil
}

type ICdstore_1 struct{}

func (*ICdstore_1) Op() ops.Op { return ops.Dstore_1 }
func (*ICdstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(1, val)
	return nil
}

type ICdstore_2 struct{}

func (*ICdstore_2) Op() ops.Op { return ops.Dstore_2 }
func (*ICdstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(2, val)
	return nil
}

type ICdstore_3 struct{}

func (*ICdstore_3) Op() ops.Op { return ops.Dstore_3 }
func (*ICdstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt64()
	stack.SetVarInt64(3, val)
	return nil
}

type ICdsub struct{}

func (*ICdsub) Op() ops.Op { return ops.Dsub }
func (*ICdsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat64()
	a := stack.PopFloat64()
	stack.PushFloat64(a - b)
	return nil
}
