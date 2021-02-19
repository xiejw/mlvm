package rng

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Defines APIs for distributions.
// ----------------------------------------------------------------------------

type DistType uint16

const (
	DistStdNorm DistType = iota
	DistTruncStdNorm
)

func FillDist(rng Rng, distType DistType, value []float32) {
	switch distType {
	case DistStdNorm:
		StdNorm(rng, value)
		return
	case DistTruncStdNorm:
		TruncStdNorm(rng, value)
		return
	default:
		panic(fmt.Sprintf("unknown distribution type: %v", distType))
	}
}
