package java_security

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/security/AccessController.getProtectionDomain(Ljava/lang/Class;)Ljava/security/ProtectionDomain;", AccessController_getProtectionDomain)
	native.RegisterDefaultNative("java/security/AccessController.ensureMaterializedForStackWalk(Ljava/lang/Object;)Z", AccessController_ensureMaterializedForStackWalk)
	native.RegisterDefaultNative("java/security/AccessController.getStackAccessControlContext()Ljava/security/AccessControlContext;", AccessController_getStackAccessControlContext)
	native.RegisterDefaultNative("java/security/AccessController.getInheritedAccessControlContext()Ljava/security/AccessControlContext;", AccessController_getInheritedAccessControlContext)
}

// private static native ProtectionDomain getProtectionDomain(final Class<?> caller);
func AccessController_getProtectionDomain(vm ir.VM) error {
	stack := vm.GetStack()
	caller := stack.GetVarRef(0)
	_ = caller
	stack.PushRef(nil)
	return nil
}

// private static native void ensureMaterializedForStackWalk(Object o);
func AccessController_ensureMaterializedForStackWalk(vm ir.VM) error {
	return nil
}

// private static native AccessControlContext getStackAccessControlContext();
func AccessController_getStackAccessControlContext(vm ir.VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}

// static native AccessControlContext getInheritedAccessControlContext();
func AccessController_getInheritedAccessControlContext(vm ir.VM) error {
	vm.GetStack().PushRef(nil)
	return nil
}
