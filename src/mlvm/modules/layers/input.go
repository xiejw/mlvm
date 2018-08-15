package layers

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

// A input layer representation.
func NewInput(ctx *c.Context, name string, shape t.Shape, dtype t.DType) Layer {
	unique_name := ctx.AssignUniqueName(name)
	return &inputLayer{
		name:   unique_name,
		inputs: nil,
		output: &outputImpl{
			dtype: dtype,
			shape: shape,
		},
	}
}

type inputLayer struct {
	name   string
	inputs Inputs
	output Output
}

func (layer *inputLayer) Weights() []w.Weight {
	return nil
}

func (layer *inputLayer) Name() string {
	return layer.name
}

func (layer *inputLayer) Inputs() Inputs {
	return layer.inputs
}

func (layer *inputLayer) Output() Output {
	return layer.output
}

func (layer *inputLayer) String() string {
	return FormatPrintString(InputLayerType, layer)
}
