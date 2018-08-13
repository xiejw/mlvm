package layers

import (
	t "mlvm/base/tensor"
)

// Creates an immutable `Output` instance.
func NewOutput(shape t.Shape, dtype t.DType) Output {
	return &outputImpl{
		shape: shape,
		dtype: dtype,
	}
}

type outputImpl struct {
	shape t.Shape
	dtype t.DType
}

func (o *outputImpl) DType() t.DType {
	return o.dtype
}

func (o *outputImpl) Shape() t.Shape {
	return o.shape
}
