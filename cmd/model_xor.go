package main

import (
	"log"
	"os"

	c "mlvm/base/context"
	t "mlvm/base/tensor"
	"mlvm/modules/layers"
	g "mlvm/modules/graph"
	"mlvm/modules/functions"
)

func main() {

	ctx :=  (&c.ContextBuilder{ }).Build()

	inputShape := t.NewShapeWithBatchSize(1)

	// Inputs
	inputX := layers.NewInput(ctx, "x", inputShape, t.Float32)
	inputY := layers.NewInput(ctx, "y", inputShape, t.Float32)
	concatLayer := functions.Concat(
		ctx, "concat_inputs", []layers.Layer{inputX, inputY})

	// NN.
	denseLayer := layers.NewDense(ctx, "layer_1", concatLayer, 3)
	activation := functions.Relu(ctx, denseLayer)

	_,err := g.NewInferenceGraph(ctx, []layers.Layer{activation}, &g.DebuggingOptions{
		PrintAllLayers: os.Stdout,
	})

	if err != nil {
		log.Fatal(err)
	}
}
