package ops

import (
	"github.com/xiejw/mlvm/vm/algorithms/rngs"
)

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

// -----------------------------------------------------------------------------
// sum option.
// -----------------------------------------------------------------------------

type SumOption struct {
	Dims []int
}

func (o *SumOption) Clone() Option {
	dims := make([]int, len(o.Dims))
	copy(dims, o.Dims)
	return &SumOption{
		Dims: dims,
	}
}
