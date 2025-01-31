package ir

import (
	"fmt"

	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRaaload struct{}

func (*IRaaload) Op() ops.Op { return ops.Aaload }
func (*IRaaload) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrRef()
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

type IRaastore struct{}

func (*IRaastore) Op() ops.Op { return ops.Aastore }
func (*IRaastore) Execute(vm VM) error {
	stack := vm.GetStack()
	arr := stack.PopArrRef()
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
	if !vm.GetClass(value).IsAssignableFrom(vm.GetArrClass(arr)) {
		return errs.ClassCastException
	}
	arr[index] = value
	return nil
}

type IRaconst_null struct{}

func (*IRaconst_null) Op() ops.Op { return ops.Aconst_null }
func (*IRaconst_null) Execute(vm VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

type IRaload struct {
	Index uint16
}

func (*IRaload) Op() ops.Op { return ops.Aload }
func (ir *IRaload) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef((uint16)(ir.Index))
	stack.PushRef(ref)
	return nil
}

type IRaload_0 struct{}

func (*IRaload_0) Op() ops.Op { return ops.Aload_0 }
func (*IRaload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(0)
	stack.PushRef(ref)
	return nil
}

type IRaload_1 struct{}

func (*IRaload_1) Op() ops.Op { return ops.Aload_1 }
func (*IRaload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	stack.PushRef(ref)
	return nil
}

type IRaload_2 struct{}

func (*IRaload_2) Op() ops.Op { return ops.Aload_2 }
func (*IRaload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(2)
	stack.PushRef(ref)
	return nil
}

type IRaload_3 struct{}

func (*IRaload_3) Op() ops.Op { return ops.Aload_3 }
func (*IRaload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(3)
	stack.PushRef(ref)
	return nil
}

type IRanewarray struct {
	Class uint16
}

func (*IRanewarray) Op() ops.Op { return ops.Anewarray }
func (ir *IRanewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	count := stack.PopInt32()
	if count < 0 {
		return errs.NegativeArraySizeException
	}
	class, err := vm.GetClassByIndex(ir.Class)
	if err != nil {
		return err
	}
	arr := vm.NewArrRef(class, count)
	stack.PushRef(arrayToRef(arr))
	return nil
}

type IRarraylength struct{}

func (*IRarraylength) Op() ops.Op { return ops.Arraylength }
func (*IRarraylength) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	count := arrayLength(ref)
	stack.PushInt32((int32)(count))
	return nil
}

type IRastore struct {
	Index uint16
}

func (*IRastore) Op() ops.Op { return ops.Astore }
func (ir *IRastore) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef((uint16)(ir.Index), ref)
	return nil
}

type IRastore_0 struct{}

func (*IRastore_0) Op() ops.Op { return ops.Astore_0 }
func (*IRastore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(0, ref)
	return nil
}

type IRastore_1 struct{}

func (*IRastore_1) Op() ops.Op { return ops.Astore_1 }
func (*IRastore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(1, ref)
	return nil
}

type IRastore_2 struct{}

func (*IRastore_2) Op() ops.Op { return ops.Astore_2 }
func (*IRastore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(2, ref)
	return nil
}

type IRastore_3 struct{}

func (*IRastore_3) Op() ops.Op { return ops.Astore_3 }
func (*IRastore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(3, ref)
	return nil
}

type IRathrow struct{}

