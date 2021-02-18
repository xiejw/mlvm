package object

import (
	//	"bytes"
	"fmt"
	//	"io"
)

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

type DType int

const (
	Int32 DType = iota
	Float32
)

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
// String Related.
// ----------------------------------------------------------------------------

// const defaultMaxNumberToPrintForArray = 9
//
// // Formats as `Array([  1.000,  2.000])`
// func (array *Array) String() string {
// 	var buf bytes.Buffer
// 	buf.WriteString("Array(")
// 	array.DebugString(&buf, defaultMaxNumberToPrintForArray)
// 	buf.WriteString(")")
// 	return buf.String()
// }
//
// // Formats as `[  1.000,  2.000]`
// func (array *Array) DebugString(w io.Writer, maxElementCountToPrint int) {
// 	size := len(array.Value)
//
// 	fmt.Fprintf(w, "[ ")
// 	for i, v := range array.Value {
// 		fmt.Fprintf(w, "%6.3f", v)
//
// 		if i < size-1 {
// 			fmt.Fprintf(w, ", ")
// 		}
//
// 		if i != size-1 && i >= maxElementCountToPrint {
// 			fmt.Fprintf(w, "... ")
// 			break
// 		}
// 	}
// 	fmt.Fprintf(w, "]")
// }
//
// // Formats as `Tensor(<@x(2), @y(3)> [  1.000,  2.000])`
// func (t *Tensor) String() string {
// 	var buf bytes.Buffer
// 	buf.WriteString("Tensor(")
// 	t.Shape.DebugString(&buf)
// 	buf.WriteString(" ")
// 	t.Array.DebugString(&buf, defaultMaxNumberToPrintForArray)
// 	buf.WriteString(")")
// 	return buf.String()
// }
//
// // Formats as `Shape(<2, 3>)`.
// func (shape *Shape) String() string {
// 	var buf bytes.Buffer
// 	fmt.Fprintf(&buf, "Shape(")
// 	shape.DebugString(&buf)
// 	fmt.Fprintf(&buf, ")")
// 	return buf.String()
// }
//
// // Formats as `<2, 3>`.
// func (shape *Shape) DebugString(w io.Writer) {
// 	rank := shape.Rank
// 	finalIndex := int(rank - 1)
// 	fmt.Fprintf(w, "<")
// 	for i, dim := range shape.Dims {
// 		fmt.Fprintf(w, "%v", dim)
// 		if i != finalIndex {
// 			fmt.Fprintf(w, ", ")
// 		}
// 	}
// 	fmt.Fprintf(w, ">")
// }
