package native

import (
	"strings"

	"github.com/LiterMC/wasm-jdk/ir"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

var defaultNatives = make(map[string]jvm.NativeMethodCallback)

func loadNative(vm ir.VM, cl ir.ClassLoader, location string, callback jvm.NativeMethodCallback) {
	cls, name, ok := strings.Cut(location, ".")
	if !ok {
		panic("no class name in location " + location)
	}
	class, err := cl.LoadClass(cls)
	if err != nil {
		panic("cannot load class " + cls + ": " + err.Error())
	}
	method := class.GetMethodByName(name)
	if method == nil {
		panic("method " + location + " is not found")
	}
	vm.LoadNativeMethod(method, callback)
}

func LoadNative(vm ir.VM, location string, callback jvm.NativeMethodCallback) {
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
		panic("method " + location + " is not found")
	}
	vm.LoadNativeMethod(method, callback)
}

func LoadDefaultNatives(vm ir.VM, cl ir.ClassLoader) {
	for loc, cb := range defaultNatives {
		loadNative(vm, cl, loc, cb)
	}
}

func RegisterDefaultNative(location string, callback jvm.NativeMethodCallback) {
	if _, ok := defaultNatives[location]; ok {
		panic("method " + location + " is already registered")
	}
	defaultNatives[location] = callback
}
