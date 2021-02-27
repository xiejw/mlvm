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
	F32 DType = iota
	I32
)

func (t DType) AllowGrad() bool {
	return t == F32
}

// ----------------------------------------------------------------------------
// tensor like.
//
// used for the case value is not ready (async schedule).
// ----------------------------------------------------------------------------

type TensorLike interface {
	Shape() *Shape
	DType() DType
}

func NewTensorLike(dtype DType, dims []int) TensorLike {
	shape := NewShape(dims)
	return &tensorShell{
		shape: shape,
		dtype: dtype,
	}
}

type tensorShell struct {
	shape *Shape
	dtype DType
}

func (te *tensorShell) Shape() *Shape {
	return te.shape
}

func (t *tensorShell) DType() DType {
	return t.dtype
}

// ----------------------------------------------------------------------------
// tensor.
// ----------------------------------------------------------------------------

type Tensor struct {
	tensorShell
	data interface{}
}

func NewTensor(dtype DType, dims []int) *Tensor {
	shape := NewShape(dims)
	var te *Tensor

	switch dtype {
	case F32:
		te = &Tensor{
			tensorShell{shape: shape, dtype: dtype},
			make([]float32, shape.Size),
		}
	case I32:
		te = &Tensor{
			tensorShell{shape: shape, dtype: dtype},
			make([]int32, shape.Size),
		}
	default:
		panic(fmt.Sprintf("unsupported dtype: %v", dtype))
	}
	return te
}

func NewTensorF32(dims []int, value []float32) *Tensor {
	shape := NewShape(dims)
	if shape.Size != len(value) {
		panic(fmt.Sprintf("dims have size %v but value has size %v", shape.Size, len(value)))
	}
	return &Tensor{tensorShell{shape, F32}, value}
}

func NewTensorI32(dims []int, value []int32) *Tensor {
	shape := NewShape(dims)
	if shape.Size != len(value) {
		panic(fmt.Sprintf("dims have size %v but value has size %v", shape.Size, len(value)))
	}
	return &Tensor{tensorShell{shape, I32}, value}
}

func (t *Tensor) Data() interface{} {
	return t.data
}

// ----------------------------------------------------------------------------
// conform String
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
	t.Shape().DebugString(&buf)
	switch t.DType() {
	case F32:
		buf.WriteString(" f32 ")
		debugStringF32(&buf, t.Data().([]float32), defaultMaxNumberToPrintForArray)
	case I32:
		buf.WriteString(" i32 ")
		debugStringI32(&buf, t.Data().([]int32), defaultMaxNumberToPrintForArray)
	default:
		panic(fmt.Sprintf("unsupported dtype: %v", t.DType()))
	}
	buf.WriteString(")")
	return buf.String()
}

// formats as `[1.000, 2.000]`
func debugStringF32(w io.Writer, array []float32, maxElementCountToPrint int) {
	size := len(array)

	fmt.Fprintf(w, "[")
	for i, v := range array {
		fmt.Fprintf(w, "%.3f", v)

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

// formats as `[1, 2]`
func debugStringI32(w io.Writer, array []int32, maxElementCountToPrint int) {
	size := len(array)

	fmt.Fprintf(w, "[")
	for i, v := range array {
		fmt.Fprintf(w, "%v", v)

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
