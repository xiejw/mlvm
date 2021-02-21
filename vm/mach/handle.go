package mach

import (
	"github.com/xiejw/mlvm/vm/object"
)

type Handle struct {
	id          int
	tensor      *object.Tensor
	vm          *VM
	requireGrad bool // this handles needs grad.
	flowGrad    bool // flows grad to the op producing the op.
	gradHandle  *Handle
	record      *Record // record produces this handle.
}

func (h *Handle) RequireGrad() {
	if h.requireGrad {
		return
	}

	if h.record != nil {
		panic("produced result cannot have grad.")
	}

	if h.flowGrad {
		panic("handle with flowGrad cannot have grad.")
	}

	// TODO alloc and add point
	h.requireGrad = true
}

func (h *Handle) ZerosGrad() {
	// TODO fill . check requireGrad
	return
}

func (h *Handle) Zeros() {
	// TODO fill .
}

func (h *Handle) Grad() *object.Tensor {
	return nil
}
