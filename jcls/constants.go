package jcls

import (
	"fmt"
	"io"

	"github.com/LiterMC/wasm-jdk/desc"
)

// https://docs.oracle.com/javase/specs/jvms/se21/html/jvms-4.html#jvms-4.4
type ConstTag byte

const (
	TagUtf8               ConstTag = 1
	TagInteger            ConstTag = 3
	TagFloat              ConstTag = 4
	TagLong               ConstTag = 5
	TagDouble             ConstTag = 6
	TagClass              ConstTag = 7
	TagString             ConstTag = 8
	TagFieldref           ConstTag = 9
	TagMethodref          ConstTag = 10
	TagInterfaceMethodref ConstTag = 11
	TagNameAndType        ConstTag = 12
	TagMethodHandle       ConstTag = 15
	TagMethodType         ConstTag = 16
	TagDynamic            ConstTag = 17
	TagInvokeDynamic      ConstTag = 18
	TagModule             ConstTag = 19
	TagPackage            ConstTag = 20
)

type ConstantInfo interface {
	Tag() ConstTag
	// If use two constant pool slots
	IsWide() bool
	Parse(io.Reader) error
	Resolve([]ConstantInfo)
}

func ParseConstant(r io.Reader) (ConstantInfo, error) {
	t, err := readUint8(r)
	if err != nil {
		return nil, err
	}
	var (
		c   ConstantInfo
		tag = (ConstTag)(t)
	)
	switch tag {
	case TagClass:
		c = &ConstantClass{}
	case TagFieldref, TagMethodref, TagInterfaceMethodref:
		c = &ConstantRef{ConstTag: tag}
	case TagString:
		c = &ConstantString{}
	case TagInteger:
		c = &ConstantInteger{}
	case TagFloat:
		c = &ConstantFloat{}
	case TagLong:
		c = &ConstantLong{}
	case TagDouble:
		c = &ConstantDouble{}
	case TagNameAndType:
		c = &ConstantNameAndType{}
	case TagUtf8:
		c = &ConstantUtf8{}
	case TagMethodHandle:
		c = &ConstantMethodHandle{}
	case TagMethodType:
		c = &ConstantMethodType{}
	case TagDynamic, TagInvokeDynamic:
		c = &ConstantDynamics{ConstTag: tag}
	case TagModule:
		c = &ConstantModule{}
	case TagPackage:
		c = &ConstantPackage{}
	default:
		return nil, fmt.Errorf("Unexpected constant tag %d", t)
	}
	if err = c.Parse(r); err != nil {
		return nil, err
	}
	return c, nil
}

type ConstantClass struct {
	NameInd uint16
	Name    string
}

func (*ConstantClass) Tag() ConstTag { return TagClass }
func (*ConstantClass) IsWide() bool  { return false }
func (c *ConstantClass) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantClass) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd-1].(*ConstantUtf8).Value
}
func (c *ConstantClass) String() string {
	return "Class: " + c.Name
}

type ConstantRef struct {
	ConstTag       ConstTag
	ClassInd       uint16
	NameAndTypeInd uint16
	Class          *ConstantClass
	NameAndType    *ConstantNameAndType
}

func (c *ConstantRef) Tag() ConstTag { return c.ConstTag }
func (*ConstantRef) IsWide() bool    { return false }
func (c *ConstantRef) Parse(r io.Reader) error {
	var err error
	if c.ClassInd, err = readUint16(r); err != nil {
		return err
	}
	if c.NameAndTypeInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantRef) Resolve(infos []ConstantInfo) {
	c.Class = infos[c.ClassInd-1].(*ConstantClass)
	c.NameAndType = infos[c.NameAndTypeInd-1].(*ConstantNameAndType)
}

type ConstantString struct {
	Utf8Ind uint16
	Utf8    string
}

func (*ConstantString) Tag() ConstTag { return TagString }
func (*ConstantString) IsWide() bool  { return false }
func (c *ConstantString) Parse(r io.Reader) error {
	var err error
	if c.Utf8Ind, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantString) Resolve(infos []ConstantInfo) {
	c.Utf8 = infos[c.Utf8Ind-1].(*ConstantUtf8).Value
}
func (c *ConstantString) String() string {
	return c.Utf8
}

type ConstantInteger struct {
	Value uint32
}

