package graph

import (
	"io"

	c "mlvm/base/context"
	"mlvm/modules/layers"

	"mlvm/internal/compilation"
)

type DebuggingOptions struct {
	// Writes all debugging information for layers. `nil` if not needed.
	LayerInfoWriter     io.Writer
	LayerDotGraphWriter io.Writer
}

// Builds an inference graph and compile.
// Graph is DAG.
func NewInferenceGraph(ctx *c.Context, outputs []layers.Layer, options *DebuggingOptions) (InferenceGraph, error) {
	g := &compilation.LayerGraph{
		Outputs: outputs,
		Options: convertOptions(options),
	}
	if err := g.Compile(); err != nil {
		return nil, err
	}
	return nil, nil
}

// Converts user provided DebuggingOptions to compilation options.
func convertOptions(options *DebuggingOptions) *compilation.Options {
	if options == nil {
		return &compilation.Options{}
	}
	// Copy cover.
	return &compilation.Options{
		LayerInfoWriter:     options.LayerInfoWriter,
		LayerDotGraphWriter: options.LayerDotGraphWriter,
	}
}
