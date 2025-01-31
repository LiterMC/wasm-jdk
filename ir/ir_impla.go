package ir

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRaaload struct{}

func (*IRaaload) Op() ops.Op            { return ops.Aaload }
func (*IRaaload) Operands() int         { return 0 }
func (*IRaaload) Parse(operands []byte) {}
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

func (*IRaastore) Op() ops.Op            { return ops.Aastore }
func (*IRaastore) Operands() int         { return 0 }
func (*IRaastore) Parse(operands []byte) {}
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

func (*IRaconst_null) Op() ops.Op            { return ops.Aconst_null }
func (*IRaconst_null) Operands() int         { return 0 }
func (*IRaconst_null) Parse(operands []byte) {}
func (*IRaconst_null) Execute(vm VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

type IRaload struct {
	index byte
}

func (*IRaload) Op() ops.Op    { return ops.Aload }
func (*IRaload) Operands() int { return 1 }
func (ir *IRaload) Parse(operands []byte) {
	ir.index = operands[0]
}
func (ir *IRaload) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef((uint16)(ir.index))
	stack.PushRef(ref)
	return nil
}

type IRaload_0 struct{}

func (*IRaload_0) Op() ops.Op            { return ops.Aload_0 }
func (*IRaload_0) Operands() int         { return 0 }
func (*IRaload_0) Parse(operands []byte) {}
func (*IRaload_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(0)
	stack.PushRef(ref)
	return nil
}

type IRaload_1 struct{}

func (*IRaload_1) Op() ops.Op            { return ops.Aload_1 }
func (*IRaload_1) Operands() int         { return 0 }
func (*IRaload_1) Parse(operands []byte) {}
func (*IRaload_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	stack.PushRef(ref)
	return nil
}

type IRaload_2 struct{}

func (*IRaload_2) Op() ops.Op            { return ops.Aload_2 }
func (*IRaload_2) Operands() int         { return 0 }
func (*IRaload_2) Parse(operands []byte) {}
func (*IRaload_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(2)
	stack.PushRef(ref)
	return nil
}

type IRaload_3 struct{}

func (*IRaload_3) Op() ops.Op            { return ops.Aload_3 }
func (*IRaload_3) Operands() int         { return 0 }
func (*IRaload_3) Parse(operands []byte) {}
func (*IRaload_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(3)
	stack.PushRef(ref)
	return nil
}

type IRanewarray struct {
	class uint16
}

func (*IRanewarray) Op() ops.Op    { return ops.Anewarray }
func (*IRanewarray) Operands() int { return 2 }
func (ir *IRanewarray) Parse(operands []byte) {
	ir.class = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRanewarray) Execute(vm VM) error {
	stack := vm.GetStack()
	count := stack.PopInt32()
	if count < 0 {
		return errs.NegativeArraySizeException
	}
	class, err := vm.GetClassByIndex(ir.class)
	if err != nil {
		return err
	}
	arr := vm.NewArrRef(class, count)
	stack.PushRef(arrayToRef(arr))
	return nil
}

type IRarraylength struct{}

func (*IRarraylength) Op() ops.Op            { return ops.Arraylength }
func (*IRarraylength) Operands() int         { return 0 }
func (*IRarraylength) Parse(operands []byte) {}
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
	index byte
}

func (*IRastore) Op() ops.Op    { return ops.Astore }
func (*IRastore) Operands() int { return 1 }
func (ir *IRastore) Parse(operands []byte) {
	ir.index = operands[0]
}
func (ir *IRastore) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef((uint16)(ir.index), ref)
	return nil
}

type IRastore_0 struct{}

func (*IRastore_0) Op() ops.Op            { return ops.Astore_0 }
func (*IRastore_0) Operands() int         { return 0 }
func (*IRastore_0) Parse(operands []byte) {}
func (*IRastore_0) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(0, ref)
	return nil
}

type IRastore_1 struct{}

func (*IRastore_1) Op() ops.Op            { return ops.Astore_1 }
func (*IRastore_1) Operands() int         { return 0 }
func (*IRastore_1) Parse(operands []byte) {}
func (*IRastore_1) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(1, ref)
	return nil
}

type IRastore_2 struct{}

func (*IRastore_2) Op() ops.Op            { return ops.Astore_2 }
func (*IRastore_2) Operands() int         { return 0 }
func (*IRastore_2) Parse(operands []byte) {}
func (*IRastore_2) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(2, ref)
	return nil
}

type IRastore_3 struct{}

