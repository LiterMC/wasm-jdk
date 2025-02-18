package cutil

import (
	"unsafe"
)

func GoString(address int64) string {
	ptr := unsafe.Pointer((uintptr)(address))
	leng := 0
	for *(*byte)(unsafe.Add(ptr, leng)) != 0 {
		leng++
	}
	bytes := unsafe.Slice((*byte)(ptr), leng)
	return (string)(bytes)
}
