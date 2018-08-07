package layers

import "fmt"

func FormatPrintString(layerType string, layer Layer) string {
	return fmt.Sprintf(" Type: %v\t Shape: %v\t Name: %v",
		layerType, layer.Output().Shape(), layer.Name())
}
