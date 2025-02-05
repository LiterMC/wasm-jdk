package java_lang

import (
	"time"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
)

func init() {
	registerDefaultNative("java/lang/System.registerNatives()V", System_registerNatives)
}

// private static native void registerNatives();
func System_registerNatives(vm ir.VM) error {
	loadNative(vm, "java/lang/System.setIn0(Ljava/io/InputStream;)V", System_setIn0)
	loadNative(vm, "java/lang/System.setOut0(Ljava/io/PrintStream;)V", System_setOut0)
	loadNative(vm, "java/lang/System.setErr0(Ljava/io/PrintStream;)V", System_setErr0)
	loadNative(vm, "java/lang/System.currentTimeMillis()J", System_currentTimeMillis)
	loadNative(vm, "java/lang/System.nanoTime()J", System_nanoTime)
	loadNative(vm, "java/lang/System.arraycopy(Ljava/lang/Object;ILjava/lang/Object;II)V", System_arraycopy)
	loadNative(vm, "java/lang/System.identityHashCode(Ljava/lang/Object;)I", System_identityHashCode)
	loadNative(vm, "java/lang/System.mapLibraryName(Ljava/lang/String;)Ljava/lang/String;", System_mapLibraryName)
	return nil
}

// private static native void setIn0(InputStream in);
func System_setIn0(vm ir.VM) error {
	stream := vm.GetStack().GetVarRef(0)
	_ = stream
	return nil
}

// private static native void setOut0(PrintStream out);
func System_setOut0(vm ir.VM) error {
	stream := vm.GetStack().GetVarRef(0)
	_ = stream
	return nil
}

// private static native void setErr0(PrintStream err);
func System_setErr0(vm ir.VM) error {
	stream := vm.GetStack().GetVarRef(0)
	_ = stream
	return nil
}

// public static native long currentTimeMillis();
func System_currentTimeMillis(vm ir.VM) error {
	vm.GetStack().PushInt64(time.Now().UnixMilli())
	return nil
}

// public static native long nanoTime();
func System_nanoTime(vm ir.VM) error {
	vm.GetStack().PushInt64(time.Now().UnixNano())
	return nil
}

// public static native void arraycopy(Object src, int srcPos, Object dest, int destPos, int length);
func System_arraycopy(vm ir.VM) error {
	stack := vm.GetStack()
	src := stack.GetVarRef(0)
	srcPos := stack.GetVarInt32(1)
	dest := stack.GetVarRef(2)
	destPos := stack.GetVarInt32(3)
	length := stack.GetVarInt32(4)
	srcTyp := src.Desc()
	if !srcTyp.Eq(dest.Desc()) {
		panic("source and destination type not match")
	}
	if srcPos < 0 || srcPos+length > src.Len() {
		panic("source length too small")
	}
	if destPos < 0 || destPos+length > dest.Len() {
		panic("destination length too small")
	}
	if length == 0 {
		return nil
	}
	elemSize := (int32)(srcTyp.Type().Size())
	srcData := unsafe.Slice((*byte)(unsafe.Add(src.Data(), srcPos*elemSize)), length*elemSize)
	destData := unsafe.Slice((*byte)(unsafe.Add(dest.Data(), destPos*elemSize)), length*elemSize)
	copy(srcData, destData)
	return nil
}

// public static native int identityHashCode(Object x);
func System_identityHashCode(vm ir.VM) error {
	stack := vm.GetStack()
	ref := stack.GetVarRef(0)
	stack.PushInt32(ref.Id())
	return nil
}

// public static native String mapLibraryName(String libname);
func System_mapLibraryName(vm ir.VM) error {
	stack := vm.GetStack()
	strRef := stack.GetVarRef(0)
	str := vm.GetString(strRef)
	stack.PushRef(vm.NewString(str + ".js"))
	return nil
}
