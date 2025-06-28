package java_lang_invoke

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/invoke/MethodHandleNatives.registerNatives()V", MethodHandleNatives_registerNatives)
}

// private static native void registerNatives();
func MethodHandleNatives_registerNatives(vm ir.VM) error {
	// native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.", MethodHandleNatives_)
	native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.init(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V", MethodHandleNatives_init)
	return nil
}

// static native void init(MemberName self, Object ref);
func MethodHandleNatives_init(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	ref := stack.GetVarRef(1)
	refClassName := ref.Class().Name()
	if refClassName == "java/lang/reflect/Field" {
		panic("TODO")
	}
	if refClassName == "java/lang/reflect/Method" {
		//
	} else if refClassName == "java/lang/reflect/Constructor" {
		panic("TODO")
	}
	_ = self
	return nil
}

// static native void expand(MemberName self);
// static native MemberName resolve(MemberName self, Class<?> caller, int lookupMode, boolean speculativeResolve) throws LinkageError, ClassNotFoundException;

// static native long objectFieldOffset(MemberName self);  // e.g., returns vmindex
func MethodHandleNatives_objectFieldOffset(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	udata := self.UserData()
	if *udata == nil {
		data := (*MemberName)(self.Data())
		class := vm.GetClass(data.Clazz)
		name := vm.GetString(data.Name)
		field := class.GetFieldByName(name)
		if field.IsStatic() {
			panic("field is static")
		}
		*udata = &MemberNameData{
			VMIndex: field.Offset(),
		}
	}
	vmindex := (*udata).(*MemberNameData).VMIndex
	stack.PushInt64(vmindex)
	return nil
}

// static native long staticFieldOffset(MemberName self);  // e.g., returns vmindex
func MethodHandleNatives_staticFieldOffset(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	udata := self.UserData()
	if *udata == nil {
		data := (*MemberName)(self.Data())
		class := vm.GetClass(data.Clazz)
		name := vm.GetString(data.Name)
		field := class.GetFieldByName(name)
		if !field.IsStatic() {
			panic("field is not static")
		}
		*udata = &MemberNameData{
			VMIndex: field.Offset(),
		}
	}
	vmindex := (*udata).(*MemberNameData).VMIndex
	stack.PushInt64(vmindex)
	return nil
}

// static native Object staticFieldBase(MemberName self);  // e.g., returns clazz
func MethodHandleNatives_staticFieldBase(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	data := (*MemberName)(self.Data())
	stack.PushRef(data.Clazz)
	return nil
}

// static native Object getMemberVMInfo(MemberName self);  // returns {vmindex,vmtarget}
// static native void setCallSiteTargetNormal(CallSite site, MethodHandle target);
// static native void setCallSiteTargetVolatile(CallSite site, MethodHandle target);
// static native void copyOutBootstrapArguments(Class<?> caller, int[] indexInfo, int start, int end, Object[] buf, int pos, boolean resolve, Object ifNotAvailable);
// private static native void clearCallSiteContext(CallSiteContext context);
// private static native int getNamedCon(int which, Object[] name);
