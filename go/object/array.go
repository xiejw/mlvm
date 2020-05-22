package object

import (
	"bytes"
	"fmt"
)

const maxNumberToPrintForArray = 9

type Array struct {
	Value []float32
}

func (array *Array) Type() ObjectType {
	return ArrayType
}

func (array *Array) String() string {
	var buf bytes.Buffer

	size := len(array.Value)

	fmt.Fprintf(&buf, "[ ")
	for i, v := range array.Value {
		fmt.Fprintf(&buf, "%6.3f", v)

		if i < size-1 {
			fmt.Fprintf(&buf, ", ")
		}

		if i != size-1 && i >= maxNumberToPrintForArray {
			fmt.Fprintf(&buf, "... ")
			break
		}
	}
	fmt.Fprintf(&buf, "]")

	return buf.String()
}
