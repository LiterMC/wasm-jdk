package vm

import (
	"sync/atomic"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/jcls"
)

type Field struct {
	*jcls.Field
	offset uintptr
	typ    ir.Class
	class  *Class
}

var _ ir.Field = (*Field)(nil)

func (f *Field) Type() ir.Class {
	return f.typ
}

func (f *Field) GetDeclaringClass() ir.Class {
	return f.class
}

func (f *Field) GetAndPush(r ir.Ref, s ir.Stack) {
	ptr := unsafe.Add(r.Data(), f.offset)
	switch f.Desc.Type() {
	case desc.Class, desc.Array:
		s.PushRef((*Ref)(atomic.LoadPointer((*unsafe.Pointer)(ptr))))
	case desc.Boolean, desc.Byte:
		s.PushInt8((int8)(atomic.LoadInt32((*int32)(ptr))))
	case desc.Char, desc.Short:
		s.PushInt16((int16)(atomic.LoadInt32((*int32)(ptr))))
	case desc.Int, desc.Float:
		s.PushInt32(atomic.LoadInt32((*int32)(ptr)))
	case desc.Long, desc.Double:
		s.PushInt64(atomic.LoadInt64((*int64)(ptr)))
	default:
		panic("unreachable")
	}
}

func (f *Field) PopAndSet(r ir.Ref, s ir.Stack) {
	ptr := unsafe.Add(r.Data(), f.offset)
	switch f.Desc.Type() {
	case desc.Class, desc.Array:
		atomic.StorePointer((*unsafe.Pointer)(ptr), (unsafe.Pointer)(s.PopRef().(*Ref)))
	case desc.Boolean, desc.Byte:
		atomic.StoreInt32((*int32)(ptr), (int32)(s.PopInt8()))
	case desc.Char, desc.Short:
		atomic.StoreInt32((*int32)(ptr), (int32)(s.PopInt16()))
	case desc.Int, desc.Float:
		atomic.StoreInt32((*int32)(ptr), s.PopInt32())
	case desc.Long, desc.Double:
		atomic.StoreInt64((*int64)(ptr), s.PopInt64())
	default:
		panic("unreachable")
	}
}
