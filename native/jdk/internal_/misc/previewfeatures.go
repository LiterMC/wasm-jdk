package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/misc/PreviewFeatures.isPreviewEnabled()Z", PreviewFeatures_isPreviewEnabled)
}

func PreviewFeatures_isPreviewEnabled(vm ir.VM) error {
	vm.GetStack().Push(0)
	return nil
}
