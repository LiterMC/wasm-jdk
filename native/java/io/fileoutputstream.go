package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/io/FileOutputStream.initIDs()V", FileOutputStream_initIDs)
}

func FileOutputStream_initIDs(vm ir.VM) error {
	return nil
}
