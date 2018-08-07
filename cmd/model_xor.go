package main

import (
	"os"
	"fmt"
	"text/tabwriter"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	_ "mlvm/base/weight"
	"mlvm/modules/layers"
	"mlvm/modules/functions"
)

func main() {

	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()

	inputShape := t.NewShapeWithBatchSize(1)

	inputLayer := layers.NewInput(ctx, "x", inputShape, t.Float32)
	denseLayer := layers.NewDense(ctx, "first_layer", inputLayer, 3)
	activation := functions.Relu(ctx, denseLayer)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, inputLayer.String())
	fmt.Fprintln(w, denseLayer.String())
	fmt.Fprintln(w, activation.String())
	w.Flush()
}
