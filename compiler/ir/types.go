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

func (t *Type) IsInt() bool   { return t.Kind == KInt }
func (t *Type) IsShape() bool { return t.Kind == KShape }

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
		object.NewShape(t.Dims).DebugString(&buf)
		fmt.Fprintf(&buf, ")")
		return buf.String()
	default:
		panic("unknown type string.")
	}
}

func (t *Type) ValidateShape() error {
	if t.Kind != KShape {
		return fmt.Errorf("Type.Kind is not KShape: %v", t.Kind)
	}
	if len(t.Dims) == 0 {
		return fmt.Errorf("Type.Dims cannot be empty for kShape")
	}
	for _, d := range t.Dims {
		if d <= 0 {
			return fmt.Errorf("All dims of Type.Dims must be positive, but got: %v", t.Dims)
		}
	}
	return nil
}

func (t *Type) ValidateArray() error {
	if t.Kind != KArray {
		return fmt.Errorf("Type.Kind is not KArray: %v", t.Kind)
	}
	if len(t.Dims) != 1 {
		return fmt.Errorf("Type.Dims must be rank 1 for KArray, got: %v", t.Dims)
	}
	if t.Dims[0] <= 0 {
		return fmt.Errorf("Dims[0] of Type.Dims must be positive, but got: %v", t.Dims)
	}
	return nil
}
