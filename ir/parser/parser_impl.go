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
	ir := new(ir.ICgoto)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICgoto_w struct{}

func (*ParserICgoto_w) Op() ops.Op { return ops.Goto_w }
func (*ParserICgoto_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICgoto_w)
	if ir.Offset, err = readInt32(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_acmpeq struct{}

func (*ParserICif_acmpeq) Op() ops.Op { return ops.If_acmpeq }
func (*ParserICif_acmpeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_acmpeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_acmpne struct{}

func (*ParserICif_acmpne) Op() ops.Op { return ops.If_acmpne }
func (*ParserICif_acmpne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_acmpne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmpeq struct{}

func (*ParserICif_icmpeq) Op() ops.Op { return ops.If_icmpeq }
func (*ParserICif_icmpeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmpeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmpge struct{}

func (*ParserICif_icmpge) Op() ops.Op { return ops.If_icmpge }
func (*ParserICif_icmpge) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmpge)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmpgt struct{}

func (*ParserICif_icmpgt) Op() ops.Op { return ops.If_icmpgt }
func (*ParserICif_icmpgt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmpgt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmple struct{}

func (*ParserICif_icmple) Op() ops.Op { return ops.If_icmple }
func (*ParserICif_icmple) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmple)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmplt struct{}

func (*ParserICif_icmplt) Op() ops.Op { return ops.If_icmplt }
func (*ParserICif_icmplt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmplt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICif_icmpne struct{}

func (*ParserICif_icmpne) Op() ops.Op { return ops.If_icmpne }
func (*ParserICif_icmpne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICif_icmpne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifeq struct{}

func (*ParserICifeq) Op() ops.Op { return ops.Ifeq }
func (*ParserICifeq) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifeq)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifge struct{}

func (*ParserICifge) Op() ops.Op { return ops.Ifge }
func (*ParserICifge) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifge)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifgt struct{}

func (*ParserICifgt) Op() ops.Op { return ops.Ifgt }
func (*ParserICifgt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifgt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifle struct{}

func (*ParserICifle) Op() ops.Op { return ops.Ifle }
func (*ParserICifle) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifle)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICiflt struct{}

func (*ParserICiflt) Op() ops.Op { return ops.Iflt }
func (*ParserICiflt) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICiflt)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifne struct{}

func (*ParserICifne) Op() ops.Op { return ops.Ifne }
func (*ParserICifne) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifne)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifnonnull struct{}

func (*ParserICifnonnull) Op() ops.Op { return ops.Ifnonnull }
func (*ParserICifnonnull) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifnonnull)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICifnull struct{}

func (*ParserICifnull) Op() ops.Op { return ops.Ifnull }
func (*ParserICifnull) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICifnull)
	if ir.Offset, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIClookupswitch struct{}

type irCaseEntry = ir.CaseEntry

func (*ParserIClookupswitch) Op() ops.Op { return ops.Lookupswitch }
func (*ParserIClookupswitch) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.IClookupswitch)
	if ir.DefaultOffset, err = readInt32(br); err != nil {
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

type ParserICtableswitch struct{}

func (*ParserICtableswitch) Op() ops.Op { return ops.Tableswitch }
func (*ParserICtableswitch) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICtableswitch)
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

type ParserICwide struct{}

func (*ParserICwide) Op() ops.Op { return ops.Wide }
func (*ParserICwide) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICwide)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.OpCode = (ops.Op)(b)
	switch ir.OpCode {
	case ops.Iload, ops.Fload, ops.Aload, ops.Lload, ops.Dload, ops.Istore, ops.Fstore, ops.Astore, ops.Lstore, ops.Dstore:
		if ir.Index, err = readUint16(br); err != nil {
			return nil, err
		}
	case ops.Iinc:
		if ir.Index, err = readUint16(br); err != nil {
			return nil, err
		}
		if ir.Const, err = readUint16(br); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("parser: wide: unexpected opcode %d", ir.OpCode)
	}
	return ir, nil
}

type ParserICaload struct{}

func (*ParserICaload) Op() ops.Op { return ops.Aload }
func (*ParserICaload) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICaload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICanewarray struct{}

func (*ParserICanewarray) Op() ops.Op { return ops.Anewarray }
func (*ParserICanewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICanewarray)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICastore struct{}

func (*ParserICastore) Op() ops.Op { return ops.Astore }
func (*ParserICastore) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICastore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICcheckcast struct{}

func (*ParserICcheckcast) Op() ops.Op { return ops.Checkcast }
func (*ParserICcheckcast) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICcheckcast)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICgetfield struct{}

func (*ParserICgetfield) Op() ops.Op { return ops.Getfield }
func (*ParserICgetfield) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICgetfield)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICgetstatic struct{}

