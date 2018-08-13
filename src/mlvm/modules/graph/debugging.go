package graph

import (
	"fmt"
	"os"
	"text/tabwriter"

	"mlvm/modules/layers"
)

func printLayersDebuggingInfo(allLayers []layers.Layer) {
	fmt.Println("## Layers:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	for _, layer := range allLayers {
		fmt.Fprintln(w, layer.String())
	}
	w.Flush()
}
