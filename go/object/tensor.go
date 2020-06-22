package object

import (
	"bytes"
)

type Tensor struct {
	Shape *Shape
	Value *Array
}

func NewTensor(dims []NamedDim, value []float32) *Tensor {
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
	var buf bytes.Buffer
	t.Shape.toHumanReadableString(&buf)
	buf.WriteString(" ")
	t.Value.toHumanReadableString(&buf, defaultMaxNumberToPrintForArray)
	return buf.String()
}

// Shortcut for t.Value.Value.
func (t *Tensor) ArrayValue() []float32 {
	return t.Value.Value
}
