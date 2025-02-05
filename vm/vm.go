package vm

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type VM struct {
	stack      *Stack
	nextPc     *ir.ICNode
	nextNative NativeMethodCallback

	loader  ir.ClassLoader
	creator *VM

	stringPool sync.Map

	preloadClasses
}

type preloadClasses struct {
	javaLangObject    ir.Class
	javaLangThrowable ir.Class

	javaLangString       ir.Class
	javaLangString_value ir.Field

	javaLangClass             ir.Class
	javaLangClass_classLoader ir.Field

	javaLangInvokeMethodHandlesLookup              ir.Class
	javaLangInvokeMethodHandlesLookup_lookupClass  ir.Field
	javaLangInvokeMethodHandlesLookup_allowedModes ir.Field

	javaLangInvokeMethodHandle      ir.Class
	javaLangInvokeMethodType        ir.Class
	javaLangInvokeMethodType_rtype  ir.Field
	javaLangInvokeMethodType_ptypes ir.Field
}

var _ ir.VM = (*VM)(nil)

func NewVM(opts *Options) *VM {
	vm := new(VM)
	vm.loader = opts.Loader
	var err error
	vm.preloadClasses.load(vm.loader)
	if opts.EntryClass != "" {
		var entryClass0 ir.Class
		if entryClass0, err = vm.loader.LoadClass(opts.EntryClass); err != nil {
			panic(err)
		}
		entryClass := entryClass0.(*Class)
		entryClass.InitBeforeUse(vm)
		entryMethod := entryClass.GetMethodByName(opts.EntryMethod).(*Method)
		vm.stack = &Stack{
			class:  entryClass,
			method: entryMethod,
		}
		vm.nextPc = entryMethod.Code.Code
	}
	return vm
}

