package ops

// Reference: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-7.html

type Op byte

const (
	// Constants

	Nop         Op = 0x00
	Aconst_null Op = 0x01
	Iconst_m1   Op = 0x02
	Iconst_0    Op = 0x03
	Iconst_1    Op = 0x04
	Iconst_2    Op = 0x05
	Iconst_3    Op = 0x06
	Iconst_4    Op = 0x07
	Iconst_5    Op = 0x08
	Lconst_0    Op = 0x09
	Lconst_1    Op = 0x0a
	Fconst_0    Op = 0x0b
	Fconst_1    Op = 0x0c
	Fconst_2    Op = 0x0d
	Dconst_0    Op = 0x0e
	Dconst_1    Op = 0x0f
	Bipush      Op = 0x10
	Sipush      Op = 0x11
	Ldc         Op = 0x12
	Ldc_w       Op = 0x13
	Ldc2_w      Op = 0x14

	// Loads

	Iload   Op = 0x15
	Lload   Op = 0x16
	Fload   Op = 0x17
	Dload   Op = 0x18
	Aload   Op = 0x19
	Iload_0 Op = 0x1a
	Iload_1 Op = 0x1b
	Iload_2 Op = 0x1c
	Iload_3 Op = 0x1d
	Lload_0 Op = 0x1e
	Lload_1 Op = 0x1f
	Lload_2 Op = 0x20
	Lload_3 Op = 0x21
	Fload_0 Op = 0x22
	Fload_1 Op = 0x23
	Fload_2 Op = 0x24
	Fload_3 Op = 0x25
	Dload_0 Op = 0x26
	Dload_1 Op = 0x27
	Dload_2 Op = 0x28
	Dload_3 Op = 0x29
	Aload_0 Op = 0x2a
	Aload_1 Op = 0x2b
	Aload_2 Op = 0x2c
	Aload_3 Op = 0x2d
	Iaload  Op = 0x2e
	Laload  Op = 0x2f
	Faload  Op = 0x30
	Daload  Op = 0x31
	Aaload  Op = 0x32
	Baload  Op = 0x33
	Caload  Op = 0x34
	Saload  Op = 0x35

	// Stores

	Istore   Op = 0x36
	Lstore   Op = 0x37
	Fstore   Op = 0x38
	Dstore   Op = 0x39
	Astore   Op = 0x3a
	Istore_0 Op = 0x3b
	Istore_1 Op = 0x3c
	Istore_2 Op = 0x3d
	Istore_3 Op = 0x3e
	Lstore_0 Op = 0x3f
	Lstore_1 Op = 0x40
	Lstore_2 Op = 0x41
	Lstore_3 Op = 0x42
	Fstore_0 Op = 0x43
	Fstore_1 Op = 0x44
	Fstore_2 Op = 0x45
	Fstore_3 Op = 0x46
	Dstore_0 Op = 0x47
	Dstore_1 Op = 0x48
	Dstore_2 Op = 0x49
	Dstore_3 Op = 0x4a
	Astore_0 Op = 0x4b
	Astore_1 Op = 0x4c
	Astore_2 Op = 0x4d
	Astore_3 Op = 0x4e
	Iastore  Op = 0x4f
	Lastore  Op = 0x50
	Fastore  Op = 0x51
	Dastore  Op = 0x52
	Aastore  Op = 0x53
	Bastore  Op = 0x54
	Castore  Op = 0x55
	Sastore  Op = 0x56

	// Stack

	Pop     Op = 0x57
	Pop2    Op = 0x58
	Dup     Op = 0x59
	Dup_x1  Op = 0x5a
	Dup_x2  Op = 0x5b
	Dup2    Op = 0x5c
	Dup2_x1 Op = 0x5d
	Dup2_x2 Op = 0x5e
	Swap    Op = 0x5f

	// Math

	Iadd  Op = 0x60
	Ladd  Op = 0x61
	Fadd  Op = 0x62
	Dadd  Op = 0x63
	Isub  Op = 0x64
	Lsub  Op = 0x65
	Fsub  Op = 0x66
	Dsub  Op = 0x67
	Imul  Op = 0x68
	Lmul  Op = 0x69
	Fmul  Op = 0x6a
	Dmul  Op = 0x6b
	Idiv  Op = 0x6c
	Ldiv  Op = 0x6d
	Fdiv  Op = 0x6e
	Ddiv  Op = 0x6f
	Irem  Op = 0x70
	Lrem  Op = 0x71
	Frem  Op = 0x72
	Drem  Op = 0x73
	Ineg  Op = 0x74
	Lneg  Op = 0x75
	Fneg  Op = 0x76
	Dneg  Op = 0x77
	Ishl  Op = 0x78
	Lshl  Op = 0x79
	Ishr  Op = 0x7a
	Lshr  Op = 0x7b
	Iushr Op = 0x7c
	Lushr Op = 0x7d
	Iand  Op = 0x7e
	Land  Op = 0x7f
	Ior   Op = 0x80
	Lor   Op = 0x81
	Ixor  Op = 0x82
	Lxor  Op = 0x83
	Iinc  Op = 0x84

	// Conversions

	I2l Op = 0x85
	I2f Op = 0x86
	I2d Op = 0x87
	L2i Op = 0x88
	L2f Op = 0x89
	L2d Op = 0x8a
	F2i Op = 0x8b
	F2l Op = 0x8c
	F2d Op = 0x8d
	D2i Op = 0x8e
	D2l Op = 0x8f
	D2f Op = 0x90
	I2b Op = 0x91
	I2c Op = 0x92
	I2s Op = 0x93

	// Comparisons

	Lcmp      Op = 0x94
	Fcmpl     Op = 0x95
	Fcmpg     Op = 0x96
	Dcmpl     Op = 0x97
	Dcmpg     Op = 0x98
	Ifeq      Op = 0x99
	Ifne      Op = 0x9a
	Iflt      Op = 0x9b
	Ifge      Op = 0x9c
	Ifgt      Op = 0x9d
	Ifle      Op = 0x9e
	If_icmpeq Op = 0x9f
	If_icmpne Op = 0xa0
	If_icmplt Op = 0xa1
	If_icmpge Op = 0xa2
	If_icmpgt Op = 0xa3
	If_icmple Op = 0xa4
	If_acmpeq Op = 0xa5
	If_acmpne Op = 0xa6

	// Control

	Goto         Op = 0xa7
	Jsr          Op = 0xa8
	Ret          Op = 0xa9
	Tableswitch  Op = 0xaa
	Lookupswitch Op = 0xab
	Ireturn      Op = 0xac
	Lreturn      Op = 0xad
	Freturn      Op = 0xae
	Dreturn      Op = 0xaf
	Areturn      Op = 0xb0
	Return       Op = 0xb1

	// References

	Getstatic       Op = 0xb2
	Putstatic       Op = 0xb3
	Getfield        Op = 0xb4
	Putfield        Op = 0xb5
	Invokevirtual   Op = 0xb6
	Invokespecial   Op = 0xb7
	Invokestatic    Op = 0xb8
	Invokeinterface Op = 0xb9
	Invokedynamic   Op = 0xba
	New             Op = 0xbb
	Newarray        Op = 0xbc
	Anewarray       Op = 0xbd
	Arraylength     Op = 0xbe
	Athrow          Op = 0xbf
	Checkcast       Op = 0xc0
	Instanceof      Op = 0xc1
	Monitorenter    Op = 0xc2
	Monitorexit     Op = 0xc3

	// Extended

	Wide           Op = 0xc4
	Multianewarray Op = 0xc5
	Ifnull         Op = 0xc6
	Ifnonnull      Op = 0xc7
	Goto_w         Op = 0xc8
	Jsr_w          Op = 0xc9

	// Reserved

	Breakpoint Op = 0xca
	Impdep1    Op = 0xfe
	Impdep2    Op = 0xff
)
