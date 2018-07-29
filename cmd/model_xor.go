package main

import (
	"fmt"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

func main() {
	inputShape := t.NewShapeWithBatchSize(1, 2)
	fmt.Printf("InputShape %v\n", inputShape)

	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()
	fmt.Printf("IsTraining %v\n", ctx.IsTraining())
}
