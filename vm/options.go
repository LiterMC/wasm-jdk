package vm

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

// VM Options
type Options struct {
	Loader          ir.ClassLoader
	EntryClass      string
	EntryMethod     string
	EntryArgs       []string
}
