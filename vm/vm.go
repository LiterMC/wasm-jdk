package vm

import (
	"fmt"
	"slices"
	"strings"
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

	opts       *Options
	loader     ir.ClassLoader
	creator    *VM
	createdMux sync.RWMutex
	created    map[*VM]struct{}

	carrierThread     *Ref
	currentThread     *Ref
	interruptNotifier chan struct{}

	stringPool sync.Map

	*preloadClasses
}

type preloadClasses struct {
	javaLangObject          ir.Class
	javaLangObject_toString ir.Method

	javaLangThrowable               ir.Class
	javaLangThrowable_backtrace     ir.Field
	javaLangThrowable_detailMessage ir.Field

	javaLangString       ir.Class
	javaLangString_value ir.Field

	javaLangClass             ir.Class
	javaLangClass_classLoader ir.Field

	javaLangThread             ir.Class
	javaLangThread_interrupted ir.Field

	javaLangSystem            ir.Class
	javaLangSystem_initPhase1 ir.Method
	javaLangSystem_initPhase2 ir.Method
	javaLangSystem_initPhase3 ir.Method

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
	vm := &VM{
		opts:              opts,
		loader:            opts.Loader,
		created:           make(map[*VM]struct{}),
		interruptNotifier: make(chan struct{}, 1),
		preloadClasses:    new(preloadClasses),
	}
	vm.stack = &Stack{}
	vm.preloadClasses.load(vm.loader)
	return vm
}

