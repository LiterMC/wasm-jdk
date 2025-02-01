// This package declares standard Java class file parser
package parser

import (
	"fmt"
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

var parsers = make(map[ops.Op]IRParser)

func RegisterParser(p IRParser) {
	op := p.Op()
	if _, ok := parsers[op]; ok {
		panic(fmt.Errorf("Parser with opcode %d is already exists", op))
	}
	parsers[op] = p
}

func GetIRParser(b ops.Op) IRParser {
	return parsers[b]
}
