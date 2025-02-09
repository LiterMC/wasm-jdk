package java_lang

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", Throwable_fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
func Throwable_fillInStackTrace(vm ir.VM) error {
	stack := vm.GetStack()
	throwable := stack.GetVarRef(0)
	vm.FillThrowableStackTrace(throwable)
	stack.PushRef(throwable)
	return nil
}
