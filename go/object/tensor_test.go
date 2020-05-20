package object

import "testing"

func assertTensorFmtEq(t *testing.T, tensor *Tensor, expected string) {
	t.Helper()
	if expected != tensor.String() {
		t.Fatalf("String Format mismatch. expected: `%v`, got: `%v`.", expected, tensor.String())
	}
}

func TestTensor(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}})
	array := &Array{[]float32{1.0, 2.0}}
	tensor := Tensor{shape, array}

	if tensor.Type() != TensorType {
		t.Fatalf("type mismatch.")
	}
}

func TestTensorStringFormatForShort(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}})
	array := &Array{[]float32{1.0, 2.0}}
	tensor := Tensor{shape, array}
	assertTensorFmtEq(t, &tensor, "< @x(2)> [  1.000,  2.000]")
}

func TestTensorComformObjectInterface(t *testing.T) {
	shape := NewShape([]NamedDimension{{"x", 2}})
	array := &Array{[]float32{1.0, 2.0}}
	tensor := &Tensor{shape, array}
	var object Object
	object = tensor
	_, ok := object.(*Tensor)
	if !ok {
		t.Errorf("cast should work.")
	}
}
