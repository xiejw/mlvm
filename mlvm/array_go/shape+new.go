package array

import (
	"fmt"
)

func NewShapeOrDie(dims []Dimension) *Shape {
	shape, err := NewShape(dims)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	return shape
}

func NewShape(dims []Dimension) (*Shape, error) {
	shape := &Shape{dims: dims}
	_, err := shape.Validate()
	return shape, err
}
