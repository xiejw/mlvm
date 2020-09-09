package kernel

import (
	"testing"

	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

func TestTensorAdd(t *testing.T) {
	tensor := tensorarray.FromTensor(
		object.NewTensor([]int{2}, []float32{1.0, 2.0}))
	o, err := TensorAdd(tensor, tensor)
	result := o.ToTensor()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.String() != "Tensor(<2> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", result)
	}
}
