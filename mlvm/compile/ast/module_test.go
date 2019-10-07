package ast

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/mlvm/array"
)

func TestNewModule(t *testing.T) {
	_ = NewModule()
}

func TestNewConstant(t *testing.T) {
	m := NewModule()
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	tensor := m.NewConstant(arr)
	if tensor.Array() != arr {
		t.Errorf("Expected same array.")
	}
}

func TestNewConstants(t *testing.T) {
	m := NewModule()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("b", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	m.NewConstant(a)
	m.NewConstant(b)
}

func TestNewConstantsWithSameArray(t *testing.T) {
	m := NewModule()
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	m.NewConstant(arr)

	defer func() {
		r := recover()
		if !strings.Contains(r.(string), "allow once") {
			t.Fatalf("Wrong error message: %v", r)
		}
	}()
	m.NewConstant(arr)
	t.Fail()
}

func TestNewConstantsWithSameNames(t *testing.T) {
	m := NewModule()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	m.NewConstant(a)

	defer func() {
		r := recover()
		if !strings.Contains(r.(string), "allow once") {
			t.Fatalf("Wrong error message: %v", r)
		}
	}()
	m.NewConstant(b)
	t.Fail()
}
