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

// Prints formatted tensor as `<@x(2), @y(3)> [  1.000,  2.000]`
func (t *Tensor) String() string {
	return fmt.Sprintf("%v %v", t.Shape, t.Value)
}

// Shortcut for t.Value.Value.
func (t *Tensor) ArrayValue() []float32 {
	return t.Value.Value
}

func (t *Tensor) DebugString(maxElementCountToPrint int) string {
	return fmt.Sprintf("%v %v", t.Shape, t.Value.DebugString(maxElementCountToPrint))
}
