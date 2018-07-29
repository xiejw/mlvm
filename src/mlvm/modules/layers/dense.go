package layers

import (
	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

func NewDenseLayer(
	ctx *c.Context, name string, input t.Tensor, output_dim t.Dimension) Layer {
	return nil
}
