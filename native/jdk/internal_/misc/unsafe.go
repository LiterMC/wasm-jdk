package jdk_internal_misc

import (
	"fmt"
	"sync/atomic"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/cutil"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/misc/Unsafe.registerNatives()V", Unsafe_registerNatives)
}

// private static native void registerNatives();
func Unsafe_registerNatives(vm ir.VM) error {
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getInt(Ljava/lang/Object;J)I", Unsafe_getInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putInt(Ljava/lang/Object;JI)V", Unsafe_putInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getReference(Ljava/lang/Object;J)Ljava/lang/Object;", Unsafe_getReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putReference(Ljava/lang/Object;JLjava/lang/Object;)V", Unsafe_putReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getBoolean(Ljava/lang/Object;J)Z", Unsafe_getBoolean)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putBoolean(Ljava/lang/Object;JZ)V", Unsafe_putBoolean)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getByte(Ljava/lang/Object;J)B", Unsafe_getByte)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putByte(Ljava/lang/Object;JB)V", Unsafe_putByte)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getShort(Ljava/lang/Object;J)S", Unsafe_getShort)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putShort(Ljava/lang/Object;JS)V", Unsafe_putShort)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getChar(Ljava/lang/Object;J)C", Unsafe_getChar)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putChar(Ljava/lang/Object;JC)V", Unsafe_putChar)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getLong(Ljava/lang/Object;J)J", Unsafe_getLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putLong(Ljava/lang/Object;JJ)V", Unsafe_putLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getFloat(Ljava/lang/Object;J)F", Unsafe_getFloat)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putFloat(Ljava/lang/Object;JF)V", Unsafe_putFloat)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getDouble(Ljava/lang/Object;J)D", Unsafe_getDouble)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putDouble(Ljava/lang/Object;JD)V", Unsafe_putDouble)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getUncompressedObject(J)Ljava/lang/Object;", Unsafe_getUncompressedObject)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.writeback0(J)V", Unsafe_writeback0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.writebackPreSync0()V", Unsafe_writebackPreSync0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.writebackPostSync0()V", Unsafe_writebackPostSync0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.defineClass0(Ljava/lang/String;[BIILjava/lang/ClassLoader;Ljava/security/ProtectionDomain;)Ljava/lang/Class;", Unsafe_defineClass0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.allocateInstance(Ljava/lang/Class;)Ljava/lang/Object;", Unsafe_allocateInstance)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.throwException(Ljava/lang/Throwable;)V", Unsafe_throwException)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndSetReference(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", Unsafe_compareAndSetReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndExchangeReference(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Ljava/lang/Object;", Unsafe_compareAndExchangeReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndSetInt(Ljava/lang/Object;JII)Z", Unsafe_compareAndSetInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndExchangeInt(Ljava/lang/Object;JII)I", Unsafe_compareAndExchangeInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndSetLong(Ljava/lang/Object;JJJ)Z", Unsafe_compareAndSetLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.compareAndExchangeLong(Ljava/lang/Object;JJJ)J", Unsafe_compareAndExchangeLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.unpark(Ljava/lang/Object;)V", Unsafe_unpark)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.park(ZJ)V", Unsafe_park)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.fullFence()V", Unsafe_fullFence)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.allocateMemory0(J)J", Unsafe_allocateMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.reallocateMemory0(JJ)J", Unsafe_reallocateMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.freeMemory0(J)V", Unsafe_freeMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.setMemory0(Ljava/lang/Object;JJB)V", Unsafe_setMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.copyMemory0(Ljava/lang/Object;JLjava/lang/Object;JJ)V", Unsafe_copyMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.copySwapMemory0(Ljava/lang/Object;JLjava/lang/Object;JJJ)V", Unsafe_copySwapMemory0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.objectFieldOffset0(Ljava/lang/reflect/Field;)J", Unsafe_objectFieldOffset0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.objectFieldOffset1(Ljava/lang/Class;Ljava/lang/String;)J", Unsafe_objectFieldOffset1)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.staticFieldOffset0(Ljava/lang/reflect/Field;)J", Unsafe_staticFieldOffset0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.staticFieldBase0(Ljava/lang/reflect/Field;)Ljava/lang/Object;", Unsafe_staticFieldBase0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.shouldBeInitialized0(Ljava/lang/Class;)Z", Unsafe_shouldBeInitialized0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.ensureClassInitialized0(Ljava/lang/Class;)V", Unsafe_ensureClassInitialized0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.arrayBaseOffset0(Ljava/lang/Class;)I", Unsafe_arrayBaseOffset0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.arrayIndexScale0(Ljava/lang/Class;)I", Unsafe_arrayIndexScale0)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getLoadAverage0([DI)I", Unsafe_getLoadAverage0)

	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getReferenceVolatile(Ljava/lang/Object;J)Ljava/lang/Object;", Unsafe_getReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putReferenceVolatile(Ljava/lang/Object;JLjava/lang/Object;)V", Unsafe_putReference)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getIntVolatile(Ljava/lang/Object;J)I", Unsafe_getInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putIntVolatile(Ljava/lang/Object;JI)V", Unsafe_putInt)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getBooleanVolatile(Ljava/lang/Object;J)Z", Unsafe_getBoolean)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putBooleanVolatile(Ljava/lang/Object;JZ)V", Unsafe_putBoolean)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getByteVolatile(Ljava/lang/Object;J)B", Unsafe_getByte)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putByteVolatile(Ljava/lang/Object;JB)V", Unsafe_putByte)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getShortVolatile(Ljava/lang/Object;J)S", Unsafe_getShort)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putShortVolatile(Ljava/lang/Object;JS)V", Unsafe_putShort)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getCharVolatile(Ljava/lang/Object;J)C", Unsafe_getChar)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putCharVolatile(Ljava/lang/Object;JC)V", Unsafe_putChar)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getLongVolatile(Ljava/lang/Object;J)J", Unsafe_getLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putLongVolatile(Ljava/lang/Object;JJ)V", Unsafe_putLong)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getFloatVolatile(Ljava/lang/Object;J)F", Unsafe_getFloat)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putFloatVolatile(Ljava/lang/Object;JF)V", Unsafe_putFloat)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.getDoubleVolatile(Ljava/lang/Object;J)D", Unsafe_getDouble)
	native.LoadNative(vm, "jdk/internal/misc/Unsafe.putDoubleVolatile(Ljava/lang/Object;JD)V", Unsafe_putDouble)
	return nil
}

