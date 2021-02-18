package object

import (
	"reflect"
	"testing"
)

// func asserArrayFmtEq(t *testing.T, array *Array, expected string) {
// 	t.Helper()
// 	if expected != array.String() {
// 		t.Fatalf("String Format mismatch. expected: `%v`, got: `%v`.", expected, array.String())
// 	}
// }

func TestShapeDimensions(t *testing.T) {
	shape := NewShape([]int{2, 3})
	assertRankEq(t, shape, 2)
	assertDimensionEq(t, shape.Dims[0], 2)
	assertDimensionEq(t, shape.Dims[1], 3)
	if shape.Size != 6 {
		t.Errorf("size mismatch.")
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

// func TestArrayStringFormatForMedium(t *testing.T) {
// 	// Must be 10 elements.
// 	array := Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}}
// 	asserArrayFmtEq(t, &array, "Array([  1.000,  2.000,  3.000,  4.000,  5.000,  6.000,  7.000,  8.000,  9.000, 10.000])")
// }
//
// func TestArrayStringFormatForLong(t *testing.T) {
// 	// Must be larger than 10 elements.
// 	array := Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0}}
// 	asserArrayFmtEq(t, &array, "Array([  1.000,  2.000,  3.000,  4.000,  5.000,  6.000,  7.000,  8.000,  9.000, 10.000, ... ])")
// }
//
// func TestArrayComformObjectInterface(t *testing.T) {
// 	array := Array{[]float32{1.0, 2.0}}
// 	var object Object
// 	object = &array
// 	_, ok := object.(*Array)
// 	if !ok {
// 		t.Errorf("cast should work.")
// 	}
// }

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
