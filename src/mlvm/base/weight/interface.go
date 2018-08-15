package weight

import (
	t "mlvm/base/tensor"
)

type Weight interface {
	Name() string
	Shape() t.Shape
}

type weightImpl struct {
	name  string
	shape t.Shape
	dtype t.DType
}

func (w *weightImpl) Name() string {
	return w.name
}

func (w *weightImpl) Shape() t.Shape {
	return w.shape
}

func NewWeight(name string, shape t.Shape, dtype t.DType) Weight {
	return &weightImpl{
		name:  name,
		shape: shape,
		dtype: dtype,
	}
}
