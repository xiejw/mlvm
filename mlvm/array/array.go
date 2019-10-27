package array

import (
	"fmt"
)

var (
	// Data related errors.
	ErrEmptyData = fmt.Errorf("Data must be non-empty.")
)

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

func (data *Data) Validate() (int, error) {
	if data.value == nil {
		return 0, ErrEmptyData
	}
	return len(data.value), nil
}
