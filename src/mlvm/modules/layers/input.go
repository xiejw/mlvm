package layers

import (
	t "mlvm/base/tensor"
)

func NewBatchedInput(name string, shape t.Shape, dtype t.DType) t.Tensor {
	dims := make([]t.Dimension, 0)
}

type inputTensor struct {
	shape t.Shape
	dtype t.DType
}

func (input *inputTensor) Shape() Shape {
	return input.shape
}

func (input *inputTensor) DType() DType {
	return input.dtype
}
