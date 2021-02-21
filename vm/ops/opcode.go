package ops

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
)

// -----------------------------------------------------------------------------
// op code.
// -----------------------------------------------------------------------------

type OpCode int

const (
	OP_RNG OpCode = iota
	OP_ADD
	OP_MUL
	OP_REDUCE
)

func (o OpCode) String() string {
	switch o {
	case OP_RNG:
		return "OP_RNG"
	case OP_ADD:
		return "OP_ADD"
	case OP_MUL:
		return "OP_MUL"
	case OP_REDUCE:
		return "OP_REDUCE"
	}
	return "(unknown)"
}

func (op OpCode) OutputTypes(operands []*object.Tensor, opt Option) (
	output_dtypes []object.DType, output_shapes [][]int, err error,
) {

	switch op {
	case OP_RNG:
		if len(operands) != 1 {
			err = errors.New("op (%v) expects only one operand; but got %v.", op, len(operands))
			return
		}
		if operands[0].DType != object.F32 {
			err = errors.New("op (%v) expects F32; but got %v.", op, operands[0].DType)
			return
		}
		if _, ok := opt.(*RngOption); !ok {
			err = errors.New("op (%v) expects RngOption; but got %v.", op, opt)
			return
		}
	default:
		err = errors.New("unsupported op (%v) for signature validation.", op)
	}
	return
}

// -----------------------------------------------------------------------------
// option.
// -----------------------------------------------------------------------------

type Option interface {
	Clone() Option
}

// -----------------------------------------------------------------------------
// rng option.
// -----------------------------------------------------------------------------

type RngDistType int

const (
	RngDistStdNorm RngDistType = iota
	RngDistTruncStdNorm
)

type RngOption struct {
	Rng      rngs.Rng
	DistType RngDistType
}

func (o *RngOption) Clone() Option {
	return &RngOption{
		Rng:      o.Rng.Clone(),
		DistType: o.DistType,
	}
}
