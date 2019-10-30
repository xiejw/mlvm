package ir

import (
	"github.com/xiejw/mlvm/mlvm/internal/names"
)

func newInstruction(name string, op *Op, operands ...*Tensor) (*Instruction, error) {
	if err := names.ValidateInstructionName(name); err != nil {
		return nil, err
	}

	ins := &Instruction{
		name:     name,
		op:       op,
		operands: operands,
	}

	resultShapes, err := op.InferShapes(operands...)
	if err != nil {
		return nil, err
	}

	results := make([]*Tensor, 0, len(resultShapes))
	for i, resultShape := range resultShapes {
		result := &Result{
			name:  names.CanonicalResultName(name, i),
			shape: resultShape,
			ins:   ins,
			index: i,
		}
		results = append(results, newResultTensor(result))
	}

	ins.results = results
	return ins, nil
}
