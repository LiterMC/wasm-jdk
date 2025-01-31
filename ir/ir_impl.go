// See: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-6.html#jvms-6.5
package ir

import (
	"fmt"
	"slices"

	"github.com/LiterMC/wasm-jdk/ops"
)

type IRareturn struct{}

func (*IRareturn) Op() ops.Op { return ops.Areturn }
func (*IRareturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRdreturn struct{}

func (*IRdreturn) Op() ops.Op { return ops.Dreturn }
func (*IRdreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRdup struct{}

func (*IRdup) Op() ops.Op { return ops.Dup }
func (*IRdup) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PeekInt32()
	stack.PushInt32(val)
	return nil
}

type IRdup_x1 struct{}

func (*IRdup_x1) Op() ops.Op { return ops.Dup_x1 }
func (*IRdup_x1) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	b := stack.PopInt32()
	stack.PushInt32(a)
	stack.PushInt32(b)
	stack.PushInt32(a)
	return nil
}

type IRdup_x2 struct{}

func (*IRdup_x2) Op() ops.Op { return ops.Dup_x2 }
func (*IRdup_x2) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	b := stack.PopInt64()
	stack.PushInt32(a)
	stack.PushInt64(b)
	stack.PushInt32(a)
	return nil
}

type IRdup2 struct{}

func (*IRdup2) Op() ops.Op { return ops.Dup2 }
func (*IRdup2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PeekInt64()
	stack.PushInt64(val)
	return nil
}

type IRdup2_x1 struct{}

func (*IRdup2_x1) Op() ops.Op { return ops.Dup2_x1 }
func (*IRdup2_x1) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt64()
	b := stack.PopInt32()
	stack.PushInt64(a)
	stack.PushInt32(b)
	stack.PushInt64(a)
	return nil
}

type IRdup2_x2 struct{}

func (*IRdup2_x2) Op() ops.Op { return ops.Dup2_x2 }
func (*IRdup2_x2) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt64()
	b := stack.PopInt64()
	stack.PushInt64(a)
	stack.PushInt64(b)
	stack.PushInt64(a)
	return nil
}

type IRfreturn struct{}

