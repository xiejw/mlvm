package ir

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/compiler/base/errors"
)

func TestModule(t *testing.T) {
	b := NewBuilder()
	m, err := b.Done()
	assertNoErr(t, err)

	if m.Fns() != nil {
		t.Errorf("expect empty module.")
	}
}

func TestIntLit(t *testing.T) {
	f, b := newMainFn(t)

	v := f.IntLiteral(12)
	f.SetOutputAndDone(v.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = IntLit(12)
  return %0
}

}`
	assertModule(t, expected, got)
}

func TestRndSeed(t *testing.T) {
	f, b := newMainFn(t)

	v := f.IntLiteral(12).GetResult()
	r := f.RngSeed(v)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = IntLit(12)
  %1 = RngSeed(%0)
  return %1
}

}`
	assertModule(t, expected, got)
}

//-----------------------------------------------------------------------------
// Helper Methods.
//-----------------------------------------------------------------------------

func newMainFn(t *testing.T) (*Fn, *Builder) {
	t.Helper()
	b := NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)
	return f, b
}

func assertModule(t *testing.T, expectedStr string, m *Module) {
	t.Helper()
	expected := strings.Trim(expectedStr, "\n")
	got := strings.Trim(m.String(), "\n")

	if expected != got {
		t.Fatalf("mismatch module:\n\n===> expected:\n%v\n\n===> got:\n%v", expected, got)
	}
}

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