func (*IRastore_3) Op() ops.Op            { return ops.Astore_3 }
func (*IRastore_3) Operands() int         { return 0 }
func (*IRastore_3) Parse(operands []byte) {}
func (*IRastore_3) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	stack.SetVarRef(3, ref)
	return nil
}

type IRathrow struct{}

func (*IRathrow) Op() ops.Op            { return ops.Athrow }
func (*IRathrow) Operands() int         { return 0 }
func (*IRathrow) Parse(operands []byte) {}
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
	class uint16
}

func (*IRcheckcast) Op() ops.Op    { return ops.Checkcast }
func (*IRcheckcast) Operands() int { return 2 }
func (ir *IRcheckcast) Parse(operands []byte) {
	ir.class = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRcheckcast) Execute(vm VM) error {
	ref := vm.GetStack().PeekRef()
	class, err := vm.GetClassByIndex(ir.class)
	if err != nil {
		return err
	}
	if ref != nil && !class.IsInstance(ref) {
		return errs.ClassCastException
	}
	return nil
}

type IRgetfield struct {
	field uint16
}

func (*IRgetfield) Op() ops.Op    { return ops.Getfield }
func (*IRgetfield) Operands() int { return 2 }
func (ir *IRgetfield) Parse(operands []byte) {
	ir.field = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRgetfield) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	if ref == nil {
		return errs.NullPointerException
	}
	field := vm.GetCurrentClass().GetField(ir.field)
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
	field uint16
}

func (*IRgetstatic) Op() ops.Op    { return ops.Getstatic }
func (*IRgetstatic) Operands() int { return 2 }
func (ir *IRgetstatic) Parse(operands []byte) {
	ir.field = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRgetstatic) Execute(vm VM) error {
	stack := vm.GetStack()
	field := vm.GetCurrentClass().GetField(ir.field)
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
	class uint16
}

func (*IRinstanceof) Op() ops.Op    { return ops.Instanceof }
func (*IRinstanceof) Operands() int { return 2 }
func (ir *IRinstanceof) Parse(operands []byte) {
	ir.class = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRinstanceof) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	class, err := vm.GetClassByIndex(ir.class)
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
	method uint16
}

func (*IRinvokedynamic) Op() ops.Op    { return ops.Invokedynamic }
func (*IRinvokedynamic) Operands() int { return 4 }
func (ir *IRinvokedynamic) Parse(operands []byte) {
	ir.method = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
	if operands[2] != 0 || operands[3] != 0 {
		panic("ir.invokedynamic: operands [2] and [3] must be 0")
	}
}
func (ir *IRinvokedynamic) Execute(vm VM) error {
	method := vm.GetCurrentClass().GetMethod(ir.method)
	if method == nil {
		// TODO: Seems not correct here
		return errs.BootstrapMethodError
	}
	vm.Invoke(method, nil)
	return nil
}

type IRinvokeinterface struct {
	method uint16
	count  byte // not in use
}

func (*IRinvokeinterface) Op() ops.Op    { return ops.Invokeinterface }
func (*IRinvokeinterface) Operands() int { return 4 }
func (ir *IRinvokeinterface) Parse(operands []byte) {
	ir.method = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
	ir.count = operands[2]
	if operands[3] != 0 {
		panic("ir.invokeinterface: operands [3] must be 0")
	}
}
func (ir *IRinvokeinterface) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.method)
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
	method uint16
}

func (*IRinvokespecial) Op() ops.Op    { return ops.Invokespecial }
func (*IRinvokespecial) Operands() int { return 2 }
func (ir *IRinvokespecial) Parse(operands []byte) {
	ir.method = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRinvokespecial) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.method)
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
	method uint16
}

func (*IRinvokestatic) Op() ops.Op    { return ops.Invokestatic }
func (*IRinvokestatic) Operands() int { return 2 }
func (ir *IRinvokestatic) Parse(operands []byte) {
	ir.method = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRinvokestatic) Execute(vm VM) error {
	method := vm.GetCurrentClass().GetMethod(ir.method)
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
	method uint16
}

func (*IRinvokevirtual) Op() ops.Op    { return ops.Invokevirtual }
func (*IRinvokevirtual) Operands() int { return 2 }
func (ir *IRinvokevirtual) Parse(operands []byte) {
	ir.method = ((uint16)(operands[0]) << 8) | (uint16)(operands[1])
}
func (ir *IRinvokevirtual) Execute(vm VM) error {
	stack := vm.GetStack()
	ref := stack.PopRef()
	method := vm.GetCurrentClass().GetMethod(ir.method)
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
