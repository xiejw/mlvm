package layers

import (
	"fmt"

	c "mlvm/base/context"
	_ "mlvm/base/weight"
)

func NewDense(
	ctx *c.Context, name string, input Layer, output_unit int) Layer {
	_ = ctx.AssignUniqueName(name)
	inputs := &inputsImpl{}
	inputs.Append(input)
	return &denseImpl{
		name:   name,
		inputs: inputs,
	}
}

type denseImpl struct {
	name   string
	inputs Inputs
}

func (layer *denseImpl) Name() string {
	return layer.name
}

func (layer *denseImpl) Inputs() Inputs {
	return layer.inputs
}

func (layer *denseImpl) String() string {
	return fmt.Sprintf("Dense (\"%v\")", layer.name)
}
