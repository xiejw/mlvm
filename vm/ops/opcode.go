package ops

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
)

// -----------------------------------------------------------------------------
// op code.
// -----------------------------------------------------------------------------

type OpCode int

const (
	OP_RNG OpCode = iota
	OP_ADD
	OP_REDUCE
)

func (o OpCode) String() string {
	switch o {
	case OP_RNG:
		return "OP_RNG"
	case OP_ADD:
		return "OP_ADD"
	case OP_REDUCE:
		return "OP_REDUCE"
	}
	return "(unknown)"
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
