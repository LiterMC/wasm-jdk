package classloader

import (
	"io/fs"
	"sync"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
	"github.com/LiterMC/wasm-jdk/vm"
)

type BasicFSClassLoader struct {
	fs     fs.FS
	loaded sync.Map
}

func NewBasicFSClassLoader(fs fs.FS) ir.ClassLoader {
	return &BasicFSClassLoader{
		fs: fs,
	}
}

func (l *BasicFSClassLoader) LoadClass(name string) (ir.Class, error) {
	loader, ok := l.loaded.Load(name)
	if !ok {
		loader, _ = l.loaded.LoadOrStore(name, sync.OnceValues(func() (*vm.Class, error) {
			return loadClassFromFS(l, l.fs, name)
		}))
	}
	class, err := loader.(func() (*vm.Class, error))()
	if err != nil {
		return nil, err
	}
	return class, nil
}

func loadClassFromFS(l ir.ClassLoader, fs fs.FS, name string) (*vm.Class, error) {
	println("loading ", name)
	fd, err := fs.Open(name + ".class")
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	cls, err := jcls.ParseClass(fd)
	if err != nil {
		return nil, err
	}
	class := vm.LoadClass(cls, l)
	return class, nil
}
