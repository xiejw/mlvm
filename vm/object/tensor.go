package object

import (
	"bytes"
	"fmt"
	"io"
)

// ----------------------------------------------------------------------------
// shape.
// ----------------------------------------------------------------------------

type Shape struct {
	Dims []int // Cannot have 0.
	Rank int   // Length of `Dims`
	Size int
}

func NewShape(dims []int) *Shape {
	var size = 1
	for _, dim := range dims {
		size *= dim
	}

	return &Shape{
		Dims: dims,
		Rank: int(len(dims)),
		Size: size,
	}
}

// ----------------------------------------------------------------------------
// dtype.
// ----------------------------------------------------------------------------

type DType int

const (
	Float32 DType = iota
	Int32
)

// ----------------------------------------------------------------------------
// tensor.
// ----------------------------------------------------------------------------

type Tensor struct {
	Shape *Shape
	DType DType
	Data  interface{}
}

func NewTensorFloat32(dims []int, value []float32) *Tensor {
	shape := NewShape(dims)
	if shape.Size != len(value) {
		panic(fmt.Sprintf("dims have size %v but value has size %v", shape.Size, len(value)))
	}
	return &Tensor{shape, Float32, value}
}

func NewTensorInt32(dims []int, value []int32) *Tensor {
	shape := NewShape(dims)
	if shape.Size != len(value) {
		panic(fmt.Sprintf("dims have size %v but value has size %v", shape.Size, len(value)))
	}
	return &Tensor{shape, Int32, value}
}

// ----------------------------------------------------------------------------
// string.
// ----------------------------------------------------------------------------

// formats as `Shape(<2, 3>)`.
func (shape *Shape) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Shape(")
	shape.DebugString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

// formats as `<2, 3>`.
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

const defaultMaxNumberToPrintForArray = 9

// formats as `Tensor(<2, 3> f32 [  1.000,  2.000])`
func (t *Tensor) String() string {
	var buf bytes.Buffer
	buf.WriteString("Tensor(")
	t.Shape.DebugString(&buf)
	switch t.DType {
	case Float32:
		buf.WriteString(" f32 ")
		debugStringFloat32(&buf, t.Data.([]float32), defaultMaxNumberToPrintForArray)
	case Int32:
		buf.WriteString(" i32 ")
	default:
		panic(fmt.Sprintf("unsupported dtype: %v", t.DType))
	}
	buf.WriteString(")")
	return buf.String()
}

// formats as `[  1.000,  2.000]`
func debugStringFloat32(w io.Writer, array []float32, maxElementCountToPrint int) {
	size := len(array)

	fmt.Fprintf(w, "[ ")
	for i, v := range array {
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
