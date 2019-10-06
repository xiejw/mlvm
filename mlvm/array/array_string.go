package array

import (
	"bytes"
	"fmt"
)

// Returns the string for debugging. See test for how it looks.
func (arr *Array) String() string {
	buf := new(bytes.Buffer)

	// First, gets how many elements in each dim. For shape [3,2,3], eleCount
	// should be [6, 3, 1].
	shape := arr.Shape()
	rank := shape.Rank()
	dims := shape.Dims()
	eleCount := make([]int, rank)
	accumulateCount := 1
	for index := rank - 1; index >= 0; index-- {
		eleCount[index] = accumulateCount
		accumulateCount *= int(dims[index])
	}

	// Next, format the data.
	if shape.Rank() > 1 {
		// Begins a new line for any rank 2+ Tensor.
		buf.WriteString("\n")
	}
	formatData(buf, 0 /* indent */, arr.data.value, arr.shape.dims, eleCount)
	if shape.Rank() > 1 {
		// Ends a new line for any rank 2+ Tensor.
		buf.WriteString("\n")
	}
	return buf.String()
}

// Recursively format the data for each trunk in the data.
func formatData(
	buf *bytes.Buffer,
	indent int,
	subdata []Float,
	subshape []Dimension,
	eleCount []int,
) {

	// Prepares the indent string for the closing bracket.
	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += "  "
	}

	// Formats the final dimension in horizantal line.
	if len(subshape) == 1 {
		buf.WriteString("[ ")
		for _, ele := range subdata {
			buf.WriteString(fmt.Sprintf("%.3f ", ele))
		}
		buf.WriteString("]")
		return
	}

	// For generate case.
	buf.WriteString("[ ")
	for x := 0; x < int(subshape[0]); x++ {
		if x > 0 {
			// For each non-first element, write indent with one more level.
			buf.WriteString(indentString + "  ")
		}
		start := x * eleCount[0]
		end := start + eleCount[0]
		formatData(buf, indent+1, subdata[start:end], subshape[1:], eleCount[1:])
		buf.WriteString("\n")
	}

	// Write indent
	for i := 0; i < indent; i++ {
		buf.WriteString(indentString)
	}
	buf.WriteString("]")
}
