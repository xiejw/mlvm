package mat

import (
	"testing"

	"github.com/xiejw/mlvm/vm/vm/tensorarray"
)

func TestReduce(t *testing.T) {
	ta := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})
	if ta.IsCompressed() {
		t.Fatalf("should not be compressed.")
	}

	o, err := Reduce(ta, MergeSum)
	assertNoErr(t, err)
	assertShape(t, []int{1}, o.Dims)
	assertAllClose(t, []float32{3.0}, o.Value, 1e-6)
}

func TestReduceWithCompressedTensor(t *testing.T) {
	ta := tensorarray.FromRaw([]int{3, 1}, []float32{2.0})
	if !ta.IsCompressed() {
		t.Fatalf("should be compressed.")
	}

	o, err := Reduce(ta, MergeSum)
	assertNoErr(t, err)
	assertShape(t, []int{1}, o.Dims)
	assertAllClose(t, []float32{6.0}, o.Value, 1e-6)
}
