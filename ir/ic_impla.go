package ir

import (
	"fmt"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICaaload struct{}

func (*ICaaload) Op() ops.Op { return ops.Aaload }
func (*ICaaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopRef().GetArrRef()
	index := stack.PopInt32()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	stack.PushRef(arr[index])
	return nil
}

type ICaastore struct{}

func (*ICaastore) Op() ops.Op { return ops.Aastore }
func (*ICaastore) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	arr := ref.GetArrRef()
	index := stack.PopInt32()
	value := stack.PopRef()
	if arr == nil {
		return errs.NullPointerException
	}
	if index < 0 || (int)(index) >= len(arr) {
		return errs.ArrayIndexOutOfBoundsException
	}
	if value == nil {
		arr[index] = value
	}
	if !vm.GetClass(value).IsAssignableFrom(ref.Class()) {
		return errs.ClassCastException
	}
	arr[index] = value
	return nil
}

type ICaconst_null struct{}

func (*ICaconst_null) Op() ops.Op { return ops.Aconst_null }
func (*ICaconst_null) Execute(vm VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

type ICaload struct {
	Index uint16
}

func (*ICaload) Op() ops.Op { return ops.Aload }
func (ic *ICaload) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef((uint16)(ic.Index))
	stack.PushRef(ref)
	return nil
}

type ICaload_0 struct{}

func (*ICaload_0) Op() ops.Op { return ops.Aload_0 }
func (*ICaload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(0)
	stack.PushRef(ref)
	return nil
}

type ICaload_1 struct{}

func (*ICaload_1) Op() ops.Op { return ops.Aload_1 }
func (*ICaload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	stack.PushRef(ref)
	return nil
}

type ICaload_2 struct{}

func (*ICaload_2) Op() ops.Op { return ops.Aload_2 }
func (*ICaload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(2)
	stack.PushRef(ref)
	return nil
}

type ICaload_3 struct{}

func (*ICaload_3) Op() ops.Op { return ops.Aload_3 }
func (*ICaload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(3)
	stack.PushRef(ref)
	return nil
}

type ICanewarray struct {
	Class uint16
}

func (*ICanewarray) Op() ops.Op { return ops.Anewarray }
func (ic *ICanewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	count := stack.PopInt32()
	if count < 0 {
		return errs.NegativeArraySizeException
	}
	class, err := vm.GetClassByIndex(ic.Class)
	if err != nil {
		return err
	}
	arr := vm.NewArrayByClass(class, count)
	stack.PushRef(arr)
	return nil
}

type ICarraylength struct{}

func (*ICarraylength) Op() ops.Op { return ops.Arraylength }
func (*ICarraylength) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	stack.PushInt32(ref.Len())
	return nil
}

type ICastore struct {
	Index uint16
}

func (*ICastore) Op() ops.Op { return ops.Astore }
func (ic *ICastore) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef((uint16)(ic.Index), ref)
	return nil
}

type ICastore_0 struct{}

func (*ICastore_0) Op() ops.Op { return ops.Astore_0 }
func (*ICastore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(0, ref)
	return nil
}

type ICastore_1 struct{}

func (*ICastore_1) Op() ops.Op { return ops.Astore_1 }
func (*ICastore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(1, ref)
	return nil
}

type ICastore_2 struct{}

func (*ICastore_2) Op() ops.Op { return ops.Astore_2 }
func (*ICastore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(2, ref)
	return nil
}

type ICastore_3 struct{}

func (*ICastore_3) Op() ops.Op { return ops.Astore_3 }
func (*ICastore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(3, ref)
	return nil
}

type ICathrow struct{}

func (*ICathrow) Op() ops.Op { return ops.Athrow }
func (*ICathrow) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	// TODO: is this necessary?
	// if !vm.GetThrowableClass().IsInstance(ref) {
	// 	return ClassCastException
	// }
	vm.Throw(ref)
	return nil
}

type ICcheckcast struct {
	Class uint16
}

func (*ICcheckcast) Op() ops.Op { return ops.Checkcast }
func (ic *ICcheckcast) Execute(vm VM) error {
	ref := vm.GetStack().PeekRef()
	class, err := vm.GetClassByIndex(ic.Class)
	if err != nil {
		return err
	}
	if ref != nil && !class.IsInstance(ref) {
		return errs.ClassCastException
	}
	return nil
}

type ICgetfield struct {
	Field uint16
}

func (*ICgetfield) Op() ops.Op { return ops.Getfield }
func (ic *ICgetfield) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	field := vm.GetCurrentClass().GetField(ic.Field)
	if field == nil {
		return errs.NoSuchFieldError
	}
	if field.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	field.GetAndPush(ref, stack)
	return nil
}

type ICgetstatic struct {
	Field uint16
}

