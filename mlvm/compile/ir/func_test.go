package ir

import (
	"reflect"
	"strings"
	"testing"

	"github.com/xiejw/mlvm/mlvm/array"
)

func TestNewFunc(t *testing.T) {
	_ = NewFunc()
}

func TestNewConstant(t *testing.T) {
	fn := NewFunc()
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	tensor := fn.NewConstantOrDie(arr)
	if tensor.Array() != arr {
		t.Errorf("Expected same array.")
	}
}

func TestNewConstants(t *testing.T) {
	fn := NewFunc()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("b", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	fn.NewConstantOrDie(a)
	fn.NewConstantOrDie(b)
}

func TestNewConstantsWithSameArray(t *testing.T) {
	fn := NewFunc()
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	fn.NewConstant(arr)

	defer func() {
		r := recover()
		if !strings.Contains(r.(error).Error(), "allow once") {
			t.Fatalf("Wrong error message: %v", r)
		}
	}()
	fn.NewConstantOrDie(arr)
	t.Fail()
}

func TestNewConstantsWithSameNames(t *testing.T) {
	fn := NewFunc()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	fn.NewConstantOrDie(a)

	defer func() {
		r := recover()
		if !strings.Contains(r.(error).Error(), "allow once") {
			t.Fatalf("Wrong error message: %v", r)
		}
	}()
	fn.NewConstantOrDie(b)
	t.Fail()
}

func TestNewInstruction(t *testing.T) {
	fn := NewFunc()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	ta := fn.NewConstantOrDie(a)

	ins := fn.NewInstructionOrDie(OpAdd(), ta, ta)
	if ins.Name() != "opAdd_001" {
		t.Fatalf("Instruction name mismatch. Got: %v.", ins.Name())
	}
	if !reflect.DeepEqual([]*Instruction{ins}, fn.Instructions()) {
		t.Fatalf("Instructions in Func mismatch.")
	}
}
