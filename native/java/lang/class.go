package java_lang

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

func init() {
	registerDefaultNative("java/lang/Class.registerNatives()V", Class_registerNatives)
}

// private static native void registerNatives();
func Class_registerNatives(vm ir.VM) error {
	loadNative(vm, "java/lang/Class.forName0(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", Class_forName0)
	loadNative(vm, "java/lang/Class.isInstance(Ljava/lang/Object;)Z", Class_isInstance)
	loadNative(vm, "java/lang/Class.isAssignableFrom(Ljava/lang/Class;)Z", Class_isAssignableFrom)
	loadNative(vm, "java/lang/Class.isInterface()Z", Class_isInterface)
	loadNative(vm, "java/lang/Class.isArray()Z", Class_isArray)
	loadNative(vm, "java/lang/Class.isPrimitive()Z", Class_isPrimitive)
	loadNative(vm, "java/lang/Class.initClassName()Ljava/lang/String;", Class_initClassName)
	loadNative(vm, "java/lang/Class.getSuperclass()Ljava/lang/Class;", Class_getSuperclass)
	loadNative(vm, "java/lang/Class.getInterfaces0()[Ljava/lang/Class;", Class_getInterfaces0)
	loadNative(vm, "java/lang/Class.getModifiers()I", Class_getModifiers)
	loadNative(vm, "java/lang/Class.getSigners()[Ljava/lang/Object;", Class_getSigners)
	loadNative(vm, "java/lang/Class.setSigners([Ljava/lang/Object;)V", Class_setSigners)
	loadNative(vm, "java/lang/Class.getEnclosingMethod0()[Ljava/lang/Object;", Class_getEnclosingMethod0)
	loadNative(vm, "java/lang/Class.getDeclaringClass0()Ljava/lang/Class;", Class_getDeclaringClass0)
	loadNative(vm, "java/lang/Class.getSimpleBinaryName0()Ljava/lang/String;", Class_getSimpleBinaryName0)
	loadNative(vm, "java/lang/Class.getProtectionDomain0()Ljava/security/ProtectionDomain;", Class_getProtectionDomain0)
	loadNative(vm, "java/lang/Class.getPrimitiveClass(Ljava/lang/String;)Ljava/lang/Class;", Class_getPrimitiveClass)
	loadNative(vm, "java/lang/Class.getGenericSignature0()Ljava/lang/String;", Class_getGenericSignature0)
	loadNative(vm, "java/lang/Class.getRawAnnotations()[B", Class_getRawAnnotations)
	loadNative(vm, "java/lang/Class.getRawTypeAnnotations()[B", Class_getRawTypeAnnotations)
	loadNative(vm, "java/lang/Class.getConstantPool()Ljdk/internal/reflect/ConstantPool;", Class_getConstantPool)
	loadNative(vm, "java/lang/Class.getDeclaredFields0(Z)[Ljava/lang/reflect/Field;", Class_getDeclaredFields0)
	loadNative(vm, "java/lang/Class.getDeclaredMethods0(Z)[Ljava/lang/reflect/Method;", Class_getDeclaredMethods0)
	loadNative(vm, "java/lang/Class.getDeclaredConstructors0(Z)[Ljava/lang/reflect/Constructor;", Class_getDeclaredConstructors0)
	loadNative(vm, "java/lang/Class.getDeclaredClasses0()[Ljava/lang/Class;", Class_getDeclaredClasses0)
	loadNative(vm, "java/lang/Class.getRecordComponents0()[Ljava/lang/reflect/RecordComponent;", Class_getRecordComponents0)
	loadNative(vm, "java/lang/Class.isRecord0()Z", Class_isRecord0)
	loadNative(vm, "java/lang/Class.desiredAssertionStatus0(Ljava/lang/Class;)Z", Class_desiredAssertionStatus0)
	loadNative(vm, "java/lang/Class.getNestHost0()Ljava/lang/Class;", Class_getNestHost0)
	loadNative(vm, "java/lang/Class.getNestMembers0()[Ljava/lang/Class;", Class_getNestMembers0)
	loadNative(vm, "java/lang/Class.isHidden()Z", Class_isHidden)
	loadNative(vm, "java/lang/Class.getPermittedSubclasses0()[Ljava/lang/Class;", Class_getPermittedSubclasses0)
	loadNative(vm, "java/lang/Class.getClassFileVersion0()I", Class_getClassFileVersion0)
	loadNative(vm, "java/lang/Class.getClassAccessFlagsRaw0()I", Class_getClassAccessFlagsRaw0)
	return nil
}

// private static native Class<?> forName0(String name, boolean initialize, ClassLoader loader, Class<?> caller);
func Class_forName0(vm ir.VM) error {
	stack := vm.GetStack()
	_ = stack
	return nil
}

