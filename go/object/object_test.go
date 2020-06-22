package object

import "testing"

func TestInteger(t *testing.T) {
	var o Object
	o = &Integer{Value: 123}
	if o.String() != "Int(123)" {
		t.Fatalf("String() method failed.")
	}
	if o.Type() != IntegerType {
		t.Fatalf("Type() method failed.")
	}
}

func TestString(t *testing.T) {
	var o Object
	o = &String{Value: "abc"}
	if o.String() != `String("abc")` {
		t.Fatalf("String() method failed.")
	}
	if o.Type() != StringType {
		t.Fatalf("Type() method failed.")
	}
}

func TestPrng(t *testing.T) {
	var o Object
	o = &Prng{}
	if o.String() != "Prng()" {
		t.Fatalf("String() method failed.")
	}
	if o.Type() != PrngType {
		t.Fatalf("Type() method failed.")
	}
}
