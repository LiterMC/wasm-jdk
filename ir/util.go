package ir

import (
	"unsafe"
)

func arrayToRef[T any](arr []T) Ref {
	return (Ref)(unsafe.SliceData(arr))
}

type refHeader struct {
	class     Ref
	arrayKind uint8
	len       int
	data      [0]byte
}

func arrayLength(ref Ref) int {
	return (*refHeader)((Ref)((uintptr)(ref) - unsafe.Offsetof(refHeader{}.data))).len
}
