package vm

import (
	"fmt"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/errs"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

var (
	VoidClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescVoid,
		},
		arrayDim: -1,
	}
	BooleanClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescBool,
		},
		arrayDim: -1,
	}
	ByteClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescInt8,
		},
		arrayDim: -1,
	}
	CharClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescChar,
		},
		arrayDim: -1,
	}
	ShortClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescInt16,
		},
		arrayDim: -1,
	}
	IntClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescInt32,
		},
		arrayDim: -1,
	}
	FloatClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescFloat32,
		},
		arrayDim: -1,
	}
	LongClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescInt64,
		},
		arrayDim: -1,
	}
	DoubleClass = &Class{
		Class: &jcls.Class{
			ThisDesc: desc.DescFloat64,
		},
		arrayDim: -1,
	}
)

func (vm *VM) GetClassByName(name string) (ir.Class, error) {
	var dc *desc.Desc
	if name[0] == '[' {
		var err error
		if dc, err = desc.ParseDesc(name); err != nil {
			return nil, err
		}
	} else {
		dc = &desc.Desc{
			EndType: desc.Class,
			Class:   name,
		}
	}
	return vm.GetClassFromDesc(dc)
}

func (vm *VM) GetClassFromDesc(dc *desc.Desc) (*Class, error) {
	var elem *Class
	switch dc.EndType {
	case desc.Class:
		if cls, err := vm.GetClassLoader().LoadClass(dc.Class); err != nil {
			return nil, &errs.ClassNotFoundException{Class: dc.Class, Cause: err}
		} else {
			elem = cls.(*Class)
		}
	case desc.Void:
		elem = VoidClass
	case desc.Boolean:
		elem = BooleanClass
	case desc.Byte:
		elem = ByteClass
	case desc.Char:
		elem = CharClass
	case desc.Short:
		elem = ShortClass
	case desc.Int:
		elem = IntClass
	case desc.Float:
		elem = FloatClass
	case desc.Long:
		elem = LongClass
	case desc.Double:
		elem = DoubleClass
	default:
		panic(fmt.Errorf("unexpected EndType: %c", dc.EndType))
	}
	if dc.ArrDim == 0 {
		return elem, nil
	}
	return elem.NewArrayClass(dc.ArrDim), nil
}

func (vm *VM) GetLoadedClassByName(name string) (ir.Class, error) {
	var dc *desc.Desc
	if name[0] == '[' {
		var err error
		if dc, err = desc.ParseDesc(name); err != nil {
			return nil, err
		}
	} else {
		dc = &desc.Desc{
			EndType: desc.Class,
			Class:   name,
		}
	}
	class := vm.GetLoadedClassFromDesc(dc)
	if class == nil {
		return nil, nil
	}
	return class, nil
}

func (vm *VM) GetLoadedClassFromDesc(dc *desc.Desc) *Class {
	var elem *Class
	switch dc.EndType {
	case desc.Class:
		cls := vm.GetClassLoader().LoadedClass(dc.Class)
		if cls == nil {
			return nil
		}
		elem = cls.(*Class)
	case desc.Void:
		elem = VoidClass
	case desc.Boolean:
		elem = BooleanClass
	case desc.Byte:
		elem = ByteClass
	case desc.Char:
		elem = CharClass
	case desc.Short:
		elem = ShortClass
	case desc.Int:
		elem = IntClass
	case desc.Float:
		elem = FloatClass
	case desc.Long:
		elem = LongClass
	case desc.Double:
		elem = DoubleClass
	default:
		panic(fmt.Errorf("unexpected EndType: %c", dc.EndType))
	}
	if dc.ArrDim == 0 {
		return elem
	}
	return elem.NewArrayClass(dc.ArrDim)
}

func (c *Class) GetConstantPool(vm *VM) ir.Ref {
	ref := vm.New(vm.jdkInternalReflectConstantPool).(*Ref)
	*ref.UserData() = &c.ConstPool
	return ref
}

