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
	outputs, err := vm.Run()
	assertNoErr(t, err)

	if len(outputs) != 0 {
		t.Fatalf("stack should be empty.")
	}
}
