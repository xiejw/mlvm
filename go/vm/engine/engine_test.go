package engine

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
)

func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("did not expect error.")
	}
}

func TestCreateEngine(t *testing.T) {
	vm := NewVM(&code.Program{})
	err := vm.Run()
	assertNil(t, err)
}
