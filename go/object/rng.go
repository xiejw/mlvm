package object

import (
	"fmt"
)

type Rng struct {
	Seed          uint64
	Gamma         uint64
	NextGammaSeed uint64
}

func (rng *Rng) Type() ObjectType {
	return RngType
}

func (rng *Rng) String() string {
	return fmt.Sprintf("Rng(%x, %x, %x)", rng.Seed, rng.Gamma, rng.NextGammaSeed)
}
