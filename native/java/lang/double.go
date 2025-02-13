package java_lang

import (
	"math"

	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("java/lang/Double.doubleToRawLongBits(D)J", Double_doubleToRawLongBits)
	native.RegisterDefaultNative("java/lang/Double.longBitsToDouble(J)D", Double_longBitsToDouble)
}

// public static native long doubleToRawLongBits(double value);
func Double_doubleToRawLongBits(vm ir.VM) error {
	stack := vm.GetStack()
	v := stack.GetVarFloat64(0)
	stack.PushInt64((int64)(math.Float64bits(v)))
	return nil
}

// public static native double longBitsToDouble(long bits);
func Double_longBitsToDouble(vm ir.VM) error {
	stack := vm.GetStack()
	v := stack.GetVarInt64(0)
	stack.PushFloat64(math.Float64frombits((uint64)(v)))
	return nil
}
