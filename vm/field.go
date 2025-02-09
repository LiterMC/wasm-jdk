package vm

import (
	"sync/atomic"
	"unsafe"

	"github.com/LiterMC/wasm-jdk/desc"
	"github.com/LiterMC/wasm-jdk/errs"
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

func (f *Field) Offset() int64 {
	return (int64)(f.offset)
}

func (f *Field) Type() ir.Class {
	return f.typ
}

func (f *Field) GetDeclaringClass() ir.Class {
	return f.class
}

func (f *Field) GetPointer(r ir.Ref) unsafe.Pointer {
	if f.IsStatic() {
		return unsafe.Add(f.class.staticData, f.offset)
	}
	return unsafe.Add(r.Data(), f.offset)
}

func (f *Field) GetAndPush(s ir.Stack) error {
	var r ir.Ref
	if !f.IsStatic() {
		if r = s.PopRef(); r == nil {
			return errs.NullPointerException
		}
	}
	ptr := f.GetPointer(r)
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
	return nil
}

func (f *Field) getPtrFromStack(s ir.Stack) unsafe.Pointer {
	if f.IsStatic() {
		return unsafe.Add(f.class.staticData, f.offset)
	}
	r := s.PopRef()
	if r == nil {
		return nil
	}
	return unsafe.Add(r.Data(), f.offset)
}

func (f *Field) PopAndSet(s ir.Stack) error {
	switch f.Desc.Type() {
	case desc.Class, desc.Array:
		v := s.PopRef()
		var w *Ref
		if v != nil {
			w = v.(*Ref)
		}
		ptr := f.getPtrFromStack(s)
		if ptr == nil {
			return errs.NullPointerException
		}
		atomic.StorePointer((*unsafe.Pointer)(ptr), (unsafe.Pointer)(w))
	case desc.Boolean, desc.Byte:
		v := (int32)(s.PopInt8())
		ptr := f.getPtrFromStack(s)
		if ptr == nil {
			return errs.NullPointerException
		}
		atomic.StoreInt32((*int32)(ptr), v)
	case desc.Char, desc.Short:
		v := (int32)(s.PopInt16())
		ptr := f.getPtrFromStack(s)
		if ptr == nil {
			return errs.NullPointerException
		}
		atomic.StoreInt32((*int32)(ptr), v)
	case desc.Int, desc.Float:
		v := s.PopInt32()
		ptr := f.getPtrFromStack(s)
		if ptr == nil {
			return errs.NullPointerException
		}
		atomic.StoreInt32((*int32)(ptr), v)
	case desc.Long, desc.Double:
		v := s.PopInt64()
		ptr := f.getPtrFromStack(s)
		if ptr == nil {
			return errs.NullPointerException
		}
		atomic.StoreInt64((*int64)(ptr), v)
	default:
		panic("unreachable")
	}
	return nil
}
