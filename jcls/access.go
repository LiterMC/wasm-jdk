package jcls

// Source: https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.5-200-A.1
type AccessFlag uint16

const (
	// Declared public; may be accessed from outside its package.
	AccPublic AccessFlag = 0x0001
	// Declared private; accessible only within the defining class and other classes belonging to the same nest (§5.4.4).
	AccPrivate AccessFlag = 0x0002
	// Declared protected; may be accessed within subclasses.
	AccProtected AccessFlag = 0x0004
	// Declared static.
	AccStatic AccessFlag = 0x0008
	// Declared final; never directly assigned to after object construction (JLS §17.5).
	AccFinal AccessFlag = 0x0010
	// Declared volatile; cannot be cached.
	AccVolatile AccessFlag = 0x0040
	// Declared transient; not written or read by a persistent object manager.
	AccTransient AccessFlag = 0x0080
	// Declared synthetic; not present in the source code.
	AccSynthetic AccessFlag = 0x1000
	// Declared as an element of an enum class.
	AccEnum AccessFlag = 0x4000
)

func (a AccessFlag) Has(f AccessFlag) bool {
	return a&f != 0
}
