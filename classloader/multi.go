package classloader

import (
	"errors"
	"io/fs"

	"github.com/LiterMC/wasm-jdk/ir"
)

type MultiClassLoader struct {
	loaders []BasicClassLoader
}

func NewMultiClassLoader(loaders ...BasicClassLoader) BasicClassLoader {
	if len(loaders) == 0 {
		panic("MultiClassLoader: need provide at least one loader")
	}
	var (
		hasMultiLoader bool
		totalLoaders   int
	)
	for _, l := range loaders {
		if ml, ok := l.(*MultiClassLoader); ok {
			hasMultiLoader = true
			totalLoaders += len(ml.loaders)
		} else {
			totalLoaders++
		}
	}
	if hasMultiLoader {
		oldLoaders := loaders
		loaders = make([]BasicClassLoader, 0, totalLoaders)
		for _, l := range oldLoaders {
			if ml, ok := l.(*MultiClassLoader); ok {
				loaders = append(loaders, ml.loaders...)
			}
		}
	}
	return &MultiClassLoader{
		loaders: loaders,
	}
}

func (l *MultiClassLoader) LoadClass(loader ir.ClassLoader, name string) (ir.Class, error) {
	var lazyErr error
	for _, ldr := range l.loaders {
		cls, err := ldr.LoadClass(loader, name)
		if err == nil {
			return cls, nil
		}
		if lazyErr == nil && !errors.Is(err, fs.ErrNotExist) {
			lazyErr = err
		}
	}
	if lazyErr != nil {
		return nil, lazyErr
	}
	return nil, fs.ErrNotExist
}

func (l *MultiClassLoader) AvaliablePackages() []string {
	packages := l.loaders[0].AvaliablePackages()
	for _, loader := range l.loaders[1:] {
		packages = append(packages, loader.AvaliablePackages()...)
	}
	return packages
}

func (l *MultiClassLoader) PackageLocation(name string) string {
	for _, loader := range l.loaders {
		if loc := loader.PackageLocation(name); loc != "" {
			return loc
		}
	}
	return ""
}
