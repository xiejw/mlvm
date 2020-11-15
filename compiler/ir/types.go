package ir

type TypeKind int

type Type struct {
	Kind TypeKind
}

const (
	KInt TypeKind = iota
	KRng
)

var (
	IntType = &Type{Kind: KInt}
	RngType = &Type{Kind: KRng}
)

func (t *Type) IsInt() bool { return t.Kind == KInt }
func (t *Type) String() string {
	switch t.Kind {
	case KInt:
		return "Int"
	case KRng:
		return "Rng"
	default:
		panic("unknown type string.")
	}
}
