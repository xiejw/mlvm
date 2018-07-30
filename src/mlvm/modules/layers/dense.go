package layers

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

func NewDense(
	ctx *c.Context, name string, input Layer, output_unit int) Layer {
	_ = ctx.AssignUniqueName(name)

	// FIXME: check input size.

	inputs := &inputsImpl{}
	inputs.Append(input)

	inputShape := input.Output().Shape().AsList()
	inputShape[1] = t.Dimension{
		Value: output_unit,
	}

	output := &outputImpl{
		shape: t.NewShapeFromDims(inputShape),
		dtype: input.Output().DType(),
	}

	return &denseImpl{
		name:   name,
		inputs: inputs,
		output: output,
	}
}

type denseImpl struct {
	name   string
	inputs Inputs
	output Output
}

func (layer *denseImpl) Name() string {
	return layer.name
}

func (layer *denseImpl) Inputs() Inputs {
	return layer.inputs
}

func (layer *denseImpl) Output() Output {
	return layer.output
}

func (layer *denseImpl) String() string {
	return fmt.Sprintf("Dense Layer\t\tshape: %v\n\t- name: %v", layer.output.Shape(), layer.name)
}
