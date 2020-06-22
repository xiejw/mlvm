package kernel

import (
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func TestTensorAdd(t *testing.T) {
	tensor := object.NewTensor([]object.NamedDim{{"x", 2}}, []float32{1.0, 2.0})
	result, err := TensorAdd(tensor, tensor)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.String() != "Tensor(<@x(2)> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", result)
	}
}
