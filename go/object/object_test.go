package object

import (
	"testing"
)

func assertStringAndType(t *testing.T, expectedStr string, expectType ObjectType, got Object) {
	t.Helper()
	if got.String() != expectedStr {
		t.Fatalf("String() method failed. expected: %v, got: %v.", expectedStr, got.String())
	}
	if got.Type() != expectType {
		t.Fatalf("Type() method failed.")
	}
}

func TestInteger(t *testing.T) {
	var o Object
	o = &Integer{Value: 123}

	assertStringAndType(t, "Integer(123)", IntegerType, o)
}

func TestString(t *testing.T) {
	var o Object
	o = &String{Value: "abc"}

	assertStringAndType(t, `String("abc")`, StringType, o)
}

func TestRng(t *testing.T) {
	var o Object
	o = &Rng{123, 456, 789}

	assertStringAndType(t, "Rng(7b, 1c8, 315)", RngType, o)
}

func TestShape(t *testing.T) {
	var o Object
	o = NewShape([]uint{2, 3})
	assertStringAndType(t, "Shape(<2, 3>)", ShapeType, o)
}

func TestArray(t *testing.T) {
	var o Object
	o = &Array{[]float32{1.0, 2.0}}
	assertStringAndType(t, "Array([  1.000,  2.000])", ArrayType, o)
}
func TestTensor(t *testing.T) {
	shape := NewShape([]uint{2})
	array := &Array{[]float32{1.0, 2.0}}

	var o Object
	o = &Tensor{shape, array}
	assertStringAndType(t, "Tensor(<2> [  1.000,  2.000])", TensorType, o)
}
