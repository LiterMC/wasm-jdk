package parser

import (
	"fmt"
	"slices"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/ops"
)

type ParserICgoto struct{}

func (*ParserICgoto) Op() ops.Op { return ops.Goto }
func (*ParserICgoto) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICgoto)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICgoto_w struct{}

func (*ParserICgoto_w) Op() ops.Op { return ops.Goto_w }
func (*ParserICgoto_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICgoto_w)
	if ic.Offset, err = readInt32(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_acmpeq struct{}

func (*ParserICif_acmpeq) Op() ops.Op { return ops.If_acmpeq }
func (*ParserICif_acmpeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_acmpeq)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_acmpne struct{}

func (*ParserICif_acmpne) Op() ops.Op { return ops.If_acmpne }
func (*ParserICif_acmpne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_acmpne)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmpeq struct{}

func (*ParserICif_icmpeq) Op() ops.Op { return ops.If_icmpeq }
func (*ParserICif_icmpeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmpeq)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmpge struct{}

func (*ParserICif_icmpge) Op() ops.Op { return ops.If_icmpge }
func (*ParserICif_icmpge) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmpge)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmpgt struct{}

func (*ParserICif_icmpgt) Op() ops.Op { return ops.If_icmpgt }
func (*ParserICif_icmpgt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmpgt)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmple struct{}

func (*ParserICif_icmple) Op() ops.Op { return ops.If_icmple }
func (*ParserICif_icmple) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmple)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmplt struct{}

func (*ParserICif_icmplt) Op() ops.Op { return ops.If_icmplt }
func (*ParserICif_icmplt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmplt)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICif_icmpne struct{}

func (*ParserICif_icmpne) Op() ops.Op { return ops.If_icmpne }
func (*ParserICif_icmpne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICif_icmpne)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifeq struct{}

func (*ParserICifeq) Op() ops.Op { return ops.Ifeq }
func (*ParserICifeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifeq)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifge struct{}

func (*ParserICifge) Op() ops.Op { return ops.Ifge }
func (*ParserICifge) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifge)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifgt struct{}

func (*ParserICifgt) Op() ops.Op { return ops.Ifgt }
func (*ParserICifgt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifgt)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifle struct{}

func (*ParserICifle) Op() ops.Op { return ops.Ifle }
func (*ParserICifle) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifle)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICiflt struct{}

func (*ParserICiflt) Op() ops.Op { return ops.Iflt }
func (*ParserICiflt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICiflt)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifne struct{}

func (*ParserICifne) Op() ops.Op { return ops.Ifne }
func (*ParserICifne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifne)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifnonnull struct{}

func (*ParserICifnonnull) Op() ops.Op { return ops.Ifnonnull }
func (*ParserICifnonnull) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifnonnull)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICifnull struct{}

func (*ParserICifnull) Op() ops.Op { return ops.Ifnull }
func (*ParserICifnull) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICifnull)
	if ic.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserIClookupswitch struct{}

func (*ParserIClookupswitch) Op() ops.Op { return ops.Lookupswitch }
func (*ParserIClookupswitch) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.IClookupswitch)
	if ic.DefaultOffset, err = readInt32(br); err != nil {
		return nil, err
	}
	indexCount, err := readInt32(br)
	if err != nil {
		return nil, err
	}
	ic.Indexes = make([]ir.CaseEntry, indexCount)
	for i := range indexCount {
		var entry ir.CaseEntry
		if entry.K, err = readInt32(br); err != nil {
			return nil, err
		}
		if entry.V, err = readInt32(br); err != nil {
			return nil, err
		}
		ic.Indexes[i] = entry
	}
	slices.SortFunc(ic.Indexes, ir.CaseEntry.Cmp)
	return ic, nil
}

type ParserICtableswitch struct{}

func (*ParserICtableswitch) Op() ops.Op { return ops.Tableswitch }
func (*ParserICtableswitch) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICtableswitch)
	if ic.DefaultOffset, err = readInt32(br); err != nil {
		return nil, err
	}
	if ic.Low, err = readInt32(br); err != nil {
		return nil, err
	}
	if ic.High, err = readInt32(br); err != nil {
		return nil, err
	}
	indexCount := ic.High - ic.Low + 1
	ic.OffsetList = make([]int32, indexCount)
	for i := range indexCount {
		if ic.OffsetList[i], err = readInt32(br); err != nil {
			return nil, err
		}
	}
	return ic, nil
}

type ParserICwide struct{}

