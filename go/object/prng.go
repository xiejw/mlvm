package object

import (
	"github.com/xiejw/mlvm/go/object/prng64"
)

type Prng struct {
	Source *prng64.Prng64
}

func (prng *Prng) Type() ObjectType {
	return PrngType
}

func (prng *Prng) String() string {
	return "Prng()"
}
