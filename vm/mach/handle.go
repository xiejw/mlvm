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
	// alloc and add point
	h.requireGrad = true
}
