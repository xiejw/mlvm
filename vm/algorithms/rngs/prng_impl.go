package rngs

/*
 * The DotMix algorithrm.
 * ----------------------
 *
 * 1. At each level in the pedigee, the coefficient `gamma` is used to advance
 *    the seed.
 *
 * 2. The addition-modulo, for the Dot in DotMix, is based on
 *
 *        George = 2^64 + 13L.
 *
 *    This is implemented as `prngUpdateSeed`. In particular,
 *
 *   a. Any new seed, after modulo, falls in [0, 2^64) is accepted.
 *   b. For [2^64, 2^64+13L), will try again. To avoid a loop, `gamma`, the
 *      coefficient, is constructed to be larger than 13L,
 *
 * 3. The Mix in DotMix is implemented in `prngMix64`.
 *
 *
 * 2. To avoid a fixed-size coefficient table.
 * Use coefficient to generate prng for coefficient.
 *
 *
 * Generating random coefficient.
 * ------------------------------
 * By using PRNG to generate coefficient for different level, a fixed table can
 * be void avoided. The algorithrm is still based on DotMix, but with a
 * length-1 pedigee.
 *
 * 1. The `gamma` coefficient for generating `gamma` is `gammaGamma`.
 *
 * 2. To improve the speed of `prngUpdateSeed`, a small Prime number (for
 *    modulo) is used for generating `gamma`, called Percy. It is the
 *    `gammaPrime` in the code below.
 *
 *    With that, the `prngUpdateSeed` has high chance take the first branch.
 *
 * 3. This gamma seed is then mixed by `prngMix56` to be _truely_ random.
 *
 * 4. The generated `gamma` should be no smaller than 13L. See DotMix
 *    algorithrm.
 */

func prngUpdateSeed(seed uint64, gamma uint64) uint64 {
	p := seed + gamma
	if p >= seed { // Most frequent branch.
		return p
	}

	if p >= 13 {
		return p - 13
	}

	return p + gamma - 13
}

func prngMix64(z uint64) uint64 {
	z = ((z ^ (z >> 33)) * 0xff51afd7ed558ccd)
	z = ((z ^ (z >> 33)) * 0xc4ceb9fe1a85ec53)
	return z ^ (z >> 33)
}

func prngMix56(z uint64) uint64 {
	z = ((z ^ (z >> 33)) * 0xff51afd7ed558ccd) & 0x00FFFFFFFFFFFFFF
	z = ((z ^ (z >> 33)) * 0xc4ceb9fe1a85ec53) & 0x00FFFFFFFFFFFFFF
	return z ^ (z >> 33)
}

func (prng *Rng64) advanceSeed() uint64 {
	// Advance one more coefficient at current level.
	newSeed := prngUpdateSeed(prng.Seed, prng.Gamma)
	prng.Seed = newSeed
	return newSeed
}
