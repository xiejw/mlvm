package kernel

import (
	"testing"

	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

func TestTensorAdd(t *testing.T) {
	tensor := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})

	o, err := BinaryOp(tensor, tensor, BinaryAdd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := o.ToTensor()
	if result.String() != "Tensor(<2> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", result)
	}
}

func TestTensorMinus(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})
	rhs := tensorarray.FromRaw([]int{2}, []float32{2.0, 3.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := o.ToTensor()
	if result.String() != "Tensor(<2> [ -1.000, -1.000])" {
		t.Errorf("value mismatch: got `%v`", result)
	}
}
