package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/io/UnixFileSystem.initIDs()V", UnixFileSystem_initIDs)
}

func UnixFileSystem_initIDs(vm ir.VM) error {
	return nil
}
