package prng64

import (
	"math"
)

const (
	twoPi float64 = 2 * math.Pi
)

var eplisonFloat32 = math.Nextafter32(0, 1)

/*
 * The implementation is based on Box–Muller transform.
 *
 * For each pair of [0, 1) uniform rn, a pair of independent, standard,
 * normally distributed rn are generated.
 */
func (prng *Prng64) Norm(value []float32) {
	size := len(value)

	numSeeds := size
	if size%2 != 0 {
		numSeeds++ // Must be even
	}

	uniforms := make([]float32, numSeeds)

	for i := 0; i < numSeeds; i++ {
		seed := prng.NextFloat32()
		// The first rn in each pair is used by log, so cannot be zero.
		if i%2 == 1 || seed >= eplisonFloat32 {
			uniforms[i] = seed
		}
	}

	for i := 0; i < size; i += 2 {
		u1 := float64(uniforms[i])
		u2 := float64(uniforms[i+1])

		r := math.Sqrt(-2.0 * math.Log(u1))
		theta := twoPi * u2

		value[i] = float32(r * math.Cos(theta))
		if i+1 < size {
			value[i+1] = float32(r * math.Sin(theta))
		}
	}
}
