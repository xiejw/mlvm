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

func TestSimpleFn(t *testing.T) {
	b := NewBuilder()

	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12)
	f.SetOutput(v)

	m, err := b.Finalize()
	assertNoErr(t, err)

	fns := m.Fns()

	if len(fns) != 1 {
		t.Errorf("expect one fn .")
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
