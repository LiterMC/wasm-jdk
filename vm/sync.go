package vm

import (
	"sync"
	"sync/atomic"
)

func OnceApply[T1, T2 any](f func(T1) T2) func(T1) T2 {
	var (
		once   sync.Once
		input  atomic.Pointer[T1]
		done   atomic.Bool
		valid  bool
		p      any
		result T2
	)
	g := func() {
		defer func() {
			p = recover()
			if !valid {
				panic(p)
			}
			done.Store(true)
		}()
		result = f(*input.Load())
		f = nil
		valid = true
	}
	return func(in T1) T2 {
		if !done.Load() {
			input.Store(&in)
			once.Do(g)
			input.Store(nil)
		}
		if !valid {
			panic(p)
		}
		return result
	}
}
