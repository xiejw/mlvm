package object

import "testing"

func assertRankEq(t *testing.T, shape *Shape, expectedRank uint) {
	if shape.Rank != expectedRank {
		t.Fatalf("Rank mismatch.")
	}
}

func assertDimensionEq(t *testing.T, dim *NamedDimension, expectedName string, expectedSize uint) {
	if expectedName != dim.Name {
		t.Fatalf("Name mismatch.")
	}
	if expectedSize != dim.Size {
		t.Fatalf("Size mismatch.")
	}
}

func TestShape(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}, {"y", 3}})
	assertRankEq(t, shape, 2)
	assertDimensionEq(t, &shape.Dimensions[0], "x", 2)
	assertDimensionEq(t, &shape.Dimensions[1], "y", 3)
}
