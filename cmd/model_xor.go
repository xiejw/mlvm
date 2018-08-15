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
	denseLayer_1 := layers.NewDense(ctx, "layer_1", concatLayer, 4)
	activation_1 := functions.Relu(ctx, denseLayer_1)
	denseLayer_2 := layers.NewDense(ctx, "layer_2", activation_1, 3)
	activation_2 := functions.Relu(ctx, denseLayer_2)

	dotFile, err := os.Create("/tmp/123.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer dotFile.Close()

	_,err= g.NewInferenceGraph(ctx, []layers.Layer{activation_2}, &g.DebuggingOptions{
		 LayerInfoWriter: os.Stdout,
		 LayerDotGraphWriter: dotFile,
	})

	if err != nil {
		log.Fatal(err)
	}
}
