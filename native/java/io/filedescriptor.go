package java_io

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/io/FileDescriptor.initIDs()V", FileDescriptor_initIDs)
	registerDefaultNative("java/io/FileDescriptor.getHandle(I)J", FileDescriptor_getHandle)
	registerDefaultNative("java/io/FileDescriptor.getAppend(I)Z", FileDescriptor_getAppend)
}

// private native void sync0() throws SyncFailedException;

// private static native void initIDs();
func FileDescriptor_initIDs(vm ir.VM) error {
	return nil
}

// private static native long getHandle(int d);
func FileDescriptor_getHandle(vm ir.VM) error {
	vm.GetStack().PushInt64(-1)
	return nil
}

// private static native boolean getAppend(int fd);
func FileDescriptor_getAppend(vm ir.VM) error {
	stack := vm.GetStack()
	fd := stack.GetVarInt32(0)
	_ = fd
	stack.Push(0)
	return nil
}

// private native void close0() throws IOException;
