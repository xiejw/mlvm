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

func (h *Handle) VM() *VM {
	return h.vm
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

	handle, err := h.vm.NewHandle(h.tensor.DType, h.tensor.Shape.Dims)
	if err != nil {
		panic("failed to alloc space for grad.")
	}

	h.requireGrad = true
	h.gradHandle = handle
}

func (h *Handle) ZerosGrad() {
	// TODO fill . check requireGrad
	return
}

func (h *Handle) Zeros() {
	// TODO fill .
}

func (h *Handle) Grad() *object.Tensor {
	if !h.requireGrad {
		panic("grad does not exist.")
	}
	h.vm.WaitBarrier()
	return h.gradHandle.tensor
}
