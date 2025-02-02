package vm

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type NativeMethodCallback = func(ir.VM) error

type Method struct {
	*jcls.Method
	class  *Class
	native NativeMethodCallback
}

var _ ir.Method = (*Method)(nil)

func (m *Method) GetDeclaringClass() ir.Class {
	return m.class
}
