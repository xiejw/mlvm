package kernel

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

type MergeType int

const (
	MergeSum MergeType = iota
)

func Reduce(ta *tensorarray.TensorArray, merge_type MergeType) (*tensorarray.TensorArray, *errors.DError) {
	if merge_type != MergeSum {
		return nil, errors.New("merge type for reduce is not supported: %v", merge_type)
	}

	// indices := make([]int, ta.Rank)
	strides = ta.Strides
	dims := ta.Dims
	dim_ptr := ta.Rank -1
	i := 0
	buf := ta.Value

	var v float32 = 0.0

	for {
		v += buf[i]
		i += strides[dim_ptr]

		if i == dims[dim_ptr] {
		}
		break
	}

	return nil, nil
}
