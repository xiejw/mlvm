package compilation

import (
	"io"

	"mlvm/modules/layers"
)

type Options struct {
	// Prints all debugging information for layers.
	LayerInfoWriter     io.Writer
	LayerDotGraphWriter io.Writer
}

// Represents a node for layer in graph.
type LayerNode struct {
	Layer    layers.Layer
	Children []*LayerNode
	IsRoot   bool // Is this node a root node.
	IsInput  bool // Is this layer input
	IsOutput bool // Is this layer output
}

type LayerGraph struct {
	// Sets by constructor.
	Options *Options
	Outputs []layers.Layer

	// Internal state.
	dag *LayersDAG
}

// A topologyci sorted sort of the layers in DAG.
type LayersDAG struct {
}

func (g *LayerGraph) Compile() error {
	// TODO: Check has not compiled yet.

	if err := g.BuildGraph(); err != nil {
		return err
	}
	// Tracs back to inputs
	// Color outputs and inputs
	// Save to graph
	return nil
}
