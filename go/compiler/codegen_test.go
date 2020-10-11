package compiler

import (
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
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
	_, err = Compile(m)
	assertNoErr(t, err)
}

//-----------------------------------------------------------------------------
// Helper Methods.
//-----------------------------------------------------------------------------

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
