package vm

import (
	"math"
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

func assertAllClose(t *testing.T, expected, got []float32, tol float64) {
	t.Helper()
	if len(expected) != len(got) {
		t.Fatalf("length mismatch. expected: %v, got: %v.", len(expected), len(got))
	}

	for i := 0; i < len(expected); i++ {
		if math.Abs(float64(expected[i]-got[i])) >= tol {
			t.Errorf("\nelement mismatch at %v: expected %v, got %v\n", i, expected[i], got[i])
		}
	}
}

func TestRunWithOpConstant(t *testing.T) {
	program := &code.Program{
		Instructions: makeOpHelper(t, code.OpConstant, 0),
		Constants:    []object.Object{&object.Integer{123}},
	}

	vm := NewVM(program)
	err := vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Integer).Value != 123 {
		t.Errorf("value mismatch.")
	}
}

func TestOpStoreAndLoad(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpStoreG, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpLoadG, 0)...)

	array := &object.Array{[]float32{1.0, 2.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{array},
	}

	vm := NewVM(program)
	err := vm.Run()
	assertNoErr(t, err)

	expected := []float32{1.0, 2.0}
	assertAllClose(t, expected, vm.StackTop().(*object.Array).Value, 1e-6)
}

func TestRunWithOpTensor(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpConstant, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpTensor)...)

	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{shape, array},
	}

	vm := NewVM(program)
	err := vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Tensor).String() != "<@x(2)> [  1.000,  2.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}

func TestRunWithOpPrng(t *testing.T) {
	seed := &object.Integer{456}
	shape := object.NewShape([]object.NamedDim{{"x", 4}})

	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpConstant, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpConstant, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpPrngNew)...)
	ins = append(ins, makeOpHelper(t, code.OpPrngDist, 0)...)

	vm := NewVM(&code.Program{
		Instructions: ins,
		Constants:    []object.Object{seed, shape},
	})
	err := vm.Run()
	assertNoErr(t, err)

	expected := []float32{1.3481823, -1.6701441, 1.4310317, 0.6320735}
	assertAllClose(t, expected, vm.StackTop().(*object.Array).Value, 1e-6)
}

func TestRunWithOpTensorAdd(t *testing.T) {
	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	ins1 := makeOpHelper(t, code.OpConstant, 0)
	ins2 := makeOpHelper(t, code.OpConstant, 1)
	ins3 := makeOpHelper(t, code.OpTensor)

	var ins code.Instructions
	// Operand 1
	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)
	// Operand 2
	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)
	// Final Add
	ins = append(ins, makeOpHelper(t, code.OpAdd)...)

	program := &code.Program{
		Instructions: ins,
		Constants:    constants,
	}

	vm := NewVM(program)
	err := vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o.(*object.Tensor).String() != "<@x(2)> [  2.000,  4.000]" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}