func (p *preloadClasses) load(loader ir.ClassLoader) {
	var err error
	if p.javaLangObject, err = loader.LoadClass("java/lang/Object"); err != nil {
		panic(err)
	}
	p.javaLangObject_toString = p.javaLangObject.GetMethodByNameAndType("toString", "()Ljava/lang/String;")

	if p.javaLangThrowable, err = loader.LoadClass("java/lang/Throwable"); err != nil {
		panic(err)
	}
	p.javaLangThrowable_backtrace = p.javaLangThrowable.GetFieldByName("backtrace")
	p.javaLangThrowable_detailMessage = p.javaLangThrowable.GetFieldByName("detailMessage")

	if p.javaLangString, err = loader.LoadClass("java/lang/String"); err != nil {
		panic(err)
	}
	p.javaLangString_value = p.javaLangString.GetFieldByName("value")

	if p.javaLangClass, err = loader.LoadClass("java/lang/Class"); err != nil {
		panic(err)
	}
	p.javaLangClass_classLoader = p.javaLangClass.GetFieldByName("classLoader")

	if p.javaLangThread, err = loader.LoadClass("java/lang/Thread"); err != nil {
		panic(err)
	}
	p.javaLangThread_interrupted = p.javaLangThread.GetFieldByName("interrupted")

	if p.javaLangSystem, err = loader.LoadClass("java/lang/System"); err != nil {
		panic(err)
	}
	p.javaLangSystem_initPhase1 = p.javaLangSystem.GetMethodByNameAndType("initPhase1", "()V")
	p.javaLangSystem_initPhase2 = p.javaLangSystem.GetMethodByNameAndType("initPhase2", "(ZZ)I")
	p.javaLangSystem_initPhase3 = p.javaLangSystem.GetMethodByNameAndType("initPhase3", "()V")

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

func (vm *VM) SetupEntryMethod() {
	if vm.opts.EntryClass == "" {
		panic("Entry class is not defined")
	}

	// TODO: initialize thread
	vm.javaLangThread.(*Class).InitBeforeUse(vm)
	vm.carrierThread = vm.New(vm.javaLangThread).(*Ref)
	vm.carrierThread.userData = &ThreadUserData{
		VM: vm,
	}
	vm.currentThread = vm.carrierThread

	vm.initSystem()

	entryClass0, err := vm.loader.LoadClass(vm.opts.EntryClass)
	if err != nil {
		panic(err)
	}
	entryClass := entryClass0.(*Class)
	entryClass.InitBeforeUse(vm)
	entryMethod := entryClass.GetMethodByName(vm.opts.EntryMethod).(*Method)
	vm.stack = &Stack{
		class:  entryClass,
		method: entryMethod,
	}
	vm.nextPc = entryMethod.Code.Code
}

func (vm *VM) newSubVM(thread *Ref) *VM {
	sub := &VM{
		opts:              vm.opts,
		loader:            vm.loader,
		created:           make(map[*VM]struct{}),
		interruptNotifier: make(chan struct{}, 1),
		preloadClasses:    vm.preloadClasses,
	}
	entryClass := thread.Class().(*Class)
	entryMethod := entryClass.GetMethodByNameAndType("run", "()V").(*Method)
	sub.stack = &Stack{
		class:  entryClass,
		method: entryMethod,
	}
	sub.nextPc = entryMethod.Code.Code

	thread.userData.(*ThreadUserData).VM = sub
	sub.carrierThread = thread
	sub.currentThread = thread

	vm.creator.createdMux.Lock()
	delete(vm.creator.created, vm)
	vm.creator.createdMux.Unlock()
	return sub
}

func (vm *VM) initSystem() {
	vm.javaLangSystem.(*Class).InitBeforeUse(vm)
	vm.InvokeStatic(vm.javaLangSystem_initPhase1)
	if err := vm.RunStack(); err != nil {
		panic(err)
	}
	vm.stack.Push(1)
	vm.stack.Push(1)
	vm.InvokeStatic(vm.javaLangSystem_initPhase2)
	if err := vm.RunStack(); err != nil {
		panic(err)
	}
	ok := vm.stack.PopInt32() == 0
	if !ok {
		panic("System.initPhase2 failed")
	}
	vm.InvokeStatic(vm.javaLangSystem_initPhase3)
	if err := vm.RunStack(); err != nil {
		panic(err)
	}
}

func (vm *VM) Running() bool {
	return vm.stack != nil
}

func (vm *VM) Step() error {
	defer func(m *Method, pc *ir.ICNode) {
		if err := recover(); err != nil {
			fmt.Println("current method:", m.class.Name()+":", m)
			for c := m.Code.Code; c != nil; c = c.Next {
				ics := fmt.Sprintf("%#v", c.IC)
				if strings.HasSuffix(ics, "(nil)") {
					ics = ""
				} else {
					_, ics, _ = strings.Cut(ics, "{")
					ics = "{" + ics
				}
				if c == pc {
					fmt.Print("-> ")
				} else {
					fmt.Print("   ")
				}
				fmt.Printf("%-16s %s\n", c.IC.Op(), ics)
			}
			fmt.Println()
			panic(err)
		}
	}(vm.stack.method.(*Method), vm.nextPc)

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
		fmt.Printf(" == step: %04x: %#v --> %#v\n", vm.stack.pc.Offset, vm.stack.pc.IC, vm.stack.pc.Next)
		// fmt.Printf("    stack: %p %#v\n", vm.stack, vm.stack)
		err = vm.stack.pc.IC.Execute(vm)
		if vm.stack == nil {
			vm.creator.createdMux.Lock()
			delete(vm.creator.created, vm)
			vm.creator.createdMux.Unlock()
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
	var arr *Ref
	if len(str) > 0 {
		arr = newRefArrayWithData(ByteArrayClass, (int32)(len(str)), (unsafe.Pointer)(unsafe.StringData(str)))
	} else {
		arr = newRefArray(ByteArrayClass, 0)
	}
	*byteArr = arr
	fmt.Printf("--- getting string %q\n", str)
	return ref
}

func (vm *VM) NewArray(dc *desc.Desc, length int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	return newRefArray(class, length)
}

func (vm *VM) NewArrayByClass(class ir.Class, length int32) ir.Ref {
	return newRefArray(class.(*Class).NewArrayClass(1), length)
}

func (vm *VM) NewArrayMultiDim(dc *desc.Desc, lengths []int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	return newMultiDimArray(class, lengths)
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
	fmt.Printf("--- getting string %q\n", str)
	strRef, _ := root.stringPool.LoadOrStore(str, ref)
	return strRef.(ir.Ref)
}

func (vm *VM) GetStringInternOrNew(str string) ir.Ref {
	root := vm.Root()
	fmt.Printf("--- getting string %q\n", str)
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
	return vm.getClassFromDescString(name)
}

func (vm *VM) GetClass(r ir.Ref) ir.Class {
	return r.(*Ref).class
}

func (vm *VM) GetClassLoader() ir.ClassLoader {
	return vm.loader
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

func (vm *VM) Invoke(method ir.Method) {
	m := method.(*Method)
	fmt.Println("==> invoking", m.Location())
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
	fmt.Println("==> invoking static " + m.Location())
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
	fmt.Println("==> invoking virtual " + m.Location())
	defer fmt.Println("   post invoke virtual " + m.Location())
	prev := vm.stack
	prev.pc = vm.nextPc
	vm.stack = &Stack{
		prev: prev,
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
	this := prev.PopRef().(*Ref)
	vm.stack.SetVarRef(0, this)
	m2 := this.class.GetMethodByDesc(method.Name(), method.Desc()).(*Method)
	vm.stack.class = m2.class
	vm.stack.method = m2
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
	var varargs []ir.Ref

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
			varargs = append(varargs, ref)
		} else {
			vm.stack.SetVarRef((uint16)(3+i), ref)
		}
	}
	if hasVarargs {
		varargsRef := vm.NewArray(inputs[inputsLen-1].Elem(), (int32)(len(varargs)))
		copy(varargsRef.GetArrRef(), varargs)
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

func (vm *VM) Return() {
	returned := vm.stack
	fmt.Println("<== returning " + returned.class.Name() + "." + returned.method.Name() + returned.method.Desc().String())
	fmt.Println()
	vm.stack = returned.prev
	if vm.stack == nil {
		return
	}
	vm.nextPc = vm.stack.pc
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
	message := vm.GetString(*(**Ref)(vm.javaLangThrowable_detailMessage.GetPointer(r)))
	backtrace := *(**Ref)(vm.javaLangThrowable_backtrace.GetPointer(r))
	fmt.Printf("Throwing: %s: %s\n", r.Class().Name(), message)
	if backtrace != nil {
		stackInfo := backtrace.userData.(*StackInfo)
		fmt.Println(stackInfo.String())
	}
	panic(r)
}

func (vm *VM) Goto(n *ir.ICNode) {
	vm.nextPc = n
}
