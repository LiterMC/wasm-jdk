package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/loader/BootLoader.getSystemPackageNames()[Ljava/lang/String;", BootLoader_getSystemPackageNames)
	native.RegisterDefaultNative("jdk/internal/loader/BootLoader.getSystemPackageLocation(Ljava/lang/String;)Ljava/lang/String;", BootLoader_getSystemPackageLocation)
	native.RegisterDefaultNative("jdk/internal/loader/BootLoader.setBootLoaderUnnamedModule0(Ljava/lang/Module;)V", BootLoader_setBootLoaderUnnamedModule0)
}

// private static native String[] getSystemPackageNames();
func BootLoader_getSystemPackageNames(vm ir.VM) error {
	systemPackages := vm.GetBootLoader().AvaliablePackages()
	packagesRef := vm.NewArray(desc.DescStringArray, (int32)(len(systemPackages)))
	packageArr := packagesRef.GetRefArr()
	for i, p := range systemPackages {
		packageArr[i] = vm.RefToPtr(vm.NewString(p))
	}
	vm.GetStack().PushRef(packagesRef)
	return nil
}

// private static native String getSystemPackageLocation(String name);
func BootLoader_getSystemPackageLocation(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	location := vm.GetBootLoader().PackageLocation(name)
	if location == "" {
		stack.PushRef(nil)
	} else {
		stack.PushRef(vm.NewString(location))
	}
	return nil
}

var bootModule ir.Ref

// private static native void setBootLoaderUnnamedModule0(Module module);
func BootLoader_setBootLoaderUnnamedModule0(vm ir.VM) error {
	stack := vm.GetStack()
	module := stack.GetVarRef(0)
	bootModule = module
	return nil
}
