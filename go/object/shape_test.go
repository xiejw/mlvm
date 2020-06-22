package object

import "testing"

func assertRankEq(t *testing.T, shape *Shape, expectedRank uint) {
	if shape.Rank != expectedRank {
		t.Fatalf("Rank mismatch.")
	}
}

func assertDimensionEq(t *testing.T, dim *NamedDim, expectedName string, expectedSize uint) {
	if expectedName != string(dim.Name) {
		t.Fatalf("Name mismatch.")
	}
	if expectedSize != dim.Size {
		t.Fatalf("Size mismatch.")
	}
}

func assertShapeFmtEq(t *testing.T, shape *Shape, expected string) {
	if expected != shape.String() {
		t.Fatalf("String Format mismatch. expected: `%v`, got: `%v`.", expected, shape.String())
	}
}

func TestShapeDimensions(t *testing.T) {
	shape := NewShape([]NamedDim{{"x", 2}, {"y", 3}})
	assertRankEq(t, shape, 2)
	assertDimensionEq(t, &shape.Dimensions[0], "x", 2)
	assertDimensionEq(t, &shape.Dimensions[1], "y", 3)
	assertShapeFmtEq(t, shape, "<@x(2), @y(3)>")
}

func TestDimensionName(t *testing.T) {
	var name DimName = "batch"
	if name.String() != "@batch" {
		t.Errorf("dimension name mismatch.")
	}
}