// public native boolean isInstance(Object obj);
func Class_isInstance(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	obj := stack.GetVarRef(1)
	if this.IsInstance(obj) {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

// public native boolean isAssignableFrom(Class<?> cls);
func Class_isAssignableFrom(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	cls := (*stack.GetVarRef(1).UserData()).(ir.Class)
	if this.IsAssignableFrom(cls) {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

// public native boolean isInterface();
func Class_isInterface(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	if this.IsInterface() {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

// public native boolean isArray();
func Class_isArray(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	if this.ArrayDim() > 0 {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

// public native boolean isPrimitive();
func Class_isPrimitive(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	dc := this.Desc()
	if dc.ArrDim == 0 && dc.EndType != desc.Array && dc.EndType != desc.Class {
		stack.PushInt32(1)
	} else {
		stack.PushInt32(0)
	}
	return nil
}

// private native String initClassName();
func Class_initClassName(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	stack.PushRef(vm.NewString(this.Name()))
	return nil
}

// public native Class<? super T> getSuperclass();
func Class_getSuperclass(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	stack.PushRef(vm.GetClassRef(this.Super()))
	return nil
}

// private native Class<?>[] getInterfaces0();
func Class_getInterfaces0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	ints := this.Interfaces()
	intsRef := vm.NewArray(desc.DescClassArray, (int32)(len(ints)))
	intsArr := intsRef.GetArrRef()
	for i, in := range ints {
		intsArr[i] = vm.RefToPtr(vm.GetClassRef(in))
	}
	stack.PushRef(intsRef)
	return nil
}

// public native int getModifiers();
func Class_getModifiers(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	stack.PushInt32(this.Modifiers())
	return nil
}

// public native Object[] getSigners();
func Class_getSigners(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	if dc := this.Desc(); dc.ArrDim == 0 && dc.EndType != desc.Array && dc.EndType != desc.Class {
		stack.PushRef(nil)
		return nil
	}
	// TODO
	stack.PushRef(nil)
	return nil
}

// native void setSigners(Object[] signers);
func Class_setSigners(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Object[] getEnclosingMethod0();
func Class_getEnclosingMethod0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Class<?> getDeclaringClass0();
func Class_getDeclaringClass0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native String getSimpleBinaryName0();
func Class_getSimpleBinaryName0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native java.security.ProtectionDomain getProtectionDomain0();
func Class_getProtectionDomain0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// static native Class<?> getPrimitiveClass(String name);
func Class_getPrimitiveClass(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	stack.PushRef(vm.GetClassRef(getPrimitiveClassByName(name)))
	return nil
}

func getPrimitiveClassByName(name string) *jvm.Class {
	switch name {
	case "void":
		return jvm.VoidClass
	case "boolean":
		return jvm.BooleanClass
	case "char":
		return jvm.CharClass
	case "byte":
		return jvm.ByteClass
	case "short":
		return jvm.ShortClass
	case "int":
		return jvm.IntClass
	case "long":
		return jvm.LongClass
	case "float":
		return jvm.FloatClass
	case "double":
		return jvm.DoubleClass
	default:
		panic("Unexpected name \"" + name + "\"")
	}
}

// private native String getGenericSignature0();
func Class_getGenericSignature0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// native byte[] getRawAnnotations();
func Class_getRawAnnotations(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// native byte[] getRawTypeAnnotations();
func Class_getRawTypeAnnotations(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// native ConstantPool getConstantPool();
func Class_getConstantPool(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Field[] getDeclaredFields0(boolean publicOnly);
func Class_getDeclaredFields0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Method[] getDeclaredMethods0(boolean publicOnly);
func Class_getDeclaredMethods0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
func Class_getDeclaredConstructors0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Class<?>[] getDeclaredClasses0();
func Class_getDeclaredClasses0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native RecordComponent[] getRecordComponents0();
func Class_getRecordComponents0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native boolean isRecord0();
func Class_isRecord0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func Class_desiredAssertionStatus0(vm ir.VM) error {
	stack := vm.GetStack()
	class := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = class
	stack.PushInt32(0)
	return nil
}

// private native Class<?> getNestHost0();
func Class_getNestHost0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Class<?>[] getNestMembers0();
func Class_getNestMembers0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// public native boolean isHidden();
func Class_isHidden(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native Class<?>[] getPermittedSubclasses0();
func Class_getPermittedSubclasses0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native int getClassFileVersion0();
func Class_getClassFileVersion0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}

// private native int getClassAccessFlagsRaw0();
func Class_getClassAccessFlagsRaw0(vm ir.VM) error {
	stack := vm.GetStack()
	this := (*stack.GetVarRef(0).UserData()).(ir.Class)
	_ = this
	return nil
}
