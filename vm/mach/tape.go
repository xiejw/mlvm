package mach

import (
	"github.com/xiejw/mlvm/vm/ops"
)

type Record struct {
	Op       ops.OpCode
	Operands []*Handle
	Outputs  []*Handle
	Option   ops.Option
	FLowGrad bool
}

type Tape struct {
	Records []*Record
}
