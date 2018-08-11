package main

import (
	"os"
	"fmt"
	"text/tabwriter"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	"mlvm/modules/layers"
	"mlvm/modules/functions"
)

func main() {

	ctx :=  (&c.ContextBuilder{
		IsTraining: false,
	}).Build()

	inputShape := t.NewShapeWithBatchSize(1)

	// Inputs
	inputX := layers.NewInput(ctx, "x", inputShape, t.Float32)
	inputY := layers.NewInput(ctx, "y", inputShape, t.Float32)
	concatLayer := functions.Concat(
		ctx, "concat_inputs", []layers.Layer{inputX, inputY})

	// NN.
	denseLayer := layers.NewDense(ctx, "layer_1", concatLayer, 3)
	activation := functions.Relu(ctx, denseLayer)

	// Output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, inputX.String())
	fmt.Fprintln(w, inputY.String())
	fmt.Fprintln(w, concatLayer.String())
	fmt.Fprintln(w, denseLayer.String())
	fmt.Fprintln(w, activation.String())
	w.Flush()
}
