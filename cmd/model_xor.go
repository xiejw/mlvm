package main

import (
	"fmt"

	c "mlvm/base/context"
	_ "mlvm/base/tensor"
	_ "mlvm/base/weight"
)

func main() {
	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()
	fmt.Printf("IsTraining %v\n", ctx.IsTraining())
}
