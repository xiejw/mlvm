// Package rng provides implementation of prng.
package rng

const (
	gammaPrime uint64  = (1 << 56) - 5      // Percy.
	gammaGamma uint64  = 0x00281E2DBA6606F3 // The coefficient to randomize gamma.
	doubleUlp  float32 = 1.0 / (1 << 53)
)

// To persist the Png64, two uint64 is enough. as the gamma can be calculated from NextGammaSeed.
// However, to trade for performance, 3 uint64 could be better.
type Prng64 struct {
	Seed          uint64
	Gamma         uint64
	NextGammaSeed uint64
}

func NewPrng64(seed uint64) *Prng64 {
	return newPrng64(seed, 0 /*gammaSeed*/)
}

func (prng *Prng64) Split() *Prng64 {
	seed := prng.advanceSeed()
	gammaSeed := prng.NextGammaSeed
	return newPrng64(seed, gammaSeed)
}

func newPrng64(seed uint64, gammaSeed uint64) *Prng64 {
	if gammaSeed >= gammaPrime {
		panic("gammaSeed passed to new prng is too large.")
	}

	// Advance the gamma seed.
	gammaSeed += gammaGamma
	if gammaSeed >= gammaPrime {
		gammaSeed -= gammaPrime // Constrain the range for gamma seed.
	}

	prng := &Prng64{
		Seed:          seed,
		Gamma:         prngMix56(gammaSeed) + 13,
		NextGammaSeed: gammaSeed,
	}
	return prng

}

func (prng *Prng64) NextInt64() uint64 {
	return prngMix64(prng.advanceSeed())
}

func (prng *Prng64) NextFloat32() float32 {
	return float32(prng.NextInt64()>>11) * doubleUlp
}
