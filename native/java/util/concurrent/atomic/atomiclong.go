package java_util_concurrent_atomic

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/util/concurrent/atomic/AtomicLong.VMSupportsCS8()Z", AtomicLong_VMSupportsCS8)
}

// private static native boolean VMSupportsCS8();
func AtomicLong_VMSupportsCS8(vm ir.VM) error {
	vm.GetStack().Push(1)
	return nil
}
