package vm

import (
	"github.com/LiterMC/wasm-jdk/desc"
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

	ByteArrayClass = ByteClass.NewArrayClass(1)
)

func (vm *VM) GetClassFromDesc(dc *desc.Desc) *Class {
	var elem *Class
	switch dc.EndType {
	case desc.Class:
		if cls, err := vm.GetClassLoader().LoadClass(dc.Class); err != nil {
			panic(err)
		} else {
			elem = cls.(*Class)
		}
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
		panic("unexpected EndType")
	}
	if dc.ArrDim == 0 {
		return elem
	}
	return elem.NewArrayClass(dc.ArrDim)
}

func (vm *VM) GetClassRef(cls ir.Class) ir.Ref {
	class := cls.(*Class)
	ref0 := class.classRef.Load()
	if ref0 == nil {
		ref := vm.New(vm.javaLangClass).(*Ref)
		classLoaderPtr := (**Ref)(vm.javaLangClass_classLoader.GetPointer(ref))
		*classLoaderPtr = nil // TODO
		*ref.UserData() = class
		class.classRef.CompareAndSwap(nil, ref)
		ref0 = class.classRef.Load()
	}
	return ref0
}

func (vm *VM) NewLookup() ir.Ref {
	ref := vm.New(vm.javaLangInvokeMethodHandlesLookup)
	lookupClassPtr := (**Ref)(vm.javaLangInvokeMethodHandlesLookup_lookupClass.GetPointer(ref))
	allowedModesPtr := (*int32)(vm.javaLangInvokeMethodHandlesLookup_allowedModes.GetPointer(ref))
	class := vm.GetClassRef(vm.GetCurrentClass())
	*lookupClassPtr = class.(*Ref)
	*allowedModesPtr = -1 // TRUSTED
	return ref
}

func (vm *VM) NewMethodHandle(method *jcls.ConstantMethodHandle) ir.Ref {
	ref := vm.New(vm.javaLangInvokeMethodHandle)
	return ref
}

func (vm *VM) NewMethodType(dc string) ir.Ref {
	md, err := desc.ParseMethodDesc(dc)
	if err != nil {
		panic(err)
	}
	ref := vm.New(vm.javaLangInvokeMethodType)
	rtypePtr := (**Ref)(vm.javaLangInvokeMethodType_rtype.GetPointer(ref))
	ptypesPtr := (**Ref)(vm.javaLangInvokeMethodType_ptypes.GetPointer(ref))

	*rtypePtr = vm.GetClassRef(vm.GetClassFromDesc(md.Output)).(*Ref)
	ptypesRef := vm.NewArray(desc.DescClassArray, (int32)(len(md.Inputs)))
	ptypesArr := ptypesRef.GetArrRef()
	for i, in := range md.Inputs {
		ptypesArr[i] = vm.GetClassRef(vm.GetClassFromDesc(in)).(*Ref)
	}
	*ptypesPtr = ptypesRef.(*Ref)
	return ref
}
