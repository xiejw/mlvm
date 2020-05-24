package prng64

import "fmt"

type DistType uint16

const (
	DistNorm DistType = iota
	DistTruncNorm
)

func (prng *Prng64) FillDist(distType DistType, value []float32) {
	switch distType {
	case DistNorm:
		prng.Norm(value)
		return
	case DistTruncNorm:
		prng.TruncNorm(value)
		return
	default:
		panic(fmt.Sprintf("unknown distribution type: %v", distType))
	}
}
