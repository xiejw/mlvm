package rng

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Defines APIs for distributions.
// ----------------------------------------------------------------------------

type DistType uint16

const (
	DistNorm DistType = iota
	DistTruncNorm
)

func FillDist(src *Prng64, distType DistType, value []float32) {
	switch distType {
	case DistNorm:
		src.Norm(value)
		return
	case DistTruncNorm:
		src.TruncNorm(value)
		return
	default:
		panic(fmt.Sprintf("unknown distribution type: %v", distType))
	}
}
