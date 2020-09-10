package tensorarray

import (
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func TestObjectInterface(t *testing.T) {
	var o object.Object

	ta := &TensorArray{}
	o = ta

	if o.String() != "TensorArray" {
		t.Fatalf("String() mistmatch.")
	}
}

func TestTA(t *testing.T) {
	ta := FromRaw([]int{2, 3}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	if ta.IsCompressed() {
		t.Errorf("expected to not be compressed.")
	}
	if ta.Size != 6 {
		t.Errorf("size mismatch.")
	}
	if ta.RealSize != 6 {
		t.Errorf("real size mismatch.")
	}
	if ta.Rank != 2 {
		t.Errorf("rank mismatch.")
	}
}

func TestCompressedTA(t *testing.T) {
	ta := FromRaw([]int{2, 3}, []float32{1.2})
	if !ta.IsCompressed() {
		t.Errorf("expected to be compressed.")
	}
	if ta.Size != 6 {
		t.Errorf("size mismatch.")
	}
	if ta.RealSize != 1 {
		t.Errorf("real size mismatch.")
	}
	if ta.Rank != 2 {
		t.Errorf("rank mismatch.")
	}
}
