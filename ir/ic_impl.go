// See: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-6.html#jvms-6.5
package ir

import (
	"fmt"
	"slices"

	"github.com/LiterMC/wasm-jdk/ops"
)

type ICareturn struct{}

func (*ICareturn) Op() ops.Op { return ops.Areturn }
func (*ICareturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICdreturn struct{}

func (*ICdreturn) Op() ops.Op { return ops.Dreturn }
func (*ICdreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICdup struct{}

func (*ICdup) Op() ops.Op { return ops.Dup }
func (*ICdup) Execute(vm VM) error {
	stack := vm.GetStack()
	if stack.IsRef() {
		val := stack.PeekRef()
		stack.PushRef(val)
	} else {
		val := stack.PeekInt32()
		stack.PushInt32(val)
	}
	return nil
}

type ICdup_x1 struct{}

func (*ICdup_x1) Op() ops.Op { return ops.Dup_x1 }
func (*ICdup_x1) Execute(vm VM) error {
	stack := vm.GetStack()
	var pushA func()
	if stack.IsRef() {
		a := stack.PopRef()
		pushA = func() { stack.PushRef(a) }
	} else {
		a := stack.PopInt32()
		pushA = func() { stack.PushInt32(a) }
	}
	if stack.IsRef() {
		b := stack.PopRef()
		pushA()
		stack.PushRef(b)
	} else {
		b := stack.PopInt32()
		pushA()
		stack.PushInt32(b)
	}
	pushA()
	return nil
}

type ICdup_x2 struct{}

func (*ICdup_x2) Op() ops.Op { return ops.Dup_x2 }
func (*ICdup_x2) Execute(vm VM) error {
	stack := vm.GetStack()
	var pushA, pushB1, pushB2 func()
	if stack.IsRef() {
		a := stack.PopRef()
		pushA = func() { stack.PushRef(a) }
	} else {
		a := stack.PopInt32()
		pushA = func() { stack.PushInt32(a) }
	}
	if stack.IsRef() {
		b2 := stack.PopRef()
		pushB2 = func() { stack.PushRef(b2) }
	} else {
		b2 := stack.PopInt32()
		pushB2 = func() { stack.PushInt32(b2) }
	}
	if stack.IsRef() {
		b1 := stack.PopRef()
		pushB1 = func() { stack.PushRef(b1) }
	} else {
		b1 := stack.PopInt32()
		pushB1 = func() { stack.PushInt32(b1) }
	}
	pushA()
	pushB1()
	pushB2()
	pushA()
	return nil
}

type ICdup2 struct{}

func (*ICdup2) Op() ops.Op { return ops.Dup2 }
func (*ICdup2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PeekInt64()
	stack.PushInt64(val)
	return nil
}

type ICdup2_x1 struct{}

func (*ICdup2_x1) Op() ops.Op { return ops.Dup2_x1 }
func (*ICdup2_x1) Execute(vm VM) error {
	stack := vm.GetStack()
	var pushA1, pushA2, pushB func()
	if stack.IsRef() {
		a2 := stack.PopRef()
		pushA2 = func() { stack.PushRef(a2) }
	} else {
		a2 := stack.PopInt32()
		pushA2 = func() { stack.PushInt32(a2) }
	}
	if stack.IsRef() {
		a1 := stack.PopRef()
		pushA1 = func() { stack.PushRef(a1) }
	} else {
		a1 := stack.PopInt32()
		pushA1 = func() { stack.PushInt32(a1) }
	}
	if stack.IsRef() {
		b := stack.PopRef()
		pushB = func() { stack.PushRef(b) }
	} else {
		b := stack.PopInt32()
		pushB = func() { stack.PushInt32(b) }
	}
	pushA1()
	pushA2()
	pushB()
	pushA1()
	pushA2()
	return nil
}

type ICdup2_x2 struct{}

func (*ICdup2_x2) Op() ops.Op { return ops.Dup2_x2 }
func (*ICdup2_x2) Execute(vm VM) error {
	stack := vm.GetStack()
	var pushA1, pushA2, pushB1, pushB2 func()
	if stack.IsRef() {
		a2 := stack.PopRef()
		pushA2 = func() { stack.PushRef(a2) }
	} else {
		a2 := stack.PopInt32()
		pushA2 = func() { stack.PushInt32(a2) }
	}
	if stack.IsRef() {
		a1 := stack.PopRef()
		pushA1 = func() { stack.PushRef(a1) }
	} else {
		a1 := stack.PopInt32()
		pushA1 = func() { stack.PushInt32(a1) }
	}
	if stack.IsRef() {
		b2 := stack.PopRef()
		pushB2 = func() { stack.PushRef(b2) }
	} else {
		b2 := stack.PopInt32()
		pushB2 = func() { stack.PushInt32(b2) }
	}
	if stack.IsRef() {
		b1 := stack.PopRef()
		pushB1 = func() { stack.PushRef(b1) }
	} else {
		b1 := stack.PopInt32()
		pushB1 = func() { stack.PushInt32(b1) }
	}
	pushA1()
	pushA2()
	pushB1()
	pushB2()
	pushA1()
	pushA2()
	return nil
}

type ICfreturn struct{}

func (*ICfreturn) Op() ops.Op { return ops.Freturn }
func (*ICfreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICgoto struct {
	Offset int16
	Node   *ICNode
}

func (*ICgoto) Op() ops.Op { return ops.Goto }
func (ir *ICgoto) Execute(vm VM) error {
	vm.Goto(ir.Node)
	return nil
}

type ICgoto_w struct {
	Offset int32
	Node   *ICNode
}

func (*ICgoto_w) Op() ops.Op { return ops.Goto_w }
func (ir *ICgoto_w) Execute(vm VM) error {
	vm.Goto(ir.Node)
	return nil
}

type ICif_acmpeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_acmpeq) Op() ops.Op { return ops.If_acmpeq }
func (ir *ICif_acmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a == b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_acmpne struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_acmpne) Op() ops.Op { return ops.If_acmpne }
func (ir *ICif_acmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a != b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmpeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpeq) Op() ops.Op { return ops.If_icmpeq }
func (ir *ICif_icmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a == b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmpge struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpge) Op() ops.Op { return ops.If_icmpge }
func (ir *ICif_icmpge) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a >= b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmpgt struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpgt) Op() ops.Op { return ops.If_icmpgt }
func (ir *ICif_icmpgt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a > b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmple struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmple) Op() ops.Op { return ops.If_icmple }
func (ir *ICif_icmple) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a <= b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmplt struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmplt) Op() ops.Op { return ops.If_icmplt }
func (ir *ICif_icmplt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a < b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICif_icmpne struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpne) Op() ops.Op { return ops.If_icmpne }
func (ir *ICif_icmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a != b {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICifeq) Op() ops.Op { return ops.Ifeq }
func (ir *ICifeq) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a == 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifge struct {
	Offset int16
	Node   *ICNode
}

func (*ICifge) Op() ops.Op { return ops.Ifge }
func (ir *ICifge) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a >= 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifgt struct {
	Offset int16
	Node   *ICNode
}

func (*ICifgt) Op() ops.Op { return ops.Ifgt }
func (ir *ICifgt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a > 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifle struct {
	Offset int16
	Node   *ICNode
}

func (*ICifle) Op() ops.Op { return ops.Ifle }
func (ir *ICifle) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a <= 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICiflt struct {
	Offset int16
	Node   *ICNode
}

func (*ICiflt) Op() ops.Op { return ops.Iflt }
func (ir *ICiflt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a < 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifne struct {
	Offset int16
	Node   *ICNode
}

func (*ICifne) Op() ops.Op { return ops.Ifne }
func (ir *ICifne) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a != 0 {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifnonnull struct {
	Offset int16
	Node   *ICNode
}

func (*ICifnonnull) Op() ops.Op { return ops.Ifnonnull }
func (ir *ICifnonnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a != nil {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICifnull struct {
	Offset int16
	Node   *ICNode
}

func (*ICifnull) Op() ops.Op { return ops.Ifnull }
func (ir *ICifnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a == nil {
		vm.Goto(ir.Node)
	}
	return nil
}

type ICireturn struct{}

func (*ICireturn) Op() ops.Op { return ops.Ireturn }
func (*ICireturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICjsr struct{}

func (*ICjsr) Op() ops.Op            { return ops.Jsr }
func (*ICjsr) Parse(operands []byte) { panic("deprecated") }
func (*ICjsr) Execute(vm VM) error   { panic("deprecated") }

type ICjsr_w struct{}

func (*ICjsr_w) Op() ops.Op            { return ops.Jsr_w }
func (*ICjsr_w) Parse(operands []byte) { panic("deprecated") }
func (*ICjsr_w) Execute(vm VM) error   { panic("deprecated") }

// A lookupswitch is a variable-length instruction.
// Immediately after the lookupswitch opcode, between zero and three bytes must act as padding,
// such that defaultbyte1 begins at an address that is a multiple of four bytes
// from the start of the current method (the opcode of its first instruction).
//
// IClookupswitch's operands' length must determined by the parser.
type IClookupswitch struct {
	DefaultOffset int32
	DefaultNode   *ICNode
	Indexes       []CaseEntry
}

type CaseEntry struct {
	K int32
	V int32
	N *ICNode
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

func (*IClookupswitch) Op() ops.Op { return ops.Lookupswitch }
func (ir *IClookupswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	node := ir.DefaultNode
	if ind, ok := slices.BinarySearchFunc(ir.Indexes, key, CaseEntry.CmpKey); ok {
		node = ir.Indexes[ind].N
	}
	vm.Goto(node)
	return nil
}

type IClreturn struct{}

func (*IClreturn) Op() ops.Op { return ops.Lreturn }
func (*IClreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICmonitorenter struct{}

func (*ICmonitorenter) Op() ops.Op { return ops.Monitorenter }
func (*ICmonitorenter) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	return vm.MonitorLock(ref)
}

type ICmonitorexit struct{}

func (*ICmonitorexit) Op() ops.Op { return ops.Monitorexit }
func (*ICmonitorexit) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	return vm.MonitorUnlock(ref)
}

type ICnop struct{}

func (*ICnop) Op() ops.Op          { return ops.Nop }
func (*ICnop) Execute(vm VM) error { return nil }

type ICpop struct{}

func (*ICpop) Op() ops.Op { return ops.Pop }
func (*ICpop) Execute(vm VM) error {
	vm.GetStack().PopInt32()
	return nil
}

type ICpop2 struct{}

func (*ICpop2) Op() ops.Op { return ops.Pop2 }
func (*ICpop2) Execute(vm VM) error {
	stack := vm.GetStack()
	stack.PopInt32()
	stack.PopInt32()
	return nil
}

type ICret struct{}

func (*ICret) Op() ops.Op            { return ops.Ret }
func (*ICret) Parse(operands []byte) { panic("deprecated") }
func (*ICret) Execute(vm VM) error   { panic("deprecated") }

type ICreturn struct{}

func (*ICreturn) Op() ops.Op { return ops.Return }
func (*ICreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type ICswap struct{}

func (*ICswap) Op() ops.Op { return ops.Swap }
func (*ICswap) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	b := stack.PopInt32()
	stack.PushInt32(a)
	stack.PushInt32(b)
	return nil
}

type ICtableswitch struct {
	DefaultOffset int32
	DefaultNode   *ICNode
	Low, High     int32
	Offsets       []int32
	Nodes         []*ICNode
}

func (*ICtableswitch) Op() ops.Op { return ops.Tableswitch }
func (ir *ICtableswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	node := ir.DefaultNode
	if ir.Low <= key && key <= ir.High {
		i := key - ir.Low
		node = ir.Nodes[i]
	}
	vm.Goto(node)
	return nil
}

type ICwide struct {
	OpCode ops.Op
	Index  uint16
	Const  uint16
}

func (*ICwide) Op() ops.Op { return ops.Wide }
func (ir *ICwide) Execute(vm VM) error {
	switch ir.OpCode {
	case ops.Iload:
		return (&ICiload{Index: ir.Index}).Execute(vm)
	case ops.Fload:
		return (&ICfload{Index: ir.Index}).Execute(vm)
	case ops.Aload:
		return (&ICaload{Index: ir.Index}).Execute(vm)
	case ops.Lload:
		return (&IClload{Index: ir.Index}).Execute(vm)
	case ops.Dload:
		return (&ICdload{Index: ir.Index}).Execute(vm)
	case ops.Istore:
		return (&ICistore{Index: ir.Index}).Execute(vm)
	case ops.Fstore:
		return (&ICfstore{Index: ir.Index}).Execute(vm)
	case ops.Astore:
		return (&ICastore{Index: ir.Index}).Execute(vm)
	case ops.Lstore:
		return (&IClstore{Index: ir.Index}).Execute(vm)
	case ops.Dstore:
		return (&ICdstore{Index: ir.Index}).Execute(vm)
	case ops.Iinc:
		return (&ICiinc{Index: ir.Index, Const: (int16)(ir.Const)}).Execute(vm)
	default:
		panic(fmt.Errorf("ir.wide: unexpected opcode %d", ir.OpCode))
	}
}
