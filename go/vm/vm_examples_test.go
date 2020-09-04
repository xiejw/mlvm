package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func TestExample1(t *testing.T) {
	var ins code.Instructions
	addIns(t, &ins, code.OpCONST, 0)
	addIns(t, &ins, code.OpCONST, 1)
	addIns(t, &ins, code.OpT)
	addIns(t, &ins, code.OpCONST, 0)
	addIns(t, &ins, code.OpCONST, 1)
	addIns(t, &ins, code.OpT)
	addIns(t, &ins, code.OpTADD)

	shape := object.NewShape([]uint{2, 3})
	array := &object.Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}}
	consts := []object.Object{shape, array}

	program := &code.Program{
		Instructions: ins,
		Constants:    consts,
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	expected := []float32{2.0, 4.0, 6.0, 8.0, 10.0, 12.0}
	assertAllClose(t, expected, o.(*object.Tensor).ArrayValue(), 1e-6)
}

func TestExample2(t *testing.T) {
	var ins code.Instructions
	addIns(t, &ins, code.OpCONST, 0)
	addIns(t, &ins, code.OpCONST, 1)
	addIns(t, &ins, code.OpRNG)
	addIns(t, &ins, code.OpRNGT, 0)
	addIns(t, &ins, code.OpSTORE, 0)
	addIns(t, &ins, code.OpLOAD, 0)
	addIns(t, &ins, code.OpLOAD, 0)
	addIns(t, &ins, code.OpTADD)

	shape := object.NewShape([]uint{2, 3})
	seed := &object.Integer{456}
	consts := []object.Object{shape, seed}

	program := &code.Program{
		Instructions: ins,
		Constants:    consts,
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)
	if o == nil {
		t.Fatalf("")
	}

	expected := []float32{1.3481823, -1.6701441, 1.4310317, 0.6320735, 0.28827125, 1.6303506}
	for i, e := range expected {
		expected[i] = 2 * e
	}
	assertAllClose(t, expected, o.(*object.Tensor).ArrayValue(), 1e-6)
}
