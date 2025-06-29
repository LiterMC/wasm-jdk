package java_lang_invoke

import (
	jvm "github.com/LiterMC/wasm-jdk/vm"
)

type ResolvedMethodNameData struct {
	VMTarget *jvm.Method
	VMHolder *jvm.Class
}

type MemberName struct {
	// private Class<?> clazz;       // class in which the member is defined
	Clazz *jvm.Ref
	// private String   name;        // may be null if not yet materialized
	Name *jvm.Ref
	// private Object   type;        // may be null if not yet materialized
	Type *jvm.Ref
	// private int      flags;       // modifier bits; see reflect.Modifier
	Flags int32
	// private ResolvedMethodName method;    // cached resolved method information
	Method *jvm.Ref
}

type MemberNameData struct {
	VMIndex int64
}
