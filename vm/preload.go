package vm

import (
	"github.com/LiterMC/wasm-jdk/ir"
)

type preloadClasses struct {
	javaLangObject          *Class
	javaLangObject_toString ir.Method

	javaLangThrowable               *Class
	javaLangThrowable_backtrace     ir.Field
	javaLangThrowable_detailMessage ir.Field

	javaLangString       *Class
	javaLangString_value ir.Field

	javaLangClass               *Class
	javaLangClass_classLoader   ir.Field
	javaLangClass_componentType ir.Field

	javaLangClassLoader *Class

	javaLangCloneable *Class

	javaLangThread             *Class
	javaLangThread_interrupted ir.Field

	javaLangThreadGroup *Class

	javaLangSystem            *Class
	javaLangSystem_initPhase1 ir.Method
	javaLangSystem_initPhase2 ir.Method
	javaLangSystem_initPhase3 ir.Method

	javaLangRefFinalizer *Class

	javaLangReflectConstructor      *Class
	javaLangReflectConstructor_init ir.Method

	javaLangReflectField      *Class
	javaLangReflectField_init ir.Method

	javaLangReflectMethod      *Class
	javaLangReflectMethod_init ir.Method

	javaLangInvokeMethodHandlesLookup              *Class
	javaLangInvokeMethodHandlesLookup_lookupClass  ir.Field
	javaLangInvokeMethodHandlesLookup_allowedModes ir.Field

	javaLangInvokeDirectMethodHandle      *Class
	javaLangInvokeDirectMethodHandle_make ir.Method

	javaLangInvokeMethodType        *Class
	javaLangInvokeMethodType_rtype  ir.Field
	javaLangInvokeMethodType_ptypes ir.Field

	javaLangInvokeMemberName      *Class
	javaLangInvokeMemberName_init ir.Method

	jdkInternalReflectConstantPool *Class
}

