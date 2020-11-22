package main

import (
	"log"

	"github.com/xiejw/mlvm/compiler/codegen"
	"github.com/xiejw/mlvm/compiler/ir"
	"github.com/xiejw/mlvm/vm/mach"
)

func main() {
	log.Printf("Hello MLVM")

	//---------------------------------------------------------------------------
	// Builds IR
	//---------------------------------------------------------------------------

	b := ir.NewBuilder()
	f, err := b.NewFn("main")
	assertNoErr(err)

	v := f.IntLiteral(12).GetResult()
	s := f.ShapeLiteral([]int{2, 3}).GetResult()
	src := f.RngSource(v)
	r := f.RngFill(s, src.GetResult())

	a := f.ArrayLiteral([]float32{1, 2}).GetResult()
	_ = f.NewTensor(s, a)
	f.SetOutputAndDone(r.GetResult())

	m, err := b.Done()
	assertNoErr(err)

	log.Printf("module: \n%v", m)

	//---------------------------------------------------------------------------
	// Compiles to Program
	//---------------------------------------------------------------------------

	p, err := codegen.Compile(m)
	assertNoErr(err)

	log.Printf("program: \n%v", p)

	//---------------------------------------------------------------------------
	// To Run with VM
	//---------------------------------------------------------------------------
	vm := mach.NewVM(p)
	outputs, err := vm.Run()
	assertNoErr(err)
	log.Printf("vm output: \n%v", outputs)
}

//------------------------------------------------------------------------------
// Helper methods.
//------------------------------------------------------------------------------
func assertNoErr(err error) {
	if err != nil {
		log.Fatalf("did not expect error: %v", err)
	}
}
