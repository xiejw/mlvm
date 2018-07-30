package layers

import (
	t "mlvm/base/tensor"
)

type Output interface {
	DType() t.DType
	Shape() t.Shape
}

type outputImpl struct {
	dtype t.DType
	shape t.Shape
}

func (o *outputImpl) DType() t.DType {
	return o.dtype
}

func (o *outputImpl) Shape() t.Shape {
	return o.shape
}
