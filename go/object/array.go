package object

import (
	"bytes"
	"fmt"
)

const defaultMaxNumberToPrintForArray = 9

type Array struct {
	Value []float32
}

func (array *Array) Type() ObjectType {
	return ArrayType
}

// Prints formatted array as `[  1.000,  2.000]`
func (array *Array) String() string {
	return array.DebugString(defaultMaxNumberToPrintForArray)
}

func (array *Array) DebugString(maxElementCountToPrint int) string {
	var buf bytes.Buffer

	size := len(array.Value)

	fmt.Fprintf(&buf, "[ ")
	for i, v := range array.Value {
		fmt.Fprintf(&buf, "%6.3f", v)

		if i < size-1 {
			fmt.Fprintf(&buf, ", ")
		}

		if i != size-1 && i >= maxElementCountToPrint {
			fmt.Fprintf(&buf, "... ")
			break
		}
	}
	fmt.Fprintf(&buf, "]")

	return buf.String()
}
