package main

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
	"mlvm/modules/layers"
)

func main() {
	inputShape := t.NewShapeWithBatchSize(1, 2)
	fmt.Printf("InputShape %v\n", inputShape)


	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()
	fmt.Printf("IsTraining %v\n", ctx.IsTraining())

	input := layers.NewInput(ctx, "x", inputShape, t.Float32)

	fmt.Printf("Input layer %v\n", input)
}
