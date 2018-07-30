package layers

import (
	_ "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

type Layer interface {
	// Weights() []w.Weight

	Name() string

	// nil if no inputs
	Inputs() Inputs

	// Output

	// Apply(args ...t.Tensor) t.Tensor

	// Backprop()
}
