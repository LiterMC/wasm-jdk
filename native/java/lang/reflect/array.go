package java_lang_reflect

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/reflect/Array.getLength(Ljava/lang/Object;)I", Array_getLength)
	native.RegisterDefaultNative("java/lang/reflect/Array.get(Ljava/lang/Object;I)Ljava/lang/Object;", Array_get)
	native.RegisterDefaultNative("java/lang/reflect/Array.getBoolean(Ljava/lang/Object;I)Z", Array_getBoolean)
	native.RegisterDefaultNative("java/lang/reflect/Array.getByte(Ljava/lang/Object;I)B", Array_getByte)
	native.RegisterDefaultNative("java/lang/reflect/Array.getChar(Ljava/lang/Object;I)C", Array_getChar)
	native.RegisterDefaultNative("java/lang/reflect/Array.getShort(Ljava/lang/Object;I)S", Array_getShort)
	native.RegisterDefaultNative("java/lang/reflect/Array.getInt(Ljava/lang/Object;I)I", Array_getInt)
	native.RegisterDefaultNative("java/lang/reflect/Array.getLong(Ljava/lang/Object;I)J", Array_getLong)
	native.RegisterDefaultNative("java/lang/reflect/Array.getFloat(Ljava/lang/Object;I)F", Array_getFloat)
	native.RegisterDefaultNative("java/lang/reflect/Array.getDouble(Ljava/lang/Object;I)D", Array_getDouble)
	native.RegisterDefaultNative("java/lang/reflect/Array.set(Ljava/lang/Object;ILjava/lang/Object;)V", Array_set)
	native.RegisterDefaultNative("java/lang/reflect/Array.setBoolean(Ljava/lang/Object;IZ)V", Array_setBoolean)
	native.RegisterDefaultNative("java/lang/reflect/Array.setByte(Ljava/lang/Object;IB)V", Array_setByte)
	native.RegisterDefaultNative("java/lang/reflect/Array.setChar(Ljava/lang/Object;IC)V", Array_setChar)
	native.RegisterDefaultNative("java/lang/reflect/Array.setShort(Ljava/lang/Object;IS)V", Array_setShort)
	native.RegisterDefaultNative("java/lang/reflect/Array.setInt(Ljava/lang/Object;II)V", Array_setInt)
	native.RegisterDefaultNative("java/lang/reflect/Array.setLong(Ljava/lang/Object;IJ)V", Array_setLong)
	native.RegisterDefaultNative("java/lang/reflect/Array.setFloat(Ljava/lang/Object;IF)V", Array_setFloat)
	native.RegisterDefaultNative("java/lang/reflect/Array.setDouble(Ljava/lang/Object;ID)V", Array_setDouble)
	native.RegisterDefaultNative("java/lang/reflect/Array.newArray(Ljava/lang/Class;I)Ljava/lang/Object;", Array_newArray)
	native.RegisterDefaultNative("java/lang/reflect/Array.multiNewArray(Ljava/lang/Class;[I)Ljava/lang/Object;", Array_multiNewArray)
}

// public static native int getLength(Object array) throws IllegalArgumentException;
func Array_getLength(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	stack.PushInt32(arrRef.Len())
	return nil
}

// public static native Object get(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_get(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushPointer(arrRef.GetRefArr()[index])
	return nil
}

// public static native boolean getBoolean(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getBoolean(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt8(arrRef.GetInt8Arr()[index])
	return nil
}

// public static native byte getByte(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getByte(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt8(arrRef.GetInt8Arr()[index])
	return nil
}

// public static native char getChar(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getChar(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt16(arrRef.GetInt16Arr()[index])
	return nil
}

// public static native short getShort(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getShort(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt16(arrRef.GetInt16Arr()[index])
	return nil
}

// public static native int getInt(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getInt(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt32(arrRef.GetInt32Arr()[index])
	return nil
}

// public static native long getLong(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getLong(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt64(arrRef.GetInt64Arr()[index])
	return nil
}

// public static native float getFloat(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getFloat(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt32(arrRef.GetInt32Arr()[index])
	return nil
}

// public static native double getDouble(Object array, int index) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_getDouble(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	stack.PushInt64(arrRef.GetInt64Arr()[index])
	return nil
}

// public static native void set(Object array, int index, Object value) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_set(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarPointer(2)
	arrRef.GetRefArr()[index] = v
	return nil
}

// public static native void setBoolean(Object array, int index, boolean z) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setBoolean(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt8(2)
	arrRef.GetInt8Arr()[index] = v
	return nil
}

// public static native void setByte(Object array, int index, byte b) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setByte(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt8(2)
	arrRef.GetInt8Arr()[index] = v
	return nil
}

// public static native void setChar(Object array, int index, char c) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setChar(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt16(2)
	arrRef.GetInt16Arr()[index] = v
	return nil
}

// public static native void setShort(Object array, int index, short s) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setShort(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt16(2)
	arrRef.GetInt16Arr()[index] = v
	return nil
}

// public static native void setInt(Object array, int index, int i) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setInt(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt32(2)
	arrRef.GetInt32Arr()[index] = v
	return nil
}

// public static native void setLong(Object array, int index, long l) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setLong(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt64(2)
	arrRef.GetInt64Arr()[index] = v
	return nil
}

// public static native void setFloat(Object array, int index, float f) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setFloat(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt32(2)
	arrRef.GetInt32Arr()[index] = v
	return nil
}

// public static native void setDouble(Object array, int index, double d) throws IllegalArgumentException, ArrayIndexOutOfBoundsException;
func Array_setDouble(vm ir.VM) error {
	stack := vm.GetStack()
	arrRef := stack.GetVarRef(0)
	index := stack.GetVarInt32(1)
	v := stack.GetVarInt64(2)
	arrRef.GetInt64Arr()[index] = v
	return nil
}

// private static native Object newArray(Class<?> componentType, int length) throws NegativeArraySizeException;
func Array_newArray(vm ir.VM) error {
	stack := vm.GetStack()
	classRef := stack.GetVarRef(0)
	length := stack.GetVarInt32(1)
	arrRef := vm.NewObjectArray((*classRef.UserData()).(ir.Class), length)
	stack.PushRef(arrRef)
	return nil
}

// private static native Object multiNewArray(Class<?> componentType, int[] dimensions) throws IllegalArgumentException, NegativeArraySizeException;
func Array_multiNewArray(vm ir.VM) error {
	stack := vm.GetStack()
	classRef := stack.GetVarRef(0)
	dimensions := stack.GetVarRef(1).GetInt32Arr()
	arrRef := vm.NewObjectMultiDimArray((*classRef.UserData()).(ir.Class), dimensions)
	stack.PushRef(arrRef)
	return nil
}
