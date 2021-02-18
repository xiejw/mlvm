package mach

import (
	"github.com/xiejw/mlvm/vm/ops"
)

type VM struct {
	Tape Tape
}

func (vm *VM) ExecOp(op ops.OpCode, operands []*Handle, opt ops.Option) (*Handle, error) {
	return nil, nil
}
