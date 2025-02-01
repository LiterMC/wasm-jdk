package jcls

const ClassMagic uint32 = 0xCAFEBABE

type Class struct {
	Magic       uint32
	Minor       uint16
	Major       uint16
	ConstPool   []ConstantInfo
	AccessFlags AccessFlag
	This        *Class
	Super       *Class
	Interfaces  []*Class
	Fields      []*Field
	Attrs       []Attribute
}
