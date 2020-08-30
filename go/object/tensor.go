package object

import (
	"bytes"
)

type Tensor struct {
	Shape *Shape
	Array *Array
}

func NewTensor(dims []uint, value []float32) *Tensor {
	return &Tensor{NewShape(dims), &Array{value}}
}

func (t *Tensor) Type() ObjectType {
	return TensorType
}

// Shortcut for t.Array.Value.
func (t *Tensor) ArrayValue() []float32 {
	return t.Array.Value
}

// Formats as `Tensor(<@x(2), @y(3)> [  1.000,  2.000])`
func (t *Tensor) String() string {
	var buf bytes.Buffer
	buf.WriteString("Tensor(")
	t.Shape.toHumanReadableString(&buf)
	buf.WriteString(" ")
	t.Array.toHumanReadableString(&buf, defaultMaxNumberToPrintForArray)
	buf.WriteString(")")
	return buf.String()
}
