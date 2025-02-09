package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("jdk/internal/misc/CDS.isDumpingClassList0()Z", CDS_isDumpingClassList0)
	registerDefaultNative("jdk/internal/misc/CDS.isDumpingArchive0()Z", CDS_isDumpingArchive0)
	registerDefaultNative("jdk/internal/misc/CDS.isSharingEnabled0()Z", CDS_isSharingEnabled0)
	registerDefaultNative("jdk/internal/misc/CDS.logLambdaFormInvoker(Ljava/lang/String;)V", CDS_logLambdaFormInvoker)
	registerDefaultNative("jdk/internal/misc/CDS.initializeFromArchive(Ljava/lang/Class;)V", CDS_initializeFromArchive)
	registerDefaultNative("jdk/internal/misc/CDS.defineArchivedModules(Ljava/lang/ClassLoader;Ljava/lang/ClassLoader;)V", CDS_defineArchivedModules)
	registerDefaultNative("jdk/internal/misc/CDS.getRandomSeedForDumping()J", CDS_getRandomSeedForDumping)
	registerDefaultNative("jdk/internal/misc/CDS.dumpClassList(Ljava/lang/String;)V", CDS_dumpClassList)
	registerDefaultNative("jdk/internal/misc/CDS.dumpDynamicArchive(Ljava/lang/String;)V", CDS_dumpDynamicArchive)
}

// private static native boolean isDumpingClassList0();
func CDS_isDumpingClassList0(vm ir.VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}

// private static native boolean isDumpingArchive0();
func CDS_isDumpingArchive0(vm ir.VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}

// private static native boolean isSharingEnabled0();
func CDS_isSharingEnabled0(vm ir.VM) error {
	vm.GetStack().PushInt32(0)
	return nil
}

// private static native void logLambdaFormInvoker(String line);
func CDS_logLambdaFormInvoker(vm ir.VM) error {
	line := vm.GetString(vm.GetStack().GetVarRef(0))
	println(line)
	return nil
}

// public static native void initializeFromArchive(Class<?> c);
func CDS_initializeFromArchive(vm ir.VM) error {
	class := vm.GetStack().GetVarRef(0)
	_ = class
	return nil
}

// public static native void defineArchivedModules(ClassLoader platformLoader, ClassLoader systemLoader);
func CDS_defineArchivedModules(vm ir.VM) error {
	stack := vm.GetStack()
	platformLoader := stack.GetVarRef(0)
	systemLoader := stack.GetVarRef(1)
	_ = platformLoader
	_ = systemLoader
	return nil
}

// public static native long getRandomSeedForDumping();
func CDS_getRandomSeedForDumping(vm ir.VM) error {
	vm.GetStack().PushInt64(1)
	return nil
}

// private static native void dumpClassList(String listFileName);
func CDS_dumpClassList(vm ir.VM) error {
	filename := vm.GetString(vm.GetStack().GetVarRef(0))
	println("TODO: dump CDS.dumpClassList:", filename)
	return nil
}

// private static native void dumpDynamicArchive(String archiveFileName);
func CDS_dumpDynamicArchive(vm ir.VM) error {
	filename := vm.GetString(vm.GetStack().GetVarRef(0))
	println("TODO: dump CDS.dumpDynamicArchive:", filename)
	return nil
}
