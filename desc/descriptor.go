package desc

type Type byte

const (
	Void      Type = 'V'
	Boolean   Type = 'Z'
	Byte      Type = 'B'
	Char      Type = 'C'
	Short     Type = 'S'
	Int       Type = 'I'
	Long      Type = 'J'
	Float     Type = 'F'
	Double    Type = 'D'
	Array     Type = '['
	Ref       Type = 'L'
	RefEnd    Type = ';'
	Method    Type = '('
	MethodEnd Type = ')'
)
