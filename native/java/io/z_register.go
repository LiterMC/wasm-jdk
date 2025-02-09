package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

func loadNative(vm ir.VM, location string, callback jvm.NativeMethodCallback) {
	native.LoadNative(vm, location, callback)
}

func registerDefaultNative(location string, callback jvm.NativeMethodCallback) {
	native.RegisterDefaultNative(location, callback)
}
