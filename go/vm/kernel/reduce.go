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
	return nil, nil
}
