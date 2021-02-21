package rngs

import (
	"math"
)

const (
	twoPi float64 = 2 * math.Pi
)

var eplisonFloat32 = math.Nextafter32(0, 1)

// -----------------------------------------------------------------------------
// design:
//
// The implementation is based on Box–Muller transform.
//
// For each pair of [0, 1) uniform rn, a pair of independent, standard,
// normally distributed rn are generated.
// -----------------------------------------------------------------------------

func BoxMullerTransform(rng Rng) (float32, float32) {
	u1 := rng.NextF32()
	u2 := rng.NextF32()

	if u1 < eplisonFloat32 {
		u1 = eplisonFloat32
	}

	r := math.Sqrt(-2.0 * math.Log(float64(u1)))
	theta := twoPi * float64(u2)

	return float32(r * math.Cos(theta)), float32(r * math.Sin(theta))
}

func StdNorm(rng Rng, value []float32) {
	size := len(value)

	for i := 0; i < size; i += 2 {
		r1, r2 := BoxMullerTransform(rng)

		value[i] = r1
		if i+1 < size {
			value[i+1] = r2
		}
	}
}
