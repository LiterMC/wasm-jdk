package classloader

import (
	"errors"
	"io/fs"
	stdpath "path"
	"strings"
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

func (l *BasicFSClassLoader) AvaliablePackages() []string {
	packages := make([]string, 0, 10)
	WalkDir(l.fs, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".class") {
			packages = append(packages, stdpath.Dir(path))
			return SkipFiles
		}
		return nil
	})
	return packages
}

func (l *BasicFSClassLoader) PackageLocation(name string) string {
	stat, err := fs.Stat(l.fs, name)
	if err != nil {
		return ""
	}
	if stat.IsDir() {
		return "file://" + name
	}
	return ""
}

func loadClassFromFS(l ir.ClassLoader, fs fs.FS, name string) (*vm.Class, error) {
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

var SkipFiles = errors.New("SkipFiles")

func WalkDir(fsys fs.FS, path string, walker fs.WalkDirFunc) error {
	stat, err := fs.Stat(fsys, path)
	if err != nil {
		err = walker(path, nil, err)
	} else {
		err = walkDir(fsys, path, fs.FileInfoToDirEntry(stat), walker)
	}
	if err == fs.SkipDir || err == fs.SkipAll {
		return nil
	}
	return err
}

func walkDir(fsys fs.FS, path string, entry fs.DirEntry, walker fs.WalkDirFunc) error {
	if err := walker(path, entry, nil); err != nil {
		if err == fs.SkipDir && entry.IsDir() {
			return nil
		}
		return err
	}
	entries, err := fs.ReadDir(fsys, path)
	if err != nil {
		if err = walker(path, entry, err); err == fs.SkipDir && entry.IsDir() {
			return nil
		}
		return err
	}
	walkFiles := true
	for _, e := range entries {
		if e.IsDir() || !walkFiles {
			err = walkDir(fsys, stdpath.Join(path, e.Name()), e, walker)
			if err == fs.SkipDir {
				break
			}
			if err == SkipFiles {
				walkFiles = false
				continue
			}
			return err
		}
	}
	return nil
}
