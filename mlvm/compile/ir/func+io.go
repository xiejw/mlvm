package ir

import (
	"github.com/xiejw/mlvm/mlvm/internal/errors"
)

var (
	errEmptyOutputs = errors.Errorf("Func should have at least one output.")
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

	return nil
}

func (fn *Func) SetOutputsOrDie(outputs []*Tensor) {
	err := fn.SetOutputs(outputs)
	if err != nil {
		panic(err)
	}
}
