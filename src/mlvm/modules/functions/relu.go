package functions

import (
	c "mlvm/base/context"
	"mlvm/modules/layers"
)

func Relu(ctx *c.Context, input layers.Layer) layers.Layer {
	inputs := &layers.InputsBuilder{}
	inputs.Append(input)
	inputs.Build()

	return &reluLayer{
		name:   ctx.AssignUniqueName(input.Name() + "_relu"),
		inputs: inputs,
		output: input.Output(),
	}
}

// Think: Is funciton a layer? or layer-like
type reluLayer struct {
	name   string
	inputs layers.Inputs
	output layers.Output
}

func (layer *reluLayer) Name() string {
	return layer.name
}

func (layer *reluLayer) Inputs() layers.Inputs {
	return layer.inputs
}

func (layer *reluLayer) Output() layers.Output {
	return layer.output
}

func (layer *reluLayer) String() string {
	return layers.FormatPrintString(layers.ReluFuncType, layer)
}
