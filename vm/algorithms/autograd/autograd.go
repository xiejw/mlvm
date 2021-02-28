// package autograd provides grad support for ops.
package autograd

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
	"github.com/xiejw/mlvm/vm/ops"
)

func Grad(
	op ops.OpCode, opt ops.Option, inputs, outputs, grads []object.TensorLike,
) ([]object.TensorLike, error) {

	switch op {
	// case OP_ADD:
	// 	return "OP_ADD"
	// case OP_MUL:
	// 	return "OP_MUL"
	// case OP_SUM:
	// 	return "OP_SUM"
	case ops.OP_RNG:
		return nil, errors.New("op (%v) is not allowed to flow grad by autograd.", op)
	default:
		return nil, errors.New("op (%v) is not supported by autograd.", op)
	}

	return nil, nil
}