func (c *Class) AsRef(vm0 ir.VM) ir.Ref {
	vm := vm0.(*VM)
	ref0 := c.classRef.Load()
	if ref0 == nil {
		ref := vm.New(vm.javaLangClass).(*Ref)
		classLoaderPtr := (**Ref)(vm.javaLangClass_classLoader.GetPointer(ref))
		componentTypePtr := (**Ref)(vm.javaLangClass_componentType.GetPointer(ref))
		*classLoaderPtr = nil // TODO
		if c.arrayDim > 0 {
			*componentTypePtr = c.elem.AsRef(vm0).(*Ref)
		}
		*ref.UserData() = c
		c.classRef.CompareAndSwap(nil, ref)
		ref0 = c.classRef.Load()
	}
	return ref0
}

func (vm *VM) NewLookup() ir.Ref {
	ref := vm.New(vm.javaLangInvokeMethodHandlesLookup)
	lookupClassPtr := (**Ref)(vm.javaLangInvokeMethodHandlesLookup_lookupClass.GetPointer(ref))
	allowedModesPtr := (*int32)(vm.javaLangInvokeMethodHandlesLookup_allowedModes.GetPointer(ref))
	class := vm.GetCurrentClass().AsRef(vm)
	*lookupClassPtr = class.(*Ref)
	*allowedModesPtr = -1 // TRUSTED
	return ref
}

func (f *Field) AsRef(vm0 ir.VM) ir.Ref {
	vm := vm0.(*VM)
	ref0 := f.fieldRef.Load()
	if ref0 == nil {
		ref := vm.New(vm.javaLangReflectField).(*Ref)
		*ref.UserData() = f
		stack := vm.GetStack()
		stack.PushRef(ref)
		stack.PushRef(f.class.AsRef(vm0))
		stack.PushRef(vm.GetStringInternOrNew(f.Name()))
		stack.PushRef(f.typ.AsRef(vm0))
		stack.PushInt32(f.Modifiers())
		stack.Push(0)
		stack.PushInt32((int32)(f.typ.Desc().Type().Slot()))
		stack.PushRef(vm.NewString(f.typ.Desc().String()))
		stack.PushRef(nil)
		vm.Invoke(vm.javaLangReflectField_init)
		if err := vm.RunStack(); err != nil {
			panic(err)
		}

		f.fieldRef.CompareAndSwap(nil, ref)
		ref0 = f.fieldRef.Load()
	}
	return ref0
}

func (m *Method) AsRef(vm0 ir.VM) ir.Ref {
	vm := vm0.(*VM)
	ref0 := m.methodRef.Load()
	if ref0 == nil {
		dc := m.Desc()
		inClsRef := vm.NewArray(desc.DescClassArray, (int32)(len(dc.Inputs)))
		inClsArr := inClsRef.GetRefArr()
		for i, in := range dc.Inputs {
			inCls, err := vm.GetClassFromDesc(in)
			if err != nil {
				panic(err)
			}
			inClsArr[i] = vm.RefToPtr(inCls.AsRef(vm0))
		}
		outCls, err := vm.GetClassFromDesc(dc.Output)
		if err != nil {
			panic(err)
		}

		isConstructor := m.IsConstructor()

		var ref *Ref
		if isConstructor {
			ref = vm0.New(vm.javaLangReflectConstructor).(*Ref)
		} else {
			ref = vm0.New(vm.javaLangReflectMethod).(*Ref)
		}
		*ref.UserData() = m

		exceptionsRef := vm0.NewObjectArray(vm.javaLangClass, (int32)(len(m.Exceptions)))
		exceptions := exceptionsRef.GetRefArr()
		for i, name := range m.Exceptions {
			cls, err := vm.GetClassLoader().LoadClass(name)
			if err != nil {
				panic(err)
			}
			exceptions[i] = vm.RefToPtr(cls.AsRef(vm))
		}

		stack := vm0.GetStack()
		stack.PushRef(ref)
		if isConstructor {
			stack.PushRef(m.class.AsRef(vm0))
			stack.PushRef(inClsRef)
			stack.PushRef(exceptionsRef)
			stack.PushInt32(m.Modifiers())
			stack.PushInt32(1)
			stack.PushRef(vm0.NewString(dc.String()))
			stack.PushRef(nil)
			stack.PushRef(nil)
			vm.Invoke(vm.javaLangReflectConstructor_init)
		} else {
			stack.PushRef(m.class.AsRef(vm0))
			stack.PushRef(vm.GetStringInternOrNew(m.Name()))
			stack.PushRef(inClsRef)
			stack.PushRef(outCls.AsRef(vm0))
			stack.PushRef(exceptionsRef)
			stack.PushInt32(m.Modifiers())
			stack.PushInt32(1)
			stack.PushRef(vm0.NewString(dc.String()))
			stack.PushRef(nil)
			stack.PushRef(nil)
			stack.PushRef(nil)
			vm.Invoke(vm.javaLangReflectMethod_init)
		}
		if err := vm.RunStack(); err != nil {
			panic(err)
		}

		m.methodRef.CompareAndSwap(nil, ref)
		ref0 = m.methodRef.Load()
	}
	return ref0
}

