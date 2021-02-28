package main

import (
	"fmt"

	"github.com/xiejw/mlvm/vm/algorithms/rngs"
	"github.com/xiejw/mlvm/vm/mach"
	"github.com/xiejw/mlvm/vm/nn"
	"github.com/xiejw/mlvm/vm/object"
)

func main() {
	fmt.Printf("hello mlvm:\n")
	rng := rngs.NewRng64(123)

	vm := new(mach.VM)

	w := nn.RngStdNorm(vm, rng, object.F32, []int{2, 3})
	w.RequireGrad()

	logits := nn.Mul(w, w)
	loss := nn.Sum(logits)
	nn.Backward(loss)

	fmt.Printf("Grad : %v\n", w.Grad())
}
