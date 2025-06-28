package jdk_internal_reflect

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getSize0(Ljava/lang/Object;)I", ConstantPool_getSize0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getClassAt0(Ljava/lang/Object;I)Ljava/lang/Class;", ConstantPool_getClassAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getClassAtIfLoaded0(Ljava/lang/Object;I)Ljava/lang/Class;", ConstantPool_getClassAtIfLoaded0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getClassRefIndexAt0(Ljava/lang/Object;I)I", ConstantPool_getClassRefIndexAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getMethodAt0(Ljava/lang/Object;I)Ljava/lang/reflect/Member;", ConstantPool_getMethodAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getMethodAtIfLoaded0(Ljava/lang/Object;I)Ljava/lang/reflect/Member;", ConstantPool_getMethodAtIfLoaded0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getFieldAt0(Ljava/lang/Object;I)Ljava/lang/reflect/Field;", ConstantPool_getFieldAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getFieldAtIfLoaded0(Ljava/lang/Object;I)Ljava/lang/reflect/Field;", ConstantPool_getFieldAtIfLoaded0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getMemberRefInfoAt0(Ljava/lang/Object;I)[Ljava/lang/String;", ConstantPool_getMemberRefInfoAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getNameAndTypeRefIndexAt0(Ljava/lang/Object;I)I", ConstantPool_getNameAndTypeRefIndexAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getNameAndTypeRefInfoAt0(Ljava/lang/Object;I)[Ljava/lang/String;", ConstantPool_getNameAndTypeRefInfoAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getIntAt0(Ljava/lang/Object;I)I", ConstantPool_getIntAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getLongAt0(Ljava/lang/Object;I)J", ConstantPool_getLongAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getFloatAt0(Ljava/lang/Object;I)F", ConstantPool_getFloatAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getDoubleAt0(Ljava/lang/Object;I)D", ConstantPool_getDoubleAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getStringAt0(Ljava/lang/Object;I)[Ljava/lang/String;", ConstantPool_getStringAt0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getUTF8At0(Ljava/lang/Object;I)[Ljava/lang/String;", ConstantPool_getUTF8At0)
	native.RegisterDefaultNative("jdk/internal/reflect/ConstantPool.getTagAt0(Ljava/lang/Object;I)B", ConstantPool_getTagAt0)
}

// private native int      getSize0            (Object constantPoolOop);
func ConstantPool_getSize0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	stack.PushInt32((int32)(len(pool)))
	return nil
}

// private native Class<?> getClassAt0         (Object constantPoolOop, int index);
func ConstantPool_getClassAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantClass)
	class, err := vm.GetClassByName(info.Name)
	if err != nil {
		return err
	}
	stack.PushRef(class.AsRef(vm))
	return nil
}

// private native Class<?> getClassAtIfLoaded0 (Object constantPoolOop, int index);
func ConstantPool_getClassAtIfLoaded0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantClass)
	class, err := vm.GetLoadedClassByName(info.Name)
	if err != nil {
		return err
	}
	if class == nil {
		stack.PushRef(nil)
	} else {
		stack.PushRef(class.AsRef(vm))
	}
	return nil
}

// private native int      getClassRefIndexAt0 (Object constantPoolOop, int index);
func ConstantPool_getClassRefIndexAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	stack.PushInt32((int32)(info.ClassInd))
	return nil
}

// private native Member   getMethodAt0        (Object constantPoolOop, int index);
func ConstantPool_getMethodAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	_ = info
	if true {
		panic("TODO")
	}
	return nil
}

// private native Member   getMethodAtIfLoaded0(Object constantPoolOop, int index);
func ConstantPool_getMethodAtIfLoaded0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	_ = info
	if true {
		panic("TODO")
	}
	return nil
}

// private native Field    getFieldAt0         (Object constantPoolOop, int index);
func ConstantPool_getFieldAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	_ = info
	if true {
		panic("TODO")
	}
	return nil
}

// private native Field    getFieldAtIfLoaded0 (Object constantPoolOop, int index);
func ConstantPool_getFieldAtIfLoaded0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	_ = info
	if true {
		panic("TODO")
	}
	return nil
}

// private native String[] getMemberRefInfoAt0 (Object constantPoolOop, int index);
func ConstantPool_getMemberRefInfoAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantRef)
	resultRef := vm.NewArray(desc.DescStringArray, 3)
	results := resultRef.GetRefArr()
	results[0] = vm.RefToPtr(vm.GetStringInternOrNew(info.Class.Name))
	results[1] = vm.RefToPtr(vm.GetStringInternOrNew(info.NameAndType.Name))
	results[2] = vm.RefToPtr(vm.GetStringInternOrNew(info.NameAndType.Desc))
	stack.PushRef(resultRef)
	return nil
}

// private native int      getNameAndTypeRefIndexAt0(Object constantPoolOop, int index);
func ConstantPool_getNameAndTypeRefIndexAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1]
	var ind uint16
	switch info := info.(type) {
	case *jcls.ConstantRef:
		ind = info.NameAndTypeInd
	case *jcls.ConstantDynamics:
		ind = info.NameAndTypeInd
	default:
		panic("TODO: Exception")
	}
	stack.PushInt32((int32)(ind))
	return nil
}

// private native String[] getNameAndTypeRefInfoAt0(Object constantPoolOop, int index);
func ConstantPool_getNameAndTypeRefInfoAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantNameAndType)
	resultRef := vm.NewArray(desc.DescStringArray, 2)
	results := resultRef.GetRefArr()
	results[0] = vm.RefToPtr(vm.GetStringInternOrNew(info.Name))
	results[1] = vm.RefToPtr(vm.GetStringInternOrNew(info.Desc))
	stack.PushRef(resultRef)
	return nil
}

// private native int      getIntAt0           (Object constantPoolOop, int index);
func ConstantPool_getIntAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantInteger)
	stack.Push(info.Value)
	return nil
}

// private native long     getLongAt0          (Object constantPoolOop, int index);
func ConstantPool_getLongAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantLong)
	stack.Push64(info.Value)
	return nil
}

// private native float    getFloatAt0         (Object constantPoolOop, int index);
func ConstantPool_getFloatAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantFloat)
	stack.Push(info.Value)
	return nil
}

// private native double   getDoubleAt0        (Object constantPoolOop, int index);
func ConstantPool_getDoubleAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantDouble)
	stack.Push64(info.Value)
	return nil
}

// private native String   getStringAt0        (Object constantPoolOop, int index);
func ConstantPool_getStringAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantString)
	stack.PushRef(vm.GetStringInternOrNew(info.Utf8))
	return nil
}

// private native String   getUTF8At0          (Object constantPoolOop, int index);
func ConstantPool_getUTF8At0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1].(*jcls.ConstantUtf8)
	stack.PushRef(vm.GetStringInternOrNew(info.Value))
	return nil
}

// private native byte     getTagAt0           (Object constantPoolOop, int index);
func ConstantPool_getTagAt0(vm ir.VM) error {
	stack := vm.GetStack()
	pool := *(*stack.GetVarRef(0).UserData()).(*[]jcls.ConstantInfo)
	index := stack.GetVarInt32(2)
	info := pool[index-1]
	stack.PushInt8((int8)(info.Tag()))
	return nil
}
