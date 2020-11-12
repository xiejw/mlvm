package main

import (
	"log"
	// "github.com/xiejw/mlvm/go/base/errors"
	// "github.com/xiejw/mlvm/go/compiler"
	// "github.com/xiejw/mlvm/go/ir"
	// vm_lib "github.com/xiejw/mlvm/go/vm"
)

// func assertNoErr(err *errors.DError) {
// 	if err != nil {
// 		log.Fatalf("did not expect error: %v", err)
// 	}
// }

func main() {
	log.Printf("Hello MLVM")

	// //---------------------------------------------------------------------------
	// // Builds IR
	// //---------------------------------------------------------------------------

	// b := ir.NewBuilder()
	// f, err := b.NewFn("main")
	// assertNoErr(err)

	// v := f.IntLiteral(12)
	// r := f.RngSeed(v)
	// f.SetOutput(r.GetResult())

	// m, err := b.Finalize()
	// assertNoErr(err)

	// log.Printf("module: \n%v", m)

	// //---------------------------------------------------------------------------
	// // Compiles to Program
	// //---------------------------------------------------------------------------

	// p, err := compiler.Compile(m)
	// assertNoErr(err)

	// log.Printf("program: \n%v", p)

	// //---------------------------------------------------------------------------
	// // To Run with VM
	// //---------------------------------------------------------------------------
	// vm := vm_lib.NewVM(p)
	// outputs, err := vm.Run()
	// assertNoErr(err)
	// log.Printf("vm output: \n%v", outputs)
}
