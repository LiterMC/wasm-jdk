// This package declares standard Java class file parser
package parser

import (
	"bytes"
	"fmt"
	"io"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ICParser interface {
	Op() ops.Op
	Parse(br *bytes.Reader) (ir.IC, error)
}

var parsers = make(map[ops.Op]ICParser)

func RegisterParser(p ICParser) {
	op := p.Op()
	if _, ok := parsers[op]; ok {
		panic(fmt.Errorf("Parser with opcode 0x%02x is already exists", op))
	}
	parsers[op] = p
}

func GetICParser(b ops.Op) ICParser {
	return parsers[b]
}

func ParseCode(buf []byte) (*ir.ICNode, error) {
	r := bytes.NewReader(buf)
	entry := new(ir.ICNode)

	type jumpRequest struct {
		i    int
		node *ir.ICNode
	}
	var (
		b     byte
		err   error
		node  = entry
		jumps = make(map[int32][]jumpRequest)
	)
	for {
		if reqs, ok := jumps[node.Offset]; ok {
			for _, req := range reqs {
				req.node.IC.(ir.ICJumpable).SetNode(req.i, node)
			}
			delete(jumps, node.Offset)
		}
		if offset, err := r.Seek(0, io.SeekCurrent); err != nil {
			return nil, err
		} else {
			node.Offset = (int32)(offset)
		}

		if b, err = r.ReadByte(); err != nil {
			return nil, err
		}
		parser := GetICParser((ops.Op)(b))
		if parser == nil {
			return nil, fmt.Errorf("parser: unknown opcode 0x%02x at 0x%04x", b, node.Offset)
		}

		if node.IC, err = parser.Parse(r); err != nil {
			return nil, err
		}

		if j, ok := node.IC.(ir.ICJumpable); ok {
			offsets := j.Offsets()
			for i, off := range offsets {
				if off == 0 {
					j.SetNode(i, node)
				} else {
					offs := node.Offset + off
					jumps[offs] = append(jumps[offs], jumpRequest{i, node})
				}
			}
		}

		if r.Len() == 0 {
			break
		}
		n := new(ir.ICNode)
		node.Next = n
		node = n
	}
	for node = entry; node != nil; node = node.Next {
		for _, req := range jumps[node.Offset] {
			req.node.IC.(ir.ICJumpable).SetNode(req.i, node)
		}
	}
	return entry, nil
}
