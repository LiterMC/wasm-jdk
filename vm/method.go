package vm

import (
	"fmt"
	"slices"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type NativeMethodCallback = func(ir.VM) error

type Method struct {
	*jcls.Method
	class  *Class
	native NativeMethodCallback
}

var _ ir.Method = (*Method)(nil)

func (m *Method) GetDeclaringClass() ir.Class {
	return m.class
}

func (m *Method) Location() string {
	return m.class.Name() + "." + m.Name() + m.Desc().String()
}

func (vm *VM) LoadNativeMethod(method ir.Method, native NativeMethodCallback) {
	m := method.(*Method)
	if !m.AccessFlags.Has(jcls.AccNative) {
		panic("method " + m.Location() + " is not native")
	}
	if m.native != nil {
		panic("method " + m.Location() + " is already loaded")
	}
	m.native = native
}

func (vm *VM) Invoke(method ir.Method) {
	m := method.(*Method)
	fmt.Println("\n==> invoking", m.Location())
	defer fmt.Println("   post invoke", m.Location())
	prev := vm.stack
	prev.pc = vm.nextPc
	isNative := m.AccessFlags.Has(jcls.AccNative)
	if isNative {
		if m.native == nil {
			panic("native method " + m.Location() + " is not loaded")
		}
		vm.nextNative = m.native
	} else {
		vm.nextPc = m.Code.Code
	}
	vm.stack = &Stack{
		prev:   prev,
		class:  m.class,
		method: m,
	}
	inputs := m.Desc().Inputs
	j := m.Desc().InputSlots() + 1
	for _, d := range slices.Backward(inputs) {
		j -= d.Type().Slot()
		switch d.Type() {
		case desc.Void:
		case desc.Class, desc.Array:
			vm.stack.SetVarRef(j, prev.PopRef())
		case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
			vm.stack.SetVar(j, prev.Pop())
		case desc.Long, desc.Double:
			vm.stack.SetVar64(j, prev.Pop64())
		default:
			panic("vm: unknown MethodDesc.Input.Type")
		}
	}
	this := prev.PopRef()
	vm.stack.SetVarRef(0, this)
}

func (vm *VM) InvokeStatic(method ir.Method) {
	m := method.(*Method)
	fmt.Println("\n==> invoking static " + m.Location())
	defer fmt.Println("   post invoke static " + m.Location())
	prev := vm.stack
	prev.pc = vm.nextPc
	if m.AccessFlags.Has(jcls.AccNative) {
		if m.native == nil {
			panic("native method " + m.Location() + " is not loaded")
		}
		vm.nextNative = m.native
	} else {
		vm.nextPc = m.Code.Code
	}
	vm.stack = &Stack{
		prev:   prev,
		class:  m.class,
		method: m,
	}
	inputs := m.Desc().Inputs
	j := m.Desc().InputSlots()
	for _, d := range slices.Backward(inputs) {
		j -= d.Type().Slot()
		switch d.Type() {
		case desc.Void:
		case desc.Class, desc.Array:
			vm.stack.SetVarRef(j, prev.PopRef())
		case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
			vm.stack.SetVar(j, prev.Pop())
		case desc.Long, desc.Double:
			vm.stack.SetVar64(j, prev.Pop64())
		default:
			panic("vm.invokestatic: unknown MethodDesc.Input.Type")
		}
	}
}

func (vm *VM) InvokeVirtual(method ir.Method) {
	m := method.(*Method)
	fmt.Println("\n==> invoking virtual " + m.Location())
	defer fmt.Println("   post invoke virtual " + m.Location())
	prev := vm.stack
	prev.pc = vm.nextPc
	newStack := &Stack{
		prev: prev,
	}
	inputs := m.Desc().Inputs
	j := m.Desc().InputSlots() + 1
	for _, d := range slices.Backward(inputs) {
		j -= d.Type().Slot()
		switch d.Type() {
		case desc.Void:
		case desc.Class, desc.Array:
			newStack.SetVarRef(j, prev.PopRef())
		case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
			newStack.SetVar(j, prev.Pop())
		case desc.Long, desc.Double:
			newStack.SetVar64(j, prev.Pop64())
		default:
			panic("vm: unknown MethodDesc.Input.Type")
		}
	}
	this := prev.PopRef().(*Ref)
	newStack.SetVarRef(0, this)
	m2 := this.class.GetMethodByDesc(method.Name(), method.Desc()).(*Method)
	newStack.class = m2.class
	newStack.method = m2
	vm.stack = newStack
	if m2.AccessFlags.Has(jcls.AccNative) {
		if m2.native == nil {
			panic("native method " + m2.Location() + " is not loaded")
		}
		vm.nextNative = m2.native
	} else {
		vm.nextPc = m2.Code.Code
	}
}

func (vm *VM) InvokeDynamic(ind uint16) error {
	info := vm.stack.class.loadedDynamics[ind]
	bootMe := info.bootstrap.Method
	if bootMe.Kind != jcls.RefInvokeStatic {
		panic("TODO: bootstrap " + bootMe.Kind.String())
	}
	bootCls, err := vm.GetClassLoader().LoadClass(bootMe.Ref.Class.Name)
	if err != nil {
		return err
	}
	bootMethod0 := bootCls.GetMethodByNameAndType(bootMe.Ref.NameAndType.Name, bootMe.Ref.NameAndType.Desc)
	if bootMethod0 == nil {
		panic("bootstrap method " + bootMe.String() + " is nil")
	}
	fmt.Println("==> invoking dynamic " + bootMe.String())
	defer fmt.Println("   post invoke dynamic " + bootMe.String())
	bootMethod := bootMethod0.(*Method)
	prev := vm.stack
	vm.stack = &Stack{
		prev:   prev,
		class:  bootCls.(*Class),
		method: bootMethod,
	}
	vm.nextPc = bootMethod.Code.Code

	hasVarargs := bootMethod.AccessFlags.Has(jcls.AccVarargs)
	inputs := bootMethod.Desc().Inputs
	inputsLen := len(inputs)
	var varargs []unsafe.Pointer

	vm.stack.SetVarRef(0, vm.NewLookup())
	vm.stack.SetVarRef(1, vm.GetStringInternOrNew(info.info.NameAndType.Name))
	vm.stack.SetVarRef(2, vm.NewMethodType(info.info.NameAndType.Desc))
	for i, arg := range info.bootstrap.Args {
		var ref ir.Ref
		switch arg := arg.(type) {
		case *jcls.ConstantString:
			ref = vm.GetStringInternOrNew(arg.Utf8)
		case *jcls.ConstantMethodHandle:
			ref = vm.NewMethodHandle(arg)
		case *jcls.ConstantMethodType:
			ref = vm.NewMethodType(arg.Desc)
		default:
			panic(fmt.Errorf("vm.invokedynamic: unknown argument type %T", arg))
		}
		if hasVarargs && i+4 >= inputsLen {
			varargs = append(varargs, vm.RefToPtr(ref))
		} else {
			vm.stack.SetVarRef((uint16)(3+i), ref)
		}
	}
	if hasVarargs {
		varargsRef := vm.NewArray(inputs[inputsLen-1], (int32)(len(varargs)))
		copy(varargsRef.GetRefArr(), varargs)
		vm.stack.SetVarRef((uint16)(inputsLen-1), varargsRef)
	}

	return vm.RunStack()
}

// RunStack steps the VM until the current stack pops
func (vm *VM) RunStack() error {
	prev := vm.stack.prev
	for vm.stack != prev {
		if err := vm.Step(); err != nil {
			return err
		}
	}
	return nil
}
