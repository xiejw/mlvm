package compiler

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/ir"
)

func TestConst(t *testing.T) {
	//--- ir
	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(t, err)

	v := f.IntLiteral(12)
	f.SetOutput(v.GetResult())

	m, err := b.Finalize()
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

//-----------------------------------------------------------------------------
// Helper Methods.
//-----------------------------------------------------------------------------

func assertProgram(t *testing.T, expectedStr string, p *code.Program) {
	t.Helper()
	expected := strings.Trim(expectedStr, "\n")
	got := strings.Trim(p.String(), "\n")

	if expected != got {
		t.Fatalf("mismatch program:\n\n===> expected:\n%v\n\n===> got:\n%v", expected, got)
	}
}

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
