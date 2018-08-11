package functions

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
	"mlvm/modules/layers"
)

// Concats the last axis.
func Concat(ctx *c.Context, name string, inputLayers []layers.Layer) layers.Layer {
	// Prepare the output shape.
	outputDims := inputLayers[0].Output().Shape().AsList()
	finalAxis := len(outputDims) - 1
	finalDim := outputDims[finalAxis]

	// Records the input layers and updaet the final dimension of output.
	inputs := &layers.InputsBuilder{}
	for i, input := range inputLayers {
		inputs.Append(input)
		if i != 0 {
			finalDim.Value += input.Output().Shape().AsList()[finalAxis].Value
		}
	}
	inputs.Build()
	outputDims[finalAxis] = finalDim

	output := layers.NewOutput(
		t.NewShapeFromDims(outputDims), inputLayers[0].Output().DType())

	return &concatLayer{
		name:   ctx.AssignUniqueName(name),
		inputs: inputs,
		output: output,
	}
}

type concatLayer struct {
	name   string
	inputs layers.Inputs
	output layers.Output
}

func (layer *concatLayer) Name() string {
	return layer.name
}

func (layer *concatLayer) Inputs() layers.Inputs {
	return layer.inputs
}

func (layer *concatLayer) Output() layers.Output {
	return layer.output
}

func (layer *concatLayer) String() string {
	return layers.FormatPrintString("Concat", layer)
}
