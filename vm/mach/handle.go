package mach

import (
	"github.com/xiejw/mlvm/vm/object"
)

type Handle struct {
	id          int
	tensor      *object.Tensor
	vm          *VM
	requireGrad bool
	gradHandle  *Handle
}

func (h *Handle) RequireGrad() {
	if h.requireGrad {
		return
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
