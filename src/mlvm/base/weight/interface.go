package weight

import (
	t "mlvm/base/tensor"
)

type Weight interface {
	Name() string
}

type weightImpl struct {
	name  string
	shape t.Shape
	dtype t.DType
}

func (w *weightImpl) Name() string {
	return w.name
}

func NewWeight(name string, shape t.Shape, dtype t.DType) Weight {
	return &weightImpl{
		name:  name,
		shape: shape,
		dtype: dtype,
	}
}
