package object

import (
	"fmt"

	"github.com/xiejw/mlvm/go/object/prng64"
)

type Prng struct {
	Source *prng64.Prng64
}

func (prng *Prng) Type() ObjectType {
	return PrngType
}

func (prng *Prng) String() string {
	return "Prng()"
}

///////////////////////////////////////////////////////////////////////////////
// Defines APIs for distributions.
///////////////////////////////////////////////////////////////////////////////

type DistType uint16

const (
	DistNorm DistType = iota
	DistTruncNorm
)

func (prng *Prng) FillDist(distType DistType, value []float32) {
	switch distType {
	case DistNorm:
		prng.Source.Norm(value)
		return
	case DistTruncNorm:
		prng.Source.TruncNorm(value)
		return
	default:
		panic(fmt.Sprintf("unknown distribution type: %v", distType))
	}
}
