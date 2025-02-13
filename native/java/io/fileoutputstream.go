package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/io/FileOutputStream.initIDs()V", FileOutputStream_initIDs)
}

func FileOutputStream_initIDs(vm ir.VM) error {
	return nil
}
