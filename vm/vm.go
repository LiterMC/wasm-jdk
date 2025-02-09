package vm

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"
	"runtime"

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
	javaLangObject          *Class
	javaLangObject_toString ir.Method

	javaLangThrowable               *Class
	javaLangThrowable_backtrace     ir.Field
	javaLangThrowable_detailMessage ir.Field

	javaLangString       *Class
	javaLangString_value ir.Field

	javaLangClass             *Class
	javaLangClass_classLoader ir.Field

	javaLangThread             *Class
	javaLangThread_interrupted ir.Field

	javaLangSystem            *Class
	javaLangSystem_initPhase1 ir.Method
	javaLangSystem_initPhase2 ir.Method
	javaLangSystem_initPhase3 ir.Method

	javaLangInvokeMethodHandlesLookup              *Class
	javaLangInvokeMethodHandlesLookup_lookupClass  ir.Field
	javaLangInvokeMethodHandlesLookup_allowedModes ir.Field

	javaLangInvokeMethodHandle      *Class
	javaLangInvokeMethodType        *Class
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
	vm.preloadClasses.load(vm)
	return vm
}

func (vm *VM) loadClass(name string) (*Class, error) {
	class, err := vm.loader.LoadClass(name)
	if err != nil {
		return nil, err
	}
	return class.(*Class), nil
}

func (p *preloadClasses) load(vm *VM) {
	var err error
	if p.javaLangObject, err = vm.loadClass("java/lang/Object"); err != nil {
		panic(err)
	}
	p.javaLangObject_toString = p.javaLangObject.GetMethodByNameAndType("toString", "()Ljava/lang/String;")

	if p.javaLangString, err = vm.loadClass("java/lang/String"); err != nil {
		panic(err)
	}
	p.javaLangString_value = p.javaLangString.GetFieldByName("value")

	if p.javaLangSystem, err = vm.loadClass("java/lang/System"); err != nil {
		panic(err)
	}
	p.javaLangSystem_initPhase1 = p.javaLangSystem.GetMethodByNameAndType("initPhase1", "()V")
	p.javaLangSystem_initPhase2 = p.javaLangSystem.GetMethodByNameAndType("initPhase2", "(ZZ)I")
	p.javaLangSystem_initPhase3 = p.javaLangSystem.GetMethodByNameAndType("initPhase3", "()V")

	if p.javaLangClass, err = vm.loadClass("java/lang/Class"); err != nil {
		panic(err)
	}
	p.javaLangClass_classLoader = p.javaLangClass.GetFieldByName("classLoader")

	if p.javaLangThread, err = vm.loadClass("java/lang/Thread"); err != nil {
		panic(err)
	}
	p.javaLangThread_interrupted = p.javaLangThread.GetFieldByName("interrupted")

	if p.javaLangThrowable, err = vm.loadClass("java/lang/Throwable"); err != nil {
		panic(err)
	}
	p.javaLangThrowable_backtrace = p.javaLangThrowable.GetFieldByName("backtrace")
	p.javaLangThrowable_detailMessage = p.javaLangThrowable.GetFieldByName("detailMessage")

	if p.javaLangInvokeMethodHandlesLookup, err = vm.loadClass("java/lang/invoke/MethodHandles$Lookup"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodHandlesLookup_lookupClass = p.javaLangInvokeMethodHandlesLookup.GetFieldByName("lookupClass")
	p.javaLangInvokeMethodHandlesLookup_allowedModes = p.javaLangInvokeMethodHandlesLookup.GetFieldByName("allowedModes")

	if p.javaLangInvokeMethodType, err = vm.loadClass("java/lang/invoke/MethodType"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodType_rtype = p.javaLangClass.GetFieldByName("rtype")
	p.javaLangInvokeMethodType_ptypes = p.javaLangClass.GetFieldByName("ptypes")
}

func (vm *VM) SetupEntryMethod() {
	vm.javaLangString.InitBeforeUse(vm)
	vm.javaLangSystem.InitBeforeUse(vm)
	vm.javaLangClass.InitBeforeUse(vm)
	vm.javaLangThread.InitBeforeUse(vm)

	if vm.opts.EntryClass == "" {
		panic("Entry class is not defined")
	}

	// TODO: initialize thread
	vm.carrierThread = vm.New(vm.javaLangThread).(*Ref)
	vm.carrierThread.userData = &ThreadUserData{
		VM: vm,
	}
	vm.currentThread = vm.carrierThread

	vm.initSystem()

	entryClass, err := vm.loadClass(vm.opts.EntryClass)
	if err != nil {
		panic(err)
	}
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

var step int = 0

func (vm *VM) Step() error {
	if step % 1000 == 0 {
		runtime.GC()
	}
	m, pc := vm.stack.method.(*Method), vm.nextPc
	printStack := func() { // early stage debug only
		fmt.Println("current method:", m.class.Name()+":", m)
		fmt.Println(NewStackInfo(vm.stack, -1).String())
		if m.AccessFlags.Has(jcls.AccNative) {
			return
		}
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
	}
	defer func() {
		if err := recover(); err != nil {
			printStack()
			panic(err)
		}
	}()

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
		step++
		fmt.Printf(" == step: %04x: %06d: %#v --> %#v\n", vm.stack.pc.Offset, step, vm.stack.pc.IC, vm.stack.pc.Next)
		// fmt.Printf("    stack: %p %#v\n", vm.stack, vm.stack)
		err = vm.stack.pc.IC.Execute(vm)
		if vm.stack == nil {
			vm.creator.createdMux.Lock()
			delete(vm.creator.created, vm)
			vm.creator.createdMux.Unlock()
		}
	}
	if err != nil {
		printStack()
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
	return ref
}

func (vm *VM) NewArray(dc *desc.Desc, length int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	class.InitBeforeUse(vm)
	return newRefArray(class, length)
}

func (vm *VM) NewArrayByClass(cls ir.Class, length int32) ir.Ref {
	class := cls.(*Class)
	class.InitBeforeUse(vm)
	return newRefArray(class.NewArrayClass(1), length)
}

func (vm *VM) NewArrayMultiDim(dc *desc.Desc, lengths []int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	class.InitBeforeUse(vm)
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