func (*IRathrow) Op() ops.Op { return ops.Athrow }
func (*IRathrow) Execute(vm VM) error {
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

type IRcheckcast struct {
	Class uint16
}

func (*IRcheckcast) Op() ops.Op { return ops.Checkcast }
func (ir *IRcheckcast) Execute(vm VM) error {
	ref := vm.GetStack().PeekRef()
	class, err := vm.GetClassByIndex(ir.Class)
	if err != nil {
		return err
	}
	if ref != nil && !class.IsInstance(ref) {
		return errs.ClassCastException
	}
	return nil
}

type IRgetfield struct {
	Field uint16
}

func (*IRgetfield) Op() ops.Op { return ops.Getfield }
func (ir *IRgetfield) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	field := vm.GetCurrentClass().GetField(ir.Field)
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

type IRgetstatic struct {
	Field uint16
}

func (*IRgetstatic) Op() ops.Op { return ops.Getstatic }
func (ir *IRgetstatic) Execute(vm VM) error {
	stack := vm.GetStack()
	field := vm.GetCurrentClass().GetField(ir.Field)
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

type IRinstanceof struct {
	Class uint16
}

func (*IRinstanceof) Op() ops.Op { return ops.Instanceof }
func (ir *IRinstanceof) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	class, err := vm.GetClassByIndex(ir.Class)
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

type IRinvokedynamic struct {
	Method uint16
}

func (*IRinvokedynamic) Op() ops.Op { return ops.Invokedynamic }
func (ir *IRinvokedynamic) Execute(vm VM) error {
	method := vm.GetCurrentClass().GetMethod(ir.Method)
	if method == nil {
		// TODO: Seems not correct here
		return errs.BootstrapMethodError
	}
	vm.Invoke(method, nil)
	return nil
}

type IRinvokeinterface struct {
	Method uint16
	Count  byte // not in use
}

func (*IRinvokeinterface) Op() ops.Op { return ops.Invokeinterface }
func (ir *IRinvokeinterface) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.Method)
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

type IRinvokespecial struct {
	Method uint16
}

func (*IRinvokespecial) Op() ops.Op { return ops.Invokespecial }
func (ir *IRinvokespecial) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.Method)
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

type IRinvokestatic struct {
	Method uint16
}

func (*IRinvokestatic) Op() ops.Op { return ops.Invokestatic }
func (ir *IRinvokestatic) Execute(vm VM) error {
	method := vm.GetCurrentClass().GetMethod(ir.Method)
	if method == nil {
		return errs.NoSuchMethodError
	}
	if !method.IsStatic() {
		return errs.IncompatibleClassChangeError
	}
	// TODO: access control
	vm.Invoke(method, nil)
	return nil
}

type IRinvokevirtual struct {
	Method uint16
}

func (*IRinvokevirtual) Op() ops.Op { return ops.Invokevirtual }
func (ir *IRinvokevirtual) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.Method)
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

type IRmultianewarray struct {
	Class      uint16
	Dimensions byte
}

func (*IRmultianewarray) Op() ops.Op { return ops.Multianewarray }
func (ir *IRmultianewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	class, err := vm.GetClassByIndex(ir.Class)
	if err != nil {
		return err
	}
	counts := make([]int32, ir.Dimensions)
	for i := range ir.Dimensions {
		count := stack.PopInt32()
		if count < 0 {
			return errs.NegativeArraySizeException
		}
		counts[i] = count
	}
	arr := vm.NewArrRefMultiDim(class, counts)
	stack.PushRef(arrayToRef(arr))
	return nil
}

type IRnew struct {
	Class uint16
}

func (*IRnew) Op() ops.Op { return ops.New }
func (ir *IRnew) Execute(vm VM) error {
	class, err := vm.GetClassByIndex(ir.Class)
	if err != nil {
		return err
	}
	ref := vm.New(class)
	vm.GetStack().PushRef(ref)
	return nil
}

type IRnewarray struct {
	Atype byte
}

func (*IRnewarray) Op() ops.Op { return ops.Newarray }
func (ir *IRnewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	count := stack.PopInt32()
	if count < 0 {
		return errs.NegativeArraySizeException
	}
	var arr Ref
	switch ir.Atype {
	case 4, 8: // T_BOOLEAN, T_BYTE
		arr = arrayToRef(vm.NewArrInt8(count))
	case 5, 9: // T_CHAR, T_SHORT
		arr = arrayToRef(vm.NewArrInt16(count))
	case 6, 10: // T_FLOAT, T_INT
		arr = arrayToRef(vm.NewArrInt32(count))
	case 7, 11: // T_DOUBLE, T_LONG
		arr = arrayToRef(vm.NewArrInt64(count))
	default:
		panic(fmt.Errorf("ir.newarray: unknown atype %d", ir.Atype))
	}
	stack.PushRef(arr)
	return nil
}

type IRputfield struct {
	Field uint16
}

func (*IRputfield) Op() ops.Op { return ops.Putfield }
func (ir *IRputfield) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	field := vm.GetCurrentClass().GetField(ir.Field)
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

type IRputstatic struct {
	Field uint16
}

func (*IRputstatic) Op() ops.Op { return ops.Putstatic }
func (ir *IRputstatic) Execute(vm VM) error {
	stack := vm.GetStack()
	field := vm.GetCurrentClass().GetField(ir.Field)
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
