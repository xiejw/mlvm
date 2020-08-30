package object

import (
	"bytes"
	"fmt"
	"io"
)

type Shape struct {
	Dims []uint // Cannot have 0.
	Rank uint   // Length of `Dims`
}

func NewShape(dims []uint) *Shape {
	return &Shape{
		Dims: dims,
		Rank: uint(len(dims)),
	}
}

type Array struct {
	Value []float32
}

func (shape *Shape) Size() uint64 {
	var size uint64 = 1
	for _, dim := range shape.Dims {
		size *= uint64(dim)
	}
	return size
}

type Tensor struct {
	Shape *Shape
	Array *Array
}

func NewTensor(dims []uint, value []float32) *Tensor {
	return &Tensor{NewShape(dims), &Array{value}}
}

// Shortcut for t.Array.Value.
func (t *Tensor) ArrayValue() []float32 { return t.Array.Value }

///////////////////////////////////////////////////////////////////////////////
// Type Related.
///////////////////////////////////////////////////////////////////////////////

func (a *Array) Type() ObjectType  { return ArrayType }
func (s *Shape) Type() ObjectType  { return ShapeType }
func (t *Tensor) Type() ObjectType { return TensorType }

///////////////////////////////////////////////////////////////////////////////
// String Related.
///////////////////////////////////////////////////////////////////////////////

const defaultMaxNumberToPrintForArray = 9

// Formats as `Array([  1.000,  2.000])`
func (array *Array) String() string {
	var buf bytes.Buffer
	buf.WriteString("Array(")
	array.toHumanReadableString(&buf, defaultMaxNumberToPrintForArray)
	buf.WriteString(")")
	return buf.String()
}

// Formats as `[  1.000,  2.000]`
func (array *Array) toHumanReadableString(w io.Writer, maxElementCountToPrint int) {
	size := len(array.Value)

	fmt.Fprintf(w, "[ ")
	for i, v := range array.Value {
		fmt.Fprintf(w, "%6.3f", v)

		if i < size-1 {
			fmt.Fprintf(w, ", ")
		}

		if i != size-1 && i >= maxElementCountToPrint {
			fmt.Fprintf(w, "... ")
			break
		}
	}
	fmt.Fprintf(w, "]")
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

// Formats as `Shape(<2, 3>)`.
func (shape *Shape) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Shape(")
	shape.toHumanReadableString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

// Formats as `<2, 3>`.
func (shape *Shape) toHumanReadableString(w io.Writer) {
	rank := shape.Rank
	finalIndex := int(rank - 1)
	fmt.Fprintf(w, "<")
	for i, dim := range shape.Dims {
		fmt.Fprintf(w, "%v", dim)
		if i != finalIndex {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ">")
}
