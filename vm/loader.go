package vm

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

type ClassLoader interface {
	LoadClass(name string) (ir.Class, error)
}
