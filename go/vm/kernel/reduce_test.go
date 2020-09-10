package kernel

import (
	"testing"

	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

func TestReduce(t *testing.T) {
	ta := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})
	if ta.IsCompressed() {
		t.Fatalf("should not be compressed.")
	}

	o, err := Reduce(ta, MergeSum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := o.ToTensor()
	expected := "Tensor(<1> [  3.000])"

	if result.String() != expected {
		t.Errorf("value mismatch: expected:\n`%v`\ngot:\n`%v`\n", expected, result)
	}
}

func TestReduceWithCompressedTensor(t *testing.T) {
	ta := tensorarray.FromRaw([]int{4}, []float32{2.0})
	if !ta.IsCompressed() {
		t.Fatalf("should be compressed.")
	}

	o, err := Reduce(ta, MergeSum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := o.ToTensor()
	expected := "Tensor(<1> [  8.000])"

	if result.String() != expected {
		t.Errorf("value mismatch: expected:\n`%v`\ngot:\n`%v`\n", expected, result)
	}
}
