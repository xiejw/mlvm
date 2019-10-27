package ir

import (
	"bytes"
	"fmt"
)

// String representation of Tensor.
//
// - Constant{"a", <2, 1>}
func (t *Tensor) String() string {
	buf := new(bytes.Buffer)

	switch t.kind {
	case KConstant:
		buf.WriteString("Constant{")
	case KResult:
		buf.WriteString("Result{")
	default:
		panic(fmt.Sprintf("Tensor kind %v is not expected.", t.kind))
	}

	buf.WriteString(fmt.Sprintf("\"%v\", %v}", t.Name(), t.Shape()))
	return buf.String()
}