func (*ICgetstatic) Op() ops.Op { return ops.Getstatic }
func (ic *ICgetstatic) Execute(vm VM) error {
	stack := vm.GetStack()
	field := vm.GetCurrentClass().GetField(ic.Field)
	if field == nil {
		return errs.NoSuchFieldError
	}
	if !field.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	field.GetAndPush(nil, stack)
	return nil
}

type ICinstanceof struct {
	Class uint16
}

func (*ICinstanceof) Op() ops.Op { return ops.Instanceof }
func (ic *ICinstanceof) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	class, err := vm.GetClassByIndex(ic.Class)
	if err != nil {
		return err
	}
	if ref != nil && class.IsInstance(ref) {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

type ICinvokedynamic struct {
	Method uint16
}

func (*ICinvokedynamic) Op() ops.Op { return ops.Invokedynamic }
func (ic *ICinvokedynamic) Execute(vm VM) error {
	return vm.InvokeDynamic(ic.Method)
}

type ICinvokeinterface struct {
	Method uint16
	Count  byte // not in use
}

func (*ICinvokeinterface) Op() ops.Op { return ops.Invokeinterface }
func (ic *ICinvokeinterface) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ic.Method)
	if method == nil {
		return errs.NoSuchMethodError
	}
	if method.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	vm.Invoke(method, ref)
	return nil
}

type ICinvokespecial struct {
	Method uint16
}

func (*ICinvokespecial) Op() ops.Op { return ops.Invokespecial }
func (ic *ICinvokespecial) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ic.Method)
	if method == nil {
		return errs.NoSuchMethodError
	}
	if method.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	vm.Invoke(method, ref)
	return nil
}

type ICinvokestatic struct {
	Method uint16
}

func (*ICinvokestatic) Op() ops.Op { return ops.Invokestatic }
func (ic *ICinvokestatic) Execute(vm VM) error {
	method := vm.GetCurrentClass().GetMethod(ic.Method)
	if method == nil {
		return errs.NoSuchMethodError
	}
	if !method.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	vm.InvokeStatic(method)
	return nil
}

type ICinvokevirtual struct {
	Method uint16
}

func (*ICinvokevirtual) Op() ops.Op { return ops.Invokevirtual }
func (ic *ICinvokevirtual) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ic.Method)
	if method == nil {
		return errs.NoSuchMethodError
	}
	if method.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	vm.Invoke(method, ref)
	return nil
}

type ICmultianewarray struct {
	Desc       uint16
	Dimensions byte
}

func (*ICmultianewarray) Op() ops.Op { return ops.Multianewarray }
func (ic *ICmultianewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	desc := vm.GetDesc(ic.Desc)
	counts := make([]int32, ic.Dimensions)
	for i := range ic.Dimensions {
		count := stack.PopInt32()
		if count < 0 {
			return errs.NegativeArraySizeException
		}
		counts[i] = count
	}
	arr := vm.NewArrayMultiDim(desc, counts)
	stack.PushRef(arr)
	return nil
}

type ICnew struct {
	Class uint16
}

func (*ICnew) Op() ops.Op { return ops.New }
func (ic *ICnew) Execute(vm VM) error {
	class, err := vm.GetClassByIndex(ic.Class)
	if err != nil {
		return err
	}
	ref := vm.New(class)
	vm.GetStack().PushRef(ref)
	return nil
}

type ICnewarray struct {
	Atype byte
}

func (*ICnewarray) Op() ops.Op { return ops.Newarray }
func (ic *ICnewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	count := stack.PopInt32()
	if count < 0 {
		return errs.NegativeArraySizeException
	}
	var arr Ref
	switch ic.Atype {
	case 4, 8: // T_BOOLEAN, T_BYTE
		arr = vm.NewArray(desc.DescInt8, count)
	case 5, 9: // T_CHAR, T_SHORT
		arr = vm.NewArray(desc.DescInt16, count)
	case 6, 10: // T_FLOAT, T_INT
		arr = vm.NewArray(desc.DescInt32, count)
	case 7, 11: // T_DOUBLE, T_LONG
		arr = vm.NewArray(desc.DescInt64, count)
	default:
		panic(fmt.Errorf("ir.newarray: unknown atype %d", ic.Atype))
	}
	stack.PushRef(arr)
	return nil
}

type ICputfield struct {
	Field uint16
}

func (*ICputfield) Op() ops.Op { return ops.Putfield }
func (ic *ICputfield) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	field := vm.GetCurrentClass().GetField(ic.Field)
	if field == nil {
		return errs.NoSuchFieldError
	}
	if field.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	field.PopAndSet(ref, stack)
	return nil
}

type ICputstatic struct {
	Field uint16
}

func (*ICputstatic) Op() ops.Op { return ops.Putstatic }
func (ic *ICputstatic) Execute(vm VM) error {
	stack := vm.GetStack()
	field := vm.GetCurrentClass().GetField(ic.Field)
	if field == nil {
		return errs.NoSuchFieldError
	}
	if !field.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	field.PopAndSet(nil, stack)
	return nil
}
