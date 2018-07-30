package main

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
	"mlvm/modules/layers"
)

func main() {

	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()

	inputShape := t.NewShapeWithBatchSize(1, 2)
	inputLayer := layers.NewInput(ctx, "x", inputShape, t.Float32)

	denseLayer := layers.NewDense(ctx, "first_layer", inputLayer, 3)

	fmt.Printf("-> %v\n", inputLayer)
	fmt.Printf("-> %v\n", denseLayer)
}
