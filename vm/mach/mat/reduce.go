package mat

import (
	"fmt"

	"github.com/xiejw/mlvm/vm/mach/tensorarray"
)

type MergeType int

const (
	MergeSum MergeType = iota
)

func Reduce(ta *tensorarray.TensorArray, merge_type MergeType) (*tensorarray.TensorArray, error) {
	if merge_type != MergeSum {
		return nil, fmt.Errorf("merge type for reduce is not supported: %v", merge_type)
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
			return nil, fmt.Errorf("unsupported merge_type: %v", merge_type)
		}
	} else {
		buf := ta.Value

		for i := 0; i < ta.Size; i++ {
			v += buf[i]
		}
	}

	return tensorarray.FromRaw([]int{1}, []float32{v}), nil
}
