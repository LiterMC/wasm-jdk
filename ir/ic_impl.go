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
		val := stack.Peek()
		stack.Push(val)
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
		a := stack.Pop()
		pushA = func() { stack.Push(a) }
	}
	if stack.IsRef() {
		b := stack.PopRef()
		pushA()
		stack.PushRef(b)
	} else {
		b := stack.Pop()
		pushA()
		stack.Push(b)
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
		a := stack.Pop()
		pushA = func() { stack.Push(a) }
	}
	if stack.IsRef() {
		b2 := stack.PopRef()
		pushB2 = func() { stack.PushRef(b2) }
	} else {
		b2 := stack.Pop()
		pushB2 = func() { stack.Push(b2) }
	}
	if stack.IsRef() {
		b1 := stack.PopRef()
		pushB1 = func() { stack.PushRef(b1) }
	} else {
		b1 := stack.Pop()
		pushB1 = func() { stack.Push(b1) }
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
	var pushA1, pushA2 func()
	if stack.IsRef() {
		a2 := stack.PopRef()
		pushA2 = func() { stack.PushRef(a2) }
	} else {
		a2 := stack.Pop()
		pushA2 = func() { stack.Push(a2) }
	}
	if stack.IsRef() {
		a1 := stack.PopRef()
		pushA1 = func() { stack.PushRef(a1) }
	} else {
		a1 := stack.Pop()
		pushA1 = func() { stack.Push(a1) }
	}
	pushA1()
	pushA2()
	pushA1()
	pushA2()
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
		a2 := stack.Pop()
		pushA2 = func() { stack.Push(a2) }
	}
	if stack.IsRef() {
		a1 := stack.PopRef()
		pushA1 = func() { stack.PushRef(a1) }
	} else {
		a1 := stack.Pop()
		pushA1 = func() { stack.Push(a1) }
	}
	if stack.IsRef() {
		b := stack.PopRef()
		pushB = func() { stack.PushRef(b) }
	} else {
		b := stack.Pop()
		pushB = func() { stack.Push(b) }
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
		a2 := stack.Pop()
		pushA2 = func() { stack.Push(a2) }
	}
	if stack.IsRef() {
		a1 := stack.PopRef()
		pushA1 = func() { stack.PushRef(a1) }
	} else {
		a1 := stack.Pop()
		pushA1 = func() { stack.Push(a1) }
	}
	if stack.IsRef() {
		b2 := stack.PopRef()
		pushB2 = func() { stack.PushRef(b2) }
	} else {
		b2 := stack.Pop()
		pushB2 = func() { stack.Push(b2) }
	}
	if stack.IsRef() {
		b1 := stack.PopRef()
		pushB1 = func() { stack.PushRef(b1) }
	} else {
		b1 := stack.Pop()
		pushB1 = func() { stack.Push(b1) }
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

func (*ICgoto) Op() ops.Op                  { return ops.Goto }
func (ic *ICgoto) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICgoto) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICgoto) Execute(vm VM) error {
	vm.Goto(ic.Node)
	return nil
}

type ICgoto_w struct {
	Offset int32
	Node   *ICNode
}

func (*ICgoto_w) Op() ops.Op                  { return ops.Goto_w }
func (ic *ICgoto_w) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICgoto_w) Offsets() []int32         { return []int32{ic.Offset} }
func (ic *ICgoto_w) Execute(vm VM) error {
	vm.Goto(ic.Node)
	return nil
}

type ICif_acmpeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_acmpeq) Op() ops.Op                  { return ops.If_acmpeq }
func (ic *ICif_acmpeq) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_acmpeq) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_acmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a == b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_acmpne struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_acmpne) Op() ops.Op                  { return ops.If_acmpne }
func (ic *ICif_acmpne) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_acmpne) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_acmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a != b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmpeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpeq) Op() ops.Op                  { return ops.If_icmpeq }
func (ic *ICif_icmpeq) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmpeq) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a == b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmpge struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpge) Op() ops.Op                  { return ops.If_icmpge }
func (ic *ICif_icmpge) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmpge) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmpge) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a >= b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmpgt struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpgt) Op() ops.Op                  { return ops.If_icmpgt }
func (ic *ICif_icmpgt) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmpgt) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmpgt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a > b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmple struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmple) Op() ops.Op                  { return ops.If_icmple }
func (ic *ICif_icmple) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmple) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmple) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a <= b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmplt struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmplt) Op() ops.Op                  { return ops.If_icmplt }
func (ic *ICif_icmplt) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmplt) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmplt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a < b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICif_icmpne struct {
	Offset int16
	Node   *ICNode
}

func (*ICif_icmpne) Op() ops.Op                  { return ops.If_icmpne }
func (ic *ICif_icmpne) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICif_icmpne) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICif_icmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a != b {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifeq struct {
	Offset int16
	Node   *ICNode
}

func (*ICifeq) Op() ops.Op                  { return ops.Ifeq }
func (ic *ICifeq) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifeq) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifeq) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a == 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifge struct {
	Offset int16
	Node   *ICNode
}

