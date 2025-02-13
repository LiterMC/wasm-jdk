package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/io/FileInputStream.initIDs()V", FileInputStream_initIDs)
}

func FileInputStream_initIDs(vm ir.VM) error {
	return nil
}
