package object

import "testing"

func assertRankEq(t *testing.T, shape *Shape, expectedRank uint) {
	if shape.Rank != expectedRank {
		t.Fatalf("Rank mismatch.")
	}
}

func assertDimensionEq(t *testing.T, dim *NamedDimension, expectedName string, expectedSize uint) {
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

func TestShape(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}, {"y", 3}})
	if shape.Type() != ShapeType {
		t.Fatalf("type mismatch.")
	}

	assertRankEq(t, shape, 2)
	assertDimensionEq(t, &shape.Dimensions[0], "x", 2)
	assertDimensionEq(t, &shape.Dimensions[1], "y", 3)
	assertShapeFmtEq(t, shape, "<@x(2), @y(3)>")
}

func TestShapeComformObjectInterface(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}, {"y", 3}})
	var object Object
	object = shape
	_, ok := object.(*Shape)
	if !ok {
		t.Errorf("cast should work.")
	}
}

func TestDimensionName(t *testing.T) {
	var name DimensionName = "batch"
	if name.String() != "@batch" {
		t.Errorf("dimension name mismatch.")
	}
}
