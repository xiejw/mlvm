package graph

import (
	"io"

	c "mlvm/base/context"
	"mlvm/modules/layers"

	"mlvm/internal/compilation"
)

type DebuggingOptions struct {
	// Prints all debugging information for layers.
	PrintAllLayers io.Writer
}

// Builds an inference graph and compile.
// Graph is DAG.
func NewInferenceGraph(ctx *c.Context, outputs []layers.Layer, options *DebuggingOptions) (InferenceGraph, error) {

	// Copy cover.
	var compilationOpt *compilation.Options
	if options == nil {
		compilationOpt = &compilation.Options{}
	} else {
		compilationOpt = &compilation.Options{
			PrintAllLayers: options.PrintAllLayers,
		}
	}

	g := &compilation.LayerGraph{
		Outputs: outputs,
		Options: compilationOpt,
	}
	if err := g.Compile(); err != nil {
		return nil, err
	}
	return nil, nil
}