// public native int getInt(Object o, long offset);
func Unsafe_getInt(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void putInt(Object o, long offset, int x);
func Unsafe_putInt(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native Object getReference(Object o, long offset);
func Unsafe_getReference(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.PushRef((*jvm.Ref)(atomic.LoadPointer((*unsafe.Pointer)(ptr))))
	return nil
}

// public native void putReference(Object o, long offset, Object x);
func Unsafe_putReference(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := asJvmRef(stack.GetVarRef(4))
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StorePointer((*unsafe.Pointer)(ptr), (unsafe.Pointer)(value))
	return nil
}

// public native boolean getBoolean(Object o, long offset);
func Unsafe_getBoolean(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void    putBoolean(Object o, long offset, boolean x);
func Unsafe_putBoolean(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native byte    getByte(Object o, long offset);
func Unsafe_getByte(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void    putByte(Object o, long offset, byte x);
func Unsafe_putByte(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native short   getShort(Object o, long offset);
func Unsafe_getShort(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void    putShort(Object o, long offset, short x);
func Unsafe_putShort(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native char    getChar(Object o, long offset);
func Unsafe_getChar(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void    putChar(Object o, long offset, char x);
func Unsafe_putChar(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native long    getLong(Object o, long offset);
func Unsafe_getLong(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push64(atomic.LoadUint64((*uint64)(ptr)))
	return nil
}

// public native void    putLong(Object o, long offset, long x);
func Unsafe_putLong(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar64(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint64((*uint64)(ptr), value)
	return nil
}

// public native float   getFloat(Object o, long offset);
func Unsafe_getFloat(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push(atomic.LoadUint32((*uint32)(ptr)))
	return nil
}

// public native void    putFloat(Object o, long offset, float x);
func Unsafe_putFloat(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint32((*uint32)(ptr), value)
	return nil
}

// public native double  getDouble(Object o, long offset);
func Unsafe_getDouble(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	ptr := unsafe.Add(getRefData(ref), offset)
	stack.Push64(atomic.LoadUint64((*uint64)(ptr)))
	return nil
}

// public native void    putDouble(Object o, long offset, double x);
func Unsafe_putDouble(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	value := stack.GetVar64(4)
	ptr := unsafe.Add(getRefData(ref), offset)
	atomic.StoreUint64((*uint64)(ptr), value)
	return nil
}

// public native Object getUncompressedObject(long address);
func Unsafe_getUncompressedObject(vm ir.VM) error {
	stack := vm.GetStack()
	address := stack.GetVar64(1)
	ptr := (unsafe.Pointer)((uintptr)(address))
	ref := (*jvm.Ref)(atomic.LoadPointer((*unsafe.Pointer)(ptr)))
	stack.PushRef(ref)
	return nil
}

// private native void writeback0(long address);
func Unsafe_writeback0(vm ir.VM) error {
	return nil
}

// private native void writebackPreSync0();
func Unsafe_writebackPreSync0(vm ir.VM) error {
	return nil
}

// private native void writebackPostSync0();
func Unsafe_writebackPostSync0(vm ir.VM) error {
	return nil
}

// public native Class<?> defineClass0(String name, byte[] b, int off, int len, ClassLoader loader, ProtectionDomain protectionDomain);
func Unsafe_defineClass0(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(1))
	if true {
		panic("TODO: Unsafe.debugClass0: creating " + name)
	}
	return nil
}

// public native Object allocateInstance(Class<?> cls) throws InstantiationException;
func Unsafe_allocateInstance(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(1).UserData()).(ir.Class)
	ref := vm.New(class)
	stack.PushRef(ref)
	return nil
}

// public native void throwException(Throwable ee);
func Unsafe_throwException(vm ir.VM) error {
	exception := vm.GetStack().GetVarRef(1)
	vm.Throw(exception)
	return nil
}

// public final native boolean compareAndSetReference(Object o, long offset, Object expected, Object x);
func Unsafe_compareAndSetReference(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := asJvmRef(stack.GetVarRef(4))
	value := asJvmRef(stack.GetVarRef(5))
	ptr := unsafe.Add(getRefData(ref), offset)
	if atomic.CompareAndSwapPointer((*unsafe.Pointer)(ptr), (unsafe.Pointer)(expected), (unsafe.Pointer)(value)) {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}

// public final native Object compareAndExchangeReference(Object o, long offset, Object expected, Object x);
func Unsafe_compareAndExchangeReference(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := asJvmRef(stack.GetVarRef(4))
	value := asJvmRef(stack.GetVarRef(5))
	ptr := (*unsafe.Pointer)(unsafe.Add(getRefData(ref), offset))
	old := atomic.LoadPointer(ptr)
	for {
		if atomic.CompareAndSwapPointer(ptr, (unsafe.Pointer)(expected), (unsafe.Pointer)(value)) {
			stack.PushRef(expected)
			return nil
		}
		o := old
		if old = atomic.LoadPointer(ptr); o == old {
			stack.PushRef((*jvm.Ref)(o))
			return nil
		}
	}
}

// public final native boolean compareAndSetInt(Object o, long offset, int expected, int x);
func Unsafe_compareAndSetInt(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := stack.GetVar(4)
	value := stack.GetVar(5)
	ptr := unsafe.Add(getRefData(ref), offset)
	if atomic.CompareAndSwapUint32((*uint32)(ptr), expected, value) {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}

// public final native int compareAndExchangeInt(Object o, long offset, int expected, int x);
func Unsafe_compareAndExchangeInt(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := stack.GetVar(4)
	value := stack.GetVar(5)
	ptr := (*uint32)(unsafe.Add(getRefData(ref), offset))
	old := atomic.LoadUint32(ptr)
	for {
		if atomic.CompareAndSwapUint32(ptr, expected, value) {
			stack.Push(expected)
			return nil
		}
		o := old
		if old = atomic.LoadUint32(ptr); o == old {
			stack.Push(o)
			return nil
		}
	}
}

// public final native boolean compareAndSetLong(Object o, long offset, long expected, long x);
func Unsafe_compareAndSetLong(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := stack.GetVar64(4)
	value := stack.GetVar64(6)
	ptr := unsafe.Add(getRefData(ref), offset)
	if atomic.CompareAndSwapUint64((*uint64)(ptr), expected, value) {
		stack.Push64(1)
	} else {
		stack.Push64(0)
	}
	return nil
}

// public final native long compareAndExchangeLong(Object o, long offset, long expected, long x);
func Unsafe_compareAndExchangeLong(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	expected := stack.GetVar64(4)
	value := stack.GetVar64(6)
	ptr := (*uint64)(unsafe.Add(getRefData(ref), offset))
	old := atomic.LoadUint64(ptr)
	for {
		if atomic.CompareAndSwapUint64(ptr, expected, value) {
			stack.Push64(expected)
			return nil
		}
		o := old
		if old = atomic.LoadUint64(ptr); o == old {
			stack.Push64(o)
			return nil
		}
	}
}

// public native void unpark(Object thread);
func Unsafe_unpark(vm ir.VM) error {
	stack := vm.GetStack()
	thread := stack.GetVarRef(1)
	_ = thread
	return nil
}

// public native void park(boolean isAbsolute, long time);
func Unsafe_park(vm ir.VM) error {
	stack := vm.GetStack()
	isAbsolute := stack.GetVar(1) != 0
	time := stack.GetVar64(2)
	_, _ = isAbsolute, time
	return nil
}

// public native void fullFence();
func Unsafe_fullFence(vm ir.VM) error {
	return nil
}

// private native long allocateMemory0(long bytes);
func Unsafe_allocateMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	bytes := stack.GetVarInt64(1)
	ptr := cutil.AllocMemory((int)(bytes))
	stack.Push64((uint64)(ptr))
	return nil
}

// private native long reallocateMemory0(long address, long bytes);
func Unsafe_reallocateMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	address := (uintptr)(stack.GetVar64(1))
	bytes := stack.GetVarInt64(2)
	address = cutil.ReallocMemory(address, (int)(bytes))
	stack.Push64((uint64)(address))
	return nil
}

// private native void freeMemory0(long address);
func Unsafe_freeMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	address := (uintptr)(stack.GetVar64(1))
	cutil.FreeMemory(address)
	return nil
}

// private native void setMemory0(Object o, long offset, long bytes, byte value);
func Unsafe_setMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(1)
	offset := stack.GetVarInt64(2)
	bytes := stack.GetVarInt64(4)
	value := (byte)(stack.GetVarInt8(6))
	if bytes == 0 {
		return nil
	}
	ptr := (*byte)(unsafe.Add(getRefData(ref), offset))
	slice := unsafe.Slice(ptr, bytes)
	memset(slice, value)
	return nil
}

// Source: https://stackoverflow.com/questions/30614165/is-there-analog-of-memset-in-go
func memset(slice []byte, v byte) {
	slice[0] = v
	for b := 1; b < len(slice); b *= 2 {
		copy(slice[b:], slice[:b])
	}
}

func getRefData(ref ir.Ref) unsafe.Pointer {
	if ref == nil {
		return nil
	}
	return ref.Data()
}

// private native void copyMemory0(Object srcBase, long srcOffset, Object destBase, long destOffset, long bytes);
func Unsafe_copyMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	srcRef := stack.GetVarRef(1)
	srcOffset := stack.GetVarInt64(2)
	destRef := stack.GetVarRef(4)
	destOffset := stack.GetVarInt64(5)
	bytes := stack.GetVarInt64(7)
	if bytes == 0 {
		return nil
	}
	srcPtr := (*byte)(unsafe.Add(getRefData(srcRef), srcOffset))
	destPtr := (*byte)(unsafe.Add(getRefData(destRef), destOffset))
	copy(unsafe.Slice(destPtr, bytes), unsafe.Slice(srcPtr, bytes))
	return nil
}

// private native void copySwapMemory0(Object srcBase, long srcOffset, Object destBase, long destOffset, long bytes, long elemSize);
func Unsafe_copySwapMemory0(vm ir.VM) error {
	stack := vm.GetStack()
	srcRef := stack.GetVarRef(1)
	srcOffset := stack.GetVarInt64(2)
	destRef := stack.GetVarRef(4)
	destOffset := stack.GetVarInt64(5)
	bytes := stack.GetVarInt64(7)
	elemSize := stack.GetVarInt64(9)
	if bytes == 0 {
		return nil
	}
	if elemSize <= 4 {
		srcPtr := unsafe.Add(getRefData(srcRef), srcOffset)
		destPtr := unsafe.Add(getRefData(destRef), destOffset)
		length := (bytes + 3) / 4
		for i := range length {
			v := (*uint32)(unsafe.Add(srcPtr, i*4))
			old := atomic.SwapUint32((*uint32)(unsafe.Add(destPtr, i*4)), atomic.LoadUint32(v))
			atomic.StoreUint32(v, old)
		}
	} else if elemSize == 8 {
		srcPtr := unsafe.Add(getRefData(srcRef), srcOffset)
		destPtr := unsafe.Add(getRefData(destRef), destOffset)
		length := (bytes + 7) / 8
		for i := range length {
			v := (*uint64)(unsafe.Add(srcPtr, i*8))
			old := atomic.SwapUint64((*uint64)(unsafe.Add(destPtr, i*8)), atomic.LoadUint64(v))
			atomic.StoreUint64(v, old)
		}
	} else {
		panic(fmt.Errorf("unsupported elemSize %d", elemSize))
	}
	return nil
}

// private native long objectFieldOffset0(Field f);
func Unsafe_objectFieldOffset0(vm ir.VM) error {
	stack := vm.GetStack()
	field := stack.GetVarRef(1)
	if true {
		_ = field
		panic("TODO: Unsafe.objectFieldOffset0")
	}
	return nil
}

// private native long objectFieldOffset1(Class<?> c, String name);
func Unsafe_objectFieldOffset1(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(1).UserData()).(ir.Class)
	name := vm.GetString(stack.GetVarRef(2))
	field := class.GetFieldByName(name)
	stack.PushInt64(field.Offset())
	return nil
}

// private native long staticFieldOffset0(Field f);
func Unsafe_staticFieldOffset0(vm ir.VM) error {
	stack := vm.GetStack()
	field := stack.GetVarRef(1)
	if true {
		_ = field
		panic("TODO: Unsafe.staticFieldOffset0")
	}
	return nil
}

// private native Object staticFieldBase0(Field f);
func Unsafe_staticFieldBase0(vm ir.VM) error {
	stack := vm.GetStack()
	field := stack.GetVarRef(1)
	if true {
		_ = field
		panic("TODO: Unsafe.staticFieldBase0")
	}
	return nil
}

// private native boolean shouldBeInitialized0(Class<?> c);
func Unsafe_shouldBeInitialized0(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(1).UserData()).(*jvm.Class)
	if class.ShouldInit() {
		stack.Push(1)
	} else {
		stack.Push(0)
	}
	return nil
}

// private native void ensureClassInitialized0(Class<?> c);
func Unsafe_ensureClassInitialized0(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(1).UserData()).(*jvm.Class)
	class.InitBeforeUse(vm.(*jvm.VM))
	return nil
}

// private native int arrayBaseOffset0(Class<?> arrayClass);
func Unsafe_arrayBaseOffset0(vm ir.VM) error {
	stack := vm.GetStack()
	stack.Push(0)
	return nil
}

// private native int arrayIndexScale0(Class<?> arrayClass);
func Unsafe_arrayIndexScale0(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(1).UserData()).(ir.Class)
	size := class.Desc().ElemType().Size()
	if size < 4 {
		stack.Push(0)
	} else if size == 4 {
		stack.Push(4)
	} else if size == 8 {
		stack.Push(8)
	} else {
		panic("unexpected array element size")
	}
	return nil
}

// private native int getLoadAverage0(double[] loadavg, int nelems);
func Unsafe_getLoadAverage0(vm ir.VM) error {
	stack := vm.GetStack()
	loadavgRef := stack.GetVarRef(1)
	nelems := stack.GetVarInt32(2)
	_, _ = loadavgRef, nelems
	stack.PushInt32(-1)
	return nil
}

func asJvmRef(ref ir.Ref) *jvm.Ref {
	if ref == nil {
		return nil
	}
	return ref.(*jvm.Ref)
}
