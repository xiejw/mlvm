package object

import (
	"reflect"
	"testing"
)

func TestTensorFields(t *testing.T) {
	shape := NewShape([]NamedDim{{"x", 2}})
	array := &Array{[]float32{1.0, 2.0}}
	tensor := Tensor{shape, array}

	if !reflect.DeepEqual(shape, tensor.Shape) {
		t.Errorf("shape mismatch.")
	}
	if !reflect.DeepEqual(array, tensor.Array) {
		t.Errorf("array mismatch.")
	}
	if !reflect.DeepEqual(array.Value, tensor.ArrayValue()) {
		t.Errorf("array value mismatch.")
	}
}
