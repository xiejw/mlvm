package object

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

const (
	sizeInt     int = int(unsafe.Sizeof(int(1)))
	sizeFloat32 int = int(unsafe.Sizeof(float32(1.0)))
)

type Shape struct {
	Dims []int // Cannot have 0.
	Rank int   // Length of `Dims`
}

func NewShape(dims []int) *Shape {
	return &Shape{
		Dims: dims,
		Rank: int(len(dims)),
	}
}

type Array struct {
	Value []float32 // Cannot be empty.
}

func (shape *Shape) Size() int {
	var size = 1
	for _, dim := range shape.Dims {
		size *= dim
	}
	return size
}

type Tensor struct {
	Shape *Shape
	Array *Array
}

func NewTensor(dims []int, value []float32) *Tensor {
	return &Tensor{NewShape(dims), &Array{value}}
}

// Shortcut for t.Array.Value.
func (t *Tensor) ArrayValue() []float32 { return t.Array.Value }

// ----------------------------------------------------------------------------
// Mem Size Related.
// ----------------------------------------------------------------------------

func (a *Array) MemSize() int  { return len(a.Value) * sizeFloat32 }
func (s *Shape) MemSize() int  { return s.Rank * sizeInt }
func (t *Tensor) MemSize() int { return t.Shape.MemSize() + t.Array.MemSize() }

// ----------------------------------------------------------------------------
// Type Related.
// ----------------------------------------------------------------------------

func (a *Array) Type() ObjectType  { return ArrayType }
func (s *Shape) Type() ObjectType  { return ShapeType }
func (t *Tensor) Type() ObjectType { return TensorType }

// ----------------------------------------------------------------------------
// String Related.
// ----------------------------------------------------------------------------

const defaultMaxNumberToPrintForArray = 9

// Formats as `Array([  1.000,  2.000])`
func (array *Array) String() string {
	var buf bytes.Buffer
	buf.WriteString("Array(")
	array.DebugString(&buf, defaultMaxNumberToPrintForArray)
	buf.WriteString(")")
	return buf.String()
}

// Formats as `[  1.000,  2.000]`
func (array *Array) DebugString(w io.Writer, maxElementCountToPrint int) {
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
	t.Shape.DebugString(&buf)
	buf.WriteString(" ")
	t.Array.DebugString(&buf, defaultMaxNumberToPrintForArray)
	buf.WriteString(")")
	return buf.String()
}

// Formats as `Shape(<2, 3>)`.
func (shape *Shape) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Shape(")
	shape.DebugString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

// Formats as `<2, 3>`.
func (shape *Shape) DebugString(w io.Writer) {
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
