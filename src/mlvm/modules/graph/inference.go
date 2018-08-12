package graph

import (
	"mlvm/modules/layers"
)

type DebuggingOptions struct {
	// Prints all debugging information for layers.
	PrintAllLayers bool
}

// Builds an inference graph and compile.
func NewInferenceGraph(outputs []layers.Layer, options *DebuggingOptions) InferenceGraph {
	if options == nil {
		options = &DebuggingOptions{}
	}

	g := &layerGraph{
		outputs: outputs,
		options: options,
	}
	g.Compile()
	return nil
}

// Represents a node for layer in graph.
type layerNode struct {
	Children []*layerNode
	IsRoot   bool // Is this node a root node.
	IsInput  bool // Is this layer input
	IsOutput bool // Is this layer output
}

type layerGraph struct {
	outputs []layers.Layer
	options *DebuggingOptions
}

func (g *layerGraph) Compile() {
	// TODO: Check has not compiled yet.

	allLayers := make([]layers.Layer, 0)

	root := &layerNode{IsRoot: true}
	root.Children = make([]*layerNode, 0, len(g.outputs))

	for _, layer := range g.outputs {
		allLayers = append(allLayers, layer)
		// FIXME: Fill nodes
		root.Children = append(root.Children, &layerNode{})
	}

	// Print layers.
	if g.options.PrintAllLayers {
		printLayersDebuggingInfo(allLayers)
	}

	// Tracs back to inputs
	// Color outputs and inputs
	// Save to graph
}