func (*ParserICwide) Op() ops.Op { return ops.Wide }
func (*ParserICwide) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICwide)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.OpCode = (ops.Op)(b)
	switch ic.OpCode {
	case ops.Iload, ops.Fload, ops.Aload, ops.Lload, ops.Dload, ops.Istore, ops.Fstore, ops.Astore, ops.Lstore, ops.Dstore:
		if ic.Index, err = readUint16(br); err != nil {
			return nil, err
		}
	case ops.Iinc:
		if ic.Index, err = readUint16(br); err != nil {
			return nil, err
		}
		if ic.Const, err = readUint16(br); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("parser: wide: unexpected opcode %d", ic.OpCode)
	}
	return ic, nil
}

type ParserICaload struct{}

func (*ParserICaload) Op() ops.Op { return ops.Aload }
func (*ParserICaload) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICaload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICanewarray struct{}

func (*ParserICanewarray) Op() ops.Op { return ops.Anewarray }
func (*ParserICanewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICanewarray)
	if ic.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICastore struct{}

func (*ParserICastore) Op() ops.Op { return ops.Astore }
func (*ParserICastore) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICastore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICcheckcast struct{}

func (*ParserICcheckcast) Op() ops.Op { return ops.Checkcast }
func (*ParserICcheckcast) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICcheckcast)
	if ic.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICgetfield struct{}

func (*ParserICgetfield) Op() ops.Op { return ops.Getfield }
func (*ParserICgetfield) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICgetfield)
	if ic.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICgetstatic struct{}

func (*ParserICgetstatic) Op() ops.Op { return ops.Getstatic }
func (*ParserICgetstatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICgetstatic)
	if ic.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICinstanceof struct{}

func (*ParserICinstanceof) Op() ops.Op { return ops.Instanceof }
func (*ParserICinstanceof) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinstanceof)
	if ic.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICinvokedynamic struct{}

func (*ParserICinvokedynamic) Op() ops.Op { return ops.Invokedynamic }
func (*ParserICinvokedynamic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinvokedynamic)
	if ic.Method, err = readUint16(br); err != nil {
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
	return ic, nil
}

type ParserICinvokeinterface struct{}

func (*ParserICinvokeinterface) Op() ops.Op { return ops.Invokeinterface }
func (*ParserICinvokeinterface) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinvokeinterface)
	if ic.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	if ic.Count, err = br.ReadByte(); err != nil {
		return nil, err
	}
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if b != 0 {
		panic("ir.invokeinterface: operands [3] must be 0")
	}
	return ic, nil
}

type ParserICinvokespecial struct{}

func (*ParserICinvokespecial) Op() ops.Op { return ops.Invokespecial }
func (*ParserICinvokespecial) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinvokespecial)
	if ic.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICinvokestatic struct{}

func (*ParserICinvokestatic) Op() ops.Op { return ops.Invokestatic }
func (*ParserICinvokestatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinvokestatic)
	if ic.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICinvokevirtual struct{}

func (*ParserICinvokevirtual) Op() ops.Op { return ops.Invokevirtual }
func (*ParserICinvokevirtual) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICinvokevirtual)
	if ic.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICmultianewarray struct{}

func (*ParserICmultianewarray) Op() ops.Op { return ops.Multianewarray }
func (*ParserICmultianewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICmultianewarray)
	if ic.Desc, err = readUint16(br); err != nil {
		return nil, err
	}
	if ic.Dimensions, err = br.ReadByte(); err != nil {
		return nil, err
	}
	if ic.Dimensions < 1 {
		panic("ir.multianewarray: dimensions is less than 1")
	}
	return ic, nil
}

type ParserICnew struct{}

func (*ParserICnew) Op() ops.Op { return ops.New }
func (*ParserICnew) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICnew)
	if ic.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICnewarray struct{}

func (*ParserICnewarray) Op() ops.Op { return ops.Newarray }
func (*ParserICnewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICnewarray)
	if ic.Atype, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICputfield struct{}

func (*ParserICputfield) Op() ops.Op { return ops.Putfield }
func (*ParserICputfield) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICputfield)
	if ic.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICputstatic struct{}

func (*ParserICputstatic) Op() ops.Op { return ops.Putstatic }
func (*ParserICputstatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICputstatic)
	if ic.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICbipush struct{}

func (*ParserICbipush) Op() ops.Op { return ops.Bipush }
func (*ParserICbipush) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICbipush)
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	ic.Value = (int8)(b)
	return ic, nil
}

type ParserICsipush struct{}

func (*ParserICsipush) Op() ops.Op { return ops.Sipush }
func (*ParserICsipush) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICsipush)
	if ic.Value, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICdload struct{}

func (*ParserICdload) Op() ops.Op { return ops.Dload }
func (*ParserICdload) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICdload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICdstore struct{}

func (*ParserICdstore) Op() ops.Op { return ops.Dstore }
func (*ParserICdstore) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICdstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICfload struct{}

