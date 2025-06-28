package vm

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

var (
	BooleanArrayClass = BooleanClass.NewArrayClass(1)
	ByteArrayClass    = ByteClass.NewArrayClass(1)
	CharArrayClass    = CharClass.NewArrayClass(1)
	ShortArrayClass   = ShortClass.NewArrayClass(1)
	IntArrayClass     = IntClass.NewArrayClass(1)
	LongArrayClass    = LongClass.NewArrayClass(1)
	FloatArrayClass   = FloatClass.NewArrayClass(1)
	DoubleArrayClass  = DoubleClass.NewArrayClass(1)
)

func (vm *VM) getArrayMethodByName(ref *jcls.ConstantRef) *Method {
	arrDesc, err := desc.ParseDesc(ref.Class.Name)
	if err != nil {
		panic(err)
	}
	arrCls, err := vm.GetClassFromDesc(arrDesc)
	if err != nil {
		panic(err)
	}
	return getArrayMethod(arrCls, ref.NameAndType.Name)
}

func getArrayMethod(cls *Class, name string) *Method {
	arrDesc := cls.Desc()
	elemTyp := arrDesc.ElemType()
	switch name {
	case "getClass":
		return &Method{
			Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "getClass", &desc.MethodDesc{Output: desc.DescClass}, nil),
			class:  cls,
			native: nativeArrayClass,
		}
	case "clone":
		switch elemTyp {
		case desc.Boolean:
			return booleanArrayCloneMethod
		case desc.Byte:
			return byteArrayCloneMethod
		case desc.Char:
			return charArrayCloneMethod
		case desc.Short:
			return shortArrayCloneMethod
		case desc.Int:
			return intArrayCloneMethod
		case desc.Long:
			return longArrayCloneMethod
		case desc.Float:
			return floatArrayCloneMethod
		case desc.Double:
			return doubleArrayCloneMethod
		case desc.Array, desc.Class:
			return &Method{
				Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: arrDesc}, nil),
				class:  cls,
				native: refArrayClone,
			}
		}
	}
	panic("unknown array method: " + arrDesc.String() + "." + name)
}

var (
	booleanArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescBooleanArray}, nil),
		class:  BooleanArrayClass,
		native: int8ArrayClone,
	}
	byteArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescByteArray}, nil),
		class:  ByteArrayClass,
		native: int8ArrayClone,
	}
	charArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescCharArray}, nil),
		class:  CharArrayClass,
		native: int16ArrayClone,
	}
	shortArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescShortArray}, nil),
		class:  ShortArrayClass,
		native: int16ArrayClone,
	}
	intArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescIntArray}, nil),
		class:  IntArrayClass,
		native: int32ArrayClone,
	}
	longArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescLongArray}, nil),
		class:  LongArrayClass,
		native: int64ArrayClone,
	}
	floatArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescFloatArray}, nil),
		class:  FloatArrayClass,
		native: int32ArrayClone,
	}
	doubleArrayCloneMethod = &Method{
		Method: jcls.NewMethod(jcls.AccPublic|jcls.AccFinal|jcls.AccNative, "clone", &desc.MethodDesc{Output: desc.DescDoubleArray}, nil),
		class:  DoubleArrayClass,
		native: int64ArrayClone,
	}
)

func nativeArrayClass(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	stack.PushRef(arrRef.Class().AsRef(vm))
	return nil
}

func int8ArrayClone(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	arr := arrRef.GetInt8Arr()
	cloned := vm.NewObjectArray(arrRef.Class().Elem(), (int32)(len(arr)))
	copy(cloned.GetInt8Arr(), arr)
	stack.PushRef(cloned)
	return nil
}

func int16ArrayClone(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	arr := arrRef.GetInt16Arr()
	cloned := vm.NewObjectArray(arrRef.Class().Elem(), (int32)(len(arr)))
	copy(cloned.GetInt16Arr(), arr)
	stack.PushRef(cloned)
	return nil
}

func int32ArrayClone(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	arr := arrRef.GetInt32Arr()
	cloned := vm.NewObjectArray(arrRef.Class().Elem(), (int32)(len(arr)))
	copy(cloned.GetInt32Arr(), arr)
	stack.PushRef(cloned)
	return nil
}

func int64ArrayClone(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	arr := arrRef.GetInt64Arr()
	cloned := vm.NewObjectArray(arrRef.Class().Elem(), (int32)(len(arr)))
	copy(cloned.GetInt64Arr(), arr)
	stack.PushRef(cloned)
	return nil
}

func refArrayClone(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	arr := arrRef.GetRefArr()
	cloned := vm.NewObjectArray(arrRef.Class().Elem(), (int32)(len(arr)))
	copy(cloned.GetRefArr(), arr)
	stack.PushRef(cloned)
	return nil
}
