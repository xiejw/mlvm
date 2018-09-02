package compilation

import (
	"fmt"
	"io"
	"strings"
)

const rootNodeName = `(Outputs)`

const (
	dotGraphLineFormat       = `  "%v" -> "%v" [label=" %v " dir=back];`
	dotGraphWeightLineFormat = `  "%v" -> "%v"  [label=" %v " style=dashed, color=grey dir=back];`
)

// Template for clusters
const (
	outputsCluster = `
  subgraph cluster_outputs {
    style=filled;
    color=lightgrey;
    node [style=filled,color=white];

		%v

    label = "";
    labelloc="t";
  }
`
	inputsCluster = `
  subgraph cluster_inputs {
    style=filled;
    color=lightgrey;
    node [style=filled,color=white];

		%v

    label = "Inputs";
    labelloc="b";
  }
`
	weightsCluster = `
  subgraph cluster_weights_%v {
		style=dotted;
    label = "Weights for %v";
    labelloc="b";

		%v

  }
`
)

// Prints the Dot graph.
func PrintLayersDotGraph(w io.Writer, dag *LayerDAG) {
	fmt.Fprintln(w, "/* Layers Dot Graph. */")
	fmt.Fprintln(w, "digraph G {")
	defer func() {
		fmt.Fprintln(w, "}")
	}()

	// Writes clusters first, then nodes connection.
	printOutputsCluster(w)
	printInputsCluster(w, dag)
	printWeightsCluster(w, dag)

	fmt.Fprintln(w)
	printNodeConnections(w, dag)
}

// Prints the outputs cluster.
func printOutputsCluster(w io.Writer) {
	outputPlaceHolder := fmt.Sprintf(`"%v";`, rootNodeName)
	fmt.Fprintf(w, outputsCluster, outputPlaceHolder)
}

// Prints the inputs cluster. It is a box surrounding all input nodes.
func printInputsCluster(w io.Writer, dag *LayerDAG) {
	inputs := make([]string, 0, len(dag.Inputs))
	for _, input := range dag.Inputs {
		inputs = append(inputs, input.Layer.Name())
	}
	fmt.Fprintf(w, inputsCluster, joinNamesAsListForCluster(inputs))
}

// Prints the weights cluster.
func printWeightsCluster(w io.Writer, dag *LayerDAG) {
	index := 0
	for layerName, weights := range dag.Weights {
		weightNames := make([]string, 0, len(weights))
		for _, w := range weights {
			weightNames = append(weightNames, w.Name())
		}
		fmt.Fprintf(w, weightsCluster, index, layerName,
			joinNamesAsListForCluster(weightNames))
		index++
	}
}

// Prints the node connections. This should be printed after clusters.
func printNodeConnections(writer io.Writer, dag *LayerDAG) {
	// Root -> Outputs
	for _, o := range dag.Outputs {
		fmt.Fprintln(writer,
			fmt.Sprintf(dotGraphLineFormat,
				rootNodeName, o.Layer.Name(), o.Layer.Output().Shape()))
	}

	// Real node connections.
	for _, node := range dag.Nodes {
		// Weights connection.
		if node.Layer.Weights() != nil {
			for _, w := range node.Layer.Weights() {
				n := w.Name()
				fmt.Fprintf(writer, dotGraphWeightLineFormat, node.Layer.Name(), n, w.Shape())
			}
		}
		// Children connection.
		for _, child := range node.Children {
			fmt.Fprintln(writer,
				fmt.Sprintf(dotGraphLineFormat,
					node.Layer.Name(), child.Layer.Name(), child.Layer.Output().Shape()))
		}
	}
}

func joinNamesAsListForCluster(names []string) string {
	formattedNames := make([]string, 0, len(names))
	for _, n := range names {
		formattedNames = append(formattedNames, fmt.Sprintf(`"%v";`, n))
	}
	return strings.Join(formattedNames, " ")
}
