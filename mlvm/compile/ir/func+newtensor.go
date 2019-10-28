package ir

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

// Register a constant (array.Array) as Tensor in Func.
func (f *Func) NewConstantOrDie(arr *array.Array) *Tensor {
	t, err := f.NewConstant(arr)
	if err != nil {
		panic(err)
	}
	return t
}

// Register a constant (array.Array) as Tensor in Func.
func (f *Func) NewConstant(arr *array.Array) (*Tensor, error) {
	err := f.mustNotFreezed()
	if err != nil {
		return nil, err
	}

	err = f.registerName(arr.Name(), arr, true /* registerOnce */)
	if err != nil {
		return nil, err
	}
	return newConstantTensor(arr), nil
}