func (p *preloadClasses) load(vm *VM) {
	var err error
	if p.javaLangObject, err = vm.loadClass("java/lang/Object"); err != nil {
		panic(err)
	}
	p.javaLangObject_toString = assertNotNil(p.javaLangObject.GetMethodByNameAndType("toString", "()Ljava/lang/String;"))

	if p.javaLangString, err = vm.loadClass("java/lang/String"); err != nil {
		panic(err)
	}
	p.javaLangString_value = assertNotNil(p.javaLangString.GetFieldByName("value"))

	if p.javaLangSystem, err = vm.loadClass("java/lang/System"); err != nil {
		panic(err)
	}
	p.javaLangSystem_initPhase1 = assertNotNil(p.javaLangSystem.GetMethodByNameAndType("initPhase1", "()V"))
	p.javaLangSystem_initPhase2 = assertNotNil(p.javaLangSystem.GetMethodByNameAndType("initPhase2", "(ZZ)I"))
	p.javaLangSystem_initPhase3 = assertNotNil(p.javaLangSystem.GetMethodByNameAndType("initPhase3", "()V"))

	if p.javaLangClass, err = vm.loadClass("java/lang/Class"); err != nil {
		panic(err)
	}
	p.javaLangClass_classLoader = assertNotNil(p.javaLangClass.GetFieldByName("classLoader"))
	p.javaLangClass_componentType = assertNotNil(p.javaLangClass.GetFieldByName("componentType"))

	if p.javaLangClassLoader, err = vm.loadClass("java/lang/ClassLoader"); err != nil {
		panic(err)
	}

	if p.javaLangCloneable, err = vm.loadClass("java/lang/Cloneable"); err != nil {
		panic(err)
	}

	if p.javaLangThread, err = vm.loadClass("java/lang/Thread"); err != nil {
		panic(err)
	}
	p.javaLangThread_interrupted = assertNotNil(p.javaLangThread.GetFieldByName("interrupted"))

	if p.javaLangThreadGroup, err = vm.loadClass("java/lang/ThreadGroup"); err != nil {
		panic(err)
	}

	if p.javaLangThrowable, err = vm.loadClass("java/lang/Throwable"); err != nil {
		panic(err)
	}
	p.javaLangThrowable_backtrace = assertNotNil(p.javaLangThrowable.GetFieldByName("backtrace"))
	p.javaLangThrowable_detailMessage = assertNotNil(p.javaLangThrowable.GetFieldByName("detailMessage"))

	if p.javaLangRefFinalizer, err = vm.loadClass("java/lang/ref/Finalizer"); err != nil {
		panic(err)
	}

	if p.javaLangReflectConstructor, err = vm.loadClass("java/lang/reflect/Constructor"); err != nil {
		panic(err)
	}
	p.javaLangReflectConstructor_init = assertNotNil(p.javaLangReflectConstructor.GetMethodByNameAndType("<init>", "(Ljava/lang/Class;[Ljava/lang/Class;[Ljava/lang/Class;IILjava/lang/String;[B[B)V"))

	if p.javaLangReflectField, err = vm.loadClass("java/lang/reflect/Field"); err != nil {
		panic(err)
	}
	p.javaLangReflectField_init = assertNotNil(p.javaLangReflectField.GetMethodByNameAndType("<init>", "(Ljava/lang/Class;Ljava/lang/String;Ljava/lang/Class;IZILjava/lang/String;[B)V"))

	if p.javaLangReflectMethod, err = vm.loadClass("java/lang/reflect/Method"); err != nil {
		panic(err)
	}
	p.javaLangReflectMethod_init = assertNotNil(p.javaLangReflectMethod.GetMethodByNameAndType("<init>", "(Ljava/lang/Class;Ljava/lang/String;[Ljava/lang/Class;Ljava/lang/Class;[Ljava/lang/Class;IILjava/lang/String;[B[B[B)V"))

	if p.javaLangInvokeMethodHandlesLookup, err = vm.loadClass("java/lang/invoke/MethodHandles$Lookup"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodHandlesLookup_lookupClass = assertNotNil(p.javaLangInvokeMethodHandlesLookup.GetFieldByName("lookupClass"))
	p.javaLangInvokeMethodHandlesLookup_allowedModes = assertNotNil(p.javaLangInvokeMethodHandlesLookup.GetFieldByName("allowedModes"))

	if p.javaLangInvokeDirectMethodHandle, err = vm.loadClass("java/lang/invoke/DirectMethodHandle"); err != nil {
		panic(err)
	}
	p.javaLangInvokeDirectMethodHandle_make = assertNotNil(p.javaLangInvokeDirectMethodHandle.GetMethodByNameAndType("make", "(Ljava/lang/invoke/MemberName;)Ljava/lang/invoke/DirectMethodHandle;"))

	if p.javaLangInvokeMethodType, err = vm.loadClass("java/lang/invoke/MethodType"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMethodType_rtype = assertNotNil(p.javaLangInvokeMethodType.GetFieldByName("rtype"))
	p.javaLangInvokeMethodType_ptypes = assertNotNil(p.javaLangInvokeMethodType.GetFieldByName("ptypes"))

	if p.javaLangInvokeMemberName, err = vm.loadClass("java/lang/invoke/MemberName"); err != nil {
		panic(err)
	}
	p.javaLangInvokeMemberName_init = assertNotNil(p.javaLangInvokeMemberName.GetMethodByNameAndType("<init>", "(Ljava/lang/reflect/Method;)V"))

	if p.jdkInternalReflectConstantPool, err = vm.loadClass("jdk/internal/reflect/ConstantPool"); err != nil {
		panic(err)
	}
}

func assertNotNil[T any](v T) T {
	if (any)(v) == nil {
		panic("unexpected nil")
	}
	return v
}

func (p *preloadClasses) JClass_JavaLangCloneable() ir.Class {
	return p.javaLangCloneable
}

func (p *preloadClasses) JClass_JavaLangReflectMethod() ir.Class {
	return p.javaLangReflectMethod
}
