package ast

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

// Register a constant (array.Array) as Tensor in Module.
func (m *Module) NewConstantOrDie(arr *array.Array) *Tensor {
	t, err := m.NewConstant(arr)
	if err != nil {
		panic(err)
	}
	return t
}

// Register a constant (array.Array) as Tensor in Module.
func (m *Module) NewConstant(arr *array.Array) (*Tensor, error) {
	err := m.mustNotFreezed()
	if err != nil {
		return nil, err
	}

	err = m.registerName(arr.Name(), arr, true /* registerOnce */)
	if err != nil {
		return nil, err
	}
	return newConstantTensor(arr), nil
}
