package classloader

import (
	"sync"

	"github.com/LiterMC/wasm-jdk/ir"
)

type BasicClassLoader interface {
	LoadClass(loader ir.ClassLoader, name string) (ir.Class, error)
	AvaliablePackages() []string
	PackageLocation(name string) string
}

type BasicSyncedClassLoader struct {
	loader BasicClassLoader
	loaded sync.Map
}

func WrapAsSyncedClassLoader(loader BasicClassLoader) ir.ClassLoader {
	return &BasicSyncedClassLoader{
		loader: loader,
	}
}

func (l *BasicSyncedClassLoader) LoadClass(name string) (ir.Class, error) {
	loader, ok := l.loaded.Load(name)
	if !ok {
		loader, _ = l.loaded.LoadOrStore(name, sync.OnceValues(func() (ir.Class, error) {
			return l.loader.LoadClass(l, name)
		}))
	}
	class, err := loader.(func() (ir.Class, error))()
	if err != nil {
		return nil, err
	}
	return class, nil
}

func (l *BasicSyncedClassLoader) LoadedClass(name string) ir.Class {
	loader, ok := l.loaded.Load(name)
	if !ok {
		return nil
	}
	class, err := loader.(func() (ir.Class, error))()
	if err != nil {
		return nil
	}
	return class
}

func (l *BasicSyncedClassLoader) AvaliablePackages() []string {
	return l.loader.AvaliablePackages()
}

func (l *BasicSyncedClassLoader) PackageLocation(name string) string {
	return l.loader.PackageLocation(name)
}