func (p *preloadClasses) load(loader ir.ClassLoader) {
	var err error
	if p.javaLangObject, err = loader.LoadClass("java/lang/Object"); err != nil {
		panic(err)
	}
	if p.javaLangThrowable, err = loader.LoadClass("java/lang/Throwable"); err != nil {
		panic(err)
	}

	if p.javaLangString, err = loader.LoadClass("java/lang/String"); err != nil {
		panic(err)
	}
	p.javaLangString_value = p.javaLangString.GetFieldByName("value")

	if p.javaLangClass, err = loader.LoadClass("java/lang/Class"); err != nil {
		panic(err)
	}
	p.javaLangClass_classLoader = p.javaLangClass.GetFieldByName("classLoader")

	if p.javaLangInvokeMethodHandlesLookup, err = loader.LoadClass("java/lang/invoke/MethodHandles$Lookup"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodHandlesLookup_lookupClass = p.javaLangInvokeMethodHandlesLookup.GetFieldByName("lookupClass")
	p.javaLangInvokeMethodHandlesLookup_allowedModes = p.javaLangInvokeMethodHandlesLookup.GetFieldByName("allowedModes")

	if p.javaLangInvokeMethodType, err = loader.LoadClass("java/lang/invoke/MethodType"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodType_rtype = p.javaLangClass.GetFieldByName("rtype")
	p.javaLangInvokeMethodType_ptypes = p.javaLangClass.GetFieldByName("ptypes")
}

func (vm *VM) Running() bool {
	return vm.stack != nil
}

func (vm *VM) Step() error {
	var err error
	if vm.nextNative != nil {
		nn := vm.nextNative
		vm.nextNative = nil
		err = nn(vm)
		if err == nil {
			vm.Return()
		}
	} else {
		vm.stack.pc = vm.nextPc
		vm.nextPc = vm.stack.pc.Next
		fmt.Printf("step: %04x: %#v\n", vm.stack.pc.Offset, vm.stack.pc.IC)
		err = vm.stack.pc.IC.Execute(vm)
		if vm.stack != nil {
			fmt.Printf("stack:\n - %d\n - %#v\n - %d\n - %#v\n", vm.stack.vars, vm.stack.varRefs, vm.stack.stack, vm.stack.stackRefs)
		}
	}
	return err
}

func (vm *VM) GetStack() ir.Stack {
	return vm.stack
}

func (vm *VM) Root() *VM {
	v := vm
	for v.creator != nil {
		v = v.creator
	}
	return v
}

func (vm *VM) New(cls ir.Class) ir.Ref {
	class := cls.(*Class)
	class.InitBeforeUse(vm)
	ref := newObjectRef(cls)
	return ref
}

func (vm *VM) NewString(str string) ir.Ref {
	ref := vm.New(vm.javaLangString)
	byteArr := (**Ref)(vm.javaLangString_value.GetPointer(ref))
	arr := newRefArrayWithData(ByteArrayClass, (int32)(len(str)), (unsafe.Pointer)(unsafe.StringData(str)))
	*byteArr = arr
	return ref
}

func (vm *VM) NewArray(dc *desc.Desc, length int32) ir.Ref {
	return newRefArray(vm.GetClassFromDesc(dc), length)
}

func (vm *VM) NewArrayByClass(class ir.Class, length int32) ir.Ref {
	return newRefArray(class.(*Class).NewArrayClass(1), length)
}

func (vm *VM) NewArrayMultiDim(dc *desc.Desc, lengths []int32) ir.Ref {
	return newMultiDimArray(vm.GetClassFromDesc(dc), lengths)
}

func (vm *VM) GetObjectClass() ir.Class {
	return vm.javaLangObject
}

func (vm *VM) GetThrowableClass() ir.Class {
	return vm.javaLangThrowable
}

func (vm *VM) GetStringClass() ir.Class {
	return vm.javaLangString
}

func (vm *VM) GetString(ref ir.Ref) string {
	if !vm.javaLangString.IsInstance(ref) {
		panic("ref is not a java/lang/String")
	}
	byteArr := *(**Ref)(vm.javaLangString_value.GetPointer(ref))
	return unsafe.String((*byte)(byteArr.Data()), byteArr.Len())
}

func (vm *VM) GetStringIntern(ref ir.Ref) ir.Ref {
	root := vm.Root()
	str := vm.GetString(ref)
	strRef, _ := root.stringPool.LoadOrStore(str, ref)
	return strRef.(ir.Ref)
}

func (vm *VM) GetStringInternOrNew(str string) ir.Ref {
	root := vm.Root()
	strRef, _ := root.stringPool.Load(str)
	if strRef == nil {
		strRef, _ = root.stringPool.LoadOrStore(str, vm.NewString(str))
	}
	return strRef.(ir.Ref)
}

func (vm *VM) getConstant(i uint16) jcls.ConstantInfo {
	return vm.stack.class.ConstPool[i-1]
}

func (vm *VM) GetDesc(i uint16) *desc.Desc {
	d, err := vm.getConstant(i).(*jcls.ConstantUtf8).AsDesc()
	if err != nil {
		panic(err)
	}
	return d
}

func (vm *VM) GetMethodDesc(i uint16) *desc.MethodDesc {
	d, err := vm.getConstant(i).(*jcls.ConstantUtf8).AsMethodDesc()
	if err != nil {
		panic(err)
	}
	return d
}

func (vm *VM) GetClassByIndex(i uint16) (ir.Class, error) {
	name := vm.getConstant(i).(*jcls.ConstantClass).Name
	return vm.GetClassLoader().LoadClass(name)
}

func (vm *VM) GetClass(r ir.Ref) ir.Class {
	return r.(*Ref).class
}

func (vm *VM) GetClassLoader() ir.ClassLoader {
	return vm.stack.class.loader
}

func (vm *VM) GetCurrentClass() ir.Class {
	return vm.stack.class
}

func (vm *VM) GetCurrentMethod() ir.Method {
	return vm.stack.method
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

func (vm *VM) Invoke(method ir.Method, this ir.Ref) {
	m := method.(*Method)
	fmt.Println("==> invoking " + m.Location())
	defer fmt.Println("   post invoke " + m.Location())
	if m.AccessFlags.Has(jcls.AccNative) {
		if m.native == nil {
			panic("native method " + m.Location() + " is not loaded")
		}
		vm.nextNative = m.native
	} else {
		vm.nextPc = m.Code.Code
	}
	prev := vm.stack
	vm.stack = &Stack{
		prev:   prev,
		class:  m.class,
		method: m,
	}
	inputs := m.Desc().Inputs
	for i := range len(inputs) {
		j := (uint16)(len(inputs) - i)
		d := inputs[j-1]
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
	vm.stack.SetVarRef(0, this)
}

func (vm *VM) InvokeStatic(method ir.Method) {
	m := method.(*Method)
	fmt.Println("==> invoking static " + m.Location())
	defer fmt.Println("   post invoke static " + m.Location())
	if m.AccessFlags.Has(jcls.AccNative) {
		if m.native == nil {
			panic("native method " + m.Location() + " is not loaded")
		}
		vm.nextNative = m.native
	} else {
		vm.nextPc = m.Code.Code
	}
	prev := vm.stack
	vm.stack = &Stack{
		prev:   prev,
		class:  m.class,
		method: m,
	}
	inputs := m.Desc().Inputs
	for i := range len(inputs) {
		j := (uint16)(len(inputs) - i - 1)
		d := inputs[j]
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

	vm.stack.SetVarRef(0, vm.NewLookup())
	vm.stack.SetVarRef(1, vm.GetStringInternOrNew(info.info.NameAndType.Name))
	vm.stack.SetVarRef(2, vm.NewMethodType(info.info.NameAndType.Desc))
	for i, arg := range info.bootstrap.Args {
		switch arg := arg.(type) {
		case *jcls.ConstantString:
			vm.stack.SetVarRef((uint16)(3+i), vm.GetStringInternOrNew(arg.Utf8))
		case *jcls.ConstantMethodHandle:
			vm.stack.SetVarRef((uint16)(3+i), vm.NewMethodHandle(arg))
		case *jcls.ConstantMethodType:
			vm.stack.SetVarRef((uint16)(3+i), vm.NewMethodType(arg.Desc))
		default:
			panic(fmt.Errorf("vm.invokedynamic: unknown argument type %T", arg))
		}
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

func (vm *VM) Return() {
	returned := vm.stack
	fmt.Println("<== returning " + vm.stack.class.Name() + "." + vm.stack.method.Name() + returned.method.Desc().String())
	vm.stack = vm.stack.prev
	if vm.stack == nil {
		return
	}
	vm.nextPc = vm.stack.pc.Next
	switch returned.method.Desc().Output.Type() {
	case desc.Void:
	case desc.Class, desc.Array:
		vm.stack.PushRef(returned.PopRef())
	case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
		vm.stack.Push(returned.Pop())
	case desc.Long, desc.Double:
		vm.stack.Push64(returned.Pop64())
	default:
		panic("vm: unknown MethodDesc.Output.Type")
	}
}

func (vm *VM) Throw(r ir.Ref) {
	// TODO: try/catch logic
	panic(r)
}

func (vm *VM) Goto(n *ir.ICNode) {
	vm.nextPc = n
}

func (vm *VM) MonitorLock(r ir.Ref) error {
	r.Lock(vm)
	return nil
}

func (vm *VM) MonitorUnlock(r ir.Ref) error {
	_, err := r.Unlock(vm)
	return err
}
