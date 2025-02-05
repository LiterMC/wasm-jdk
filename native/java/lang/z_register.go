package java_lang

import (
	"strings"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

func loadNative(vm ir.VM, location string, callback jvm.NativeMethodCallback) {
	cls, name, ok := strings.Cut(location, ".")
	if !ok {
		panic("no class name in location " + location)
	}
	class, err := vm.GetClassLoader().LoadClass(cls)
	if err != nil {
		panic("cannot load class " + cls + ": " + err.Error())
	}
	method := class.GetMethodByName(name)
	if method == nil {
		panic("method " + location + "is not found")
	}
	vm.LoadNativeMethod(method, callback)
}

func registerDefaultNative(location string, callback jvm.NativeMethodCallback) {
	native.RegisterDefaultNative(location, callback)
}