func (*ParserICgetstatic) Op() ops.Op { return ops.Getstatic }
func (*ParserICgetstatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICgetstatic)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICinstanceof struct{}

func (*ParserICinstanceof) Op() ops.Op { return ops.Instanceof }
func (*ParserICinstanceof) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinstanceof)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICinvokedynamic struct{}

func (*ParserICinvokedynamic) Op() ops.Op { return ops.Invokedynamic }
func (*ParserICinvokedynamic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinvokedynamic)
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

type ParserICinvokeinterface struct{}

func (*ParserICinvokeinterface) Op() ops.Op { return ops.Invokeinterface }
func (*ParserICinvokeinterface) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinvokeinterface)
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

type ParserICinvokespecial struct{}

func (*ParserICinvokespecial) Op() ops.Op { return ops.Invokespecial }
func (*ParserICinvokespecial) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinvokespecial)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICinvokestatic struct{}

func (*ParserICinvokestatic) Op() ops.Op { return ops.Invokestatic }
func (*ParserICinvokestatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinvokestatic)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICinvokevirtual struct{}

func (*ParserICinvokevirtual) Op() ops.Op { return ops.Invokevirtual }
func (*ParserICinvokevirtual) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICinvokevirtual)
	if ir.Method, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICmultianewarray struct{}

func (*ParserICmultianewarray) Op() ops.Op { return ops.Multianewarray }
func (*ParserICmultianewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICmultianewarray)
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

type ParserICnew struct{}

func (*ParserICnew) Op() ops.Op { return ops.New }
func (*ParserICnew) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICnew)
	if ir.Class, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICnewarray struct{}

func (*ParserICnewarray) Op() ops.Op { return ops.Newarray }
func (*ParserICnewarray) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICnewarray)
	if ir.Atype, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICputfield struct{}

func (*ParserICputfield) Op() ops.Op { return ops.Putfield }
func (*ParserICputfield) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICputfield)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICputstatic struct{}

func (*ParserICputstatic) Op() ops.Op { return ops.Putstatic }
func (*ParserICputstatic) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICputstatic)
	if ir.Field, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICbipush struct{}

func (*ParserICbipush) Op() ops.Op { return ops.Bipush }
func (*ParserICbipush) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICbipush)
	var b byte
	if b, err = br.ReadByte(); err != nil {
		return nil, err
	}
	ir.Value = (int8)(b)
	return ir, nil
}

type ParserICsipush struct{}

func (*ParserICsipush) Op() ops.Op { return ops.Sipush }
func (*ParserICsipush) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICsipush)
	if ir.Value, err = readInt16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICdload struct{}

func (*ParserICdload) Op() ops.Op { return ops.Dload }
func (*ParserICdload) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICdload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICdstore struct{}

func (*ParserICdstore) Op() ops.Op { return ops.Dstore }
func (*ParserICdstore) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICdstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICfload struct{}

func (*ParserICfload) Op() ops.Op { return ops.Fload }
func (*ParserICfload) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICfload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICfstore struct{}

func (*ParserICfstore) Op() ops.Op { return ops.Fstore }
func (*ParserICfstore) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICfstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICiinc struct{}

func (*ParserICiinc) Op() ops.Op { return ops.Iinc }
func (*ParserICiinc) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICiinc)
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

type ParserICiload struct{}

func (*ParserICiload) Op() ops.Op { return ops.Iload }
func (*ParserICiload) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICiload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICistore struct{}

func (*ParserICistore) Op() ops.Op { return ops.Istore }
func (*ParserICistore) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.ICistore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserICldc struct{}

func (*ParserICldc) Op() ops.Op { return ops.Ldc }
func (*ParserICldc) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICldc)
	if ir.Index, err = br.ReadByte(); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICldc_w struct{}

func (*ParserICldc_w) Op() ops.Op { return ops.Ldc_w }
func (*ParserICldc_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICldc_w)
	if ir.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserICldc2_w struct{}

func (*ParserICldc2_w) Op() ops.Op { return ops.Ldc2_w }
func (*ParserICldc2_w) Parse(br ByteReader) (ir.IC, error) {
	var err error
	ir := new(ir.ICldc2_w)
	if ir.Index, err = readUint16(br); err != nil {
		return nil, err
	}
	return ir, nil
}

type ParserIClload struct{}

func (*ParserIClload) Op() ops.Op { return ops.Lload }
func (*ParserIClload) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.IClload)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
}

type ParserIClstore struct{}

func (*ParserIClstore) Op() ops.Op { return ops.Lstore }
func (*ParserIClstore) Parse(br ByteReader) (ir.IC, error) {
	ir := new(ir.IClstore)
	b, err := br.ReadByte()
	if err != nil {
		return nil, err
	}
	ir.Index = (uint16)(b)
	return ir, nil
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
