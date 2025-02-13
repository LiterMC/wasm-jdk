package java_lang

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/String.intern()Ljava/lang/String;", String_intern)
}

// public native String intern();
func String_intern(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	ref := vm.GetStringIntern(this)
	stack.PushRef(ref)
	return nil
}
