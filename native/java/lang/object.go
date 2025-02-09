package java_lang

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/lang/Object.getClass()Ljava/lang/Class;", Object_getClass)
	registerDefaultNative("java/lang/Object.hashCode()I", Object_hashCode)
	registerDefaultNative("java/lang/Object.clone()Ljava/lang/Object;", Object_clone)
	registerDefaultNative("java/lang/Object.notify()V", Object_notify)
	registerDefaultNative("java/lang/Object.notifyAll()V", Object_notifyAll)
	registerDefaultNative("java/lang/Object.wait0(J)V", Object_wait0)
}

// public final native Class<?> getClass();
func Object_getClass(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	stack.PushRef(vm.GetClassRef(this.Class()))
	return nil
}

// public native int hashCode();
func Object_hashCode(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	stack.PushInt32(this.Id())
	return nil
}

// protected native Object clone() throws CloneNotSupportedException;
func Object_clone(vm ir.VM) error {
	panic("CloneNotSupportedException")
	// return nil
}

// public final native void notify();
func Object_notify(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	return this.Notify(vm)
}

// public final native void notifyAll();
func Object_notifyAll(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	return this.NotifyAll(vm)
}

// private final native void wait0(long timeoutMillis) throws InterruptedException;
func Object_wait0(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	millis := stack.GetVarInt64(1)
	return this.Wait(vm, millis)
}
