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

func TestArrayLit(t *testing.T) {
	f, b := newMainFn(t)

	v := f.ArrayLiteral([]float32{1, 2})
	f.SetOutputAndDone(v.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ArrayLit([  1.000,  2.000])
  return %0
}

}`
	assertModule(t, expected, got)
}

func TestTensorNew(t *testing.T) {
	f, b := newMainFn(t)

	s := f.ShapeLiteral([]int{1, 2}).GetResult()
	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	te := f.TensorNew(s, a)
	f.SetOutputAndDone(te.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  %1 = ArrayLit([  1.000,  2.000])
  %2 = TensorNew(%0, %1)
  return %2
}

}`
	assertModule(t, expected, got)
}

func TestTensorAdd(t *testing.T) {
	f, b := newMainFn(t)

	s1 := f.ShapeLiteral([]int{1, 2}).GetResult()
	s2 := f.ShapeLiteral([]int{2}).GetResult()
	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	t1 := f.TensorNew(s1, a).GetResult()
	t2 := f.TensorNew(s2, a).GetResult()
	r := f.TensorAdd(t1, t2)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  %1 = ShapeLit(<2>)
  %2 = ArrayLit([  1.000,  2.000])
  %3 = TensorNew(%0, %2)
  %4 = TensorNew(%1, %2)
  %5 = TensorAdd(%3, %4)
  return %5
}

}`
	assertModule(t, expected, got)
}

func TestTensorMinus(t *testing.T) {
	f, b := newMainFn(t)

	s1 := f.ShapeLiteral([]int{1, 2}).GetResult()
	s2 := f.ShapeLiteral([]int{2}).GetResult()
	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	t1 := f.TensorNew(s1, a).GetResult()
	t2 := f.TensorNew(s2, a).GetResult()
	r := f.TensorMinus(t1, t2)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  %1 = ShapeLit(<2>)
  %2 = ArrayLit([  1.000,  2.000])
  %3 = TensorNew(%0, %2)
  %4 = TensorNew(%1, %2)
  %5 = TensorMinus(%3, %4)
  return %5
}

}`
	assertModule(t, expected, got)
}

func TestTensorMul(t *testing.T) {
	f, b := newMainFn(t)

	s1 := f.ShapeLiteral([]int{1, 2}).GetResult()
	s2 := f.ShapeLiteral([]int{2}).GetResult()
	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	t1 := f.TensorNew(s1, a).GetResult()
	t2 := f.TensorNew(s2, a).GetResult()
	r := f.TensorMul(t1, t2)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  %1 = ShapeLit(<2>)
  %2 = ArrayLit([  1.000,  2.000])
  %3 = TensorNew(%0, %2)
  %4 = TensorNew(%1, %2)
  %5 = TensorMul(%3, %4)
  return %5
}

}`
	assertModule(t, expected, got)
}

func TestTensorDiv(t *testing.T) {
	f, b := newMainFn(t)

	s1 := f.ShapeLiteral([]int{1, 2}).GetResult()
	s2 := f.ShapeLiteral([]int{2}).GetResult()
	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	t1 := f.TensorNew(s1, a).GetResult()
	t2 := f.TensorNew(s2, a).GetResult()
	r := f.TensorDiv(t1, t2)
	f.SetOutputAndDone(r.GetResult())

	got, err := b.Done()
	assertNoErr(t, err)

	expected := `
module {

fn main() {
  %0 = ShapeLit(<1, 2>)
  %1 = ShapeLit(<2>)
  %2 = ArrayLit([  1.000,  2.000])
  %3 = TensorNew(%0, %2)
  %4 = TensorNew(%1, %2)
  %5 = TensorDiv(%3, %4)
  return %5
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
