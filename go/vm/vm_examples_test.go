package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func TestExample1(t *testing.T) {
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

