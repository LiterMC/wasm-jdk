
package errs

import (
	"errors"
)

var (
	ArrayIndexOutOfBoundsException = errors.New("ArrayIndexOutOfBoundsException")
	BootstrapMethodError = errors.New("BootstrapMethodError")
	ClassCastException = errors.New("ClassCastException")
	IllegalMonitorStateException = errors.New("IllegalMonitorStateException")
	IncompatibleClassChangeError = errors.New("IncompatibleClassChangeError")
	NegativeArraySizeException = errors.New("NegativeArraySizeException")
	NoSuchFieldError = errors.New("NoSuchFieldError")
	NoSuchMethodError = errors.New("NoSuchMethodError")
	NullPointerException = errors.New("NullPointerException")
)
