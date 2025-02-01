package jcls

import (
	"io"
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
	Parse(io.Reader) error
	Resolve([]ConstantInfo)
}

type ConstantClass struct {
	NameInd uint16
	Name    string
}

func (c *ConstantClass) Tag() ConstTag { return TagClass }
func (c *ConstantClass) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantClass) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd].(*ConstantUtf8).Value
}

type ConstantRef struct {
	ConstTag       ConstTag
	ClassInd       uint16
	NameAndTypeInd uint16
	Class          *ConstantClass
	NameAndType    *ConstantNameAndType
}

func (c *ConstantRef) Tag() ConstTag { return c.ConstTag }
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
	c.Class = infos[c.ClassInd].(*ConstantClass)
	c.NameAndType = infos[c.NameAndTypeInd].(*ConstantNameAndType)
}

type ConstantString struct {
	StringInd uint16
	String    string
}

func (c *ConstantString) Tag() ConstTag { return TagString }
func (c *ConstantString) Parse(r io.Reader) error {
	var err error
	if c.StringInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantString) Resolve(infos []ConstantInfo) {
	c.String = infos[c.StringInd].(*ConstantUtf8).Value
}

type ConstantInteger struct {
	Value uint32
}

func (c *ConstantInteger) Tag() ConstTag { return TagInteger }
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

func (c *ConstantFloat) Tag() ConstTag { return TagFloat }
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

func (c *ConstantLong) Tag() ConstTag { return TagLong }
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

func (c *ConstantDouble) Tag() ConstTag { return TagDouble }
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

func (c *ConstantNameAndType) Tag() ConstTag { return TagNameAndType }
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
	c.Name = infos[c.NameInd].(*ConstantUtf8).Value
	c.Desc = infos[c.DescInd].(*ConstantUtf8).Value
}

type ConstantUtf8 struct {
	Value string
}

func (c *ConstantUtf8) Tag() ConstTag { return TagUtf8 }
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

type ConstantMethodHandle struct {
	RefKind uint8
	RefInd  uint16
	Ref     *ConstantRef
}

func (c *ConstantMethodHandle) Tag() ConstTag { return TagMethodHandle }
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
	c.Ref = infos[c.RefInd].(*ConstantRef)
}

type ConstantMethodType struct {
	DescInd uint16
	Desc    string
}

func (c *ConstantMethodType) Tag() ConstTag { return TagMethodType }
func (c *ConstantMethodType) Parse(r io.Reader) error {
	var err error
	if c.DescInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantMethodType) Resolve(infos []ConstantInfo) {
	c.Desc = infos[c.DescInd].(*ConstantUtf8).Value
}

type ConstantDynamics struct {
	ConstTag        ConstTag
	BootstrapMethod uint16
	NameAndTypeInd  uint16
	NameAndType     *ConstantNameAndType
}

func (c *ConstantDynamics) Tag() ConstTag { return c.ConstTag }
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
	c.NameAndType = infos[c.NameAndTypeInd].(*ConstantNameAndType)
}

type ConstantModule struct {
	NameInd uint16
	Name    string
}

func (c *ConstantModule) Tag() ConstTag { return TagModule }
func (c *ConstantModule) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantModule) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd].(*ConstantUtf8).Value
}

type ConstantPackage struct {
	NameInd uint16
	Name    string
}

func (c *ConstantPackage) Tag() ConstTag { return TagPackage }
func (c *ConstantPackage) Parse(r io.Reader) error {
	var err error
	if c.NameInd, err = readUint16(r); err != nil {
		return err
	}
	return nil
}
func (c *ConstantPackage) Resolve(infos []ConstantInfo) {
	c.Name = infos[c.NameInd].(*ConstantUtf8).Value
}
