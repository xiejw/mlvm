package rngs

type Rng interface {
	Split() Rng
	NextUI64() uint64
	NextF32() float32
}
