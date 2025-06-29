package java_lang

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
	"github.com/LiterMC/wasm-jdk/native/helper"
)

func init() {
	native.RegisterDefaultNative("java/lang/Object.getClass()Ljava/lang/Class;", Object_getClass)
	native.RegisterDefaultNative("java/lang/Object.hashCode()I", Object_hashCode)
	native.RegisterDefaultNative("java/lang/Object.clone()Ljava/lang/Object;", Object_clone)
	native.RegisterDefaultNative("java/lang/Object.notify()V", Object_notify)
	native.RegisterDefaultNative("java/lang/Object.notifyAll()V", Object_notifyAll)
	native.RegisterDefaultNative("java/lang/Object.wait0(J)V", Object_wait0)
}

// public final native Class<?> getClass();
func Object_getClass(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	stack.PushRef(this.Class().AsRef(vm))
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
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	class := this.Class()
	if class.ArrayDim() <= 0 && !vm.(helper.VMHelper).JClass_javaLangCloneable().IsAssignableFrom(class) {
		return errs.CloneNotSupportedException
	}
	cloned := this.Clone(vm)
	stack.PushRef(cloned)
	return nil
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
