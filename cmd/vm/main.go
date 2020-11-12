package main

import (
	"fmt"
	"log"

	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/code"
	"github.com/xiejw/mlvm/vm/mach"
	"github.com/xiejw/mlvm/vm/object"
)

func main() {
	seed := &object.Integer{456}
	shape := object.NewShape([]int{4})

	var ins code.Instructions
	ins = append(ins, makeOpHelper(code.OpCONST, 1)...)
	ins = append(ins, makeOpHelper(code.OpCONST, 0)...)
	ins = append(ins, makeOpHelper(code.OpRNG)...)
	ins = append(ins, makeOpHelper(code.OpRNGT, 0)...)

	vm := mach.NewVM(&code.Program{
		Instructions: ins,
		Constants:    []object.Object{seed, shape},
	})

	outputs, err := vm.Run()
	assertNoErr(err)

	fmt.Printf("output: %v", outputs[0])

}

// ----------------------------------------------------------------------------
// Helper methods
// ----------------------------------------------------------------------------

func assertNoErr(err *errors.DError) {
	if err != nil {
		log.Fatalf("did not expect error. got: %v", err)
	}
}

func makeOpHelper(op code.Opcode, args ...int) []byte {
	ins, err := code.MakeOp(op, args...)
	if err != nil {
		log.Fatalf("unxpected make op error: %v", err)
	}
	return ins
}
