package ast

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

// Creates a new constant Tensor from array.
func newConstantTensor(arr *array.Array) *Tensor {
	return &Tensor{
		kind: KConstant,
		arr:  arr,
	}
}
