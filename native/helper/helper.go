package helper

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

type VMHelper interface {
	JClass_JavaLangCloneable() ir.Class
	JClass_JavaLangReflectMethod() ir.Class
}
