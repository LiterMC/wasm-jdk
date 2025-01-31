// See: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-6.html#jvms-6.5
package ir

import (
	"github.com/LiterMC/wasm-jdk/ops"
)

type IRareturn struct{}

func (*IRareturn) Op() ops.Op            { return ops.Areturn }
func (*IRareturn) Operands() int         { return 0 }
func (*IRareturn) Parse(operands []byte) {}
func (*IRareturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRdreturn struct{}

func (*IRdreturn) Op() ops.Op            { return ops.Dreturn }
func (*IRdreturn) Operands() int         { return 0 }
func (*IRdreturn) Parse(operands []byte) {}
func (*IRdreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRdup struct{}

func (*IRdup) Op() ops.Op            { return ops.Dup }
func (*IRdup) Operands() int         { return 0 }
func (*IRdup) Parse(operands []byte) {}
func (*IRdup) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PeekInt32()
	stack.PushInt32(val)
	return nil
}

type IRdup_x1 struct{}

func (*IRdup_x1) Op() ops.Op            { return ops.Dup_x1 }
func (*IRdup_x1) Operands() int         { return 0 }
func (*IRdup_x1) Parse(operands []byte) {}
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

func (*IRdup_x2) Op() ops.Op            { return ops.Dup_x2 }
func (*IRdup_x2) Operands() int         { return 0 }
func (*IRdup_x2) Parse(operands []byte) {}
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

func (*IRdup2) Op() ops.Op            { return ops.Dup2 }
func (*IRdup2) Operands() int         { return 0 }
func (*IRdup2) Parse(operands []byte) {}
func (*IRdup2) Execute(vm VM) error {
	stack := vm.GetStack()
	val := stack.PeekInt64()
	stack.PushInt64(val)
	return nil
}

type IRdup2_x1 struct{}

func (*IRdup2_x1) Op() ops.Op            { return ops.Dup2_x1 }
func (*IRdup2_x1) Operands() int         { return 0 }
func (*IRdup2_x1) Parse(operands []byte) {}
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

func (*IRdup2_x2) Op() ops.Op            { return ops.Dup2_x2 }
func (*IRdup2_x2) Operands() int         { return 0 }
func (*IRdup2_x2) Parse(operands []byte) {}
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

func (*IRfreturn) Op() ops.Op            { return ops.Freturn }
func (*IRfreturn) Operands() int         { return 0 }
func (*IRfreturn) Parse(operands []byte) {}
func (*IRfreturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRgoto struct {
	offset int16
}

func (*IRgoto) Op() ops.Op    { return ops.Goto }
func (*IRgoto) Operands() int { return 2 }
func (ir *IRgoto) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRgoto) Execute(vm VM) error {
	vm.Goto((int32)(ir.offset))
	return nil
}

type IRgoto_w struct {
	offset int32
}

func (*IRgoto_w) Op() ops.Op    { return ops.Goto_w }
func (*IRgoto_w) Operands() int { return 2 }
func (ir *IRgoto_w) Parse(operands []byte) {
	ir.offset = ((int32)(operands[0]) << 24) | ((int32)(operands[1]) << 16) | ((int32)(operands[2]) << 8) | (int32)(operands[3])
}
func (ir *IRgoto_w) Execute(vm VM) error {
	vm.Goto(ir.offset)
	return nil
}

type IRif_acmpeq struct {
	offset int16
}

func (*IRif_acmpeq) Op() ops.Op    { return ops.If_acmpeq }
func (*IRif_acmpeq) Operands() int { return 2 }
func (ir *IRif_acmpeq) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_acmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a == b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_acmpne struct {
	offset int16
}

func (*IRif_acmpne) Op() ops.Op    { return ops.If_acmpne }
func (*IRif_acmpne) Operands() int { return 2 }
func (ir *IRif_acmpne) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_acmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopRef()
	a := stack.PopRef()
	if a != b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmpeq struct {
	offset int16
}

func (*IRif_icmpeq) Op() ops.Op    { return ops.If_icmpeq }
func (*IRif_icmpeq) Operands() int { return 2 }
func (ir *IRif_icmpeq) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmpeq) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a == b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmpge struct {
	offset int16
}

