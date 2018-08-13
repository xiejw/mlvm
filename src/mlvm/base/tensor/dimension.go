package tensor

import (
	"fmt"
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
