package ast

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
)

type TensorKind int

const (
	KConstant TensorKind = iota + 1 // 0 is invalid
	KResult
)

// Operands/Results of Instructions in Module.
type Tensor struct {
	kind TensorKind

	arr *array.Array
}

func (t *Tensor) Kind() TensorKind {
	return t.kind
}

func (t *Tensor) Name() string {
	switch t.kind {
	case KConstant:
		return t.arr.Name()
	default:
		panic(fmt.Sprintf("Tensor kind %v is not expected.", t.kind))
	}
}

func (t *Tensor) Shape() *array.Shape {
	return t.arr.Shape()
}

func (t *Tensor) Array() *array.Array {
	if t.kind != KConstant {
		panic("Array is allowed only for KConstant.")
	}
	return t.arr
}

func newConstantTensor(arr *array.Array) *Tensor {
	return &Tensor{
		kind: KConstant,
		arr:  arr,
	}
}
