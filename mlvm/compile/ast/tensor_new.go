package ast

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

// Creates a new KConstant Tensor from array.
func newConstantTensor(arr *array.Array) *Tensor {
	return &Tensor{
		kind: KConstant,
		arr:  arr,
	}
}

// Creates a new KResult Tensor from `Result`.
func newResultTensor(result *Result) *Tensor {
	return &Tensor{
		kind:   KResult,
		result: result,
	}
}
