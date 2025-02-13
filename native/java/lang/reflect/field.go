package java_lang_reflect

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/reflect/Field.getTypeAnnotationBytes0()[B", Field_getTypeAnnotationBytes0)
}

// private native byte[] getTypeAnnotationBytes0();
func Field_getTypeAnnotationBytes0(vm ir.VM) error {
	stack := vm.GetStack()
	this := stack.GetVarRef(0)
	_ = this
	annotationBytes := vm.NewArray(desc.DescByteArray, 0)
	stack.PushRef(annotationBytes)
	return nil
}
