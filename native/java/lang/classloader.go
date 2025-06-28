package java_lang

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/ClassLoader.registerNatives()V", ClassLoader_registerNatives)
}

// private static native void registerNatives();
func ClassLoader_registerNatives(vm ir.VM) error {
	native.LoadNative(vm, "java/lang/ClassLoader.defineClass1(Ljava/lang/ClassLoader;Ljava/lang/String;[BIILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;", ClassLoader_defineClass1)
	native.LoadNative(vm, "java/lang/ClassLoader.defineClass2(Ljava/lang/ClassLoader;Ljava/lang/String;Ljava/nio/ByteBuffer;IILjava/security/ProtectionDomain;Ljava/lang/String;)Ljava/lang/Class;", ClassLoader_defineClass2)
	native.LoadNative(vm, "java/lang/ClassLoader.defineClass0(Ljava/lang/ClassLoader;Ljava/lang/Class;Ljava/lang/String;[BIILjava/security/ProtectionDomain;ZILjava/lang/Object;)Ljava/lang/Class;", ClassLoader_defineClass0)
	native.LoadNative(vm, "java/lang/ClassLoader.findBootstrapClass(Ljava/lang/String;)Ljava/lang/Class;", ClassLoader_findBootstrapClass)
	native.LoadNative(vm, "java/lang/ClassLoader.findLoadedClass0(Ljava/lang/String;)Ljava/lang/Class;", ClassLoader_findLoadedClass0)
	return nil
}

// static native Class<?> defineClass1(ClassLoader loader, String name, byte[] b, int off, int len, ProtectionDomain pd, String source);
func ClassLoader_defineClass1(vm ir.VM) error {
	panic("TODO")
}

// static native Class<?> defineClass2(ClassLoader loader, String name, java.nio.ByteBuffer b, int off, int len, ProtectionDomain pd, String source);
func ClassLoader_defineClass2(vm ir.VM) error {
	panic("TODO")
}

// /**
//   - Defines a class of the given flags via Lookup.defineClass.
//     *
//   - @param loader the defining loader
//   - @param lookup nest host of the Class to be defined
//   - @param name the binary name or {@code null} if not findable
//   - @param b class bytes
//   - @param off the start offset in {@code b} of the class bytes
//   - @param len the length of the class bytes
//   - @param pd protection domain
//   - @param initialize initialize the class
//   - @param flags flags
//   - @param classData class data
//     */
//
// static native Class<?> defineClass0(
//
//	ClassLoader loader, Class<?> lookup, String name, byte[] b, int off, int len,
//	ProtectionDomain pd, boolean initialize, int flags, Object classData);
func ClassLoader_defineClass0(vm ir.VM) error {
	panic("TODO")
}

// private static native Class<?> findBootstrapClass(String name);
func ClassLoader_findBootstrapClass(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	class, err := vm.GetBootLoader().LoadClass(name)
	if err != nil {
		stack.PushRef(nil)
		return nil
	}
	stack.PushRef(class.AsRef(vm))
	return nil
}

// private final native Class<?> findLoadedClass0(String name);
func ClassLoader_findLoadedClass0(vm ir.VM) error {
	stack := vm.GetStack()
	loader := (*stack.GetVarRef(0).UserData()).(ir.ClassLoader)
	name := vm.GetString(stack.GetVarRef(1))
	class := loader.LoadedClass(name)
	if class == nil {
		stack.PushRef(nil)
		return nil
	}
	stack.PushRef(class.AsRef(vm))
	return nil
}

// private static native AssertionStatusDirectives retrieveDirectives();
