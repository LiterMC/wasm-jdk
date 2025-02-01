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

type ICParser interface {
	Op() ops.Op
	Parse(br ByteReader) (ir.IC, error)
}

var parsers = make(map[ops.Op]ICParser)

func RegisterParser(p ICParser) {
	op := p.Op()
	if _, ok := parsers[op]; ok {
		panic(fmt.Errorf("Parser with opcode %d is already exists", op))
	}
	parsers[op] = p
}

func GetICParser(b ops.Op) ICParser {
	return parsers[b]
}
