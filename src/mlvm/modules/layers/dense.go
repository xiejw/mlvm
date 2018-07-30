package layers

import (
	c "mlvm/base/context"
	_ "mlvm/base/weight"
)

func NewDenseLayer(
	ctx *c.Context, name string, input Layer, output_unit int) Layer {
	_ = ctx.AssignUniqueName(name)
	return nil
}
