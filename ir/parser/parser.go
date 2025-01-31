// This package declares standard Java class file parser
package parser

import (
	"io"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ByteReader interface {
	io.ByteReader
	io.Reader
}

type IRParser interface {
	Op() ops.Op
	Parse(br ByteReader) (ir.IR, error)
}
