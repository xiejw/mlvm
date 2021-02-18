package mach

import (
	"github.com/xiejw/mlvm/vm/ops"
)

type Handle struct {
}

type Record struct {
	Op       ops.OpCode
	Operands []*Handle
	Outputs  []*Handle
	Option   ops.Option
}

type Tape struct {
	Records []*Record
}
