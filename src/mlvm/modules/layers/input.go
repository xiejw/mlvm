package layers

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
)

func NewInput(ctx *c.Context, name string, shape t.Shape, dtype t.DType) Layer {
	unique_name := ctx.AssignUniqueName(name)
	return &inputLayer{
		name:   unique_name,
		inputs: nil,
	}
}

type inputLayer struct {
	name   string
	inputs Inputs
}

func (layer *inputLayer) Name() string {
	return layer.name
}

func (layer *inputLayer) Inputs() Inputs {
	return layer.inputs
}

func (layer *inputLayer) String() string {
	return fmt.Sprintf("Input (\"%v\")", layer.name)
}
