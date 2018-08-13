package graph

import (
	"errors"

	"mlvm/modules/layers"
)

// TODO: dedup layer based on name.
func (g *layerGraph) BuildGraph() error {
	g.allLayers = make([]layers.Layer, 0)

	outputs := g.outputs
	if len(outputs) == 0 {
		return errors.New("Outpus of Graph should not be empty.")
	}

	// Builds root node.
	root := &layerNode{IsRoot: true}
	root.Children = make([]*layerNode, 0, len(outputs))

	for _, outputLayer := range outputs {
		node := g.buildNodeForLayer(outputLayer)
		node.IsOutput = true
		root.Children = append(root.Children, node)
	}

	// Print layers.
	if g.options.PrintAllLayers {
		printLayersDebuggingInfo(g.allLayers)
	}
	return nil
}

// Builds a node for the graph. Registers layers.
func (g *layerGraph) buildNodeForLayer(layer layers.Layer) *layerNode {
	// TODO(xiejw): dedup.
	g.allLayers = append(g.allLayers, layer)

	node := &layerNode{
		Layer: layer,
	}

	layerInputs := layer.Inputs()
	if layerInputs == nil {
		return node
	}

	node.Children = make([]*layerNode, 0, layerInputs.Count())

	for childLayer := range layerInputs.Iterator() {
		childNode := g.buildNodeForLayer(childLayer)
		node.Children = append(node.Children, childNode)
	}

	return node
}