func (*ConstantInteger) Tag() ConstTag { return TagInteger }
func (*ConstantInteger) IsWide() bool  { return false }
func (c *ConstantInteger) Parse(r io.Reader) error {
	var err error
	if c.Value, err = readUint32(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantInteger) Resolve(infos []ConstantInfo) {}

type ConstantFloat struct {
	Value uint32
}

func (*ConstantFloat) Tag() ConstTag { return TagFloat }
func (*ConstantFloat) IsWide() bool  { return false }
func (c *ConstantFloat) Parse(r io.Reader) error {
	var err error
	if c.Value, err = readUint32(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantFloat) Resolve(infos []ConstantInfo) {}

type ConstantLong struct {
	Value uint64
}

func (*ConstantLong) Tag() ConstTag { return TagLong }
func (*ConstantLong) IsWide() bool  { return true }
func (c *ConstantLong) Parse(r io.Reader) error {
	var err error
	if c.Value, err = readUint64(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantLong) Resolve(infos []ConstantInfo) {}

type ConstantDouble struct {
	Value uint64
}

func (*ConstantDouble) Tag() ConstTag { return TagDouble }
func (*ConstantDouble) IsWide() bool  { return true }
func (c *ConstantDouble) Parse(r io.Reader) error {
	var err error
	if c.Value, err = readUint64(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantDouble) Resolve(infos []ConstantInfo) {}

type ConstantNameAndType struct {
	NameInd uint16
	DescInd uint16
	Name    string
	Desc    string
}

func (*ConstantNameAndType) Tag() ConstTag { return TagNameAndType }
func (*ConstantNameAndType) IsWide() bool  { return false }
func (c *ConstantNameAndType) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	if c.DescInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantNameAndType) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd-1].(*ConstantUtf8).Value
	c.Desc = infos[c.DescInd-1].(*ConstantUtf8).Value
}
func (c *ConstantNameAndType) String() string {
	return c.Name + " " + c.Desc
}

type ConstantUtf8 struct {
	Value string

	desc       *desc.Desc
	methodDesc *desc.MethodDesc
}

func (*ConstantUtf8) Tag() ConstTag { return TagUtf8 }
func (*ConstantUtf8) IsWide() bool  { return false }
func (c *ConstantUtf8) Parse(r io.Reader) error {
	size, err := readUint16(r)
	if err != nil {
		return err
	}
	buf := make([]byte, size)
	if _, err = io.ReadFull(r, buf); err != nil {
		return err
	}
	c.Value = (string)(buf)
	return nil
}
func (c *ConstantUtf8) Resolve(infos []ConstantInfo) {}

func (c *ConstantUtf8) AsDesc() (*desc.Desc, error) {
	if c.desc == nil {
		var err error
		if c.desc, err = desc.ParseDesc(c.Value); err != nil {
			return nil, err
		}
	}
	return c.desc, nil
}

func (c *ConstantUtf8) AsMethodDesc() (*desc.MethodDesc, error) {
	if c.methodDesc == nil {
		var err error
		if c.methodDesc, err = desc.ParseMethodDesc(c.Value); err != nil {
			return nil, err
		}
	}
	return c.methodDesc, nil
}

type ConstantMethodHandle struct {
	RefKind uint8
	RefInd  uint16
	Ref     *ConstantRef
}

func (*ConstantMethodHandle) Tag() ConstTag { return TagMethodHandle }
func (*ConstantMethodHandle) IsWide() bool  { return false }
func (c *ConstantMethodHandle) Parse(r io.Reader) error {
	var err error
	if c.RefKind, err = readUint8(r); err != nil {
		return err
	}
	if c.RefInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantMethodHandle) Resolve(infos []ConstantInfo) {
	c.Ref = infos[c.RefInd-1].(*ConstantRef)
}

type ConstantMethodType struct {
	DescInd uint16
	Desc    string
}

func (*ConstantMethodType) Tag() ConstTag { return TagMethodType }
func (*ConstantMethodType) IsWide() bool  { return false }
func (c *ConstantMethodType) Parse(r io.Reader) error {
	var err error
	if c.DescInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantMethodType) Resolve(infos []ConstantInfo) {
	c.Desc = infos[c.DescInd-1].(*ConstantUtf8).Value
}

type ConstantDynamics struct {
	ConstTag        ConstTag
	BootstrapMethod uint16
	NameAndTypeInd  uint16
	NameAndType     *ConstantNameAndType
}

func (c *ConstantDynamics) Tag() ConstTag { return c.ConstTag }
func (*ConstantDynamics) IsWide() bool    { return false }
func (c *ConstantDynamics) Parse(r io.Reader) error {
	var err error
	if c.BootstrapMethod, err = readUint16(r); err != nil {
		return err
	}
	if c.NameAndTypeInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantDynamics) Resolve(infos []ConstantInfo) {
	c.NameAndType = infos[c.NameAndTypeInd-1].(*ConstantNameAndType)
}

type ConstantModule struct {
	NameInd uint16
	Name    string
}

func (*ConstantModule) Tag() ConstTag { return TagModule }
func (*ConstantModule) IsWide() bool  { return false }
func (c *ConstantModule) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantModule) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd-1].(*ConstantUtf8).Value
}

type ConstantPackage struct {
	NameInd uint16
	Name    string
}

func (*ConstantPackage) Tag() ConstTag { return TagPackage }
func (*ConstantPackage) IsWide() bool  { return false }
func (c *ConstantPackage) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantPackage) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd-1].(*ConstantUtf8).Value
}
