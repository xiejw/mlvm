package object

import (
	"fmt"
)

type Tensor struct {
	Shape *Shape
	Value *Array
}

func NewTensor(dims []NamedDimension, value []float32) *Tensor {
	return &Tensor{
		NewShape(dims),
		&Array{value},
	}
}

func (t *Tensor) Type() ObjectType {
	return TensorType
}

func (t *Tensor) String() string {
	return fmt.Sprintf("%v %v", t.Shape, t.Value)
}
