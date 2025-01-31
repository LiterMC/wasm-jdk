package parser

import (
	"io"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

func (ir *IRgoto) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRgoto_w) Parse(operands []byte) {
	ir.Offset = bytesToInt32(operands)
}

func (ir *IRif_acmpeq) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_acmpne) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmpeq) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmpge) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmpgt) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmple) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmplt) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRif_icmpne) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifeq) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifge) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifgt) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifle) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRiflt) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifne) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifnonnull) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRifnull) Parse(operands []byte) {
	ir.Offset = bytesToInt16(operands)
}

func (ir *IRlookupswitch) Parse(operands []byte) {
	ir.DefaultOffset = bytesToInt32(operands[0:4])
	indexCount := bytesToInt32(operands[4:8])
	ir.Indexes = make([]ir.CaseEntry, indexCount)
	for i := range indexCount {
		j := 8 + 8*i
		k := bytesToInt32(operands[j : j+4])
		v := bytesToInt32(operands[j+4 : j+8])
		ir.Indexes[i] = ir.CaseEntry{K: k, V: v}
	}
	slices.SortFunc(ir.Indexes, ir.CaseEntry.Cmp)
}

func (ir *IRtableswitch) Parse(operands []byte) {
	ir.DefaultOffset = bytesToInt32(operands[0:4])
	ir.Low = bytesToInt32(operands[4:8])
	ir.High = bytesToInt32(operands[8:12])
	indexCount := ir.High - ir.Low + 1
	ir.Offsets = make([]int32, indexCount)
	for i := range indexCount {
		j := 12 + 4*i
		v := bytesToInt32(operands[j : j+4])
		ir.Offsets[i] = v
	}
}

func (ir *IRaload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRanewarray) Parse(operands []byte) {
	ir.Class = bytesToUint16(operands)
}

func (ir *IRastore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRcheckcast) Parse(operands []byte) {
	ir.Class = bytesToUint16(operands)
}

func (ir *IRgetfield) Parse(operands []byte) {
	ir.Field = bytesToUint16(operands)
}

func (ir *IRgetstatic) Parse(operands []byte) {
	ir.Field = bytesToUint16(operands)
}

func (ir *IRinstanceof) Parse(operands []byte) {
	ir.Class = bytesToUint16(operands)
}

func (ir *IRinvokedynamic) Parse(operands []byte) {
	ir.Method = bytesToUint16(operands)
	if operands[2] != 0 || operands[3] != 0 {
		panic("ir.invokedynamic: operands [2] and [3] must be 0")
	}
}

func (ir *IRinvokeinterface) Parse(operands []byte) {
	ir.Method = bytesToUint16(operands)
	ir.Count = operands[2]
	if operands[3] != 0 {
		panic("ir.invokeinterface: operands [3] must be 0")
	}
}

func (ir *IRinvokespecial) Parse(operands []byte) {
	ir.Method = bytesToUint16(operands)
}

func (ir *IRinvokestatic) Parse(operands []byte) {
	ir.Method = bytesToUint16(operands)
}

func (ir *IRinvokevirtual) Parse(operands []byte) {
	ir.Method = bytesToUint16(operands)
}

func (ir *IRmultianewarray) Parse(operands []byte) {
	ir.Class = bytesToUint16(operands)
	ir.Dimensions = operands[2]
	if ir.Dimensions < 1 {
		panic("ir.multianewarray: dimensions is less than 1")
	}
}

func (ir *IRnew) Parse(operands []byte) {
	ir.Class = bytesToUint16(operands)
}

func (ir *IRnewarray) Parse(operands []byte) {
	ir.atype = operands[0]
}

func (ir *IRputfield) Parse(operands []byte) {
	ir.Field = bytesToUint16(operands)
}

func (ir *IRputstatic) Parse(operands []byte) {
	ir.Field = bytesToUint16(operands)
}

func (ir *IRbipush) Parse(operands []byte) {
	ir.value = (int8)(operands[0])
}

func (ir *IRsipush) Parse(operands []byte) {
	ir.value = bytesToInt16(operands)
}

func (ir *IRdload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRdstore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRfload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRfstore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRiinc) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
	ir.value = (int16)((int8)(operands[1]))
}

func (ir *IRiload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRistore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRldc) Parse(operands []byte) {
	ir.Index = operands[0]
}

func (ir *IRldc_w) Parse(operands []byte) {
	ir.Index = bytesToUint16(operands)
}

func (ir *IRldc2_w) Parse(operands []byte) {
	ir.Index = bytesToUint16(operands)
}

func (ir *IRlload) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}

func (ir *IRlstore) Parse(operands []byte) {
	ir.Index = (uint16)(operands[0])
}
