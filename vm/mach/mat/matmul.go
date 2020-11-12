package mat

import (
	"log"

	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/mach/tensorarray"
)

type MatmulTransType int

const (
	MatMulNoTrans MatmulTransType = iota
)

func Matmul(lhs, rhs *tensorarray.TensorArray, trans_type MatmulTransType,
) (*tensorarray.TensorArray, *errors.DError) {
	/////////////////////////////////////////////////////////////////////////////
	// Error checking.
	/////////////////////////////////////////////////////////////////////////////
	if lhs.Rank != 2 {
		return nil, errors.New("matmul requires rank 2 lhs operand; got: %v", lhs.Rank)
	}
	if rhs.Rank != 2 {
		return nil, errors.New("matmul requires rank 2 rhs operand; got: %v", rhs.Rank)
	}
	switch trans_type {
	case MatMulNoTrans:
		if lhs.Dims[1] != rhs.Dims[0] {
			return nil, errors.New(
				"operand shape incompatable: lhs 1-th dim %v, rhs 0-th dim %v", lhs.Dims[1], rhs.Dims[0])
		}
	default:
		return nil, errors.New("trans type %v is not supported for matmul.", trans_type)
	}
	if lhs.IsCompressed() {
		log.Printf("matmul lhs compressed. try to optimize")
	}
	if rhs.IsCompressed() {
		log.Printf("matmul rhs compressed. try to optimize")
	}

	/////////////////////////////////////////////////////////////////////////////
	// Dummy Implementation.
	/////////////////////////////////////////////////////////////////////////////
	lhs = lhs.ToFullArray()
	rhs = rhs.ToFullArray()

	lhs_dim_i := lhs.Dims[0]
	lhs_dim_j := lhs.Dims[1]
	rhs_dim_k := rhs.Dims[1]

	size := lhs_dim_i * rhs_dim_k
	v := make([]float32, size)
	buf1 := lhs.Value
	buf2 := rhs.Value

	for i := 0; i < lhs_dim_i; i++ {
		for k := 0; k < rhs_dim_k; k++ {
			v_index := i*rhs_dim_k + k
			for j := 0; j < lhs_dim_j; j++ {
				v[v_index] += buf1[i*lhs_dim_j+j] * buf2[j*rhs_dim_k+k]
			}
		}
	}

	return &tensorarray.TensorArray{
		Dims:     []int{lhs_dim_i, rhs_dim_k},
		Rank:     2,
		Size:     size,
		RealSize: size,
		Value:    v,
	}, nil
}
