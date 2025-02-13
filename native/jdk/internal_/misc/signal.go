package jdk_internal_misc

import (
	"github.com/LiterMC/wasm-jdk/ir"
	"github.com/LiterMC/wasm-jdk/native"
)

func init() {
	native.RegisterDefaultNative("jdk/internal/misc/Signal.findSignal0(Ljava/lang/String;)I", Signal_findSignal0)
	native.RegisterDefaultNative("jdk/internal/misc/Signal.handle0(IJ)J", Signal_handle0)
	native.RegisterDefaultNative("jdk/internal/misc/Signal.raise0(I)V", Signal_raise0)
}

const (
	SIGHUP  = 0x01
	SIGINT  = 0x02
	SIGKILL = 0x03
	SIGQUIT = 0x04
	SIGTERM = 0x05
	SIGTRAP = 0x06
	SIGUSR1 = 0x10
	SIGUSR2 = 0x11
)

// private static native int findSignal0(String sigName);
func Signal_findSignal0(vm ir.VM) error {
	stack := vm.GetStack()
	name := vm.GetString(stack.GetVarRef(0))
	var sig int32 = -1
	switch name {
	case "HUP":
		sig = SIGHUP
	case "INT":
		sig = SIGINT
	case "KILL":
		sig = SIGKILL
	case "QUIT":
		sig = SIGQUIT
	case "TERM":
		sig = SIGTERM
	case "TRAP":
		sig = SIGTRAP
	case "USR1":
		sig = SIGUSR1
	case "USR2":
		sig = SIGUSR2
	default:
		panic("getting signal " + name)
	}
	stack.PushInt32(sig)
	return nil
}

// private static native long handle0(int sig, long nativeH);
func Signal_handle0(vm ir.VM) error {
	stack := vm.GetStack()
	sig := stack.GetVarInt32(0)
	handler := stack.GetVarInt64(1)
	_ = sig
	switch handler {
	case 0:
		stack.PushInt64(1)
	case 1:
		stack.PushInt64(1)
	case 2:
		stack.PushInt64(1)
	default:
		panic("unexpected signal handle value")
	}
	return nil
}

// private static native void raise0(int sig);
func Signal_raise0(vm ir.VM) error {
	stack := vm.GetStack()
	sig := stack.GetVarInt32(0)
	_ = sig
	return nil
}
