package array

import (
	"fmt"
)

var (
	// Shape related errors.
	ErrEmptyShape     = fmt.Errorf("Shape must be non-empty (rank >= 1).")
	ErrNonPositiveDim = fmt.Errorf("Dimension must be positive.")
)

type Dimension int

type Shape struct {
	dims []Dimension
}

func (shape *Shape) Validate() (int, error) {
	if len(shape.dims) == 0 {
		return 0, ErrEmptyShape
	}
	count := 1
	for _, dim := range shape.dims {
		if int(dim) <= 0 {
			return 0, ErrNonPositiveDim
		}
		count *= int(dim)
	}
	return count, nil
}

func (shape *Shape) IsScalar() bool {
	return len(shape.dims) == 1 && shape.dims[0] == 1
}

func (shape *Shape) Rank() int {
	return len(shape.dims)
}

func (shape *Shape) Dims() []Dimension {
	return shape.dims
}

func (shape *Shape) ElementCount() int {
	count := 1
	for _, dim := range shape.dims {
		count *= int(dim)
	}
	return count
}

// Checks whether two shapes are equal.
func ShapesEqual(a, b *Shape) bool {
	if len(a.dims) != len(b.dims) {
		return false
	}

	for i, d := range a.dims {
		if b.dims[i] != d {
			return false
		}
	}

	return true
}
