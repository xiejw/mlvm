package object

import (
	"bytes"
	"fmt"
	"io"
)

const defaultMaxNumberToPrintForArray = 9

type Array struct {
	Value []float32
}

func (array *Array) Type() ObjectType {
	return ArrayType
}

// Formats as `Array([  1.000,  2.000])`
func (array *Array) String() string {
	var buf bytes.Buffer
	buf.WriteString("Array(")
	array.toHumanReadableString(&buf, defaultMaxNumberToPrintForArray)
	buf.WriteString(")")
	return buf.String()
}

// Formats as `[  1.000,  2.000]`
func (array *Array) toHumanReadableString(w io.Writer, maxElementCountToPrint int) {
	size := len(array.Value)

	fmt.Fprintf(w, "[ ")
	for i, v := range array.Value {
		fmt.Fprintf(w, "%6.3f", v)

		if i < size-1 {
			fmt.Fprintf(w, ", ")
		}

		if i != size-1 && i >= maxElementCountToPrint {
			fmt.Fprintf(w, "... ")
			break
		}
	}
	fmt.Fprintf(w, "]")
}
