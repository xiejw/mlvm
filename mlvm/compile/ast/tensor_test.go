package ast

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/mlvm/array"
)

func TestNewConstantTensor(t *testing.T) {
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	tensor := newConstantTensor(arr)

	if tensor.Name() != arr.Name() {
		t.Fatalf("Tensor name mismatch.")
	}
	if !reflect.DeepEqual(arr.Shape().Dims(), tensor.Shape().Dims()) {
		t.Fatalf("Tensor shape mismatch. Expected: %v, Got: %v.",
			arr.Shape().Dims(), tensor.Shape().Dims())
	}
	if tensor.Kind() != KConstant {
		t.Fatalf("Tensor kind mismatch.")
	}
	if tensor.Array() != arr {
		t.Fatalf("Tensor arry instance mismatch.")
	}
}
