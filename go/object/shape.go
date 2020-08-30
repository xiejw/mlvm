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

func (shape *Shape) Type() ObjectType {
	return ShapeType
}

func (shape *Shape) Size() uint64 {
	var size uint64 = 1
	for _, dim := range shape.Dims {
		size *= uint64(dim)
	}
	return size
}

// Formats as `Shape(<2, 3>)`.
func (shape *Shape) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Shape(")
	shape.toHumanReadableString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

// Formats as `<@x(2), @y(3)>`.
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
