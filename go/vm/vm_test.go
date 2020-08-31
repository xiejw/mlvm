package vm

import (
	"math"
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}

func TestCreateVM(t *testing.T) {
	vm := NewVM(&code.Program{})
	outputs, err := vm.Run()
	assertNoErr(t, err)

	if len(outputs) != 0 {
		t.Fatalf("stack should be empty.")
	}
}

func makeOpHelper(t *testing.T, op code.Opcode, args ...int) []byte {
	t.Helper()
	ins, err := code.MakeOp(op, args...)
	if err != nil {
		t.Fatalf("unxpected make op error: %v", err)
	}

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

func assertSingleOutput(t *testing.T, outputs Outputs, err *errors.DError) object.Object {
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

func TestOpIOR(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpIOR, 1)...)

	obj := &object.Integer{123}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{},
	}

	vm := NewVM(program)
	c := vm.InfeedChan()
	go func() {
		c <- obj
	}()
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	if o.(*object.Integer).Value != obj.Value {
		t.Errorf("value mismatch.")
	}
}

func TestRunWithOpTensor(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpT)...)

	shape := object.NewShape([]uint{2})
	array := &object.Array{[]float32{1.0, 2.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{shape, array},
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	if o.(*object.Tensor).String() != "Tensor(<2> [  1.000,  2.000])" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}

func TestRunWithOpPrng(t *testing.T) {
	seed := &object.Integer{456}
	shape := object.NewShape([]uint{4})

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
	shape := object.NewShape([]uint{2})
	array := &object.Array{[]float32{1.0, 2.0}}

	var constants []object.Object
	constants = append(constants, shape)
	constants = append(constants, array)

	ins1 := makeOpHelper(t, code.OpCONST, 0)
	ins2 := makeOpHelper(t, code.OpCONST, 1)
	ins3 := makeOpHelper(t, code.OpT)

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

	if o.(*object.Tensor).String() != "Tensor(<2> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", o.(*object.Tensor).String())
	}
}
