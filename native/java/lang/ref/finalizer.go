package java_lang_ref

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/ref/Finalizer.isFinalizationEnabled()Z", Finalizer_isFinalizationEnabled)
	native.RegisterDefaultNative("java/lang/ref/Finalizer.reportComplete(Ljava/lang/Object;)V", Finalizer_reportComplete)
}

// private static native boolean isFinalizationEnabled();
func Finalizer_isFinalizationEnabled(vm ir.VM) error {
	vm.GetStack().Push(1)
	return nil
}

// private static native void reportComplete(Object finalizee);
func Finalizer_reportComplete(vm ir.VM) error {
	return nil
}
