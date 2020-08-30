package object

import "testing"

func assertRankEq(t *testing.T, shape *Shape, expectedRank uint) {
	if shape.Rank != expectedRank {
		t.Fatalf("Rank mismatch.")
	}
}

func assertDimensionEq(t *testing.T, dim uint, expectedSize uint) {
	if expectedSize != dim {
		t.Fatalf("Size mismatch.")
	}
}

func TestShapeDimensions(t *testing.T) {
	shape := NewShape([]uint{2, 3})
	assertRankEq(t, shape, 2)
	assertDimensionEq(t, shape.Dims[0], 2)
	assertDimensionEq(t, shape.Dims[1], 3)
}
