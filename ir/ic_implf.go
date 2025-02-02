package ir

import (
	"math"

	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICf2d struct{}

func (*ICf2d) Op() ops.Op { return ops.F2d }
func (*ICf2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	stack.PushFloat64((float64)(value))
	return nil
}

type ICf2i struct{}

func (*ICf2i) Op() ops.Op { return ops.F2i }
func (*ICf2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	if value != value {
		stack.PushInt32(0)
	} else {
		stack.PushInt32((int32)(value))
	}
	return nil
}

type ICf2l struct{}

func (*ICf2l) Op() ops.Op { return ops.F2l }
func (*ICf2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	if value != value {
		stack.PushInt64(0)
	} else {
		stack.PushInt64((int64)(value))
	}
	return nil
}

type ICfadd struct{}

func (*ICfadd) Op() ops.Op { return ops.Fadd }
func (*ICfadd) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat32()
	b := stack.PopFloat32()
	stack.PushFloat32(a + b)
	return nil
}

type ICfaload struct{}

func (*ICfaload) Op() ops.Op { return ops.Faload }
func (*ICfaload) Execute(vm VM) error {
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

type ICfastore struct{}

func (*ICfastore) Op() ops.Op { return ops.Fastore }
func (*ICfastore) Execute(vm VM) error {
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

type ICfcmpg struct{}

func (*ICfcmpg) Op() ops.Op { return ops.Fcmpg }
func (*ICfcmpg) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
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

type ICfcmpl struct{}

func (*ICfcmpl) Op() ops.Op { return ops.Fcmpl }
func (*ICfcmpl) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
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

type ICfconst_0 struct{}

func (*ICfconst_0) Op() ops.Op { return ops.Fconst_0 }
func (*ICfconst_0) Execute(vm VM) error {
	vm.GetStack().PushFloat32(0)
	return nil
}

type ICfconst_1 struct{}

func (*ICfconst_1) Op() ops.Op { return ops.Fconst_1 }
func (*ICfconst_1) Execute(vm VM) error {
	vm.GetStack().PushFloat32(1)
	return nil
}

type ICfconst_2 struct{}

func (*ICfconst_2) Op() ops.Op { return ops.Fconst_2 }
func (*ICfconst_2) Execute(vm VM) error {
	vm.GetStack().PushFloat32(2)
	return nil
}

type ICfdiv struct{}

func (*ICfdiv) Op() ops.Op { return ops.Fdiv }
func (*ICfdiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a / b)
	return nil
}

type ICfload struct {
	Index uint16
}

func (*ICfload) Op() ops.Op { return ops.Fload }
func (ic *ICfload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32((uint16)(ic.Index))
	stack.PushInt32(val)
	return nil
}

type ICfload_0 struct{}

func (*ICfload_0) Op() ops.Op { return ops.Fload_0 }
func (*ICfload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(0)
	stack.PushInt32(val)
	return nil
}

type ICfload_1 struct{}

func (*ICfload_1) Op() ops.Op { return ops.Fload_1 }
func (*ICfload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(1)
	stack.PushInt32(val)
	return nil
}

type ICfload_2 struct{}

func (*ICfload_2) Op() ops.Op { return ops.Fload_2 }
func (*ICfload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(2)
	stack.PushInt32(val)
	return nil
}

type ICfload_3 struct{}

func (*ICfload_3) Op() ops.Op { return ops.Fload_3 }
func (*ICfload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(3)
	stack.PushInt32(val)
	return nil
}

type ICfmul struct{}

func (*ICfmul) Op() ops.Op { return ops.Fmul }
func (*ICfmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a * b)
	return nil
}

type ICfneg struct{}

func (*ICfneg) Op() ops.Op { return ops.Fneg }
func (*ICfneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat32()
	stack.PushFloat32(-a)
	return nil
}

type ICfrem struct{}

func (*ICfrem) Op() ops.Op { return ops.Frem }
func (*ICfrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32((float32)(math.Mod((float64)(a), (float64)(b))))
	return nil
}

type ICfstore struct {
	Index uint16
}

func (*ICfstore) Op() ops.Op { return ops.Fstore }
func (ic *ICfstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32((uint16)(ic.Index), val)
	return nil
}

type ICfstore_0 struct{}

func (*ICfstore_0) Op() ops.Op { return ops.Fstore_0 }
func (*ICfstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(0, val)
	return nil
}

type ICfstore_1 struct{}

func (*ICfstore_1) Op() ops.Op { return ops.Fstore_1 }
func (*ICfstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(1, val)
	return nil
}

type ICfstore_2 struct{}

func (*ICfstore_2) Op() ops.Op { return ops.Fstore_2 }
func (*ICfstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(2, val)
	return nil
}

type ICfstore_3 struct{}

func (*ICfstore_3) Op() ops.Op { return ops.Fstore_3 }
func (*ICfstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(3, val)
	return nil
}

type ICfsub struct{}

func (*ICfsub) Op() ops.Op { return ops.Fsub }
func (*ICfsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a - b)
	return nil
}