func (*IRif_icmpge) Op() ops.Op    { return ops.If_icmpge }
func (*IRif_icmpge) Operands() int { return 2 }
func (ir *IRif_icmpge) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmpge) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a >= b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmpgt struct {
	offset int16
}

func (*IRif_icmpgt) Op() ops.Op    { return ops.If_icmpgt }
func (*IRif_icmpgt) Operands() int { return 2 }
func (ir *IRif_icmpgt) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmpgt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a > b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmple struct {
	offset int16
}

func (*IRif_icmple) Op() ops.Op    { return ops.If_icmple }
func (*IRif_icmple) Operands() int { return 2 }
func (ir *IRif_icmple) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmple) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a <= b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmplt struct {
	offset int16
}

func (*IRif_icmplt) Op() ops.Op    { return ops.If_icmplt }
func (*IRif_icmplt) Operands() int { return 2 }
func (ir *IRif_icmplt) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmplt) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a < b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRif_icmpne struct {
	offset int16
}

func (*IRif_icmpne) Op() ops.Op    { return ops.If_icmpne }
func (*IRif_icmpne) Operands() int { return 2 }
func (ir *IRif_icmpne) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRif_icmpne) Execute(vm VM) error {
	stack := vm.GetStack()
	b := stack.PopInt32()
	a := stack.PopInt32()
	if a != b {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifeq struct {
	offset int16
}

func (*IRifeq) Op() ops.Op    { return ops.Ifeq }
func (*IRifeq) Operands() int { return 2 }
func (ir *IRifeq) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifeq) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a == 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifge struct {
	offset int16
}

func (*IRifge) Op() ops.Op    { return ops.Ifge }
func (*IRifge) Operands() int { return 2 }
func (ir *IRifge) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifge) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a >= 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifgt struct {
	offset int16
}

func (*IRifgt) Op() ops.Op    { return ops.Ifgt }
func (*IRifgt) Operands() int { return 2 }
func (ir *IRifgt) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifgt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a > 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifle struct {
	offset int16
}

func (*IRifle) Op() ops.Op    { return ops.Ifle }
func (*IRifle) Operands() int { return 2 }
func (ir *IRifle) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifle) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a <= 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRiflt struct {
	offset int16
}

func (*IRiflt) Op() ops.Op    { return ops.Iflt }
func (*IRiflt) Operands() int { return 2 }
func (ir *IRiflt) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRiflt) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a < 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifne struct {
	offset int16
}

func (*IRifne) Op() ops.Op    { return ops.Ifne }
func (*IRifne) Operands() int { return 2 }
func (ir *IRifne) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifne) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopInt32()
	if a != 0 {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifnonnull struct {
	offset int16
}

func (*IRifnonnull) Op() ops.Op    { return ops.Ifnonnull }
func (*IRifnonnull) Operands() int { return 2 }
func (ir *IRifnonnull) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifnonnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a != nil {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRifnull struct {
	offset int16
}

func (*IRifnull) Op() ops.Op    { return ops.Ifnull }
func (*IRifnull) Operands() int { return 2 }
func (ir *IRifnull) Parse(operands []byte) {
	ir.offset = ((int16)(operands[0]) << 8) | (int16)(operands[1])
}
func (ir *IRifnull) Execute(vm VM) error {
	stack := vm.GetStack()
	a := stack.PopRef()
	if a == nil {
		vm.Goto((int32)(ir.offset))
	}
	return nil
}

type IRireturn struct{}

func (*IRireturn) Op() ops.Op            { return ops.Ireturn }
func (*IRireturn) Operands() int         { return 0 }
func (*IRireturn) Parse(operands []byte) {}
func (*IRireturn) Execute(vm VM) error {
	vm.Return()
	return nil
}

type IRjsr struct{}

func (*IRjsr) Op() ops.Op            { return ops.Jsr }
func (*IRjsr) Operands() int         { return 0 }
func (*IRjsr) Parse(operands []byte) { panic("deprecated") }
func (*IRjsr) Execute(vm VM) error   { panic("deprecated") }

type IRjsr_w struct{}

func (*IRjsr_w) Op() ops.Op            { return ops.Jsr_w }
func (*IRjsr_w) Operands() int         { return 0 }
func (*IRjsr_w) Parse(operands []byte) { panic("deprecated") }
func (*IRjsr_w) Execute(vm VM) error   { panic("deprecated") }
