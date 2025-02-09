package java_io

import (
	"errors"

	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/io/FileInputStream.initIDs()V", FileInputStream_initIDs)
}

func FileInputStream_initIDs(vm ir.VM) error {
	return errors.New("TODO")
}
