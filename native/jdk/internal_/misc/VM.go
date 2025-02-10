package jdk_internal_misc

import (
	"os"
	"time"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("jdk/internal/misc/VM.latestUserDefinedLoader0()V", VM_latestUserDefinedLoader0)
	registerDefaultNative("jdk/internal/misc/VM.getuid()J", VM_getuid)
	registerDefaultNative("jdk/internal/misc/VM.geteuid()J", VM_geteuid)
	registerDefaultNative("jdk/internal/misc/VM.getgid()J", VM_getgid)
	registerDefaultNative("jdk/internal/misc/VM.getegid()J", VM_getegid)
	registerDefaultNative("jdk/internal/misc/VM.getNanoTimeAdjustment(J)J", VM_getNanoTimeAdjustment)
	registerDefaultNative("jdk/internal/misc/VM.getRuntimeArguments()Ljava/lang/String;", VM_getRuntimeArguments)
	registerDefaultNative("jdk/internal/misc/VM.initialize()V", VM_initialize)
}

// private static native ClassLoader latestUserDefinedLoader0();
func VM_latestUserDefinedLoader0(vm ir.VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

// public static native long getuid();
func VM_getuid(vm ir.VM) error {
	vm.GetStack().PushInt64(-1)
	return nil
}

// public static native long geteuid();
func VM_geteuid(vm ir.VM) error {
	vm.GetStack().PushInt64(-1)
	return nil
}

// public static native long getgid();
func VM_getgid(vm ir.VM) error {
	vm.GetStack().PushInt64(-1)
	return nil
}

// public static native long getegid();
func VM_getegid(vm ir.VM) error {
	vm.GetStack().PushInt64(-1)
	return nil
}

// public static native long getNanoTimeAdjustment(long offsetInSeconds);
func VM_getNanoTimeAdjustment(vm ir.VM) error {
	stack := vm.GetStack()
	offsetInSeconds := stack.GetVarInt64(0)
	ns := time.Since(time.Unix(offsetInSeconds, 0)).Nanoseconds()
	stack.PushInt64(ns)
	return nil
}

// public static native String[] getRuntimeArguments();
func VM_getRuntimeArguments(vm ir.VM) error {
	argumentsRef := vm.NewArray(desc.DescStringArray, (int32)(len(os.Args)-1))
	arguments := argumentsRef.GetArrRef()
	for i, a := range os.Args[1:] {
		arguments[i] = vm.RefToPtr(vm.GetStringInternOrNew(a))
	}
	vm.GetStack().PushRef(argumentsRef)
	return nil
}

// private static native void initialize();
func VM_initialize(vm ir.VM) error {
	// Used to register the performance-critical methods, e.g. getNanoTimeAdjustment
	return nil
}
