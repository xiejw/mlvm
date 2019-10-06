package array

import (
	_ "fmt"
)

// Returns an Array for the data with given shape.
//
// All data passed in should not be not mutated in future.
func NewArray(name string, dims []Dimension, value []Float) (*Array, error) {
	// // Check shape and dimensions.
	// eleCount := Dimension(1)
	// for _, dim := range shape {
	// 	if dim < 1 {
	// 		return nil, fmt.Errorf("Tensor (%v) should have all positive dimension, but got %v", name, dim)
	// 	}
	// 	eleCount *= dim
	// }
	// if int(eleCount) != len(value) {
	// 	return nil, fmt.Errorf("Tensor (%v) should have %v elements, but got %v", name, eleCount, len(value))
	// }

	// Now construct.
	return &Array{
		name:  name,
		shape: &Shape{Dims: dims},
		data:  &Data{Value: value},
	}, nil
}