func (*IRfreturn) Op() ops.Op { return ops.Freturn }
func (*IRfreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRgoto struct {
	Offset int16
}

func (*IRgoto) Op() ops.Op { return ops.Goto }
func (ir *IRgoto) Execute(vm VM) error {
	vm.Goto((int32)(ir.Offset))
	return nil
}

type IRgoto_w struct {
	Offset int32
}

func (*IRgoto_w) Op() ops.Op { return ops.Goto_w }
func (ir *IRgoto_w) Execute(vm VM) error {
	vm.Goto(ir.Offset)
	return nil
}

type IRif_acmpeq struct {
	Offset int16
}

func (*IRif_acmpeq) Op() ops.Op { return ops.If_acmpeq }
func (ir *IRif_acmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a == b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_acmpne struct {
	Offset int16
}

func (*IRif_acmpne) Op() ops.Op { return ops.If_acmpne }
func (ir *IRif_acmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a != b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmpeq struct {
	Offset int16
}

func (*IRif_icmpeq) Op() ops.Op { return ops.If_icmpeq }
func (ir *IRif_icmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a == b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmpge struct {
	Offset int16
}

func (*IRif_icmpge) Op() ops.Op { return ops.If_icmpge }
func (ir *IRif_icmpge) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a >= b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmpgt struct {
	Offset int16
}

func (*IRif_icmpgt) Op() ops.Op { return ops.If_icmpgt }
func (ir *IRif_icmpgt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a > b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmple struct {
	Offset int16
}

func (*IRif_icmple) Op() ops.Op { return ops.If_icmple }
func (ir *IRif_icmple) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a <= b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmplt struct {
	Offset int16
}

func (*IRif_icmplt) Op() ops.Op { return ops.If_icmplt }
func (ir *IRif_icmplt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a < b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRif_icmpne struct {
	Offset int16
}

func (*IRif_icmpne) Op() ops.Op { return ops.If_icmpne }
func (ir *IRif_icmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a != b {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifeq struct {
	Offset int16
}

func (*IRifeq) Op() ops.Op { return ops.Ifeq }
func (ir *IRifeq) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a == 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifge struct {
	Offset int16
}

func (*IRifge) Op() ops.Op { return ops.Ifge }
func (ir *IRifge) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a >= 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifgt struct {
	Offset int16
}

func (*IRifgt) Op() ops.Op { return ops.Ifgt }
func (ir *IRifgt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a > 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifle struct {
	Offset int16
}

func (*IRifle) Op() ops.Op { return ops.Ifle }
func (ir *IRifle) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a <= 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRiflt struct {
	Offset int16
}

func (*IRiflt) Op() ops.Op { return ops.Iflt }
func (ir *IRiflt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a < 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifne struct {
	Offset int16
}

func (*IRifne) Op() ops.Op { return ops.Ifne }
func (ir *IRifne) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a != 0 {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifnonnull struct {
	Offset int16
}

func (*IRifnonnull) Op() ops.Op { return ops.Ifnonnull }
func (ir *IRifnonnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a != nil {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRifnull struct {
	Offset int16
}

func (*IRifnull) Op() ops.Op { return ops.Ifnull }
func (ir *IRifnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a == nil {
		vm.Goto((int32)(ir.Offset))
	}
	return nil
}

type IRireturn struct{}

func (*IRireturn) Op() ops.Op { return ops.Ireturn }
func (*IRireturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRjsr struct{}

func (*IRjsr) Op() ops.Op            { return ops.Jsr }
func (*IRjsr) Parse(operands []byte) { panic("deprecated") }
func (*IRjsr) Execute(vm VM) error   { panic("deprecated") }

type IRjsr_w struct{}

func (*IRjsr_w) Op() ops.Op            { return ops.Jsr_w }
func (*IRjsr_w) Parse(operands []byte) { panic("deprecated") }
func (*IRjsr_w) Execute(vm VM) error   { panic("deprecated") }

// A lookupswitch is a variable-length instruction.
// Immediately after the lookupswitch opcode, between zero and three bytes must act as padding,
// such that defaultbyte1 begins at an address that is a multiple of four bytes
// from the start of the current method (the opcode of its first instruction).
//
// IRlookupswitch's operands' length must determined by the parser.
type IRlookupswitch struct {
	DefaultOffset int32
	Indexes       []CaseEntry
}

type CaseEntry struct {
	K int32
	V int32
}

func (c CaseEntry) CmpKey(k int32) int {
	if c.K > k {
		return 1
	}
	if c.K < k {
		return -1
	}
	return 0
}

func (c CaseEntry) Cmp(o CaseEntry) int {
	return c.CmpKey(o.K)
}

func (*IRlookupswitch) Op() ops.Op { return ops.Lookupswitch }
func (ir *IRlookupswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	offset := ir.DefaultOffset
	if ind, ok := slices.BinarySearchFunc(ir.Indexes, key, CaseEntry.CmpKey); ok {
		offset = ir.Indexes[ind].V
	}
	vm.Goto(offset)
	return nil
}

type IRlreturn struct{}

func (*IRlreturn) Op() ops.Op { return ops.Lreturn }
func (*IRlreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRmonitorenter struct{}

func (*IRmonitorenter) Op() ops.Op { return ops.Monitorenter }
func (*IRmonitorenter) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	return vm.MonitorLock(ref)
}

type IRmonitorexit struct{}

func (*IRmonitorexit) Op() ops.Op { return ops.Monitorexit }
func (*IRmonitorexit) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	return vm.MonitorUnlock(ref)
}

type IRnop struct{}

func (*IRnop) Op() ops.Op          { return ops.Nop }
func (*IRnop) Execute(vm VM) error { return nil }

type IRpop struct{}

func (*IRpop) Op() ops.Op { return ops.Pop }
func (*IRpop) Execute(vm VM) error {
	vm.GetStack().PopInt32()
	return nil
}

type IRpop2 struct{}

func (*IRpop2) Op() ops.Op { return ops.Pop2 }
func (*IRpop2) Execute(vm VM) error {
	stack := vm.GetStack()
	stack.PopInt32()
	stack.PopInt32()
	return nil
}

type IRret struct{}

func (*IRret) Op() ops.Op            { return ops.Ret }
func (*IRret) Parse(operands []byte) { panic("deprecated") }
func (*IRret) Execute(vm VM) error   { panic("deprecated") }

type IRreturn struct{}

func (*IRreturn) Op() ops.Op { return ops.Return }
func (*IRreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRswap struct{}

func (*IRswap) Op() ops.Op { return ops.Swap }
func (*IRswap) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	b := stack.PopInt32()
	stack.PushInt32(a)
	stack.PushInt32(b)
	return nil
}

type IRtableswitch struct {
	DefaultOffset int32
	Low, High     int32
	Offsets       []int32
}

func (*IRtableswitch) Op() ops.Op { return ops.Tableswitch }
func (ir *IRtableswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	offset := ir.DefaultOffset
	if ir.Low <= key && key <= ir.High {
		i := key - ir.Low
		offset = ir.Offsets[i]
	}
	vm.Goto(offset)
	return nil
}

type IRwide struct {
	OpCode ops.Op
	Index  uint16
	Const  uint16
}

func (*IRwide) Op() ops.Op { return ops.Wide }
func (ir *IRwide) Execute(vm VM) error {
	switch ir.OpCode {
	case ops.Iload:
		return (&IRiload{Index: ir.Index}).Execute(vm)
	case ops.Fload:
		return (&IRfload{Index: ir.Index}).Execute(vm)
	case ops.Aload:
		return (&IRaload{Index: ir.Index}).Execute(vm)
	case ops.Lload:
		return (&IRlload{Index: ir.Index}).Execute(vm)
	case ops.Dload:
		return (&IRdload{Index: ir.Index}).Execute(vm)
	case ops.Istore:
		return (&IRistore{Index: ir.Index}).Execute(vm)
	case ops.Fstore:
		return (&IRfstore{Index: ir.Index}).Execute(vm)
	case ops.Astore:
		return (&IRastore{Index: ir.Index}).Execute(vm)
	case ops.Lstore:
		return (&IRlstore{Index: ir.Index}).Execute(vm)
	case ops.Dstore:
		return (&IRdstore{Index: ir.Index}).Execute(vm)
	case ops.Iinc:
		return (&IRiinc{Index: ir.Index, Const: (int16)(ir.Const)}).Execute(vm)
	default:
		panic(fmt.Errorf("ir.wide: unexpected opcode %d", ir.OpCode))
	}
}
