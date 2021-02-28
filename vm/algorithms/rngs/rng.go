package rngs

type Rng interface {
	Clone() Rng
	Split() Rng
	NextUI64() uint64
	NextF32() float32
}
