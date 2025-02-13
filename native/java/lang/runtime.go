package java_lang

import (
	"runtime"
	"runtime/debug"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/Runtime.availableProcessors()I", Runtime_availableProcessors)
	native.RegisterDefaultNative("java/lang/Runtime.freeMemory()J", Runtime_freeMemory)
	native.RegisterDefaultNative("java/lang/Runtime.totalMemory()J", Runtime_totalMemory)
	native.RegisterDefaultNative("java/lang/Runtime.maxMemory()J", Runtime_maxMemory)
	native.RegisterDefaultNative("java/lang/Runtime.gc()V", Runtime_gc)
}

// public native int availableProcessors();
func Runtime_availableProcessors(vm ir.VM) error {
	vm.GetStack().PushInt32((int32)(runtime.GOMAXPROCS(-1)))
	return nil
}

// public native long freeMemory();
func Runtime_freeMemory(vm ir.VM) error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	vm.GetStack().PushInt64((int64)(memStats.HeapSys - memStats.HeapAlloc))
	return nil
}

// public native long totalMemory();
func Runtime_totalMemory(vm ir.VM) error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	vm.GetStack().PushInt64((int64)(memStats.HeapSys))
	return nil
}

// public native long maxMemory();
func Runtime_maxMemory(vm ir.VM) error {
	vm.GetStack().PushInt64(debug.SetMemoryLimit(-1))
	return nil
}

// public native void gc();
func Runtime_gc(vm ir.VM) error {
	runtime.GC()
	return nil
}
