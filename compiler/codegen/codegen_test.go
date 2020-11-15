package codegen

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/compiler/base/errors"
	"github.com/xiejw/mlvm/compiler/ir"
	"github.com/xiejw/mlvm/vm/code"
)

func TestConst(t *testing.T) {
	//--- ir
	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12)
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

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
