package code

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/go/object"
)

type Constants []object.Object

func (cs Constants) String() string {
	var buf bytes.Buffer
	cs.ToHumanReadableString(&buf)
	return buf.String()
}

func (cs Constants) ToHumanReadableString(w io.Writer) {
	if len(cs) == 0 {
		fmt.Fprintf(w, "(empty)")
		return
	}

	fmt.Fprint(w, "[\n")
	for i, c := range cs {
		fmt.Fprintf(w, "  %3d: %v\n", i, c)
	}
	fmt.Fprint(w, "]")
}