func (vm *VM) NewMethodHandle(method *jcls.ConstantMethodHandle) ir.Ref {
	class, err := vm.GetClassByName(method.Ref.Class.Name)
	if err != nil {
		panic(err)
	}
	method0 := class.GetMethodByNameAndType(method.Ref.NameAndType.Name, method.Ref.NameAndType.Desc)
	methodRef := method0.AsRef(vm)

	memberNameRef := vm.New(vm.javaLangInvokeMemberName)
	vm.stack.PushRef(memberNameRef)
	vm.stack.PushRef(methodRef)
	vm.Invoke(vm.javaLangInvokeMemberName_init)
	if err := vm.RunStack(); err != nil {
		panic(err)
	}

	vm.stack.PushRef(memberNameRef)
	vm.InvokeStatic(vm.javaLangInvokeDirectMethodHandle_make)
	if err := vm.RunStack(); err != nil {
		panic(err)
	}
	ref := vm.stack.PopRef()

	return ref
}

func (vm *VM) NewMethodType(dc string) ir.Ref {
	md, err := desc.ParseMethodDesc(dc)
	if err != nil {
		panic(err)
	}
	ref := vm.New(vm.javaLangInvokeMethodType)
	rtypePtr := (*unsafe.Pointer)(vm.javaLangInvokeMethodType_rtype.GetPointer(ref))
	ptypesPtr := (*unsafe.Pointer)(vm.javaLangInvokeMethodType_ptypes.GetPointer(ref))

	outCls, err := vm.GetClassFromDesc(md.Output)
	if err != nil {
		panic(err)
	}
	*rtypePtr = vm.RefToPtr(outCls.AsRef(vm))
	ptypesRef := vm.NewArray(desc.DescClassArray, (int32)(len(md.Inputs)))
	ptypesArr := ptypesRef.GetRefArr()
	for i, in := range md.Inputs {
		inCls, err := vm.GetClassFromDesc(in)
		if err != nil {
			panic(err)
		}
		ptypesArr[i] = vm.RefToPtr(inCls.AsRef(vm))
	}
	*ptypesPtr = vm.RefToPtr(ptypesRef)
	return ref
}

func (vm *VM) FillThrowableStackTrace(throwable ir.Ref) {
	st := vm.stack.Prev().Prev()
	backtrace := vm.New(vm.GetObjectClass()).(*Ref)
	*backtrace.UserData() = NewStackInfo(vm, st, -1)
	*(**Ref)(vm.javaLangThrowable_backtrace.GetPointer(throwable)) = backtrace
}
