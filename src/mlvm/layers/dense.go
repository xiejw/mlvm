package layers

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

func NewInput(name string, shape t.Shape) t.Tensor {
}

func NewDenseLayer(
	ctx *c.Context, name string, input t.Tensor, output_dim Dimension) Layer {
}
