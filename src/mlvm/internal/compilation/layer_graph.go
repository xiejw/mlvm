package compilation

import (
	"mlvm/modules/layers"
)

type Options struct {
	// Prints all debugging information for layers.
	PrintAllLayers bool
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
	allLayers []layers.Layer
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
