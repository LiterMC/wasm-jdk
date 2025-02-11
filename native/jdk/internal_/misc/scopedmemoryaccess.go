package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("jdk/internal/misc/ScopedMemoryAccess.registerNatives()V", ScopedMemoryAccess_registerNatives)
}

func ScopedMemoryAccess_registerNatives(vm ir.VM) error {
	loadNative(vm, "jdk/internal/misc/ScopedMemoryAccess.closeScope0(Ljdk/internal/foreign/MemorySessionImpl;)Z", ScopedMemoryAccess_closeScope0)
	return nil
}

// native boolean closeScope0(MemorySessionImpl session);
func ScopedMemoryAccess_closeScope0(vm ir.VM) error {
	// TODO: not sure if this is necessary
	vm.GetStack().Push(1)
	return nil
}
