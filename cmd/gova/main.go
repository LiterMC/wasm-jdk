package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/LiterMC/wasm-jdk/classloader"
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

var stringArrayDesc = &desc.Desc{
	ArrDim:  1,
	EndType: desc.Class,
	Class:   "java/lang/String",
}

func main() {
	class := os.Args[1]
	method := "main"
	class = strings.ReplaceAll(class, ".", "/")
	ws, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("EntryClass:", class)
	fmt.Println("EntryMethod:", method)
	cl := classloader.NewBasicFSClassLoader(os.DirFS(ws))
	vm := jvm.NewVM(&jvm.Options{
		Loader:      cl,
		EntryClass:  class,
		EntryMethod: "main([Ljava/lang/String;)V",
	})
	{
		arr := vm.NewArray(stringArrayDesc, (int32)(len(os.Args)-2))
		refs := arr.GetArrRef()
		for i, arg := range os.Args[2:] {
			refs[i] = vm.NewString(arg)
		}
		vm.GetStack().SetVarRef(0, arr)
	}
	fmt.Println("Loading native library ...")
	loadSystem(cl, vm)
	fmt.Println("Running ...")
	for vm.Running() {
		fmt.Println("step")
		if err := vm.Step(); err != nil {
			fmt.Println("VM error:", err)
			break
		}
	}
}

func loadSystem(cl ir.ClassLoader, vm *jvm.VM) {
	loadNativeMethod(cl, vm, "java/lang/System.registerNatives()V", func(vm ir.VM) error {
		fmt.Println("registering system natives")
		return nil
	})
}

func loadNativeMethod(cl ir.ClassLoader, vm *jvm.VM, location string, native jvm.NativeMethodCallback) {
	cls, location, ok := strings.Cut(location, ".")
	if !ok {
		panic("no class name in location")
	}
	class, err := cl.LoadClass(cls)
	if err != nil {
		panic(err)
	}
	fmt.Println("getting method:", location)
	method := class.GetMethodByName(location)
	vm.LoadNativeMethod(method, native)
}
