package layers

import (
	t "mlvm/base/tensor"
)

// Output of the layer, which only records the shape and dtype.
type Output interface {
	Shape() t.Shape
	DType() t.DType
}

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
