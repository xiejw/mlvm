package engine

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
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
		Constants:    []code.Object{&code.Integer{1}},
	}
	vm := NewVM(program)
	err = vm.Run()
	assertNil(t, err)
}
