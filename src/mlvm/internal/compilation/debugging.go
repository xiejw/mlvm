package compilation

import (
	"fmt"
	"io"
	"text/tabwriter"

	"mlvm/modules/layers"
)

func printLayersDebuggingInfo(w io.Writer, allLayers []layers.Layer) {
	fmt.Fprintln(w, "## Layers:")
	tabw := tabwriter.NewWriter(w, 0, 0, 1, ' ', tabwriter.Debug)
	for _, layer := range allLayers {
		fmt.Fprintln(tabw, layer.String())
	}
	tabw.Flush()
}
