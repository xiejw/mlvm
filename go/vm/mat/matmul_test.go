package mat

import (
	"testing"

	"github.com/xiejw/mlvm/go/vm/tensorarray"
)

func TestMatMul(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0, 2.0})
	rhs := tensorarray.FromRaw([]int{2, 3}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})

	o, err := Matmul(lhs, rhs, MatMulNoTrans)
	assertNoErr(t, err)
	assertShape(t, []int{3, 3}, o.Dims)
	assertAllClose(t, []float32{11., 16., 21., 11., 16., 21., 11., 16., 21.}, o.Value, 1e-6)
}
