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
