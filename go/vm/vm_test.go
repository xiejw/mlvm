package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}

func TestCreateEngine(t *testing.T) {
	vm := NewVM(&code.Program{})
	err := vm.Run()
	assertNil(t, err)
}

func TestRunWithOpConstant(t *testing.T) {
	ins, err := code.MakeOp(code.OpConstant, 0)
	assertNil(t, err)

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{&object.Integer{123}},
	}

	vm := NewVM(program)
	o := vm.StackTop()
	if o != nil {
		t.Fatalf("stack should be empty.")
	}

	err = vm.Run()
	assertNil(t, err)

	o = vm.StackTop()
	if o.(*object.Integer).Value != 123 {
		t.Errorf("value mismatch.")
	}
}

func TestRunWithOpTensor(t *testing.T) {

	ins1, err := code.MakeOp(code.OpConstant, 0)
	assertNil(t, err)

	ins2, err := code.MakeOp(code.OpConstant, 1)
	assertNil(t, err)

	ins3, err := code.MakeOp(code.OpTensor)
	assertNil(t, err)

	var ins code.Instructions
	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)

	shape := object.NewShape([]object.NamedDimension{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	program := &code.Program{
		Instructions: ins,
		Constants:    constants,
	}

	vm := NewVM(program)
	o := vm.StackTop()
	if o != nil {
		t.Fatalf("stack should be empty.")
	}

	err = vm.Run()
	assertNil(t, err)

	o = vm.StackTop()
	if o.(*object.Tensor).String() != "< @x(2)> [  1.000,  2.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}

func TestRunWithOpTensorAdd(t *testing.T) {
	shape := object.NewShape([]object.NamedDimension{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	ins1, err := code.MakeOp(code.OpConstant, 0)
	assertNil(t, err)

	ins2, err := code.MakeOp(code.OpConstant, 1)
	assertNil(t, err)

	ins3, err := code.MakeOp(code.OpTensor)
	assertNil(t, err)

	var ins code.Instructions
	// Operand 1
	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)
	// Operand 2
	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)

	insAdd, err := code.MakeOp(code.OpAdd)
	assertNil(t, err)
	ins = append(ins, insAdd...)

	program := &code.Program{
		Instructions: ins,
		Constants:    constants,
	}

	vm := NewVM(program)
	o := vm.StackTop()
	if o != nil {
		t.Fatalf("stack should be empty.")
	}

	err = vm.Run()
	assertNil(t, err)

	o = vm.StackTop()
	if o.(*object.Tensor).String() != "< @x(2)> [  1.000,  2.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}
