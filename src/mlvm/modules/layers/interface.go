package layers

import (
	_ "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

type Layer interface {
	// Weights() []w.Weight

	// Name of the layer. Unique in graph.
	Name() string

	// Inputs of current layer. `nil` if absent.
	Inputs() Inputs

	// Single output of current layer.
	Output() Output

	// String representation.
	String() string
}
