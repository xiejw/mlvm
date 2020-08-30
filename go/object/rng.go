package object

import (
	"fmt"

	"github.com/xiejw/mlvm/go/object/prng64"
)

type Rng struct {
	Source *prng64.Prng64
}

func (rng *Rng) Type() ObjectType {
	return RngType
}

func (rng *Rng) String() string {
	src := rng.Source
	return fmt.Sprintf("Rng(%x, %x, %x)", src.Seed, src.Gamma, src.NextGammaSeed)
}

///////////////////////////////////////////////////////////////////////////////
// Defines APIs for distributions.
///////////////////////////////////////////////////////////////////////////////

type DistType uint16

const (
	DistNorm DistType = iota
	DistTruncNorm
)

func (rng *Rng) FillDist(distType DistType, value []float32) {
	switch distType {
	case DistNorm:
		rng.Source.Norm(value)
		return
	case DistTruncNorm:
		rng.Source.TruncNorm(value)
		return
	default:
		panic(fmt.Sprintf("unknown distribution type: %v", distType))
	}
}
