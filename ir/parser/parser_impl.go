package parser

import (
	"slices"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ParserIRgoto struct{}

func (*ParserIRgoto) Op() ops.Op { return ops.Goto }
func (*ParserIRgoto) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRgoto)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRgoto_w struct{}

func (*ParserIRgoto_w) Op() ops.Op { return ops.Goto_w }
func (*ParserIRgoto_w) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRgoto_w)
	if ir.Offset, err = readInt32(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_acmpeq struct{}

func (*ParserIRif_acmpeq) Op() ops.Op { return ops.If_acmpeq }
func (*ParserIRif_acmpeq) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_acmpeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_acmpne struct{}

func (*ParserIRif_acmpne) Op() ops.Op { return ops.If_acmpne }
func (*ParserIRif_acmpne) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_acmpne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmpeq struct{}

func (*ParserIRif_icmpeq) Op() ops.Op { return ops.If_icmpeq }
func (*ParserIRif_icmpeq) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmpeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmpge struct{}

func (*ParserIRif_icmpge) Op() ops.Op { return ops.If_icmpge }
func (*ParserIRif_icmpge) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmpge)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmpgt struct{}

func (*ParserIRif_icmpgt) Op() ops.Op { return ops.If_icmpgt }
func (*ParserIRif_icmpgt) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmpgt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmple struct{}

func (*ParserIRif_icmple) Op() ops.Op { return ops.If_icmple }
func (*ParserIRif_icmple) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmple)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmplt struct{}

func (*ParserIRif_icmplt) Op() ops.Op { return ops.If_icmplt }
func (*ParserIRif_icmplt) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmplt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRif_icmpne struct{}

func (*ParserIRif_icmpne) Op() ops.Op { return ops.If_icmpne }
func (*ParserIRif_icmpne) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRif_icmpne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifeq struct{}

func (*ParserIRifeq) Op() ops.Op { return ops.Ifeq }
func (*ParserIRifeq) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifge struct{}

func (*ParserIRifge) Op() ops.Op { return ops.Ifge }
func (*ParserIRifge) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifge)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifgt struct{}

func (*ParserIRifgt) Op() ops.Op { return ops.Ifgt }
func (*ParserIRifgt) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifgt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifle struct{}

func (*ParserIRifle) Op() ops.Op { return ops.Ifle }
func (*ParserIRifle) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifle)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRiflt struct{}

func (*ParserIRiflt) Op() ops.Op { return ops.Iflt }
func (*ParserIRiflt) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRiflt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifne struct{}

func (*ParserIRifne) Op() ops.Op { return ops.Ifne }
func (*ParserIRifne) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifnonnull struct{}

func (*ParserIRifnonnull) Op() ops.Op { return ops.Ifnonnull }
func (*ParserIRifnonnull) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifnonnull)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRifnull struct{}

func (*ParserIRifnull) Op() ops.Op { return ops.Ifnull }
func (*ParserIRifnull) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRifnull)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRlookupswitch struct{}

type irCaseEntry = ir.CaseEntry

func (*ParserIRlookupswitch) Op() ops.Op { return ops.Lookupswitch }
func (*ParserIRlookupswitch) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRlookupswitch)
	if ir.DefaultOffset, err = readInt32(br); err != nil{
		return nil, err
	}
	indexCount, err := readInt32(br)
	if err != nil {
		return nil, err
	}
	ir.Indexes = make([]irCaseEntry, indexCount)
	for i := range indexCount {
		var entry irCaseEntry
		if entry.K, err = readInt32(br); err != nil {
			return nil, err
		}
		if entry.V, err = readInt32(br); err != nil {
			return nil, err
		}
		ir.Indexes[i] = entry
	}
	slices.SortFunc(ir.Indexes, irCaseEntry.Cmp)
	return ir, nil
}

type ParserIRtableswitch struct{}

func (*ParserIRtableswitch) Op() ops.Op { return ops.Tableswitch }
func (*ParserIRtableswitch) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRtableswitch)
	if ir.DefaultOffset, err = readInt32(br); err != nil {
		return nil, err
	}
	if ir.Low, err = readInt32(br); err != nil {
		return nil, err
	}
	if ir.High, err = readInt32(br); err != nil {
		return nil, err
	}
	indexCount := ir.High - ir.Low + 1
	ir.Offsets = make([]int32, indexCount)
	for i := range indexCount {
		if ir.Offsets[i], err = readInt32(br); err != nil {
			return nil, err
		}
	}
	return ir, nil
}

type ParserIRaload struct{}

func (*ParserIRaload) Op() ops.Op { return ops.Aload }
func (*ParserIRaload) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRaload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRanewarray struct{}

func (*ParserIRanewarray) Op() ops.Op { return ops.Anewarray }
func (*ParserIRanewarray) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRanewarray)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRastore struct{}

func (*ParserIRastore) Op() ops.Op { return ops.Astore }
func (*ParserIRastore) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRastore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRcheckcast struct{}

func (*ParserIRcheckcast) Op() ops.Op { return ops.Checkcast }
func (*ParserIRcheckcast) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRcheckcast)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRgetfield struct{}

