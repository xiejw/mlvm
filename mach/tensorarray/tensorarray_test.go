package tensorarray

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/vm/object"
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

func TestTAPanic(t *testing.T) {
	defer func() { recover() }()
	_ = FromRaw([]int{2, 3}, []float32{1.0, 2.0, 3.0, 4.0, 5.0})
	t.FailNow()
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

func TestToFullArray(t *testing.T) {
	ta := FromRaw([]int{2, 3}, []float32{1.2, 2.3, 3.4})
	if !ta.IsCompressed() {
		t.Errorf("expected to be compressed.")
	}

	ta = ta.ToFullArray()

	if ta.Size != 6 {
		t.Errorf("size mismatch.")
	}
	if ta.RealSize != 6 {
		t.Errorf("real size mismatch.")
	}
	if ta.Rank != 2 {
		t.Errorf("rank mismatch.")
	}
	if !reflect.DeepEqual([]float32{1.2, 2.3, 3.4, 1.2, 2.3, 3.4}, ta.Value) {
		t.Errorf("value mismatch.")
	}
}

func TestToFullArrayWithSingleV(t *testing.T) {
	ta := FromRaw([]int{2, 3}, []float32{1.2})
	if !ta.IsCompressed() {
		t.Errorf("expected to be compressed.")
	}

	ta = ta.ToFullArray()

	if ta.Size != 6 {
		t.Errorf("size mismatch.")
	}
	if ta.RealSize != 6 {
		t.Errorf("real size mismatch.")
	}
	if ta.Rank != 2 {
		t.Errorf("rank mismatch.")
	}
	if !reflect.DeepEqual([]float32{1.2, 1.2, 1.2, 1.2, 1.2, 1.2}, ta.Value) {
		t.Errorf("value mismatch.")
	}
}

func TestMemSize(t *testing.T) {
	ta := FromRaw([]int{2, 3}, []float32{1.2})
	if ta.MemSize() != (2+3)*sizeInt+sizeFloat32 {
		t.Errorf("mem size mismatch.")
	}
}
