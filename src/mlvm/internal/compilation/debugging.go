package compilation

import (
	"fmt"
	"io"
	"text/tabwriter"

	"mlvm/modules/layers"
)

func printLayersDebuggingInfo(writer io.Writer, allLayers []layers.Layer) {
	fmt.Println("## Layers:")
	w := tabwriter.NewWriter(writer, 0, 0, 1, ' ', tabwriter.Debug)
	for _, layer := range allLayers {
		fmt.Fprintln(w, layer.String())
	}
	w.Flush()
}
