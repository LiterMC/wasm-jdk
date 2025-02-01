package vm

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type VM struct {
	stack             *Stack
	javaLangObject    *jcls.Class
	javaLangThrowable *jcls.Class
}

var _ ir.VM = (*VM)(nil)

func (vm *VM) GetStack() ir.Stack {
	return vm.stack
}

func (vm *VM) New(desc *desc.Desc) ir.Ref {
	return newObjectRef(desc)
}

func (vm *VM) NewArray(desc *desc.Desc, length int32) ir.Ref {
	return newRefArray(desc, length)
}

func (vm *VM) NewArrayMultiDim(desc *desc.Desc, lengths []int32) ir.Ref {
	return newMultiDimArray(desc, lengths)
}

func (vm *VM) GetObjectClass() ir.Class {
	return vm.javaLangObject
}

func (vm *VM) GetThrowableClass() ir.Class {
	return vm.javaLangThrowable
}

func (vm *VM) getConstant(i uint16) jcls.ConstantInfo {
	return vm.stack.class.ConstPool[i-1]
}

func (vm *VM) GetDesc(i uint16) *desc.Desc {
	d, err := vm.getConstant(i).(*jcls.ConstantUtf8).AsDesc()
	if err != nil {
		panic(err)
	}
	return d
}

func (vm *VM) GetMethodDesc(i uint16) *desc.MethodDesc {
	d, err := vm.getConstant(i).(*jcls.ConstantUtf8).AsMethodDesc()
	if err != nil {
		panic(err)
	}
	return d
}

func (vm *VM) GetClassByIndex(i uint16) (ir.Class, error) {
	name := vm.getConstant(i).(*jcls.ConstantClass).Name
	return vm.LoadClass(name)
}

func (vm *VM) GetClass(r ir.Ref) ir.Class {
	ref := r.(*Ref)
	if ref.desc.EndType != desc.Class {
		return nil
	}
	cls, err := vm.LoadClass(ref.desc.Class)
	if err != nil {
		panic(err)
	}
	return cls
}

// GetArrClass([]ir.Ref) ir.Class

// GetCurrentClass() ir.Class
// GetCurrentMethod() ir.Method
// Invoke(ir.Method, ir.Ref)
// InvokeStatic(ir.Method)
// Return()
// Throw(ir.Ref)
// Goto(*ir.ICNode)

// MonitorLock(ir.Ref) error
// MonitorUnlock(ir.Ref) error
