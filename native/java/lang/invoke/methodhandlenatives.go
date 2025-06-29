package java_lang_invoke

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
	"github.com/LiterMC/wasm-jdk/native/helper"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

func init() {
	native.RegisterDefaultNative("java/lang/invoke/MethodHandleNatives.registerNatives()V", MethodHandleNatives_registerNatives)
}

const (
	MN_IS_METHOD            = 0x00010000 // method (not constructor)
	MN_IS_CONSTRUCTOR       = 0x00020000 // constructor
	MN_IS_FIELD             = 0x00040000 // field
	MN_IS_TYPE              = 0x00080000 // nested type
	MN_CALLER_SENSITIVE     = 0x00100000 // @CallerSensitive annotation detected
	MN_TRUSTED_FINAL        = 0x00200000 // trusted final field
	MN_REFERENCE_KIND_SHIFT = 24         // refKind
	MN_REFERENCE_KIND_MASK  = 0x0F000000 >> MN_REFERENCE_KIND_SHIFT
)

const (
	REF_NONE             = 0 // null value
	REF_getField         = 1
	REF_getStatic        = 2
	REF_putField         = 3
	REF_putStatic        = 4
	REF_invokeVirtual    = 5
	REF_invokeStatic     = 6
	REF_invokeSpecial    = 7
	REF_newInvokeSpecial = 8
	REF_invokeInterface  = 9
)

// private static native void registerNatives();
func MethodHandleNatives_registerNatives(vm ir.VM) error {
	// native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.", MethodHandleNatives_)
	native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.init(Ljava/lang/invoke/MemberName;Ljava/lang/Object;)V", MethodHandleNatives_init)
	native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.expand(Ljava/lang/invoke/MemberName;)V", MethodHandleNatives_expand)
	native.LoadNative(vm, "java/lang/invoke/MethodHandleNatives.resolve(Ljava/lang/invoke/MemberName;Ljava/lang/Class;IZ)Ljava/lang/invoke/MemberName;", MethodHandleNatives_resolve)
	return nil
}

// static native void init(MemberName self, Object ref);
func MethodHandleNatives_init(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	ref := stack.GetVarRef(1)
	refClassName := ref.Class().Name()
	data := (*MemberName)(self.Data())
	vmHelper := vm.(helper.VMHelper)
	if refClassName == "java/lang/reflect/Field" {
		panic("TODO")
	}
	if refClassName == "java/lang/reflect/Method" {
		data.Clazz = *(**jvm.Ref)(vmHelper.JField_javaLangReflectMethod_clazz().GetPointer(ref))
		flags := *(*int32)(vmHelper.JField_javaLangReflectMethod_modifiers().GetPointer(ref))
		flags |= MN_IS_METHOD
		flags |= REF_invokeVirtual << MN_REFERENCE_KIND_SHIFT
		data.Flags = flags
	} else if refClassName == "java/lang/reflect/Constructor" {
		panic("TODO")
	}
	_ = self
	return nil
}

// static native void expand(MemberName self);
func MethodHandleNatives_expand(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	if true {
		panic("TODO")
	}
	_ = self
	return nil
}

// static native MemberName resolve(MemberName self, Class<?> caller, int lookupMode, boolean speculativeResolve) throws LinkageError, ClassNotFoundException;
func MethodHandleNatives_resolve(vm ir.VM) error {
	stack := vm.GetStack()
	self := stack.GetVarRef(0)
	// caller := stack.GetVarRef(1)
	// lookupMode := stack.GetVarInt32(2)
	// speculativeResolve := stack.GetVar(3) != 0
	stack.PushRef(self)
	return nil
}

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
