package object

import (
	"bytes"
	"fmt"
)

// For a `@batch` dimension name, the DimensionName is batch. It is global in the script.
type DimensionName string

// Prints the canonical name with leading `@`.
func (name DimensionName) String() string {
	return "@" + string(name)
}

type NamedDimension struct {
	Name DimensionName // Dimension string name (without `@`)
	Size uint          // Static dimension size. Cannot be zero.
}

type Shape struct {
	Dimensions []NamedDimension // Cannot have dup names.
	Rank       uint             // Length of `Dimensions`
}

func NewShape(dims []NamedDimension) *Shape {
	return &Shape{
		Dimensions: dims,
		Rank:       uint(len(dims)),
	}
}

func (shape *Shape) Type() ObjectType {
	return ShapeType
}

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
