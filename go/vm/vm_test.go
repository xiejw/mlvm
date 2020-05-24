package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
)

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}

func TestCreateVM(t *testing.T) {
	vm := NewVM(&code.Program{})
	err := vm.Run()
	assertNoErr(t, err)

	o := vm.StackTop()
	if o != nil {
		t.Fatalf("stack should be empty.")
	}
}
