package layers

import (
	"fmt"
	"strings"
)

func FormatPrintString(layerType string, layer Layer) string {
	// Major information
	layerInfo := fmt.Sprintf(" Type: %v\t Shape: %v\t Name: %v",
		layerType, layer.Output().Shape(), layer.Name())

	// Optional inputs for the layer.
	var inputsInfo string
	if inputs := layer.Inputs(); inputs != nil {
		inputNames := make([]string, 0, inputs.Count())
		for inputLayer := range inputs.Iterator() {
			inputNames = append(inputNames, inputLayer.Name())
		}
		inputsInfo = fmt.Sprintf("\t Inputs: %v",
			strings.Join(inputNames, ", "))
	}

	return fmt.Sprintf("%v%v", layerInfo, inputsInfo)
}
