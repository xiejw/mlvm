package ir

import (
	"bytes"
	"strings"
	"testing"
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

func TestShapeLit(t *testing.T) {
	f, b := newMainFn(t)

	v := f.ShapeLiteral([]int{1, 2})
	f.SetOutputAndDone(v.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  return %0
}

}`
	assertModule(t, expected, got)
}

func TestRndSeed(t *testing.T) {
	f, b := newMainFn(t)

	v := f.IntLiteral(12).GetResult()
	r := f.RngSource(v)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = IntLit(12)
  %1 = RngSource(%0)
  return %1
}

}`
	assertModule(t, expected, got)
}

//func TestRndSeed(t *testing.T) {

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

	var buf bytes.Buffer
	m.DebugString(&buf, false /* printType*/)
	got := strings.Trim(buf.String(), "\n")

	if expected != got {
		t.Fatalf("mismatch module:\n\n===> expected:\n%v\n\n===> got:\n%v", expected, got)
	}
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
