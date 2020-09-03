package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func TestExample1(t *testing.T) {
	var ins code.Instructions
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpT)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(t, code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(t, code.OpT)...)
	ins = append(ins, makeOpHelper(t, code.OpTADD)...)

	shape := object.NewShape([]uint{2, 3})
	array := &object.Array{[]float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}}

	program := &code.Program{
		Instructions: ins,
		Constants:    []object.Object{shape, array},
	}

	vm := NewVM(program)
	outputs, err := vm.Run()
	o := assertSingleOutput(t, outputs, err)

	expected := []float32{2.0, 4.0, 6.0, 8.0, 10.0, 12.0}
	assertAllClose(t, expected, o.(*object.Tensor).ArrayValue(), 1e-6)
}
