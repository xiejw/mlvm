package mat

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/vm/tensorarray"
)

type MergeType int

const (
	MergeSum MergeType = iota
)

func Reduce(ta *tensorarray.TensorArray, merge_type MergeType) (*tensorarray.TensorArray, *errors.DError) {
	if merge_type != MergeSum {
		return nil, errors.New("merge type for reduce is not supported: %v", merge_type)
	}

	var v float32 = 0.0
	if ta.IsCompressed() {
		buf := ta.Value
		real_size := ta.RealSize

		switch merge_type {
		case MergeSum:
			for i := 0; i < real_size; i++ {
				v += buf[i]
			}
			v *= float32(ta.Size / real_size)
		default:
			return nil, errors.New("unsupported merge_type: %v", merge_type)
		}
	} else {
		buf := ta.Value

		for i := 0; i < ta.Size; i++ {
			v += buf[i]
		}
	}

	return tensorarray.FromRaw([]int{1}, []float32{v}), nil
}
