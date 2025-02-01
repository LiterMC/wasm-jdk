package vm

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

type VM struct {
	stack *Stack
}

var _ ir.VM = (*VM)(nil)

func (vm *VM) GetStack() ir.Stack {
	return vm.stack
}
