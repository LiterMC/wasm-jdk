package ir

import (
	"math"

	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRf2d struct{}

func (*IRf2d) Op() ops.Op            { return ops.F2d }
func (*IRf2d) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	stack.PushFloat64((float64)(value))
	return nil
}

type IRf2i struct{}

func (*IRf2i) Op() ops.Op            { return ops.F2i }
func (*IRf2i) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	if value != value {
		stack.PushInt32(0)
	} else {
		stack.PushInt32((int32)(value))
	}
	return nil
}

type IRf2l struct{}

func (*IRf2l) Op() ops.Op            { return ops.F2l }
func (*IRf2l) Execute(vm VM) error {
	stack := vm.GetStack()
	value := stack.PopFloat32()
	if value != value {
		stack.PushInt64(0)
	} else {
		stack.PushInt64((int64)(value))
	}
	return nil
}

type IRfadd struct{}

func (*IRfadd) Op() ops.Op            { return ops.Fadd }
func (*IRfadd) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat32()
	b := stack.PopFloat32()
	stack.PushFloat32(a + b)
	return nil
}

type IRfaload struct{}

func (*IRfaload) Op() ops.Op            { return ops.Faload }
func (*IRfaload) Execute(vm VM) error {
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

type IRfastore struct{}

func (*IRfastore) Op() ops.Op            { return ops.Fastore }
func (*IRfastore) Execute(vm VM) error {
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

type IRfcmpg struct{}

func (*IRfcmpg) Op() ops.Op            { return ops.Fcmpg }
func (*IRfcmpg) Execute(vm VM) error {
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

type IRfcmpl struct{}

func (*IRfcmpl) Op() ops.Op            { return ops.Fcmpl }
func (*IRfcmpl) Execute(vm VM) error {
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

type IRfconst_0 struct{}

func (*IRfconst_0) Op() ops.Op            { return ops.Fconst_0 }
func (*IRfconst_0) Execute(vm VM) error {
	vm.GetStack().PushFloat32(0)
	return nil
}

type IRfconst_1 struct{}

func (*IRfconst_1) Op() ops.Op            { return ops.Fconst_1 }
func (*IRfconst_1) Execute(vm VM) error {
	vm.GetStack().PushFloat32(1)
	return nil
}

type IRfconst_2 struct{}

func (*IRfconst_2) Op() ops.Op            { return ops.Fconst_2 }
func (*IRfconst_2) Execute(vm VM) error {
	vm.GetStack().PushFloat32(2)
	return nil
}

type IRfdiv struct{}

func (*IRfdiv) Op() ops.Op            { return ops.Fdiv }
func (*IRfdiv) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a / b)
	return nil
}

type IRfload struct {
	Index uint16
}

func (*IRfload) Op() ops.Op    { return ops.Fload }
func (ir *IRfload) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32((uint16)(ir.Index))
	stack.PushInt32(val)
	return nil
}

type IRfload_0 struct{}

func (*IRfload_0) Op() ops.Op            { return ops.Fload_0 }
func (*IRfload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(0)
	stack.PushInt32(val)
	return nil
}

type IRfload_1 struct{}

func (*IRfload_1) Op() ops.Op            { return ops.Fload_1 }
func (*IRfload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(1)
	stack.PushInt32(val)
	return nil
}

type IRfload_2 struct{}

func (*IRfload_2) Op() ops.Op            { return ops.Fload_2 }
func (*IRfload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(2)
	stack.PushInt32(val)
	return nil
}

type IRfload_3 struct{}

func (*IRfload_3) Op() ops.Op            { return ops.Fload_3 }
func (*IRfload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.GetVarInt32(3)
	stack.PushInt32(val)
	return nil
}

type IRfmul struct{}

func (*IRfmul) Op() ops.Op            { return ops.Fmul }
func (*IRfmul) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a * b)
	return nil
}

type IRfneg struct{}

func (*IRfneg) Op() ops.Op            { return ops.Fneg }
func (*IRfneg) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopFloat32()
	stack.PushFloat32(-a)
	return nil
}

type IRfrem struct{}

func (*IRfrem) Op() ops.Op            { return ops.Frem }
func (*IRfrem) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32((float32)(math.Mod((float64)(a), (float64)(b))))
	return nil
}

type IRfstore struct {
	Index uint16
}

func (*IRfstore) Op() ops.Op    { return ops.Fstore }
func (ir *IRfstore) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32((uint16)(ir.Index), val)
	return nil
}

type IRfstore_0 struct{}

func (*IRfstore_0) Op() ops.Op            { return ops.Fstore_0 }
func (*IRfstore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(0, val)
	return nil
}

type IRfstore_1 struct{}

func (*IRfstore_1) Op() ops.Op            { return ops.Fstore_1 }
func (*IRfstore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(1, val)
	return nil
}

type IRfstore_2 struct{}

func (*IRfstore_2) Op() ops.Op            { return ops.Fstore_2 }
func (*IRfstore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(2, val)
	return nil
}

type IRfstore_3 struct{}

func (*IRfstore_3) Op() ops.Op            { return ops.Fstore_3 }
func (*IRfstore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PopInt32()
	stack.SetVarInt32(3, val)
	return nil
}

type IRfsub struct{}

func (*IRfsub) Op() ops.Op            { return ops.Fsub }
func (*IRfsub) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopFloat32()
	a := stack.PopFloat32()
	stack.PushFloat32(a - b)
	return nil
}
