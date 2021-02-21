package ops

import (
	"reflect"

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

	case OP_MUL, OP_ADD:

		if len(operands) != 2 {
			err = errors.New("op (%v) expects two operands; but got %v.", op, len(operands))
			return
		}
		if operands[0].DType != object.F32 {
			err = errors.New("op (%v) expects F32 for operand; but got %v.", op, operands[0].DType)
			return
		}
		if operands[1].DType != object.F32 {
			err = errors.New("op (%v) expects F32 for operand; but got %v.", op, operands[1].DType)
			return
		}
		if opt != nil {
			err = errors.New("op (%v) expects nil Option; but got %v.", op, opt)
			return
		}

		if !reflect.DeepEqual(operands[0].Shape, operands[1].Shape) {
			err = errors.New("op (%v) operands' shape mismatch.")
			return
		}

		output_dtypes = append(output_dtypes, object.F32)
		output_shapes = append(output_shapes, operands[0].Shape.Dims)

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
