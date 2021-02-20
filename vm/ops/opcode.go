package ops

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
)

type OpCode int

const (
	OP_RNG OpCode = iota
	OP_ADD
	OP_REDUCE
)

type Option interface {
	Clone() Option
}

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
