package layers

import (
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

type Layer interface {
	Weights() []w.Weight

	Apply(args ...t.Tensor) t.Tensor

	Backprop()
}
