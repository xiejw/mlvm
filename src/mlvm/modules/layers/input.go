package layers

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
)

func NewInput(ctx *c.Context, name string, shape t.Shape, dtype t.DType) Layer {
	unique_name := ctx.AssignUniqueName(name)
	return &inputLayer{
		name:  unique_name,
		shape: shape,
		dtype: dtype,
	}
}

type inputLayer struct {
	name  string
	shape t.Shape
	dtype t.DType
}

func (input *inputLayer) Name() string {
	return input.name
}

func (input *inputLayer) Inputs() Inputs {
	return nil
}

func (input *inputLayer) Shape() t.Shape {
	return input.shape
}

func (input *inputLayer) DType() t.DType {
	return input.dtype
}

func (input *inputLayer) String() string {
	return fmt.Sprintf("Input (\"%v\"), shape: %v", input.name, input.shape)
}
