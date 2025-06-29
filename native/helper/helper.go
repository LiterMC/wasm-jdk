package helper

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

type VMHelper interface {
	JClass_javaLangCloneable() ir.Class
	JClass_javaLangReflectMethod() ir.Class
	JField_javaLangClass_classData() ir.Field
	JField_javaLangReflectMethod_clazz() ir.Field
	JField_javaLangReflectMethod_modifiers() ir.Field
}
