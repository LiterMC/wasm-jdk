package errs

import (
	"errors"
	"fmt"
)

var (
	ArrayIndexOutOfBoundsException = errors.New("ArrayIndexOutOfBoundsException")
	BootstrapMethodError           = errors.New("BootstrapMethodError")
	CloneNotSupportedException     = errors.New("CloneNotSupportedException")
	IllegalMonitorStateException   = errors.New("IllegalMonitorStateException")
	IncompatibleClassChangeError   = errors.New("IncompatibleClassChangeError")
	InterruptedException           = errors.New("InterruptedException")
	NegativeArraySizeException     = errors.New("NegativeArraySizeException")
	NoSuchFieldError               = errors.New("NoSuchFieldError")
	NoSuchMethodError              = errors.New("NoSuchMethodError")
	NullPointerException           = errors.New("NullPointerException")
)

type ClassCastException struct {
	Have string
	Want string
}

func (e *ClassCastException) Error() string {
	return fmt.Sprintf("ClassCastException: have %s, want %s", e.Have, e.Want)
}

type ClassNotFoundException struct {
	Class string
	Cause error
}

func (e *ClassNotFoundException) Error() string {
	return fmt.Sprintf("ClassNotFoundException: %s: %v", e.Class, e.Cause)
}

type UnsatisfiedLinkError struct {
	Name string
}

func (e *UnsatisfiedLinkError) Error() string {
	return fmt.Sprintf("UnsatisfiedLinkError: %s is not found", e.Name)
}
