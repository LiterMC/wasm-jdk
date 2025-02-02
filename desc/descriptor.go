package desc

import (
	"errors"
	"reflect"
	"strings"
	"unsafe"
)

type Type byte

const (
	Void    Type = 'V'
	Boolean Type = 'Z'
	Byte    Type = 'B'
	Char    Type = 'C'
	Short   Type = 'S'
	Int     Type = 'I'
	Long    Type = 'J'
	Float   Type = 'F'
	Double  Type = 'D'
	Class   Type = 'L'
	Array   Type = '['

	ClassEnd  byte = ';'
	Method    byte = '('
	MethodEnd byte = ')'
)

var (
	typePointer = reflect.TypeOf((unsafe.Pointer)(nil))
	typeInt8    = reflect.TypeOf((int8)(0))
	typeInt16   = reflect.TypeOf((int16)(0))
	typeInt32   = reflect.TypeOf((int32)(0))
	typeInt64   = reflect.TypeOf((int64)(0))
)

func (t Type) IsRef() bool {
	return t == Class || t == Array
}

func (t Type) Size() uintptr {
	switch t {
	case Void:
		return 0
	case Class, Array:
		return unsafe.Sizeof((unsafe.Pointer)(nil))
	case Boolean, Byte:
		return 1
	case Char, Short:
		return 2
	case Int, Float:
		return 4
	case Long, Double:
		return 8
	default:
		panic("unknown desc.Type")
	}
}

func (t Type) AsReflect() reflect.Type {
	switch t {
	case Class, Array:
		return typePointer
	case Boolean, Byte:
		return typeInt8
	case Char, Short:
		return typeInt16
	case Int, Float:
		return typeInt32
	case Long, Double:
		return typeInt64
	default:
		panic("unexpected desc.Type")
	}
}

var (
	ErrInvalid     = errors.New("Invalid descriptor")
	ErrEndTooLate  = errors.New("Invalid descriptor: end too late")
	ErrEndTooEarly = errors.New("Invalid descriptor: end too early")
)

type Desc struct {
	ArrDim  int
	EndType Type
	Class   string
}

var (
	DescInt8  = &Desc{EndType: Byte}
	DescInt16 = &Desc{EndType: Short}
	DescInt32 = &Desc{EndType: Int}
	DescInt64 = &Desc{EndType: Long}
)

func ParseDesc(s string) (*Desc, error) {
	s, d := parseDesc(s)
	if d == nil {
		return nil, ErrInvalid
	}
	if s != "" {
		return nil, ErrEndTooLate
	}
	return d, nil
}

func parseDesc(s string) (string, *Desc) {
	if len(s) == 0 {
		return "", nil
	}
	t := (Type)(s[0])
	switch t {
	case Void, Boolean, Byte, Char, Short, Int, Long, Float, Double:
		return s[1:], &Desc{EndType: t}
	case Class:
		i := strings.IndexByte(s, ClassEnd)
		if i == -1 {
			return "", nil
		}
		return s[i+1:], &Desc{
			EndType: Class,
			Class:   s[1:i],
		}
	case Array:
		var desc *Desc
		s, desc = parseDesc(s[1:])
		if desc != nil {
			desc.ArrDim++
		}
		return s, desc
	default:
		return "", nil
	}
}

func (d *Desc) Clone() *Desc {
	o := new(Desc)
	*o = *d
	return o
}

func (d *Desc) IsArray() bool {
	return d.ArrDim > 0
}

func (d *Desc) Type() Type {
	if d.ArrDim > 0 {
		return Array
	}
	return d.EndType
}

func (d *Desc) ElemType() Type {
	if d.ArrDim == 0 {
		panic("the descriptor is not an array")
	}
	if d.ArrDim > 1 {
		return Array
	}
	return d.EndType
}

func (d *Desc) Elem() *Desc {
	o := d.Clone()
	o.ArrDim--
	if o.ArrDim < 0 {
		panic("the descriptor is not an array")
	}
	return o
}

func (d *Desc) AsReflect() reflect.Type {
	return d.Type().AsReflect()
}

func (d *Desc) Eq(o *Desc) bool {
	return d.ArrDim == o.ArrDim && d.EndType == o.EndType && d.Class == o.Class
}

func (d *Desc) String() string {
	var s strings.Builder
	for range d.ArrDim {
		s.WriteByte('[')
	}
	s.WriteByte((byte)(d.EndType))
	if d.EndType == Class {
		s.WriteString(d.Class)
		s.WriteByte(ClassEnd)
	}
	return s.String()
}

type MethodDesc struct {
	Inputs []*Desc
	Output *Desc
}

func ParseMethodDesc(s string) (*MethodDesc, error) {
	if s[0] != Method {
		return nil, ErrInvalid
	}
	s = s[1:]
	md := new(MethodDesc)
	for {
		if len(s) <= 1 {
			return nil, ErrEndTooEarly
		}
		if s[0] == MethodEnd {
			s = s[1:]
			break
		}
		var d *Desc
		if s, d = parseDesc(s); d == nil {
			return nil, ErrInvalid
		}
		md.Inputs = append(md.Inputs, d)
	}
	if s, md.Output = parseDesc(s); md.Output == nil {
		return nil, ErrInvalid
	}
	if len(s) != 0 {
		return nil, ErrEndTooLate
	}
	return md, nil
}

func (d *MethodDesc) Clone() *MethodDesc {
	o := new(MethodDesc)
	o.Inputs = make([]*Desc, len(d.Inputs))
	for i, dc := range d.Inputs {
		o.Inputs[i] = dc.Clone()
	}
	o.Output = d.Output.Clone()
	return o
}

func (d *MethodDesc) EqInputs(o *MethodDesc) bool {
	if len(d.Inputs) != len(o.Inputs) {
		return false
	}
	for i, in := range d.Inputs {
		if !in.Eq(o.Inputs[i]) {
			return false
		}
	}
	return true
}

func (d *MethodDesc) String() string {
	var s strings.Builder
	s.WriteByte(Method)
	for _, in := range d.Inputs {
		s.WriteString(in.String())
	}
	s.WriteByte(MethodEnd)
	s.WriteString(d.Output.String())
	return s.String()
}

func (d *MethodDesc) AsReflect() reflect.Type {
	inputs := make([]reflect.Type, len(d.Inputs))
	output := d.Output.AsReflect()
	for i, in := range d.Inputs {
		inputs[i] = in.AsReflect()
	}
	return reflect.FuncOf(inputs, []reflect.Type{output}, false)
}
