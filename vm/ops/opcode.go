package ops

import (
	"reflect"

	"github.com/xiejw/mlvm/vm/algorithms/linalg"
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
	OP_SUM
)

func (o OpCode) String() string {
	switch o {
	case OP_RNG:
		return "OP_RNG"
	case OP_ADD:
		return "OP_ADD"
	case OP_MUL:
		return "OP_MUL"
	case OP_SUM:
		return "OP_SUM"
	}
	return "(unknown)"
}

// execs on allocated buffer, i.e., Tensor.
func (op OpCode) Exec(operands []*object.Tensor, outputs []*object.Tensor, opt Option) error {
	switch op {
	case OP_RNG:
		value := operands[0].Data().([]float32)
		rng_opt := opt.(*RngOption)

		switch rng_opt.DistType {
		case RngDistStdNorm:
			rngs.StdNorm(rng_opt.Rng, value)
		case RngDistTruncStdNorm:
			rngs.TruncStdNorm(rng_opt.Rng, value)
		default:
			return errors.New("failed to execute rng: unknown distribution type: %v", rng_opt.DistType)
		}
		return nil

	case OP_ADD:
		err := linalg.Add(&linalg.Context{},
			operands[0].Data().([]float32),
			operands[1].Data().([]float32),
			outputs[0].Data().([]float32))
		if err != nil {
			return errors.WrapNote(err, "failed to execute linalg.Add.")
		}
		return nil

	case OP_MUL:
		err := linalg.Mul(&linalg.Context{},
			operands[0].Data().([]float32),
			operands[1].Data().([]float32),
			outputs[0].Data().([]float32))
		if err != nil {
			return errors.WrapNote(err, "failed to execute linalg.Mul.")
		}
		return nil

	case OP_SUM:
		err := linalg.Sum(&linalg.Context{},
			operands[0].Data().([]float32),
			operands[0].Shape().Dims,
			opt.(*SumOption).Dims,
			outputs[0].Data().([]float32))
		if err != nil {
			return errors.WrapNote(err, "failed to execute linalg.Mul.")
		}
		return nil

	default:
		return errors.New("unsupported op (%v) for execution", op)
	}
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
		if operands[0].DType() != object.F32 {
			err = errors.New("op (%v) expects F32; but got %v.", op, operands[0].DType())
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
		if operands[0].DType() != object.F32 {
			err = errors.New("op (%v) expects F32 for operand; but got %v.", op, operands[0].DType())
			return
		}
		if operands[1].DType() != object.F32 {
			err = errors.New("op (%v) expects F32 for operand; but got %v.", op, operands[1].DType())
			return
		}
		if opt != nil {
			err = errors.New("op (%v) expects nil Option; but got %v.", op, opt)
			return
		}

		if !reflect.DeepEqual(operands[0].Shape(), operands[1].Shape()) {
			err = errors.New("op (%v) operands' shape mismatch.", op)
			return
		}

		output_dtypes = append(output_dtypes, object.F32)
		output_shapes = append(output_shapes, operands[0].Shape().Dims)

	case OP_SUM:
		if len(operands) != 1 {
			err = errors.New("op (%v) expects only one operand; but got %v.", op, len(operands))
			return
		}
		if operands[0].DType() != object.F32 {
			err = errors.New("op (%v) expects F32; but got %v.", op, operands[0].DType())
			return
		}
		if o, ok := opt.(*SumOption); !ok {
			err = errors.New("op (%v) expects SumOption; but got %v.", op, opt)
			return
		} else if !reflect.DeepEqual(operands[0].Shape().Dims, o.Dims) {
			err = errors.New(
				"op (%v) expects reducing all dims; but got operand dims %v, reducing dims %v.", op,
				operands[0].Shape().Dims, o.Dims)
			return
		}

		output_dtypes = append(output_dtypes, object.F32)
		output_shapes = append(output_shapes, []int{1})

	default:
		err = errors.New("unsupported op (%v) for signature validation.", op)
	}
	return
}

func (op OpCode) AllowGrad(operands []object.TensorLike, opt Option) error {
	switch op {
	case OP_RNG:
		return errors.New("op (%v) is not allowed to flow grad.", op)
	case OP_MUL, OP_ADD, OP_SUM:
		if !operands[0].DType().AllowGrad() {
			return errors.New("op (%v) is not allowed to flow grad for dtype %v.", op, operands[0].DType())
		}
		return nil
	default:
		return errors.New("unsupported op (%v) for allowing grad.", op)
	}
}
