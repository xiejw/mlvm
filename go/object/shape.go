package object

import (
	"bytes"
	"fmt"
)

type DimensionName string

type NamedDimension struct {
	Name DimensionName
	Size uint // Noozero
}

type Shape struct {
	Dimensions []NamedDimension // Cannot have dup names.
	Rank       uint
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
	fmt.Fprintf(&buf, "< ")
	for i, dim := range shape.Dimensions {
		fmt.Fprintf(&buf, "@%v(%v)", dim.Name, dim.Size)
		if i != finalIndex {
			fmt.Fprintf(&buf, ", ")
		}
	}
	fmt.Fprintf(&buf, ">")

	return buf.String()
}
