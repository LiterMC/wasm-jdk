package cutil

import (
	"runtime"
	"sync"
	"unsafe"
)

var (
	pinnerMux   sync.RWMutex
	pinners     = make([]*runtime.Pinner, 0, 7)
	pinnerSlots = make([]int, 0, 3)
)

func allocPinner() (pinner *runtime.Pinner, slot int) {
	pinnerMux.Lock()
	defer pinnerMux.Unlock()

	pinner = new(runtime.Pinner)

	if i := len(pinnerSlots) - 1; i >= 0 {
		pinnerSlots, slot = pinnerSlots[:i], pinnerSlots[i]
		pinners[slot] = pinner
	} else {
		slot = len(pinners)
		pinners = append(pinners, pinner)
	}
	return
}

func getPinner(slot int) *runtime.Pinner {
	pinnerMux.RLock()
	defer pinnerMux.RUnlock()
	return pinners[slot]
}

func freePinner(slot int) {
	pinnerMux.Lock()
	defer pinnerMux.Unlock()

	var pinner *runtime.Pinner
	pinner, pinners[slot] = pinners[slot], nil
	pinnerSlots = append(pinnerSlots, slot)

	pinner.Unpin()
}

type pinnedMemoryHeader struct {
	pinnerInd int
	length    int
	data      [0]byte
}

const memoryHeaderOffset = (int)(unsafe.Offsetof(pinnedMemoryHeader{}.data))

func AllocMemory(length int) uintptr {
	ptr := (unsafe.Pointer)(unsafe.SliceData(make([]byte, length+memoryHeaderOffset)))
	header := (*pinnedMemoryHeader)(ptr)
	header.length = length

	var pinner *runtime.Pinner
	pinner, header.pinnerInd = allocPinner()
	pinner.Pin(header)

	println("AllocMemory:", ptr, (uintptr)(unsafe.Add(ptr, memoryHeaderOffset)))
	return (uintptr)(unsafe.Add(ptr, memoryHeaderOffset))
}

func ReallocMemory(address uintptr, length int) uintptr {
	ptr := unsafe.Add((unsafe.Pointer)(address), -memoryHeaderOffset)
	pinner := getPinner((*pinnedMemoryHeader)(ptr).pinnerInd)
	dataLength := (*pinnedMemoryHeader)(ptr).length
	need := length - dataLength
	if need <= 0 {
		return address
	}
	slice := unsafe.Slice((*byte)(ptr), dataLength)
	slice = append(slice, make([]byte, need)...)
	ptr = (unsafe.Pointer)(unsafe.SliceData(slice))
	pinner.Unpin()
	pinner.Pin((*pinnedMemoryHeader)(ptr))
	return (uintptr)(unsafe.Add(ptr, memoryHeaderOffset))
}

func FreeMemory(address uintptr) {
	ptr := unsafe.Add((unsafe.Pointer)(address), -memoryHeaderOffset)
	header := (*pinnedMemoryHeader)(ptr)
	freePinner(header.pinnerInd)
}
