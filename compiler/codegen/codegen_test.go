package codegen

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/compiler/ir"
	"github.com/xiejw/mlvm/vm/code"
)

func TestConst(t *testing.T) {
	//--- ir
	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12)
	_ = f.ShapeLiteral([]int{1, 2})
	_ = f.ArrayLiteral([]float32{2})
	f.SetOutputAndDone(v.GetResult())

	m, err := b.Done()
	assertNoErr(t, err)

	//--- compile
	got, err := Compile(m)
	assertNoErr(t, err)

	expected := `
-> Constants:

[
    0: Integer(12)
    1: Shape(<1, 2>)
    2: Array([  2.000])
]

-> Instruction:

000000 OpCONST    0
`
	assertProgram(t, expected, got)
}

func TestRngSeed(t *testing.T) {
	//--- ir
	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12)
	r := f.RngSource(v.GetResult())
	f.SetOutputAndDone(r.GetResult())

	m, err := b.Done()
	assertNoErr(t, err)

	//--- compile
	got, err := Compile(m)
	assertNoErr(t, err)

	expected := `
-> Constants:

[
    0: Integer(12)
]

-> Instruction:

000000 OpCONST    0
000003 OpRNG
000004 OpSTORE    0
000007 OpLOAD     0
`
	assertProgram(t, expected, got)
}

func TestRngTensor(t *testing.T) {
	//--- ir
	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12).GetResult()
	s := f.ShapeLiteral([]int{2, 3}).GetResult()
	src := f.RngSource(v)
	r := f.RngTensor(s, src.GetResult())
	f.SetOutputAndDone(r.GetResult())

	m, err := b.Done()
	assertNoErr(t, err)

	//--- compile
	got, err := Compile(m)
	assertNoErr(t, err)

	expected := `
-> Constants:

[
    0: Integer(12)
    1: Shape(<2, 3>)
]

-> Instruction:

000000 OpCONST    0
000003 OpRNG
000004 OpSTORE    0
000007 OpCONST    1
000010 OpLOAD     0
000013 OpRNGT     0
000016 OpSTORE    1
000019 OpLOAD     1
`
	assertProgram(t, expected, got)
}

//-----------------------------------------------------------------------------
// Helper Methods.
//-----------------------------------------------------------------------------

func assertProgram(t *testing.T, expectedStr string, p *code.Program) {
	t.Helper()
	expected := strings.Trim(expectedStr, "\n\a")
	got := strings.Trim(p.String(), "\n\a")

	if expected != got {
		t.Errorf("mismatch program:\n\n===> expected:\n%v\n\n===> got:\n%v", expected, got)
		for i, c := range expected {
			if byte(c) != got[i] {
				t.Fatalf("the %d-th char is different: `%v` vs `%v`", i, c, got[i])
			}
		}
	}
}

func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
