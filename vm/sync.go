package vm

import (
	"sync"
	"sync/atomic"
)

func OnceApply[T1, T2 any](f func(T1) T2) func(T1) T2 {
	var (
		done    atomic.Bool
		mutex   sync.Mutex
		succeed bool
		err     any
		result  T2
	)
	slow := func(input T1) {
		mutex.Lock()
		defer mutex.Unlock()
		if done.Load() {
			return
		}
		defer func() {
			f = nil
			if err = recover(); !succeed {
				panic(err)
			}
		}()
		defer done.Store(true)
		result = f(input)
		succeed = true
	}
	return func(input T1) T2 {
		if !done.Load() {
			slow(input)
		}
		if !succeed {
			panic(err)
		}
		return result
	}
}
