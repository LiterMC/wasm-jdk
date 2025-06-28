package jdk_internal_perf

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/perf/Perf.registerNatives()V", Perf_registerNatives)
}

// private static native void registerNatives();
func Perf_registerNatives(vm ir.VM) error {
	native.LoadNative(vm, "jdk/internal/perf/Perf.attach0(I)Ljava/nio/ByteBuffer;", Perf_attach0)
	native.LoadNative(vm, "jdk/internal/perf/Perf.detach(Ljava/nio/ByteBuffer;)V", Perf_detach)
	native.LoadNative(vm, "jdk/internal/perf/Perf.createLong(Ljava/lang/String;IIJ)Ljava/nio/ByteBuffer;", Perf_createLong)
	native.LoadNative(vm, "jdk/internal/perf/Perf.createByteArray(Ljava/lang/String;II[BI)Ljava/nio/ByteBuffer;", Perf_createByteArray)
	native.LoadNative(vm, "jdk/internal/perf/Perf.highResCounter()J", Perf_highResCounter)
	native.LoadNative(vm, "jdk/internal/perf/Perf.highResFrequency()J", Perf_highResFrequency)
	return nil
}

// private native ByteBuffer attach0(int lvmid) throws IOException;
func Perf_attach0(vm ir.VM) error {
	stack := vm.GetStack()
	lvmid := stack.GetVarInt32(1)
	if true {
		println("lvmid:", lvmid)
		panic("not implemented attach0")
	}
	return nil
}

// private native void detach(ByteBuffer bb);
func Perf_detach(vm ir.VM) error {
	stack := vm.GetStack()
	buffer := stack.GetVarRef(1)
	if true {
		_ = buffer
		panic("not implemented detach")
	}
	return nil
}

// public native ByteBuffer createLong(String name, int variability, int units, long value);
func Perf_createLong(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(1))
	variability := stack.GetVarInt32(2)
	units := stack.GetVarInt32(3)
	value := stack.GetVarInt64(4)

	_, _, _ = name, variability, units

	jByteBuffer, err := vm.GetClassByName("java/nio/ByteBuffer")
	if err != nil {
		return err
	}
	method := jByteBuffer.GetMethodByName("allocateDirect(I)Ljava/nio/ByteBuffer;")
	stack.PushInt32(8)
	vm.InvokeStatic(method)
	if err := vm.RunStack(); err != nil {
		return err
	}
	bufferRef := stack.PeekRef()
	addressPtr := jByteBuffer.GetFieldByName("address").GetPointer(bufferRef)
	*(*int64)(addressPtr) = value
	return nil
}

// public native ByteBuffer createByteArray(String name, int variability, int units, byte[] value, int maxLength);
func Perf_createByteArray(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(1))
	variability := stack.GetVarInt32(2)
	units := stack.GetVarInt32(3)
	value := stack.GetVarRef(4).GetByteArr()
	maxLength := stack.GetVarInt32(5)
	if true {
		_, _, _, _ = variability, units, value, maxLength
		panic("not implemented createByteArray: " + name)
	}
	return nil
}

// public native long highResCounter();
func Perf_highResCounter(vm ir.VM) error {
	panic("Unsupported operation")
}

// public native long highResFrequency();
func Perf_highResFrequency(vm ir.VM) error {
	panic("Unsupported operation")
}
