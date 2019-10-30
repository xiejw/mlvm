package ir

import (
	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/internal/errors"
	"github.com/xiejw/mlvm/mlvm/internal/shapes"
)

type OpKind int

type ResultShapesInferenceFn func([]*array.Shape) ([]*array.Shape, error)

const (
	OpKAdd OpKind = iota + 1 // 0 Is invalid.
)

var (
	opConstAdd = &Op{kind: OpKAdd}
)

// Oper
type Op struct {
	kind OpKind
}

func (op *Op) Kind() OpKind {
	return op.kind
}

func (op *Op) BaseName() string {
	switch op.kind {
	case OpKAdd:
		return "opAdd"
	default:
		panic("Op Kind is not expected.")
	}

}

func OpAdd() *Op {
	return opConstAdd
}

func (op *Op) InferShapes(operands ...*Tensor) ([]*array.Shape, error) {
	// Convert operands array to shape array.
	operandShapes := make([]*array.Shape, 0, len(operands))
	for _, operand := range operands {
		operandShapes = append(operandShapes, operand.Shape())
	}

	var resultShapes []*array.Shape
	var err error

	switch op.kind {

	case OpKAdd:
		resultShapes, err = shapes.InferResultShapesForElementWiseOp(operandShapes)
	default:
		panic("Op Kind is not expected.")
	}

	if err != nil {
		return nil, errors.ErrorfW(
			err, "Fail to infer result shapes for Op kind `%v`", op.BaseName())
	}

	return resultShapes, nil
}
