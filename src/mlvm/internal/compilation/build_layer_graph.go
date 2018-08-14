package compilation

import (
	"errors"

	"mlvm/modules/layers"
)

// TODO: dedup layer based on name.
func (g *LayerGraph) BuildGraph() error {
	g.allLayers = make([]layers.Layer, 0)

	outputs := g.Outputs
	if len(outputs) == 0 {
		return errors.New("Outpus of Graph should not be empty.")
	}

	// Builds root node.
	root := &LayerNode{IsRoot: true}
	root.Children = make([]*LayerNode, 0, len(outputs))

	for _, outputLayer := range outputs {
		node := g.buildNodeForLayerGraph(outputLayer)
		node.IsOutput = true
		root.Children = append(root.Children, node)
	}

	// Print layers.
	if g.Options.PrintAllLayers != nil {
		printLayersDebuggingInfo(g.Options.PrintAllLayers, g.allLayers)
	}
	return nil
}

// Builds a node for the graph. Registers layers.
func (g *LayerGraph) buildNodeForLayerGraph(layer layers.Layer) *LayerNode {
	// TODO(xiejw): dedup.
	g.allLayers = append(g.allLayers, layer)

	node := &LayerNode{
		Layer: layer,
	}

	layerInputs := layer.Inputs()
	if layerInputs == nil {
		return node
	}

	node.Children = make([]*LayerNode, 0, layerInputs.Count())

	for childLayer := range layerInputs.Iterator() {
		childNode := g.buildNodeForLayerGraph(childLayer)
		node.Children = append(node.Children, childNode)
	}

	return node
}
