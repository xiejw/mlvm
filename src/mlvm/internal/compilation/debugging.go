package compilation

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"mlvm/modules/layers"
)

const (
	dotGraphLineFormat = `  "%v" -> "%v" [label=" %v " dir=back];`
	rootNodeName       = `(Root)`
)

func PrintLayersDebuggingInfo(w io.Writer, allLayers []layers.Layer) {
	fmt.Fprintln(w, "## Layers:")
	tabw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	for _, layer := range allLayers {
		fmt.Fprintln(tabw, layer.String())
	}
	tabw.Flush()
}

const (
	outputsCluster = `
  subgraph cluster_outputs {
    style=filled;
    color=lightgrey;
    node [style=filled,color=white];

		%v

    label = "Outpus";
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
)

func PrintLayersDotGraph(w io.Writer, root *LayerNode) {
	fmt.Fprintln(w, "/* Layers Dot Graph. */")
	fmt.Fprintln(w, "digraph G {")
	defer func() {
		fmt.Fprintln(w, "}")
	}()

	outputs := make([]string, 0, len(root.Children))
	inputs := make([]string, 0)

	var emitFn func(io.Writer, *LayerNode)
	emitFn = func(writer io.Writer, node *LayerNode) {
		if node.Children == nil {
			// Use attributes, instead this logic.
			inputs = append(inputs, fmt.Sprintf(`"%v";`, node.Layer.Name()))
		}
		for _, child := range node.Children {
			if node.IsRoot {
				outputs = append(outputs, fmt.Sprintf(`"%v";`, child.Layer.Name()))
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
	fmt.Fprintf(w, outputsCluster, strings.Join(outputs, " "))
	fmt.Fprintf(w, inputsCluster, strings.Join(inputs, " "))
	fmt.Fprintln(w)
	fmt.Fprintf(w, buf.String())
}
