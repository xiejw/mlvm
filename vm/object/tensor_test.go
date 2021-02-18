package object

import (
	"reflect"
	"testing"
)

func TestShapeDimensions(t *testing.T) {
	shape := NewShape([]int{2, 3})
	assertRankEq(t, shape, 2)
	assertDimensionEq(t, shape.Dims[0], 2)
	assertDimensionEq(t, shape.Dims[1], 3)

	if shape.Size != 6 {
		t.Errorf("size mismatch.")
	}
	if shape.String() != "Shape(<2, 3>)" {
		t.Errorf("string mismatch.")
	}
}

func TestTensorFloat32(t *testing.T) {
	tensor := NewTensorFloat32([]int{2}, []float32{1.0, 2.0})

	if !reflect.DeepEqual(NewShape([]int{2}), tensor.Shape) {
		t.Errorf("shape mismatch.")
	}
	if !reflect.DeepEqual([]float32{1.0, 2.0}, tensor.Data) {
		t.Errorf("data mismatch.")
	}
	if Float32 != tensor.DType {
		t.Errorf("dtype mismatch.")
	}
}

func TestTensorInt32(t *testing.T) {
	tensor := NewTensorInt32([]int{2}, []int32{1, 2})

	if !reflect.DeepEqual(NewShape([]int{2}), tensor.Shape) {
		t.Errorf("shape mismatch.")
	}
	if !reflect.DeepEqual([]int32{1, 2}, tensor.Data) {
		t.Errorf("data mismatch.")
	}
	if Int32 != tensor.DType {
		t.Errorf("dtype mismatch.")
	}
}

func TestTensorI32Fmt(t *testing.T) {
	te := NewTensorInt32([]int{2}, []int32{1, 2})
	assertTensorFmtEq(t, te, "Tensor(<2> i32 [1, 2])")
}

func TestTensorF32FmtForMedium(t *testing.T) {
	// Must be 10 elements.
	te := NewTensorFloat32([]int{5, 2}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0})
	assertTensorFmtEq(t, te, "Tensor(<5, 2> f32 [1.000, 2.000, 3.000, 4.000, 5.000, 6.000, 7.000, 8.000, 9.000, 10.000])")
}

func TestTensorF32FmtForLong(t *testing.T) {
	// Must be larger than 10 elements.
	te := NewTensorFloat32([]int{6, 2}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11, 12})
	assertTensorFmtEq(t, te, "Tensor(<6, 2> f32 [1.000, 2.000, 3.000, 4.000, 5.000, 6.000, 7.000, 8.000, 9.000, 10.000, ... ])")
}

// -----------------------------------------------------------------------------
// helper methods.
// -----------------------------------------------------------------------------
func assertRankEq(t *testing.T, shape *Shape, expectedRank int) {
	if shape.Rank != expectedRank {
		t.Fatalf("Rank mismatch.")
	}
}

func assertDimensionEq(t *testing.T, dim int, expectedSize int) {
	if expectedSize != dim {
		t.Fatalf("Size mismatch.")
	}
}

func assertTensorFmtEq(t *testing.T, te *Tensor, expected string) {
	t.Helper()
	if expected != te.String() {
		t.Fatalf("String Format mismatch. expected:\n`%v`\ngot:\n`%v`\n", expected, te.String())
	}
}
