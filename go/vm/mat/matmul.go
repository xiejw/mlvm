package mat

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

type MatmulTransType int

const (
	MatMulNoTrans MatmulTransType = iota
)

func Matmul(lhs, rhs *tensorarray.TensorArray, trans_type MatmulTransType,
) (*tensorarray.TensorArray, *errors.DError) {
	return nil, nil
}
