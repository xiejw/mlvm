package ir

import (
	"bytes"
	"fmt"
)

func (m *Module) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("Module {\n")

	instructions := m.Instructions()

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
