package ast

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

// Register a constant (array.Array) as Tensor in Module.
func (m *Module) NewConstant(arr *array.Array) *Tensor {
	m.registerName(arr.Name(), arr, true /* registerOnce */)
	return newConstantTensor(arr)
}


