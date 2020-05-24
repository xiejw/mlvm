package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func makeOpHelper(t *testing.T, op code.Opcode, args ...int) []byte {
	t.Helper()
	ins, err := code.MakeOp(op, args...)
	assertNoErr(t, err)

	return ins
}

func TestRunWithOpConstant(t *testing.T) {
	ins, err := code.MakeOp(code.OpConstant, 0)
	assertNoErr(t, err)

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{&object.Integer{123}},
	}

	vm := NewVM(program)
	err = vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Integer).Value != 123 {
		t.Errorf("value mismatch.")
	}
}

func TestOpStoreAndLoad(t *testing.T) {
}

func TestRunWithOpTensor(t *testing.T) {

	ins1, err := code.MakeOp(code.OpConstant, 0)
	assertNoErr(t, err)

	ins2, err := code.MakeOp(code.OpConstant, 1)
	assertNoErr(t, err)

	ins3, err := code.MakeOp(code.OpTensor)
	assertNoErr(t, err)

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
	err = vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Tensor).String() != "<@x(2)> [  1.000,  2.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}

func TestRunWithOpPrng(t *testing.T) {
	seed := &object.Integer{456}
	shape := object.NewShape([]object.NamedDimension{{"x", 2}})
	// value = &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, seed)
	constants = append(constants, shape)

	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpPrngNew)...)
	ins = append(ins, makeOpHelper(t, code.OpPrngDist, 0)...)
}

func TestRunWithOpTensorAdd(t *testing.T) {
	shape := object.NewShape([]object.NamedDimension{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	ins1, err := code.MakeOp(code.OpConstant, 0)
	assertNoErr(t, err)

	ins2, err := code.MakeOp(code.OpConstant, 1)
	assertNoErr(t, err)

	ins3, err := code.MakeOp(code.OpTensor)
	assertNoErr(t, err)

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
	assertNoErr(t, err)
	ins = append(ins, insAdd...)

	program := &code.Program{
		Instructions: ins,
		Constants:    constants,
	}

	vm := NewVM(program)
	err = vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Tensor).String() != "<@x(2)> [  2.000,  4.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}
