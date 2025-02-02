package vm

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type VM struct {
	stack  *Stack
	nextPc *ir.ICNode

	loader ClassLoader

	javaLangObject    ir.Class
	javaLangThrowable ir.Class
}

var _ ir.VM = (*VM)(nil)

var mainMethodDesc = &desc.MethodDesc{
	Inputs: []*desc.Desc{
		&desc.Desc{
			ArrDim:  1,
			EndType: desc.Class,
			Class:   "java/lang/String",
		},
	},
	Output: &desc.Desc{EndType: desc.Void},
}

func NewVM(opts *Options) *VM {
	vm := new(VM)
	vm.loader = opts.Loader
	var err error
	if vm.javaLangObject, err = vm.loader.LoadClass("java/lang/Object"); err != nil {
		panic(err)
	}
	if vm.javaLangThrowable, err = vm.loader.LoadClass("java/lang/Throwable"); err != nil {
		panic(err)
	}
	var entryClass0 ir.Class
	if entryClass0, err = vm.loader.LoadClass(opts.EntryClass); err != nil {
		panic(err)
	}
	entryClass := entryClass0.(*Class)
	entryMethod := entryClass.GetMethodByName(opts.EntryMethod, mainMethodDesc).(*Method)
	vm.stack = &Stack{
		class:  entryClass,
		method: entryMethod,
	}
	return vm
}

func (vm *VM) Step() error {
	vm.nextPc = vm.stack.pc.Next
	err := vm.stack.pc.IC.Execute(vm)
	vm.stack.pc = vm.nextPc
	return err
}

func (vm *VM) GetStack() ir.Stack {
	return vm.stack
}

func (vm *VM) New(cls ir.Class) ir.Ref {
	ref := newObjectRef(cls)
	return ref
}

func (vm *VM) NewArray(dc *desc.Desc, length int32) ir.Ref {
	return newRefArray(dc, length)
}

func (vm *VM) NewArrayMultiDim(dc *desc.Desc, lengths []int32) ir.Ref {
	return newMultiDimArray(dc, lengths)
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
	return vm.loader.LoadClass(name)
}

func (vm *VM) GetClass(r ir.Ref) ir.Class {
	return r.(*Ref).class
}

func (vm *VM) GetCurrentClass() ir.Class {
	return vm.stack.class
}

func (vm *VM) GetCurrentMethod() ir.Method {
	return vm.stack.method
}

func (vm *VM) Invoke(m ir.Method, this ir.Ref) {
	switch m := m.(type) {
	case *Method:
		vm.nextPc = m.Code.Code
		prev := vm.stack
		vm.stack = &Stack{
			prev:   prev,
			class:  m.class,
			method: m,
		}
		for i := range len(m.Desc.Inputs) {
			j := (uint16)(len(m.Desc.Inputs) - i)
			d := m.Desc.Inputs[j-1]
			switch d.Type() {
			case desc.Void:
			case desc.Class, desc.Array:
				vm.stack.SetVarRef(j, prev.PopRef())
			case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
				vm.stack.SetVar(j, prev.Pop())
			case desc.Long, desc.Double:
				vm.stack.SetVar64(j, prev.Pop64())
			default:
				panic("vm: unknown MethodDesc.Input.Type")
			}
		}
		vm.stack.SetVarRef(0, this)
	default:
		panic("vm: unknown method type")
	}
}

func (vm *VM) InvokeStatic(m ir.Method) {
	switch m := m.(type) {
	case *Method:
		vm.nextPc = m.Code.Code
		prev := vm.stack
		vm.stack = &Stack{
			prev:   prev,
			class:  m.class,
			method: m,
		}
		for i := range len(m.Desc.Inputs) {
			j := (uint16)(len(m.Desc.Inputs) - i - 1)
			d := m.Desc.Inputs[j]
			switch d.Type() {
			case desc.Void:
			case desc.Class, desc.Array:
				vm.stack.SetVarRef(j, prev.PopRef())
			case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
				vm.stack.SetVar(j, prev.Pop())
			case desc.Long, desc.Double:
				vm.stack.SetVar64(j, prev.Pop64())
			default:
				panic("vm: unknown MethodDesc.Input.Type")
			}
		}
	default:
		panic("vm: unknown method type")
	}
}

func (vm *VM) Return() {
	returned := vm.stack
	vm.stack = vm.stack.prev
	if vm.stack != nil {
		vm.nextPc = vm.stack.pc.Next
		switch returned.method.Desc.Output.Type() {
		case desc.Void:
		case desc.Class, desc.Array:
			vm.stack.PushRef(returned.PopRef())
		case desc.Boolean, desc.Byte, desc.Char, desc.Short, desc.Int, desc.Float:
			vm.stack.Push(returned.Pop())
		case desc.Long, desc.Double:
			vm.stack.Push64(returned.Pop64())
		default:
			panic("vm: unknown MethodDesc.Output.Type")
		}
	}
}

func (vm *VM) Throw(r ir.Ref) {
	// TODO: try/catch logic
	panic(r)
}

func (vm *VM) Goto(n *ir.ICNode) {
	vm.nextPc = n
}

func (vm *VM) MonitorLock(r ir.Ref) error {
	r.Lock(vm)
	return nil
}

func (vm *VM) MonitorUnlock(r ir.Ref) error {
	_, err := r.Unlock(vm)
	return err
}
