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

func assertSingleOutput(t *testing.T, outputs Outputs, err error) object.Object {
	t.Helper()
	assertNoErr(t, err)
	if len(outputs) != 1 {
		t.Fatalf("unexpected single output, got: %v", outputs)
	}
	return outputs[0]
}

func TestRunWithOpCONST(t *testing.T) {
	program := &code.Program{
		Instructions: makeOpHelper(t, code.OpCONST, 0),
		Constants:    []object.Object{&object.Integer{123}},
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	if o.(*object.Integer).Value != 123 {
		t.Errorf("value mismatch.")
	}
}

func TestOpStoreAndLoad(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpSTORE, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpLOAD, 0)...)

	array := &object.Array{[]float32{1.0, 2.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{array},
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	expected := []float32{1.0, 2.0}
	assertAllClose(t, expected, o.(*object.Array).Value, 1e-6)
}

func TestRunWithOpTensor(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpTensor)...)

	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{shape, array},
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	if o.(*object.Tensor).String() != "Tensor(<@x(2)> [  1.000,  2.000])" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}

func TestRunWithOpPrng(t *testing.T) {
	seed := &object.Integer{456}
	shape := object.NewShape([]object.NamedDim{{"x", 4}})

	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpRNG)...)
	ins = append(ins, makeOpHelper(t, code.OpRNGT, 0)...)

	vm := NewVM(&code.Program{
		Instructions: ins,
		Constants:    []object.Object{seed, shape},
	})

	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	expected := []float32{1.3481823, -1.6701441, 1.4310317, 0.6320735}
	assertAllClose(t, expected, o.(*object.Array).Value, 1e-6)
}

func TestRunWithOpTensorAdd(t *testing.T) {
	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	ins1 := makeOpHelper(t, code.OpCONST, 0)
	ins2 := makeOpHelper(t, code.OpCONST, 1)
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
	ins = append(ins, makeOpHelper(t, code.OpTADD)...)

	program := &code.Program{
		Instructions: ins,
		Constants:    constants,
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	if o.(*object.Tensor).String() != "Tensor(<@x(2)> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}
