package object

import "testing"

func asserArrayFmtEq(t *testing.T, array *Array, expected string) {
	t.Helper()
	if expected != array.String() {
		t.Fatalf("String Format mismatch. expected: `%v`, got: `%v`.", expected, array.String())
	}
}

func TestArray(t *testing.T) {
	array := Array{[]float32{1.0, 2.0}}
	if array.Type() != ArrayType {
		t.Fatalf("type mismatch.")
	}
}

func TestArrayStringFormatForShort(t *testing.T) {
	array := Array{[]float32{1.0, 2.0}}
	asserArrayFmtEq(t, &array, "[  1.000,  2.000]")
}

func TestArrayStringFormatForMedium(t *testing.T) {
	// Must be 10 elements.
	array := Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}}
	asserArrayFmtEq(t, &array, "[  1.000,  2.000,  3.000,  4.000,  5.000,  6.000,  7.000,  8.000,  9.000, 10.000]")
}

func TestArrayStringFormatForLong(t *testing.T) {
	// Must be larger than 10 elements.
	array := Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0}}
	asserArrayFmtEq(t, &array, "[  1.000,  2.000,  3.000,  4.000,  5.000,  6.000,  7.000,  8.000,  9.000, 10.000, ... ]")
}

func TestArrayComformObjectInterface(t *testing.T) {
	array := Array{[]float32{1.0, 2.0}}
	var object Object
	object = &array
	_, ok := object.(*Array)
	if !ok {
		t.Errorf("cast should work.")
	}
}
