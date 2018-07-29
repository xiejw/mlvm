package layers

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
)

func NewInput(ctx *c.Context, name string, shape t.Shape, dtype t.DType) t.Tensor {
	unique_name := ctx.GetUniqueNameForTensor(name)
	return &inputTensor{
		name:  unique_name,
		shape: shape,
		dtype: dtype,
	}
}

type inputTensor struct {
	name  string
	shape t.Shape
	dtype t.DType
}

func (input *inputTensor) Name() string {
	return input.name
}

func (input *inputTensor) Shape() t.Shape {
	return input.shape
}

func (input *inputTensor) DType() t.DType {
	return input.dtype
}

func (input *inputTensor) String() string {
	return fmt.Sprintf("Input (\"%v\"), shape: %v", input.name, input.shape)
}