func (*ICifge) Op() ops.Op                  { return ops.Ifge }
func (ic *ICifge) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifge) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifge) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a >= 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifgt struct {
	Offset int16
	Node   *ICNode
}

func (*ICifgt) Op() ops.Op                  { return ops.Ifgt }
func (ic *ICifgt) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifgt) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifgt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a > 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifle struct {
	Offset int16
	Node   *ICNode
}

func (*ICifle) Op() ops.Op                  { return ops.Ifle }
func (ic *ICifle) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifle) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifle) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a <= 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICiflt struct {
	Offset int16
	Node   *ICNode
}

func (*ICiflt) Op() ops.Op                  { return ops.Iflt }
func (ic *ICiflt) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICiflt) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICiflt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a < 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifne struct {
	Offset int16
	Node   *ICNode
}

func (*ICifne) Op() ops.Op                  { return ops.Ifne }
func (ic *ICifne) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifne) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifne) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a != 0 {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifnonnull struct {
	Offset int16
	Node   *ICNode
}

func (*ICifnonnull) Op() ops.Op                  { return ops.Ifnonnull }
func (ic *ICifnonnull) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifnonnull) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifnonnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a != nil {
		vm.Goto(ic.Node)
	}
	return nil
}

type ICifnull struct {
	Offset int16
	Node   *ICNode
}

func (*ICifnull) Op() ops.Op                  { return ops.Ifnull }
func (ic *ICifnull) SetNode(i int, n *ICNode) { ic.Node = n }
func (ic *ICifnull) Offsets() []int32         { return []int32{(int32)(ic.Offset)} }
func (ic *ICifnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a == nil {
		vm.Goto(ic.Node)
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
func (ic *IClookupswitch) SetNode(i int, n *ICNode) {
	if i == 0 {
		ic.DefaultNode = n
	} else {
		ic.Indexes[i-1].N = n
	}
}
func (ic *IClookupswitch) Offsets() []int32 {
	offsets := make([]int32, len(ic.Indexes)+1)
	offsets[0] = ic.DefaultOffset
	for i, e := range ic.Indexes {
		offsets[i+1] = e.V
	}
	return offsets
}
func (ic *IClookupswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	node := ic.DefaultNode
	if ind, ok := slices.BinarySearchFunc(ic.Indexes, key, CaseEntry.CmpKey); ok {
		node = ic.Indexes[ind].N
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
	ref.Lock(vm)
	return nil
}

type ICmonitorexit struct{}

func (*ICmonitorexit) Op() ops.Op { return ops.Monitorexit }
func (*ICmonitorexit) Execute(vm VM) error {
	ref := vm.GetStack().PopRef()
	_, err := ref.Unlock(vm)
	return err
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
	OffsetList    []int32
	Nodes         []*ICNode
}

func (*ICtableswitch) Op() ops.Op { return ops.Tableswitch }
func (ic *ICtableswitch) SetNode(i int, n *ICNode) {
	if i == 0 {
		ic.DefaultNode = n
	} else {
		if ic.Nodes == nil {
			ic.Nodes = make([]*ICNode, len(ic.OffsetList))
		}
		ic.Nodes[i-1] = n
	}
}
func (ic *ICtableswitch) Offsets() []int32 {
	offsets := make([]int32, len(ic.OffsetList)+1)
	offsets[0] = ic.DefaultOffset
	copy(offsets[1:], ic.OffsetList)
	return offsets
}
func (ic *ICtableswitch) Execute(vm VM) error {
	key := vm.GetStack().PopInt32()
	node := ic.DefaultNode
	if ic.Low <= key && key <= ic.High {
		i := key - ic.Low
		node = ic.Nodes[i]
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
func (ic *ICwide) Execute(vm VM) error {
	switch ic.OpCode {
	case ops.Iload:
		return (&ICiload{Index: ic.Index}).Execute(vm)
	case ops.Fload:
		return (&ICfload{Index: ic.Index}).Execute(vm)
	case ops.Aload:
		return (&ICaload{Index: ic.Index}).Execute(vm)
	case ops.Lload:
		return (&IClload{Index: ic.Index}).Execute(vm)
	case ops.Dload:
		return (&ICdload{Index: ic.Index}).Execute(vm)
	case ops.Istore:
		return (&ICistore{Index: ic.Index}).Execute(vm)
	case ops.Fstore:
		return (&ICfstore{Index: ic.Index}).Execute(vm)
	case ops.Astore:
		return (&ICastore{Index: ic.Index}).Execute(vm)
	case ops.Lstore:
		return (&IClstore{Index: ic.Index}).Execute(vm)
	case ops.Dstore:
		return (&ICdstore{Index: ic.Index}).Execute(vm)
	case ops.Iinc:
		return (&ICiinc{Index: ic.Index, Const: (int16)(ic.Const)}).Execute(vm)
	default:
		panic(fmt.Errorf("ir.wide: unexpected opcode %d", ic.OpCode))
	}
}
