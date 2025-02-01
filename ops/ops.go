package ops

// Reference: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-7.html

type Op byte

//go:generate stringer -linecomment -type=Op
const (
	// Constants

	Nop         Op = 0x00 // nop
	Aconst_null Op = 0x01 // aconst_null
	Iconst_m1   Op = 0x02 // iconst_m1
	Iconst_0    Op = 0x03 // iconst_0
	Iconst_1    Op = 0x04 // iconst_1
	Iconst_2    Op = 0x05 // iconst_2
	Iconst_3    Op = 0x06 // iconst_3
	Iconst_4    Op = 0x07 // iconst_4
	Iconst_5    Op = 0x08 // iconst_5
	Lconst_0    Op = 0x09 // lconst_0
	Lconst_1    Op = 0x0a // lconst_1
	Fconst_0    Op = 0x0b // fconst_0
	Fconst_1    Op = 0x0c // fconst_1
	Fconst_2    Op = 0x0d // fconst_2
	Dconst_0    Op = 0x0e // dconst_0
	Dconst_1    Op = 0x0f // dconst_1
	Bipush      Op = 0x10 // bipush
	Sipush      Op = 0x11 // sipush
	Ldc         Op = 0x12 // ldc
	Ldc_w       Op = 0x13 // ldc_w
	Ldc2_w      Op = 0x14 // ldc2_w

	// Loads

	Iload   Op = 0x15 // iload
	Lload   Op = 0x16 // lload
	Fload   Op = 0x17 // fload
	Dload   Op = 0x18 // dload
	Aload   Op = 0x19 // aload
	Iload_0 Op = 0x1a // iload_0
	Iload_1 Op = 0x1b // iload_1
	Iload_2 Op = 0x1c // iload_2
	Iload_3 Op = 0x1d // iload_3
	Lload_0 Op = 0x1e // lload_0
	Lload_1 Op = 0x1f // lload_1
	Lload_2 Op = 0x20 // lload_2
	Lload_3 Op = 0x21 // lload_3
	Fload_0 Op = 0x22 // fload_0
	Fload_1 Op = 0x23 // fload_1
	Fload_2 Op = 0x24 // fload_2
	Fload_3 Op = 0x25 // fload_3
	Dload_0 Op = 0x26 // dload_0
	Dload_1 Op = 0x27 // dload_1
	Dload_2 Op = 0x28 // dload_2
	Dload_3 Op = 0x29 // dload_3
	Aload_0 Op = 0x2a // aload_0
	Aload_1 Op = 0x2b // aload_1
	Aload_2 Op = 0x2c // aload_2
	Aload_3 Op = 0x2d // aload_3
	Iaload  Op = 0x2e // iaload
	Laload  Op = 0x2f // laload
	Faload  Op = 0x30 // faload
	Daload  Op = 0x31 // daload
	Aaload  Op = 0x32 // aaload
	Baload  Op = 0x33 // baload
	Caload  Op = 0x34 // caload
	Saload  Op = 0x35 // saload

	// Stores

	Istore   Op = 0x36 // istore
	Lstore   Op = 0x37 // lstore
	Fstore   Op = 0x38 // fstore
	Dstore   Op = 0x39 // dstore
	Astore   Op = 0x3a // astore
	Istore_0 Op = 0x3b // istore_0
	Istore_1 Op = 0x3c // istore_1
	Istore_2 Op = 0x3d // istore_2
	Istore_3 Op = 0x3e // istore_3
	Lstore_0 Op = 0x3f // lstore_0
	Lstore_1 Op = 0x40 // lstore_1
	Lstore_2 Op = 0x41 // lstore_2
	Lstore_3 Op = 0x42 // lstore_3
	Fstore_0 Op = 0x43 // fstore_0
	Fstore_1 Op = 0x44 // fstore_1
	Fstore_2 Op = 0x45 // fstore_2
	Fstore_3 Op = 0x46 // fstore_3
	Dstore_0 Op = 0x47 // dstore_0
	Dstore_1 Op = 0x48 // dstore_1
	Dstore_2 Op = 0x49 // dstore_2
	Dstore_3 Op = 0x4a // dstore_3
	Astore_0 Op = 0x4b // astore_0
	Astore_1 Op = 0x4c // astore_1
	Astore_2 Op = 0x4d // astore_2
	Astore_3 Op = 0x4e // astore_3
	Iastore  Op = 0x4f // iastore
	Lastore  Op = 0x50 // lastore
	Fastore  Op = 0x51 // fastore
	Dastore  Op = 0x52 // dastore
	Aastore  Op = 0x53 // aastore
	Bastore  Op = 0x54 // bastore
	Castore  Op = 0x55 // castore
	Sastore  Op = 0x56 // sastore

	// Stack

	Pop     Op = 0x57 // pop
	Pop2    Op = 0x58 // pop2
	Dup     Op = 0x59 // dup
	Dup_x1  Op = 0x5a // dup_x1
	Dup_x2  Op = 0x5b // dup_x2
	Dup2    Op = 0x5c // dup2
	Dup2_x1 Op = 0x5d // dup2_x1
	Dup2_x2 Op = 0x5e // dup2_x2
	Swap    Op = 0x5f // swap

	// Math

	Iadd  Op = 0x60 // iadd
	Ladd  Op = 0x61 // ladd
	Fadd  Op = 0x62 // fadd
	Dadd  Op = 0x63 // dadd
	Isub  Op = 0x64 // isub
	Lsub  Op = 0x65 // lsub
	Fsub  Op = 0x66 // fsub
	Dsub  Op = 0x67 // dsub
	Imul  Op = 0x68 // imul
	Lmul  Op = 0x69 // lmul
	Fmul  Op = 0x6a // fmul
	Dmul  Op = 0x6b // dmul
	Idiv  Op = 0x6c // idiv
	Ldiv  Op = 0x6d // ldiv
	Fdiv  Op = 0x6e // fdiv
	Ddiv  Op = 0x6f // ddiv
	Irem  Op = 0x70 // irem
	Lrem  Op = 0x71 // lrem
	Frem  Op = 0x72 // frem
	Drem  Op = 0x73 // drem
	Ineg  Op = 0x74 // ineg
	Lneg  Op = 0x75 // lneg
	Fneg  Op = 0x76 // fneg
	Dneg  Op = 0x77 // dneg
	Ishl  Op = 0x78 // ishl
	Lshl  Op = 0x79 // lshl
	Ishr  Op = 0x7a // ishr
	Lshr  Op = 0x7b // lshr
	Iushr Op = 0x7c // iushr
	Lushr Op = 0x7d // lushr
	Iand  Op = 0x7e // iand
	Land  Op = 0x7f // land
	Ior   Op = 0x80 // ior
	Lor   Op = 0x81 // lor
	Ixor  Op = 0x82 // ixor
	Lxor  Op = 0x83 // lxor
	Iinc  Op = 0x84 // iinc

	// Conversions

	I2l Op = 0x85 // i2l
	I2f Op = 0x86 // i2f
	I2d Op = 0x87 // i2d
	L2i Op = 0x88 // l2i
	L2f Op = 0x89 // l2f
	L2d Op = 0x8a // l2d
	F2i Op = 0x8b // f2i
	F2l Op = 0x8c // f2l
	F2d Op = 0x8d // f2d
	D2i Op = 0x8e // d2i
	D2l Op = 0x8f // d2l
	D2f Op = 0x90 // d2f
	I2b Op = 0x91 // i2b
	I2c Op = 0x92 // i2c
	I2s Op = 0x93 // i2s

	// Comparisons

	Lcmp      Op = 0x94 // lcmp
	Fcmpl     Op = 0x95 // fcmpl
	Fcmpg     Op = 0x96 // fcmpg
	Dcmpl     Op = 0x97 // dcmpl
	Dcmpg     Op = 0x98 // dcmpg
	Ifeq      Op = 0x99 // ifeq
	Ifne      Op = 0x9a // ifne
	Iflt      Op = 0x9b // iflt
	Ifge      Op = 0x9c // ifge
	Ifgt      Op = 0x9d // ifgt
	Ifle      Op = 0x9e // ifle
	If_icmpeq Op = 0x9f // if_icmpeq
	If_icmpne Op = 0xa0 // if_icmpne
	If_icmplt Op = 0xa1 // if_icmplt
	If_icmpge Op = 0xa2 // if_icmpge
	If_icmpgt Op = 0xa3 // if_icmpgt
	If_icmple Op = 0xa4 // if_icmple
	If_acmpeq Op = 0xa5 // if_acmpeq
	If_acmpne Op = 0xa6 // if_acmpne

	// Control

	Goto         Op = 0xa7 // goto
	Jsr          Op = 0xa8 // jsr
	Ret          Op = 0xa9 // ret
	Tableswitch  Op = 0xaa // tableswitch
	Lookupswitch Op = 0xab // lookupswitch
	Ireturn      Op = 0xac // ireturn
	Lreturn      Op = 0xad // lreturn
	Freturn      Op = 0xae // freturn
	Dreturn      Op = 0xaf // dreturn
	Areturn      Op = 0xb0 // areturn
	Return       Op = 0xb1 // return

	// References

	Getstatic       Op = 0xb2 // getstatic
	Putstatic       Op = 0xb3 // putstatic
	Getfield        Op = 0xb4 // getfield
	Putfield        Op = 0xb5 // putfield
	Invokevirtual   Op = 0xb6 // invokevirtual
	Invokespecial   Op = 0xb7 // invokespecial
	Invokestatic    Op = 0xb8 // invokestatic
	Invokeinterface Op = 0xb9 // invokeinterface
	Invokedynamic   Op = 0xba // invokedynamic
	New             Op = 0xbb // new
	Newarray        Op = 0xbc // newarray
	Anewarray       Op = 0xbd // anewarray
	Arraylength     Op = 0xbe // arraylength
	Athrow          Op = 0xbf // athrow
	Checkcast       Op = 0xc0 // checkcast
	Instanceof      Op = 0xc1 // instanceof
	Monitorenter    Op = 0xc2 // monitorenter
	Monitorexit     Op = 0xc3 // monitorexit

	// Extended

	Wide           Op = 0xc4 // wide
	Multianewarray Op = 0xc5 // multianewarray
	Ifnull         Op = 0xc6 // ifnull
	Ifnonnull      Op = 0xc7 // ifnonnull
	Goto_w         Op = 0xc8 // goto_w
	Jsr_w          Op = 0xc9 // jsr_w

	// Reserved

	Breakpoint Op = 0xca // breakpoint
	Impdep1    Op = 0xfe // impdep1
	Impdep2    Op = 0xff // impdep2
)
