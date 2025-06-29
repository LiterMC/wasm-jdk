package vm

import (
	"fmt"
	"strings"
	"sync"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
	"github.com/LiterMC/wasm-jdk/native/helper"
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

	step              uint64
	carrierThread     *Ref
	currentThread     *Ref
	interruptNotifier chan struct{}
	throwing          ir.Ref

	stringPool sync.Map

	*preloadClasses
}

var _ ir.VM = (*VM)(nil)
var _ helper.VMHelper = (*VM)(nil)

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

func (vm *VM) SetupEntryMethod() {
	vm.javaLangString.InitBeforeUse(vm)
	vm.javaLangSystem.InitBeforeUse(vm)
	vm.javaLangClass.InitBeforeUse(vm)
	vm.javaLangThreadGroup.InitBeforeUse(vm)
	vm.javaLangThread.InitBeforeUse(vm)

	vm.initSystemThread()
	vm.initSystem()

	if vm.opts.EntryClass == "" {
		panic("Entry class is not defined")
	}
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

func (vm *VM) NewSubVM(thread0 ir.Ref) ir.VM {
	sub := &VM{
		opts:              vm.opts,
		loader:            vm.loader,
		creator:           vm,
		created:           make(map[*VM]struct{}),
		interruptNotifier: make(chan struct{}, 1),
		preloadClasses:    vm.preloadClasses,
	}
	thread := thread0.(*Ref)
	entryClass := thread.Class().(*Class)
	entryMethod := entryClass.GetMethodByNameAndType("run", "()V").(*Method)
	sub.stack = &Stack{
		class:  entryClass,
		method: entryMethod,
	}
	sub.nextPc = entryMethod.Code.Code
	sub.stack.SetVarRef(0, thread)

	if thread.userData == nil {
		thread.userData = new(ThreadUserData)
	}
	thread.userData.(*ThreadUserData).VM = sub
	sub.carrierThread = thread
	sub.currentThread = thread

	vm.createdMux.Lock()
	vm.created[sub] = struct{}{}
	vm.createdMux.Unlock()
	return sub
}

func (vm *VM) initSystemThread() {
	systemThreadGroup := vm.New(vm.javaLangThreadGroup)
	vm.stack.PushRef(systemThreadGroup)
	vm.Invoke(vm.javaLangThreadGroup.GetMethodByNameAndType("<init>", "()V"))
	if err := vm.RunStack(); err != nil {
		panic(err)
	}

	vm.carrierThread = vm.New(vm.javaLangThread).(*Ref)
	vm.carrierThread.userData = &ThreadUserData{
		VM: vm,
	}
	vm.currentThread = vm.carrierThread
	vm.stack.PushRef(vm.carrierThread)
	vm.stack.PushRef(systemThreadGroup)
	vm.stack.PushRef(vm.GetStringInternOrNew("Main Thread"))
	vm.stack.PushInt32(0)
	vm.stack.PushRef(nil)
	vm.stack.PushInt64(0)
	vm.stack.PushRef(nil)
	vm.Invoke(vm.javaLangThread.GetMethodByNameAndType("<init>", "(Ljava/lang/ThreadGroup;Ljava/lang/String;ILjava/lang/Runnable;JLjava/security/AccessControlContext;)V"))
	if err := vm.RunStack(); err != nil {
		panic(err)
	}
}

func (vm *VM) initSystem() {
	vm.javaLangRefFinalizer.InitBeforeUse(vm)
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

func debugFormatIC(ic ir.IC) string {
	ics := fmt.Sprintf("%#v", ic)
	if strings.HasSuffix(ics, "(nil)") {
		return ic.Op().String()
	}
	i := strings.IndexByte(ics, '{')
	return fmt.Sprintf("%-16s %s", ic.Op(), ics[i:])
}

func (vm *VM) Step() error {
	m, pc := vm.stack.method.(*Method), vm.nextPc
	printStack := func() { // early stage debug only
		fmt.Println("current method:", m.class.Name()+":", m)
		fmt.Println(NewStackInfo(vm, vm.stack, -1).String())
		if m.AccessFlags.Has(jcls.AccNative) {
			return
		}
		for c := m.Code.Code; c != nil; c = c.Next {
			if c == pc {
				fmt.Print("-> ")
			} else {
				fmt.Print("   ")
			}
			fmt.Println(debugFormatIC(c.IC))
		}
		fmt.Println()
	}
	defer func() {
		if vm.Throwing() != nil {
			return
		}
		if err := recover(); err != nil {
			if vm.creator == nil {
				printStack()
			}
			vm.throwing = (*Ref)(nil)
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
		vm.step++
		if vm.creator == nil {
			fmt.Println(vm.stack.GoString())
			fmt.Printf(" == step: %04x: %06d: %s --> %#v\n", vm.stack.pc.Offset, vm.step, debugFormatIC(vm.stack.pc.IC), vm.stack.pc.Next)
		}
		err = vm.stack.pc.IC.Execute(vm)
		if vm.stack == nil && vm.creator != nil {
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

// Alloc an array with the descriptor as the array's type
func (vm *VM) NewArray(dc *desc.Desc, length int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	class.InitBeforeUse(vm)
	return newRefArray(class, length)
}

func (vm *VM) NewArrayMultiDim(dc *desc.Desc, lengths []int32) ir.Ref {
	class, err := vm.GetClassFromDesc(dc)
	if err != nil {
		panic(err)
	}
	class.InitBeforeUse(vm)
	return newMultiDimArray(class, lengths)
}

// Alloc an array with the class as the array's type
func (vm *VM) NewArrayMultiDimWithClass(cls ir.Class, lengths []int32) ir.Ref {
	class := cls.(*Class)
	class.InitBeforeUse(vm)
	return newMultiDimArray(class, lengths)
}

// Alloc an array with the class as the elements' type
func (vm *VM) NewObjectArray(cls ir.Class, length int32) ir.Ref {
	class := cls.(*Class)
	class.InitBeforeUse(vm)
	return newRefArray(class.NewArrayClass(1), length)
}

// Alloc an array with the class as the elements' type
func (vm *VM) NewObjectMultiDimArray(cls ir.Class, lengths []int32) ir.Ref {
	class := cls.(*Class)
	class.InitBeforeUse(vm)
	return newMultiDimArray(class.NewArrayClass(len(lengths)), lengths)
}

func (vm *VM) RefToPtr(ref ir.Ref) unsafe.Pointer {
	if ref == nil {
		return nil
	}
	return (unsafe.Pointer)(ref.(*Ref))
}
func (vm *VM) PtrToRef(ptr unsafe.Pointer) ir.Ref {
	if ptr == nil {
		return nil
	}
	return (*Ref)(ptr)
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
	if ref == nil || ref == (*Ref)(nil) {
		return "<null>"
	}
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
	return vm.GetClassByName(name)
}

func (vm *VM) GetClass(r ir.Ref) ir.Class {
	return r.(*Ref).class
}

func (vm *VM) GetBootLoader() ir.ClassLoader {
	return vm.opts.Loader
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
	if vm.creator == nil {
		fmt.Println("<== returning", returned.class.Name()+"."+returned.method.Name()+returned.method.Desc().String())
		fmt.Println()
	}
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
	vm.throwing = r
	message := vm.GetString(*(**Ref)(vm.javaLangThrowable_detailMessage.GetPointer(r)))
	backtrace := *(**Ref)(vm.javaLangThrowable_backtrace.GetPointer(r))
	fmt.Printf("Throwing: %s: %s\n", r.Class().Name(), message)
	if backtrace != nil {
		stackInfo := backtrace.userData.(*StackInfo)
		fmt.Println(stackInfo.String())
	}
	panic(r)
}

func (vm *VM) Throwing() ir.Ref {
	return vm.throwing
}

func (vm *VM) Goto(n *ir.ICNode) {
	vm.nextPc = n
}
