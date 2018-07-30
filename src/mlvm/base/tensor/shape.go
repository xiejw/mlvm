package tensor

import (
	"bytes"
	"fmt"
	"strings"
)

type DimensionType int

const (
	IntType       DimensionType = 0
	BatchSizeType DimensionType = 1
)

type Dimension struct {
	Type  DimensionType
	Value int
}

func (d *Dimension) String() string {
	switch d.Type {
	case IntType:
		return fmt.Sprintf("%v", d.Value)
	case BatchSizeType:
		return "<BS>"
	default:
		panic(fmt.Sprintf("unknown dimension type: %v", d.Type))
	}
}

// Immutable
type Shape interface {
	Dim(index int) Dimension
	Rank() int
	AsList() []Dimension
}

func NewShape(dimArgs ...int) Shape {
	dims := make([]*Dimension, 0, len(dimArgs))
	dims = fillDims(dims, dimArgs...)
	return &shapeImpl{dims: dims}
}

func NewShapeFromDims(dimArgs []Dimension) Shape {
	dims := make([]*Dimension, 0, len(dimArgs))
	for _, dim := range dimArgs {
		var dimCopy Dimension
		dimCopy = dim
		dims = append(dims, &dimCopy)
	}
	return &shapeImpl{dims: dims}
}

func NewShapeWithBatchSize(dimArgs ...int) Shape {
	dims := make([]*Dimension, 0, len(dimArgs)+1)
	dims = append(dims, &Dimension{
		Type: BatchSizeType,
	})
	dims = fillDims(dims, dimArgs...)
	return &shapeImpl{dims: dims}
}

// Fills `dims` with int dimension specified by `dimArgs`.
func fillDims(dims []*Dimension, dimArgs ...int) []*Dimension {
	for _, dim := range dimArgs {
		dims = append(dims, &Dimension{
			Type:  IntType,
			Value: dim,
		})
	}
	return dims
}

type shapeImpl struct {
	dims []*Dimension
}

func (s *shapeImpl) Dim(index int) Dimension {
	return *s.dims[index]
}

func (s *shapeImpl) Rank() int {
	return len(s.dims)
}

func (s *shapeImpl) AsList() []Dimension {
	dimsCopy := make([]Dimension, 0, len(s.dims))
	for _, dim := range s.dims {
		dimsCopy = append(dimsCopy, *dim)
	}
	return dimsCopy
}

func (s *shapeImpl) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")

	dimStrs := make([]string, 0, s.Rank())
	for _, dim := range s.dims {
		dimStrs = append(dimStrs, dim.String())
	}
	buf.WriteString(strings.Join(dimStrs, ", "))

	buf.WriteString("]")
	return buf.String()
}
