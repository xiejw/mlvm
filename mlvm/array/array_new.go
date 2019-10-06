package array

import (
	"fmt"
)

// Returns an Array for the data with given shape.
//
// All data passed in should not be not mutated in future.
func NewArrayOrDie(name string, dims []Dimension, value []Float) *Array {
	arr, err := NewArray(name, dims, value)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	return arr
}

// Returns an Array for the data with given shape.
//
// All data passed in should not be not mutated in future.
func NewArray(name string, dims []Dimension, value []Float) (*Array, error) {

	shape := &Shape{dims: dims}
	eleCount, err := shape.Validate()
	if err != nil {
		return nil, err
	}

	data := &Data{value: value}
	valueCount, err := data.Validate()
	if err != nil {
		return nil, err
	}

	if eleCount != valueCount {
		return nil, fmt.Errorf(
			"Array (%v) should have %v elements, but got %v",
			name, eleCount, valueCount)
	}

	// Now construct.
	return &Array{
		name:  name,
		shape: shape,
		data:  data,
	}, nil
}
