package object

import (
	"bytes"
	"fmt"
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
	t.Shape.CompactString(&buf)
	fmt.Fprintf(&buf, " %v", t.Value)
	return buf.String()
}

// Shortcut for t.Value.Value.
func (t *Tensor) ArrayValue() []float32 {
	return t.Value.Value
}

func (t *Tensor) DebugString(maxElementCountToPrint int) string {
	return fmt.Sprintf("%v %v", t.Shape, t.Value.DebugString(maxElementCountToPrint))
}
