package compilation

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"mlvm/modules/layers"
)

// Prints all debugging information about layers in lines (no hierarchy).
func PrintLayersDebuggingInfo(w io.Writer, allLayers []layers.Layer) {
	fmt.Fprintln(w, "## Layers:")
	tabw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	for _, layer := range allLayers {
		fmt.Fprintln(tabw, layer.String())
	}
	tabw.Flush()
}

const (
	dotGraphLineFormat       = `  "%v" -> "%v" [label=" %v " dir=back];`
	dotGraphWeightLineFormat = `  "%v" -> "%v"  [style=dashed, color=grey dir=back];`
	rootNodeName             = `(Outputs)`

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
  subgraph cluster_weights {
		style=dotted;
    label = "Weights";

		%v

  }
`
)

func joinNamesAsListForCluster(names []string) string {
	formattedNames := make([]string, 0, len(names))
	for _, n := range names {
		formattedNames = append(formattedNames, fmt.Sprintf(`"%v";`, n))
	}
	return strings.Join(formattedNames, " ")
}

func PrintLayersDotGraph(w io.Writer, root *LayerNode) {
	fmt.Fprintln(w, "/* Layers Dot Graph. */")
	fmt.Fprintln(w, "digraph G {")
	defer func() {
		fmt.Fprintln(w, "}")
	}()

	outputPlaceHolder := fmt.Sprintf(`"%v";`, rootNodeName)
	inputs := make([]string, 0)
	weights := make([]string, 0)

	var emitFn func(io.Writer, *LayerNode)
	emitFn = func(writer io.Writer, node *LayerNode) {

		// Write weights.
		if node.Layer != nil && node.Layer.Weights() != nil {
			for _, w := range node.Layer.Weights() {
				n := w.Name()
				weights = append(weights, n)
				fmt.Fprintln(writer, dotGraphWeightLineFormat, node.Layer.Name(), w)
			}
		}

		if node.Children == nil {
			// Use attributes, instead this logic.
			inputs = append(inputs, node.Layer.Name())
		}

		// Write children.
		for _, child := range node.Children {
			if node.IsRoot {
				fmt.Fprintln(writer,
					fmt.Sprintf(dotGraphLineFormat,
						rootNodeName, child.Layer.Name(), child.Layer.Output().Shape()))
			} else {
				fmt.Fprintln(writer,
					fmt.Sprintf(dotGraphLineFormat,
						node.Layer.Name(), child.Layer.Name(), child.Layer.Output().Shape()))
			}
			emitFn(writer, child)
		}
	}

	// Walks over the tree to record information.
	var buf bytes.Buffer
	emitFn(&buf, root)

	// Writes clusters first, then nodes connection.
	fmt.Fprintf(w, outputsCluster, outputPlaceHolder)
	fmt.Fprintf(w, inputsCluster, joinNamesAsListForCluster(inputs))
	fmt.Fprintf(w, weightsCluster, joinNamesAsListForCluster(weights))
	fmt.Fprintln(w)
	fmt.Fprintf(w, buf.String())
}
