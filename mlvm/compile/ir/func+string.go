package ir

import (
	"bytes"
	"fmt"
)

func (f *Func) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("Func {\n")

	instructions := f.Instructions()

	if len(instructions) == 0 {
		buf.WriteString("  (empty)\n")
	} else {
		for i, ins := range instructions {
			buf.WriteString(fmt.Sprintf(
				"  %03d: %v\n", i+1, ins))
		}
	}

	buf.WriteString("}\n")
	return buf.String()
}