func (*ParserICfload) Op() ops.Op { return ops.Fload }
func (*ParserICfload) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICfload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICfstore struct{}

func (*ParserICfstore) Op() ops.Op { return ops.Fstore }
func (*ParserICfstore) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICfstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICiinc struct{}

func (*ParserICiinc) Op() ops.Op { return ops.Iinc }
func (*ParserICiinc) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICiinc)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	if ic.Const, err = readInt16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICiload struct{}

func (*ParserICiload) Op() ops.Op { return ops.Iload }
func (*ParserICiload) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICiload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICistore struct{}

func (*ParserICistore) Op() ops.Op { return ops.Istore }
func (*ParserICistore) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.ICistore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserICldc struct{}

func (*ParserICldc) Op() ops.Op { return ops.Ldc }
func (*ParserICldc) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICldc)
	if ic.Index, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICldc_w struct{}

func (*ParserICldc_w) Op() ops.Op { return ops.Ldc_w }
func (*ParserICldc_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICldc_w)
	if ic.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserICldc2_w struct{}

func (*ParserICldc2_w) Op() ops.Op { return ops.Ldc2_w }
func (*ParserICldc2_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ic := new(ir.ICldc2_w)
	if ic.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ic, nil
}

type ParserIClload struct{}

func (*ParserIClload) Op() ops.Op { return ops.Lload }
func (*ParserIClload) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.IClload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

type ParserIClstore struct{}

func (*ParserIClstore) Op() ops.Op { return ops.Lstore }
func (*ParserIClstore) Parse(br ByteReader) (ir.IC, error) {
	ic := new(ir.IClstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ic.Index = (uint16)(b)
	return ic, nil
}

func init() {
	RegisterParser((*ParserICgoto)(nil))
	RegisterParser((*ParserICgoto_w)(nil))
	RegisterParser((*ParserICif_acmpeq)(nil))
	RegisterParser((*ParserICif_acmpne)(nil))
	RegisterParser((*ParserICif_icmpeq)(nil))
	RegisterParser((*ParserICif_icmpge)(nil))
	RegisterParser((*ParserICif_icmpgt)(nil))
	RegisterParser((*ParserICif_icmple)(nil))
	RegisterParser((*ParserICif_icmplt)(nil))
	RegisterParser((*ParserICif_icmpne)(nil))
	RegisterParser((*ParserICifeq)(nil))
	RegisterParser((*ParserICifge)(nil))
	RegisterParser((*ParserICifgt)(nil))
	RegisterParser((*ParserICifle)(nil))
	RegisterParser((*ParserICiflt)(nil))
	RegisterParser((*ParserICifne)(nil))
	RegisterParser((*ParserICifnonnull)(nil))
	RegisterParser((*ParserICifnull)(nil))
	RegisterParser((*ParserIClookupswitch)(nil))
	RegisterParser((*ParserICtableswitch)(nil))
	RegisterParser((*ParserICwide)(nil))
	RegisterParser((*ParserICaload)(nil))
	RegisterParser((*ParserICanewarray)(nil))
	RegisterParser((*ParserICastore)(nil))
	RegisterParser((*ParserICcheckcast)(nil))
	RegisterParser((*ParserICgetfield)(nil))
	RegisterParser((*ParserICgetstatic)(nil))
	RegisterParser((*ParserICinstanceof)(nil))
	RegisterParser((*ParserICinvokedynamic)(nil))
	RegisterParser((*ParserICinvokeinterface)(nil))
	RegisterParser((*ParserICinvokespecial)(nil))
	RegisterParser((*ParserICinvokestatic)(nil))
	RegisterParser((*ParserICinvokevirtual)(nil))
	RegisterParser((*ParserICmultianewarray)(nil))
	RegisterParser((*ParserICnew)(nil))
	RegisterParser((*ParserICnewarray)(nil))
	RegisterParser((*ParserICputfield)(nil))
	RegisterParser((*ParserICputstatic)(nil))
	RegisterParser((*ParserICbipush)(nil))
	RegisterParser((*ParserICsipush)(nil))
	RegisterParser((*ParserICdload)(nil))
	RegisterParser((*ParserICdstore)(nil))
	RegisterParser((*ParserICfload)(nil))
	RegisterParser((*ParserICfstore)(nil))
	RegisterParser((*ParserICiinc)(nil))
	RegisterParser((*ParserICiload)(nil))
	RegisterParser((*ParserICistore)(nil))
	RegisterParser((*ParserICldc)(nil))
	RegisterParser((*ParserICldc_w)(nil))
	RegisterParser((*ParserICldc2_w)(nil))
	RegisterParser((*ParserIClload)(nil))
	RegisterParser((*ParserIClstore)(nil))
}
