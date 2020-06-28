package code

import (
	"fmt"
	"bytes"

	"github.com/xiejw/mlvm/go/object"
)

type Constants []object.Object

func (cs Constants) String() string {
	if len(cs) == 0 {
		return "(empty)"
	}

	var buf bytes.Buffer
	buf.WriteString("[\n")
	for i, c := range cs {
		fmt.Fprintf(&buf, "  %3d: %v\n", i, c)
	}
	buf.WriteString("]")
	return buf.String()
}
