package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/LiterMC/wasm-jdk/classloader"
	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/properties"
	jvm "github.com/LiterMC/wasm-jdk/vm"

	"github.com/LiterMC/wasm-jdk/native"
	_ "github.com/LiterMC/wasm-jdk/native/init_all"
	misc "github.com/LiterMC/wasm-jdk/native/jdk/internal_/misc"
)

func main() {
	class := os.Args[1]
	method := "main"
	class = strings.ReplaceAll(class, ".", "/")

	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	properties.SetProp("java.home", workingDir)

	fmt.Println("EntryClass:", class)
	fmt.Println("EntryMethod:", method)

	modulesDir := filepath.Join(workingDir, "modules")

	cl := classloader.WrapAsSyncedClassLoader(classloader.NewMultiClassLoader(
		classloader.NewExplodeModuleClassLoader(os.DirFS(modulesDir), "file://"+modulesDir),
		classloader.NewBasicFSClassLoader(os.DirFS(workingDir), "file://"+workingDir),
	))
	vm := jvm.NewVM(&jvm.Options{
		Loader:      cl,
		EntryClass:  class,
		EntryMethod: "main([Ljava/lang/String;)V",
	})

	fmt.Println("Loading native library ...")
	misc.InitUnsafeConstants(vm)
	native.LoadDefaultNatives(vm, cl)
	vm.SetupEntryMethod()

	{
		arr := vm.NewArray(desc.DescStringArray, (int32)(len(os.Args)-2))
		refs := arr.GetRefArr()
		for i, arg := range os.Args[2:] {
			refs[i] = vm.RefToPtr(vm.NewString(arg))
		}
		vm.GetStack().SetVarRef(0, arr)
	}
	fmt.Println("Running ...")
	for vm.Running() {
		if err := vm.Step(); err != nil {
			fmt.Println("VM error:", err)
			break
		}
	}
}
