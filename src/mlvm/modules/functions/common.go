// Layer-like function without states.
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

func (r *reluLayer) Name() string {
	return r.name
}

func (r *reluLayer) Inputs() layers.Inputs {
	return r.inputs
}

func (r *reluLayer) Output() layers.Output {
	return r.output
}

func (r *reluLayer) String() string {
	return layers.FormatPrintString("Relu", r)
}
