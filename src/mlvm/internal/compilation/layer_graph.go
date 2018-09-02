package compilation

import (
	"io"

	"mlvm/modules/layers"
)

type Options struct {
	// Prints all debugging information for layers.
	LayerInfoWriter     io.Writer
	LayerDotGraphWriter io.Writer
}

type LayerGraph struct {
	// Sets by constructor.
	Options *Options
	Outputs []layers.Layer

	// Internal state.
	dag *LayerDAG
}

func (g *LayerGraph) Compile() error {
	// TODO: Check has not compiled yet.
	// TODO(xiejw): Verify layer names are different.

	dag := new(LayerDAG)
	if err := dag.Build(g.Outputs); err != nil {
		return err
	}

	// Debugging information.
	if g.Options.LayerInfoWriter != nil {
		PrintLayersDebuggingInfo(g.Options.LayerInfoWriter, dag)
	}

	if g.Options.LayerDotGraphWriter != nil {
		PrintLayersDotGraph(g.Options.LayerDotGraphWriter, dag)
	}

	// Tracs back to inputs
	// Color outputs and inputs
	// Save to graph
	return nil
}
