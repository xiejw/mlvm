package layers

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
)

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
	return FormatPrintString("Input Layer", layer)
}
