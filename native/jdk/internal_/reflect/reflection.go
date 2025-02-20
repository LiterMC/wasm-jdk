package jdk_internal_reflect

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/reflect/Reflection.getCallerClass()Ljava/lang/Class;", Reflection_getCallerClass)
	native.RegisterDefaultNative("jdk/internal/reflect/Reflection.getClassAccessFlags(Ljava/lang/Class;)I", Reflection_getClassAccessFlags)
}

// public static native Class<?> getCallerClass();
func Reflection_getCallerClass(vm ir.VM) error {
	stack := vm.GetStack()
	for st := stack.Prev(); st != nil; st = st.Prev() {
		method := st.Method()
		class := method.GetDeclaringClass()
		if class.Name() != "java/lang/reflect/Method" || method.Name() != "invoke" {
			stack.PushRef(class.AsRef(vm))
			return nil
		}
	}
	stack.PushRef(stack.Method().GetDeclaringClass().AsRef(vm))
	return nil
}

// public static native int getClassAccessFlags(Class<?> c);
func Reflection_getClassAccessFlags(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(0).UserData()).(ir.Class)
	stack.PushInt32(class.Modifiers())
	return nil
}

// public static native boolean areNestMates(Class<?> currentClass, Class<?> memberClass);
func Reflection_areNestMates(vm ir.VM) error {
	stack := vm.GetStack()
	// TODO
	if true {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}
