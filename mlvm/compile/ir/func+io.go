package ir

import (
	"fmt"
)

var (
	errEmptyOutputs = fmt.Errorf("Func should have at least one output.")
)

// Set the outputs.
//
// After this, the Func is frozen.
func (fn *Func) SetOutputs(outputs []*Tensor) error {
	if len(outputs) == 0 {
		return errEmptyOutputs
	}

	err := fn.freeze()
	if err != nil {
		return err
	}

	fmt.Printf("Compiling: %v\n", fn)

	return nil
}

func (fn *Func) SetOutputsOrDie(outputs []*Tensor) {
	err := fn.SetOutputs(outputs)
	if err != nil {
		panic(err)
	}
}
