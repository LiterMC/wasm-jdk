package ir

import (
	"unsafe"
)

func arrayToRef[T any](arr []T) Ref {
	return *((*Ref)((unsafe.Pointer)(&arr)))
}

type sliceHeader struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func arrayLength(ref Ref) int {
	return (*sliceHeader)((unsafe.Pointer)(&ref)).len
}
