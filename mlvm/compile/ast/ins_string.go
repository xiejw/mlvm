package ast

import (
	"bytes"
	"fmt"
)

// String representation for Instruction
//
// Ins{"op_add_001", (Constant{"a", <2, 1>}, Constant{"b", <2, 1>}) -> ()}.
func (ins *Instruction) String() string {
	buf := new(bytes.Buffer)

	buf.WriteString(fmt.Sprintf(
		"Ins{\"%v\", (", ins.Name()))

	operands := ins.Operands()
	for i, t := range operands {
		buf.WriteString(fmt.Sprintf("%v", t))
		if i != len(operands)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString(") -> ()")
	buf.WriteString("}")

	return buf.String()
}
