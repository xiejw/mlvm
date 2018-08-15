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
	denseLayer_1 := layers.NewDense(ctx, "common_layer", concatLayer, 4)
	activation_1 := functions.Relu(ctx, denseLayer_1)

	// Output 1 For XOR
	denseLayer_2 := layers.NewDense(ctx, "layer_xor", activation_1, 3)
	output_1 := functions.Relu(ctx, denseLayer_2)

	// OUtput 2 For AND
	denseLayer_3 := layers.NewDense(ctx, "layer_and", activation_1, 3)
	output_2 := functions.Relu(ctx, denseLayer_3)

	dotFile, err := os.Create("/tmp/123.dot")
	if err != nil {
		log.Fatal(err)
	}
	defer dotFile.Close()

	_,err= g.NewInferenceGraph(ctx, []layers.Layer{output_1,output_2}, &g.DebuggingOptions{
		 LayerInfoWriter: os.Stdout,
		 LayerDotGraphWriter: dotFile,
	})

	if err != nil {
		log.Fatal(err)
	}
}
