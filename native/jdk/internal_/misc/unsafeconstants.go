package jdk_internal_misc

import (
	"unsafe"

	"github.com/LiterMC/wasm-jdk/ir"
)

var (
	AddressSize            int32 = (int32)(unsafe.Sizeof((unsafe.Pointer)(nil)))
	PageSize               int32 = 65536
	BigEndian              bool
	UnalignedAccess        bool  = false
	DataCacheLineFlushSize int32 = 0
)

func init() {
	i := (uint16)(0x0124)
	v := (*(*[2]byte)((unsafe.Pointer)(&i)))[0]
	if v == 0x01 {
		BigEndian = true
	} else if v == 0x24 {
		BigEndian = false
	} else {
		panic("unexpected int encoding")
	}
}

func InitUnsafeConstants(vm ir.VM) {
	cls, err := vm.GetClassLoader().LoadClass("jdk/internal/misc/UnsafeConstants")
	if err != nil {
		panic(err)
	}
	addrSizePtr := (*int32)(cls.GetFieldByName("ADDRESS_SIZE0").GetPointer(nil))
	pageSizePtr := (*int32)(cls.GetFieldByName("PAGE_SIZE").GetPointer(nil))
	bigEndianPtr := (*int32)(cls.GetFieldByName("BIG_ENDIAN").GetPointer(nil))
	unalignedAccessPtr := (*int32)(cls.GetFieldByName("UNALIGNED_ACCESS").GetPointer(nil))
	dataCacheLineFlushSizePtr := (*int32)(cls.GetFieldByName("DATA_CACHE_LINE_FLUSH_SIZE").GetPointer(nil))
	*addrSizePtr = AddressSize
	*pageSizePtr = PageSize
	if BigEndian {
		*bigEndianPtr = 1
	} else {
		*bigEndianPtr = 0
	}
	if UnalignedAccess {
		*unalignedAccessPtr = 1
	} else {
		*unalignedAccessPtr = 0
	}
	*dataCacheLineFlushSizePtr = DataCacheLineFlushSize
}
