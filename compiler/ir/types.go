package ir

import (
	"github.com/xiejw/mlvm/vm/object"
)

type TypeKind int

type Type struct {
	Kind TypeKind
	Dims []int // KShape
}

const (
	KInt TypeKind = iota
	KRng
	KShape
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
	case KShape:
		return object.NewShape(t.Dims).String()
	default:
		panic("unknown type string.")
	}
}
