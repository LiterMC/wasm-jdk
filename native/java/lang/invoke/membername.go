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
	// private String   name;        // may be null if not yet materialized
	// private Object   type;        // may be null if not yet materialized
	// private int      flags;       // modifier bits; see reflect.Modifier
	// private ResolvedMethodName method;    // cached resolved method information
	Clazz  *jvm.Ref
	Name   *jvm.Ref
	Type   *jvm.Ref
	Flags  int32
	Method *jvm.Ref
}

type MemberNameData struct {
	VMIndex int64
}
