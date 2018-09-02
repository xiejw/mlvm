package compilation

import (
	"errors"

	"mlvm/modules/layers"
)

// Represents a node for layer in graph.
type LayerNode struct {
	Layer    layers.Layer
	Children []*LayerNode
	// delete?
	IsInput  bool // Is this layer input
	IsOutput bool // Is this layer output
}

// A topological sorted layer nodes in DAG.
type LayerDAG struct {
	Outputs []*LayerNode
	Inputs  []*LayerNode
	Layers  []*LayerNode // sorted in topological order.
}

func (dag *LayerDAG) Build(outputs []layers.Layer) error {
	// For permanent marker, the node's name is added into nodesPool. For
	// temporary marker, the node's name is added into visited.t
	nodesPool := make(map[string]*LayerNode)
	visited := make(map[string]bool)

	if len(outputs) == 0 {
		return errors.New("Outpus of Graph should not be empty.")
	}

	dag.Layers = make([]*LayerNode, 0)
	dag.Outputs = make([]*LayerNode, 0, len(outputs))
	for _, o := range outputs {
		node := visit(o, dag, nodesPool, visited)
		node.IsOutput = true
		dag.Outputs = append(dag.Outputs, node)
	}
	return nil
}

func visit(layer layers.Layer, dag *LayerDAG, nodesPool map[string]*LayerNode, visited map[string]bool) *LayerNode {
	name := layer.Name()
	if node, existed := nodesPool[name]; existed {
		return node // Permanent.
	}

	if visited[name] {
		panic("Not a DAG.") // Meet a temporary.
	}

	// Mark temporary.
	visited[name] = true

	inputs := layer.Inputs()
	node := &LayerNode{
		Layer:    layer,
		Children: make([]*LayerNode, 0, inputs.Count()),
	}
	for inputLayer := range inputs.Iterator() {
		inputNode := visit(inputLayer, dag, nodesPool, visited)
		node.Children = append(node.Children, inputNode)
	}

	// Mark permanent.
	nodesPool[name] = node
	dag.Layers = append(dag.Layers, node)
	return node
}
