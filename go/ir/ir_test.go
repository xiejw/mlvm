package ir

import (
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
)

func TestModule(t *testing.T) {
	b := NewBuilder()
	m, err := b.Finalize()
	assertNoErr(t, err)

	if m.Fns() != nil {
		t.Errorf("expect empty module.")
	}
}

///////////////////////////////////////////////////////////////////////////////
// Helper Methods.
///////////////////////////////////////////////////////////////////////////////

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
