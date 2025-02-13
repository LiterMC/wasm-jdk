package java_lang_ref

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/ref/Reference.getAndClearReferencePendingList()Ljava/lang/ref/Reference;", Reference_getAndClearReferencePendingList)
	native.RegisterDefaultNative("java/lang/ref/Reference.hasReferencePendingList()Z", Reference_hasReferencePendingList)
	native.RegisterDefaultNative("java/lang/ref/Reference.waitForReferencePendingList()V", Reference_waitForReferencePendingList)
	native.RegisterDefaultNative("java/lang/ref/Reference.refersTo0(Ljava/lang/Object;)Z", Reference_refersTo0)
	native.RegisterDefaultNative("java/lang/ref/Reference.clear0()V", Reference_clear0)
}

var (
	referenceCh   = make(chan ir.Ref, 8)
	referenceCond = make(chan struct{}, 1)
)

func QueueReference(ref ir.Ref) {
	referenceCh <- ref
	select {
	case referenceCond <- struct{}{}:
	default:
	}
}

// private static native Reference<?> getAndClearReferencePendingList();
func Reference_getAndClearReferencePendingList(vm ir.VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

// private static native boolean hasReferencePendingList();
func Reference_hasReferencePendingList(vm ir.VM) error {
	stack := vm.GetStack()
	if len(referenceCond) > 0 {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}

// private static native void waitForReferencePendingList();
func Reference_waitForReferencePendingList(vm ir.VM) error {
	if len(referenceCh) > 0 {
		if len(referenceCh) > 0 {
			<-referenceCond
		}
		return nil
	}
	<-referenceCond
	return nil
}

// private native boolean refersTo0(Object o);
func Reference_refersTo0(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	obj := stack.GetVarRef(1)
	userData := *this.UserData()
	if userData != nil && obj != nil && userData == obj {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}

// private native void clear0();
func Reference_clear0(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	*this.UserData() = nil
	return nil
}
