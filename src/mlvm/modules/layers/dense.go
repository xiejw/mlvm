package layers

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

func NewDense(
	ctx *c.Context, name string, input Layer, output_unit int) Layer {
	unique_name := ctx.AssignUniqueName(name)

	// FIXME: check input size. Must be rank 2.

	inputs := &InputsBuilder{}
	inputs.Append(input)
	inputs.Build()

	outputShape := input.Output().Shape().AsList()
	outputShape[1] = t.Dimension{
		Value: output_unit,
	}

	output := &outputImpl{
		shape: t.NewShapeFromDims(outputShape...),
		dtype: input.Output().DType(),
	}

	weights := make([]w.Weight, 0, 2)

	kernel := ctx.NewWeight(unique_name+"_weights",
		t.NewShape(input.Output().Shape().Dim(1).Value, output_unit), t.Float32)
	weights = append(weights, kernel)

	bias := ctx.NewWeight(unique_name+"_bias",
		t.NewShape(output_unit), t.Float32)

	weights = append(weights, bias)

	return &denseImpl{
		name:    unique_name,
		inputs:  inputs,
		output:  output,
		weights: weights,
	}
}

type denseImpl struct {
	name    string
	inputs  Inputs
	output  Output
	weights []w.Weight
}

func (layer *denseImpl) Weights() []w.Weight {
	return layer.weights
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
	return FormatPrintString(DenseLayerType, layer)
}
