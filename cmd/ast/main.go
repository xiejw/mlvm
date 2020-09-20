package main

import (
	"log"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/ir"
)

func assertNoErr(err *errors.DError) {
	if err != nil {
		log.Fatalf("did not expect error: %v", err)
	}
}

func main() {
	log.Printf("Hello MLVM")

	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(err)

	v := f.IntLiteral(12)
	f.SetOutput(v)

	m, err := b.Finalize()
	assertNoErr(err)

	fns := m.Fns()
	log.Printf("Fns: %v", fns)
}
