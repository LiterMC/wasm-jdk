package classloader

import (
	"errors"
	"io/fs"
	stdpath "path"
	"strings"

	"github.com/LiterMC/wasm-jdk/ir"
)

type ExplodeModuleClassLoader struct {
	fs       fs.FS
	location string
}

func NewExplodeModuleClassLoader(fs fs.FS, location string) BasicClassLoader {
	return &ExplodeModuleClassLoader{
		fs:       fs,
		location: location,
	}
}

func (l *ExplodeModuleClassLoader) LoadClass(loader ir.ClassLoader, name string) (ir.Class, error) {
	entires, err := fs.ReadDir(l.fs, ".")
	if err != nil {
		return nil, err
	}
	var lazyErr error
	for _, e := range entires {
		if e.IsDir() {
			subFS, err := fs.Sub(l.fs, e.Name())
			if err != nil {
				panic(err) // Should never happen
			}
			cls, err := loadClassFromFS(loader, subFS, name)
			if err == nil {
				return cls, nil
			}
			if lazyErr == nil && !errors.Is(err, fs.ErrNotExist) {
				lazyErr = err
			}
		}
	}
	if lazyErr != nil {
		return nil, lazyErr
	}
	return nil, fs.ErrNotExist
}

func (l *ExplodeModuleClassLoader) AvaliablePackages() []string {
	packages := make([]string, 0, 10)
	for _, subFS := range l.forEachModules {
		WalkDir(subFS, ".", func(path string, entry fs.DirEntry, err error) error {
			if err != nil || entry.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".class") {
				packages = append(packages, stdpath.Dir(path))
				return SkipFiles
			}
			return nil
		})
	}
	return packages
}

func (l *ExplodeModuleClassLoader) PackageLocation(name string) string {
	for module, subFS := range l.forEachModules {
		stat, err := fs.Stat(subFS, name)
		if err == nil && stat.IsDir() {
			return l.location + stdpath.Join(module, name)
		}
	}
	return ""
}

func (l *ExplodeModuleClassLoader) forEachModules(iter func(module string, subFS fs.FS) bool) {
	entires, err := fs.ReadDir(l.fs, ".")
	if err != nil {
		return
	}
	for _, e := range entires {
		if !e.IsDir() {
			continue
		}
		subFS, err := fs.Sub(l.fs, e.Name())
		if err != nil {
			continue
		}
		if !iter(e.Name(), subFS) {
			break
		}
	}
}
