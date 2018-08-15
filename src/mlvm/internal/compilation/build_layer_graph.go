package compilation

import (
	"errors"

	"mlvm/modules/layers"
)

func (g *LayerGraph) BuildGraph() error {
	g.allLayers = make([]layers.Layer, 0)

	outputs := g.Outputs
	if len(outputs) == 0 {
		return errors.New("Outpus of Graph should not be empty.")
	}

	// Builds root node.
	root := &LayerNode{IsRoot: true}
	root.Children = make([]*LayerNode, 0, len(outputs))

	visitedLayers := make(map[string]*LayerNode)
	for _, outputLayer := range outputs {
		node := g.buildNodeForLayerGraph(outputLayer, visitedLayers)
		node.IsOutput = true
		root.Children = append(root.Children, node)
	}

	if g.Options.LayerInfoWriter != nil {
		PrintLayersDebuggingInfo(g.Options.LayerInfoWriter, g.allLayers)
	}
	if g.Options.LayerDotGraphWriter != nil {
		PrintLayersDotGraph(g.Options.LayerDotGraphWriter, root)
	}
	return nil
}

// Builds a node for the graph. Registers layers.
func (g *LayerGraph) buildNodeForLayerGraph(layer layers.Layer, visitedLayers map[string]*LayerNode) *LayerNode {
	if oldNode, existed := visitedLayers[layer.Name()]; existed {
		return oldNode
	}

	g.allLayers = append(g.allLayers, layer)

	node := &LayerNode{
		Layer: layer,
	}

	visitedLayers[layer.Name()] = node

	layerInputs := layer.Inputs()
	if layerInputs == nil {
		return node
	}

	node.Children = make([]*LayerNode, 0, layerInputs.Count())

	for childLayer := range layerInputs.Iterator() {
		childNode := g.buildNodeForLayerGraph(childLayer, visitedLayers)
		node.Children = append(node.Children, childNode)
	}

	return node
}
