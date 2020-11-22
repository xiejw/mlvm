package ir

import (
	"bytes"
	"fmt"

	"github.com/xiejw/mlvm/vm/object"
)

type TypeKind int

type Type struct {
	Kind TypeKind
	Dims []int // KShape, KTensor, 1-D for KArray
}

const (
	KInt TypeKind = iota
	KRng
	KShape
	KArray
	KTensor
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
	case KArray:
		return fmt.Sprintf("Array(<%v>)", t.Dims[0])
	case KTensor:
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "Tensor(")
		object.NewShape(t.Dims).ToHumanReadableString(&buf)
		fmt.Fprintf(&buf, ")")
		return buf.String()
	default:
		panic("unknown type string.")
	}
}
