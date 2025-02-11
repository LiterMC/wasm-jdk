package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("jdk/internal/loader/NativeLibraries.load(Ljdk/internal/loader/NativeLibraries$NativeLibraryImpl;Ljava/lang/String;ZZ)Z", NativeLibraries_load)
	registerDefaultNative("jdk/internal/loader/NativeLibraries.unload(Ljava/lang/String;ZJ)V", NativeLibraries_unload)
	registerDefaultNative("jdk/internal/loader/NativeLibraries.findBuiltinLib(Ljava/lang/String;)Ljava/lang/String;", NativeLibraries_findBuiltinLib)
}

// private static native boolean load(NativeLibraryImpl impl, String name, boolean isBuiltin, boolean throwExceptionIfFail);
func NativeLibraries_load(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(1))
	isBuiltin := stack.GetVar(2) != 0
	throwExceptionIfFail := stack.GetVar(3) != 0
	_ = isBuiltin
	if false && throwExceptionIfFail {
		return &errs.UnsatisfiedLinkError{name}
	}
	stack.Push(0)
	return nil
}

// private static native void unload(String name, boolean isBuiltin, long handle);
func NativeLibraries_unload(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	isBuiltin := stack.GetVar(1) != 0
	handle := stack.GetVarInt64(2)
	_, _, _ = name, isBuiltin, handle
	return nil
}

// private static native String findBuiltinLib(String name);
func NativeLibraries_findBuiltinLib(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	if false && true {
		panic("finding builtin lib: " + name)
	}
	stack.PushRef(vm.NewString(name))
	return nil
}
