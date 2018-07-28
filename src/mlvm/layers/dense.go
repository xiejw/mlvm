package layers

import(
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
	c "mlvm/base/context"
)

func NewInput(name string, shape t.Shape) t.Tensor {
}

func NewDenseLayer(
	ctx *c.Context, name string, input t.Tensor, output_dim Dimension) Layer {
}
