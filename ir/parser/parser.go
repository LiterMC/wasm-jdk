package parser

import (
	"io"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRParser interface {
	Op() ops.Op
	Parse(br io.ByteReader) (ir.IR, error)
}
