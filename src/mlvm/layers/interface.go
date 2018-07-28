package layers

import (
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

struct Layer interface {
	func Weights() []w.Weight

	func Apply(args ...t.Tensor) t.Tensor

	func Backprop()
}
