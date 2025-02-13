package java_lang

import (
	"math"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/Float.floatToRawIntBits(F)I", Float_floatToRawIntBits)
	native.RegisterDefaultNative("java/lang/Float.intBitsToFloat(I)F", Float_intBitsToFloat)
}

// public static native int floatToRawIntBits(float value);
func Float_floatToRawIntBits(vm ir.VM) error {
	stack := vm.GetStack()
	v := stack.GetVarFloat32(0)
	stack.PushInt32((int32)(math.Float32bits(v)))
	return nil
}

// public static native float intBitsToFloat(int bits);
func Float_intBitsToFloat(vm ir.VM) error {
	stack := vm.GetStack()
	v := stack.GetVarInt32(0)
	stack.PushFloat32(math.Float32frombits((uint32)(v)))
	return nil
}
