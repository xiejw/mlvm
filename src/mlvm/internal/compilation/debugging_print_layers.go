package compilation

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// Prints all debugging information about layers in lines, from inputs to
// outputs.
func PrintLayersDebuggingInfo(w io.Writer, dag *LayerDAG) {
	fmt.Fprintln(w, "## Layers:")
	tabw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	for _, node := range dag.Nodes {
		fmt.Fprintln(tabw, node.Layer.String())
	}
	tabw.Flush()
}
