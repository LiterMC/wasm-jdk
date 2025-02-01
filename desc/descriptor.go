package desc

import (
	"errors"
	"strings"
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
	Ref     Type = 'L'
	Array   Type = '['

	RefEnd    byte = ';'
	Method    byte = '('
	MethodEnd byte = ')'
)

var (
	ErrInvalid = errors.New("Invalid descriptor")
	ErrEndTooLate = errors.New("Invalid descriptor: end too late")
	ErrEndTooEarly = errors.New("Invalid descriptor: end too early")
)

type Desc struct {
	ArrDim  int
	Type    Type
	RefName string
}

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
		return s[1:], &Desc{Type: t}
	case Ref:
		i := strings.IndexByte(s, RefEnd)
		if i == -1 {
			return "", nil
		}
		return s[i+1:], &Desc{
			Type:    Ref,
			RefName: s[1:i],
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