func (*ParserIRgetfield) Op() ops.Op { return ops.Getfield }
func (*ParserIRgetfield) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRgetfield)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRgetstatic struct{}

func (*ParserIRgetstatic) Op() ops.Op { return ops.Getstatic }
func (*ParserIRgetstatic) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRgetstatic)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRinstanceof struct{}

func (*ParserIRinstanceof) Op() ops.Op { return ops.Instanceof }
func (*ParserIRinstanceof) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinstanceof)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRinvokedynamic struct{}

func (*ParserIRinvokedynamic) Op() ops.Op { return ops.Invokedynamic }
func (*ParserIRinvokedynamic) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinvokedynamic)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if b != 0 {
		panic("ir.invokedynamic: operands [2] must be 0")
	}
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if b != 0 {
		panic("ir.invokedynamic: operands [3] must be 0")
	}
	return ir, nil
}

type ParserIRinvokeinterface struct{}

func (*ParserIRinvokeinterface) Op() ops.Op { return ops.Invokeinterface }
func (*ParserIRinvokeinterface) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinvokeinterface)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	if ir.Count, err = br.ReadByte(); err != nil {
		return nil, err
	}
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if b != 0 {
		panic("ir.invokeinterface: operands [3] must be 0")
	}
	return ir, nil
}

type ParserIRinvokespecial struct{}

func (*ParserIRinvokespecial) Op() ops.Op { return ops.Invokespecial }
func (*ParserIRinvokespecial) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinvokespecial)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRinvokestatic struct{}

func (*ParserIRinvokestatic) Op() ops.Op { return ops.Invokestatic }
func (*ParserIRinvokestatic) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinvokestatic)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRinvokevirtual struct{}

func (*ParserIRinvokevirtual) Op() ops.Op { return ops.Invokevirtual }
func (*ParserIRinvokevirtual) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRinvokevirtual)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRmultianewarray struct{}

func (*ParserIRmultianewarray) Op() ops.Op { return ops.Multianewarray }
func (*ParserIRmultianewarray) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRmultianewarray)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	if ir.Dimensions, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if ir.Dimensions < 1 {
		panic("ir.multianewarray: dimensions is less than 1")
	}
	return ir, nil
}

type ParserIRnew struct{}

func (*ParserIRnew) Op() ops.Op { return ops.New }
func (*ParserIRnew) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRnew)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRnewarray struct{}

func (*ParserIRnewarray) Op() ops.Op { return ops.Newarray }
func (*ParserIRnewarray) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRnewarray)
	if ir.Atype, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRputfield struct{}

func (*ParserIRputfield) Op() ops.Op { return ops.Putfield }
func (*ParserIRputfield) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRputfield)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRputstatic struct{}

func (*ParserIRputstatic) Op() ops.Op { return ops.Putstatic }
func (*ParserIRputstatic) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRputstatic)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRbipush struct{}

func (*ParserIRbipush) Op() ops.Op { return ops.Bipush }
func (*ParserIRbipush) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRbipush)
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	ir.Value = (int8)(b)
	return ir, nil
}

type ParserIRsipush struct{}

func (*ParserIRsipush) Op() ops.Op { return ops.Sipush }
func (*ParserIRsipush) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRsipush)
	if ir.Value, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRdload struct{}

func (*ParserIRdload) Op() ops.Op { return ops.Dload }
func (*ParserIRdload) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRdload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRdstore struct{}

func (*ParserIRdstore) Op() ops.Op { return ops.Dstore }
func (*ParserIRdstore) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRdstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRfload struct{}

func (*ParserIRfload) Op() ops.Op { return ops.Fload }
func (*ParserIRfload) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRfload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRfstore struct{}

func (*ParserIRfstore) Op() ops.Op { return ops.Fstore }
func (*ParserIRfstore) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRfstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRiinc struct{}

func (*ParserIRiinc) Op() ops.Op { return ops.Iinc }
func (*ParserIRiinc) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRiinc)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	if ir.Const, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRiload struct{}

func (*ParserIRiload) Op() ops.Op { return ops.Iload }
func (*ParserIRiload) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRiload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRistore struct{}

func (*ParserIRistore) Op() ops.Op { return ops.Istore }
func (*ParserIRistore) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRistore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRldc struct{}

func (*ParserIRldc) Op() ops.Op { return ops.Ldc }
func (*ParserIRldc) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRldc)
	if ir.Index, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRldc_w struct{}

func (*ParserIRldc_w) Op() ops.Op { return ops.Ldc_w }
func (*ParserIRldc_w) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRldc_w)
	if ir.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRldc2_w struct{}

func (*ParserIRldc2_w) Op() ops.Op { return ops.Ldc2_w }
func (*ParserIRldc2_w) Parse(br ByteReader) (ir.IR, error) {
	var err error
	ir := new(ir.IRldc2_w)
	if ir.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIRlload struct{}

func (*ParserIRlload) Op() ops.Op { return ops.Lload }
func (*ParserIRlload) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRlload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIRlstore struct{}

func (*ParserIRlstore) Op() ops.Op { return ops.Lstore }
func (*ParserIRlstore) Parse(br ByteReader) (ir.IR, error) {
	ir := new(ir.IRlstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}
