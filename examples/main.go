package main

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/compile/ir"
)

func main() {
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("b", []array.Dimension{2, 1}, []array.Float{2.1, 3.2})
	fmt.Printf("Array %v: %v", a.Name(), a)
	fmt.Printf("Array %v: %v", b.Name(), b)

	m := ir.NewModule()
	ta := m.NewConstantOrDie(a)
	tb := m.NewConstantOrDie(b)
	fmt.Printf("Tensor %v: %v\n", ta.Name(), ta)
	fmt.Printf("Tensor %v: %v\n", tb.Name(), tb)

	m.NewInstructionOrDie(ir.OpAdd(), ta, tb)

	m.CompileOrDie()
}
