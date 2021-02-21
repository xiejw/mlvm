// Package rngs provides implementation of rng.
package rngs

const (
	gammaPrime uint64  = (1 << 56) - 5      // Percy.
	gammaGamma uint64  = 0x00281E2DBA6606F3 // The coefficient to randomize gamma.
	doubleUlp  float32 = 1.0 / (1 << 53)
)

// To persist the Rng64, two uint64 is enough. as the gamma can be calculated from NextGammaSeed.
// However, to trade for performance, 3 uint64 could be better.
type Rng64 struct {
	Seed          uint64
	Gamma         uint64
	NextGammaSeed uint64
}

// -----------------------------------------------------------------------------
// constructor.
// -----------------------------------------------------------------------------

func NewRng64(seed uint64) *Rng64 {
	return newRng64(seed, 0 /*gammaSeed*/)
}

func (prng *Rng64) Split() Rng {
	seed := prng.advanceSeed()
	gammaSeed := prng.NextGammaSeed
	return newRng64(seed, gammaSeed)
}

func (r *Rng64) Clone() Rng {
	return &Rng64{
		Seed:          r.Seed,
		Gamma:         r.Gamma,
		NextGammaSeed: r.NextGammaSeed,
	}
}

// -----------------------------------------------------------------------------
// internal constructor.
// -----------------------------------------------------------------------------

func newRng64(seed uint64, gammaSeed uint64) *Rng64 {
	if gammaSeed >= gammaPrime {
		panic("gammaSeed passed to new prng is too large.")
	}

	// Advance the gamma seed.
	gammaSeed += gammaGamma
	if gammaSeed >= gammaPrime {
		gammaSeed -= gammaPrime // Constrain the range for gamma seed.
	}

	prng := &Rng64{
		Seed:          seed,
		Gamma:         rng64Mix56(gammaSeed) + 13,
		NextGammaSeed: gammaSeed,
	}
	return prng

}

// -----------------------------------------------------------------------------
// conform Rng interface.
// -----------------------------------------------------------------------------

func (prng *Rng64) NextUI64() uint64 {
	return rng64Mix64(prng.advanceSeed())
}

func (prng *Rng64) NextF32() float32 {
	return float32(prng.NextUI64()>>11) * doubleUlp
}

