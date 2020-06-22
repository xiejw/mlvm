package object

import (
	"bytes"
	"fmt"
)

// For a `@batch` dimension name, the DimName is batch. It is global in the script.
type DimName string

// Prints the canonical name with leading `@`.
func (name DimName) String() string {
	return "@" + string(name)
}

type NamedDim struct {
	Name DimName // Dimension string name (without `@`)
	Size uint    // Static dimension size. Cannot be zero.
}

type Shape struct {
	Dimensions []NamedDim // Cannot have dup names.
	Rank       uint       // Length of `Dimensions`
}

func NewShape(dims []NamedDim) *Shape {
	return &Shape{
		Dimensions: dims,
		Rank:       uint(len(dims)),
	}
}

func (shape *Shape) Type() ObjectType {
	return ShapeType
}

// Prints formatted shape as `<@x(2), @y(3)>`.
func (shape *Shape) String() string {
	var buf bytes.Buffer

	rank := shape.Rank
	finalIndex := int(rank - 1)
	fmt.Fprintf(&buf, "<")
	for i, dim := range shape.Dimensions {
		fmt.Fprintf(&buf, "%v(%v)", dim.Name, dim.Size)
		if i != finalIndex {
			fmt.Fprintf(&buf, ", ")
		}
	}
	fmt.Fprintf(&buf, ">")

	return buf.String()
}

func (shape *Shape) Size() uint64 {
	var size uint64 = 1
	for _, dim := range shape.Dimensions {
		size *= uint64(dim.Size)
	}
	return size
}
