package mach

import (
	"fmt"
	"os"

	//	"github.com/xiejw/mlvm/vm/algorithms/autograd"
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/ops"
)

// whether enable debug.
var debugGradTape bool

func init() {
	debugGradTape = len(os.Getenv("MLVM_DEBUG_GRAD_TAPE")) != 0
}

// most are expected to be immutable. few exceptions are
//
// - handle passed to non-grad-flow op which later becomes required grad.
//   handle -> rng -> require_grad.
type Record struct {
	Op       ops.OpCode
	Operands []*Handle
	Outputs  []*Handle
	Option   ops.Option
}

type Tape struct {
	Records []*Record
	GradDAG []*Record
}

func (t *Tape) RecordOpAndGradDAG(
	op ops.OpCode, operands []*Handle, outputs []*Handle, opt ops.Option, flowGrad bool,
) error {
	r := &Record{
		Op:       op,
		Operands: operands,
		Outputs:  outputs,
		Option:   opt,
	}

	for _, o := range outputs {
		o.record = r
		o.flowGrad = flowGrad
	}

	t.Records = append(t.Records, r)
	if flowGrad {
		t.GradDAG = append(t.GradDAG, r)
	}
	return nil
}

// algorithrm
// - prepare the ones tensor matching the shape of passed in x.
// - prepare a map, called handle_grad, from handle to current grad. start with the x -> ones
//   created in last step.
// - walk through dag in topological reverse order. iterate on each Op.
//   - look up the gradients for outputs handles, via handle_grad. error if cannot find.
//   - call autograd to get the gradients.
//   - for each operand
//     - discard gradient if neither require grad nor flow grad.
//     - if flow grad, add the current grad on top of the old one in handle_grad and save new result
//       to handle_grad
//     - if require_grad, add the current grad on top of the handle'grad handle.
func (t *Tape) BProp(x *Handle) error {
	if !x.flowGrad {
		return errors.New(
			"Handle (%v) has not grad flowing through it; so cannot be used for Backward", x)
	}

	vm := x.VM()

	// print debug info if requested.
	if debugGradTape {
		fmt.Printf("--> tape: \n")
		for _, r := range t.Records {
			fmt.Printf("  %+v\n", r)
		}
		fmt.Printf("--> dag: \n")
		for _, r := range t.GradDAG {
			fmt.Printf("  %+v\n", r)
		}
	}

	gradX, err := vm.ones(x.DType(), x.Shape().Dims)
	if err != nil {
		return err
	}

	var debugGrads []*Handle

	grads := make(map[*Handle]*Handle)
	grads[x] = gradX
	if debugGradTape {
		debugGrads = append(debugGrads, x)
	}

	// dagSize := len(t.GradDAG)

	// for _, r := range t.GradDAG {
	// 	_, err := autograd.Grad(
	// 		r.Op,
	// 		r.Option,
	// 		handlesToTensorLikes(r.Operands),
	// 		handlesToTensorLikes(r.Outputs),
	// 		nil)
	// 	if err != nil {
	// 		return err
	// 	}

	// }

	// print debug info if requested.
	if debugGradTape {
		fmt.Printf("--> dag grads: \n")
		for _, k := range debugGrads {
			v := grads[k]
			fmt.Printf("  handle/grad: %v / %v\n", k, v)
			fmt.Printf("    value : %v\n", k.tensor)
			fmt.Printf("    grad  : %v\n", v.tensor)
		}
	}
	return nil
}
