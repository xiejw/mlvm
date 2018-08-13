package graph

import (
	c "mlvm/base/context"
	"mlvm/modules/layers"
)

type DebuggingOptions struct {
	// Prints all debugging information for layers.
	PrintAllLayers bool
}

// Builds an inference graph and compile.
// Graph is DAG.
func NewInferenceGraph(ctx *c.Context, outputs []layers.Layer, options *DebuggingOptions) (InferenceGraph, error) {
	if options == nil {
		options = &DebuggingOptions{}
	}

	g := &layerGraph{
		outputs: outputs,
		options: options,
	}
	if err := g.Compile(); err != nil {
		return nil, err
	}
	return nil, nil
}

// Represents a node for layer in graph.
type layerNode struct {
	Layer    layers.Layer
	Children []*layerNode
	IsRoot   bool // Is this node a root node.
	IsInput  bool // Is this layer input
	IsOutput bool // Is this layer output
}

type layerGraph struct {
	// Sets by constructor.
	options *DebuggingOptions
	outputs []layers.Layer

	// Internal state.
	allLayers []layers.Layer
}

func (g *layerGraph) Compile() error {
	// TODO: Check has not compiled yet.

	if err := g.BuildGraph(); err != nil {
		return err
	}
	// Tracs back to inputs
	// Color outputs and inputs
	// Save to graph
	return nil
}
