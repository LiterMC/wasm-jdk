package java_lang

import (
	"runtime"
	"time"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

var threadDesc = &desc.Desc{
	EndType: desc.Class,
	Class:   "java/lang/Thread",
}

func init() {
	registerDefaultNative("java/lang/Thread.registerNatives()V", Thread_registerNatives)
}

// private static native void registerNatives();
func Thread_registerNatives(vm ir.VM) error {
	return nil
}

// static native Object findScopedValueBindings();
func Thread_findScopedValueBindings(vm ir.VM) error {
	panic("TODO: Thread.findScopedValueBindings")
	// return nil
}

// static native Thread currentCarrierThread();
func Thread_currentCarrierThread(vm ir.VM) error {
	vm.GetStack().PushRef(vm.GetCarrierThread())
	return nil
}

// public static native Thread currentThread();
func Thread_currentThread(vm ir.VM) error {
	vm.GetStack().PushRef(vm.GetCurrentThread())
	return nil
}

// native void setCurrentThread(Thread thread);
func Thread_setCurrentThread(vm ir.VM) error {
	thread := vm.GetStack().GetVarRef(0)
	vm.SetCurrentThread(thread)
	return nil
}

// static native Object[] scopedValueCache();
func Thread_scopedValueCache(vm ir.VM) error {
	cache := (*vm.GetCurrentThread().UserData()).(*jvm.ThreadUserData).ScopedValueCache
	vm.GetStack().PushRef(cache)
	return nil
}

// static native void setScopedValueCache(Object[] cache);
func Thread_setScopedValueCache(vm ir.VM) error {
	cache := vm.GetStack().GetVarRef(0)
	(*vm.GetCurrentThread().UserData()).(*jvm.ThreadUserData).ScopedValueCache = cache
	return nil
}

// static native void ensureMaterializedForStackWalk(Object o);
func Thread_ensureMaterializedForStackWalk(vm ir.VM) error {
	// TODO: noop?
	return nil
}

// private static native void yield0();
func Thread_yield0(vm ir.VM) error {
	runtime.Gosched()
	return nil
}

// private static native void sleep0(long nanos) throws InterruptedException;
func Thread_sleep0(vm ir.VM) error {
	nanos := vm.GetStack().GetVarInt64(0)
	time.Sleep(time.Nanosecond * (time.Duration)(nanos))
	return nil
}

// private native void start0();
func Thread_start0(vm ir.VM) error {
	panic("TODO: Thread.start0")
	// return nil
}

// public static native boolean holdsLock(Object obj);
func Thread_holdsLock(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(0)
	if ref.IsLocked(vm) == 0 {
		stack.PushInt32(0)
	} else {
		stack.PushInt32(1)
	}
	return nil
}

// private native Object getStackTrace0();
func Thread_getStackTrace0(vm ir.VM) error {
	panic("TODO: Thread.getStackTrace0")
	// return nil
}

// private static native StackTraceElement[][] dumpThreads(Thread[] threads);
func Thread_dumpThreads(vm ir.VM) error {
	panic("TODO: Thread.dumpThreads")
	// return nil
}

// private static native Thread[] getThreads();
func Thread_getThreads(vm ir.VM) error {
	panic("TODO: Thread.getThreads")
	// threads := vm.GetThreads()
	// ref := vm.NewArray(threadDesc, len(threads))
	// copy(ref.GetArrRef(), threads)
	// vm.GetStack().PushRef(ref)
	// return nil
}

// private native void setPriority0(int newPriority);
func Thread_setPriority0(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	newPriority := stack.GetVarInt32(1)
	(*this.UserData()).(*jvm.ThreadUserData).Priority = newPriority
	return nil
}

// private native void interrupt0();
func Thread_interrupt0(vm ir.VM) error {
	this := vm.GetStack().GetVarRef(0)
	vm.Interrupt(this)
	return nil
}

// private static native void clearInterruptEvent();
func Thread_clearInterruptEvent(vm ir.VM) error {
	vm.ClearInterrupt()
	return nil
}

// private native void setNativeName(String name);
func Thread_setNativeName(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	name := vm.GetString(stack.GetVarRef(1))
	(*this.UserData()).(*jvm.ThreadUserData).Name = name
	return nil
}

// private static native long getNextThreadIdOffset();
func Thread_getNextThreadIdOffset(vm ir.VM) error {
	vm.GetStack().PushInt64(jvm.GetNextThreadIdAddress())
	return nil
}
