package object

import "testing"

func assertStringAndType(t *testing.T, expectedStr string, expectType ObjectType, got Object) {
	t.Helper()
	if got.String() != expectedStr {
		t.Fatalf("String() method failed.")
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

func TestPrng(t *testing.T) {
	var o Object
	o = &Prng{}

	assertStringAndType(t, "Prng()", PrngType, o)
}

func TestShape(t *testing.T) {
	var o Object
	o = NewShape([]NamedDim{{"x", 2}, {"y", 3}})
	assertStringAndType(t, "Shape(<@x(2), @y(3)>)", ShapeType, o)
}
