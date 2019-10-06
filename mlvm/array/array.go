package array

import (
	"fmt"
)

var (
	ErrEmptyShape     = fmt.Errorf("Shape must be non-empty (rank >= 1).")
	ErrNonPositiveDim = fmt.Errorf("Dimension must be positive.")

	ErrEmptyData = fmt.Errorf("Data must be non-empty.")
)

type Dimension int

type Shape struct {
	dims []Dimension
}

type Float float32

type Data struct {
	value []Float
}

type Array struct {
	name  string // Name of the Array
	shape *Shape // The real shape.
	data  *Data  // The backing data.
}

func (arr *Array) Name() string {
	return arr.name
}

func (arr *Array) Shape() *Shape {
	return arr.shape
}

func (arr *Array) Data() *Data {
	return arr.data
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

func (data *Data) Validate() (int, error) {
	if data.value == nil {
		return 0, ErrEmptyData
	}
	return len(data.value), nil
}
