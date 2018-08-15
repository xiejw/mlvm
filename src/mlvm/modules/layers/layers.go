package layers

import (
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

// Layer is a single-output multi-inputs building block in neural network.
type Layer interface {
	Weights() []w.Weight

	// Name of the layer. Must be unique in graph. Typically generated by
	// `Context.AssignUniqueName'.
	Name() string

	// Inputs of current layer. `nil` if absent.
	Inputs() Inputs

	// Single output of current layer.
	Output() Output

	// String debugging information of the layer. Typically generated by
	// `layers.FormatPrintString` utility funciton.
	String() string
}

// Output of the layer, which only records the shape and dtype.
type Output interface {
	Shape() t.Shape
	DType() t.DType
}

type Inputs interface {
	inputsTemplate // Generated by tools.
}
