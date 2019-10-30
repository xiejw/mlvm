package main

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/compile/ir"
)

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("\x1B[31mFatal error encoutered:\x1B[0m\n%v", r)
		}
	}()
	a := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	b := array.NewArrayOrDie("b", []array.Dimension{2, 1}, []array.Float{2.1, 3.2})
	fmt.Printf("Array %v: %v", a.Name(), a)
	fmt.Printf("Array %v: %v", b.Name(), b)

	fn := ir.NewFunc()
	ta := fn.NewConstantOrDie(a)
	tb := fn.NewConstantOrDie(b)
	fmt.Printf("Tensor %v: %v\n", ta.Name(), ta)
	fmt.Printf("Tensor %v: %v\n", tb.Name(), tb)

	ins := fn.NewInstructionOrDie(ir.OpAdd(), ta, tb)

	fn.SetOutputsOrDie([]*ir.Tensor{ins.OnlyResult()})
}
