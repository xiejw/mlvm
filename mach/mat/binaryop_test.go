package mat

import (
	"math"
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/vm/mach/tensorarray"
)

// ----------------------------------------------------------------------------
// Basic Operations.
// ----------------------------------------------------------------------------

func TestBinaryAdd(t *testing.T) {
	tensor := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})

	o, err := BinaryOp(tensor, tensor, BinaryAdd)
	assertNoErr(t, err)

	result := o.ToTensor()
	if result.String() != "Tensor(<2> [  2.000,  4.000])" {
		t.Errorf("value mismatch: got `%v`", result)
	}
}

func TestBinaryMinus(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})
	rhs := tensorarray.FromRaw([]int{2}, []float32{2.0, 3.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)

	assertShape(t, []int{2}, o.Dims)
	assertAllClose(t, []float32{-1.0, -1.0}, o.Value, 1e-6)
}

func TestBinaryMul(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{2}, []float32{1.0, 2.0})
	rhs := tensorarray.FromRaw([]int{2}, []float32{2.0, 3.0})

	o, err := BinaryOp(lhs, rhs, BinaryMul)
	assertNoErr(t, err)

	assertShape(t, []int{2}, o.Dims)
	assertAllClose(t, []float32{2., 6.0}, o.Value, 1e-6)
}

func TestBinaryDiv(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{2}, []float32{6.0, 9.0})
	rhs := tensorarray.FromRaw([]int{2}, []float32{2.0, 3.0})

	o, err := BinaryOp(lhs, rhs, BinaryDiv)
	assertNoErr(t, err)

	assertShape(t, []int{2}, o.Dims)
	assertAllClose(t, []float32{3., 3.0}, o.Value, 1e-6)
}

// ----------------------------------------------------------------------------
// Broadcasting.
// ----------------------------------------------------------------------------

func TestBinaryOpWithSingleElementInRHS(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{1.0, 2.0})
	rhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)
	assertShape(t, []int{3, 2}, o.Dims)
	assertAllClose(t, []float32{-2.0, -1.0}, o.Value, 1e-6)
}

func TestBinaryOpWithSingleElementInLHS(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0})
	rhs := tensorarray.FromRaw([]int{3, 2}, []float32{1.0, 2.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)
	assertShape(t, []int{3, 2}, o.Dims)
	assertAllClose(t, []float32{2.0, 1.0}, o.Value, 1e-6)
}

func TestBinaryOpWithBothBroadcastOperands(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{1.0, 2.0})
	rhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0, 2.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)
	assertShape(t, []int{3, 2}, o.Dims)
	assertAllClose(t, []float32{-2.0, 0.0}, o.Value, 1e-6)
}

func TestBinaryOpWithBroadcastOperandtInRHS(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	rhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0, 2.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)
	assertShape(t, []int{3, 2}, o.Dims)
	assertAllClose(t, []float32{-2.0, 0.0, 0.0, 2.0, 2.0, 4.0}, o.Value, 1e-6)
}

func TestBinaryOpWithBroadcastOperandtInLHS(t *testing.T) {
	lhs := tensorarray.FromRaw([]int{3, 2}, []float32{3.0, 2.0})
	rhs := tensorarray.FromRaw([]int{3, 2}, []float32{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})

	o, err := BinaryOp(lhs, rhs, BinaryMinus)
	assertNoErr(t, err)
	assertShape(t, []int{3, 2}, o.Dims)
	assertAllClose(t, []float32{2.0, 0.0, 0.0, -2.0, -2.0, -4.0}, o.Value, 1e-6)
}

// ----------------------------------------------------------------------------
// Helper Method.
// ----------------------------------------------------------------------------
func assertNoErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}

func assertShape(t *testing.T, expected, got []int) {
	t.Helper()
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("shape mismatch. expected: %v, got: %v", expected, got)
	}
}

func assertAllClose(t *testing.T, expected, got []float32, tol float64) {
	t.Helper()
	if len(expected) != len(got) {
		t.Fatalf("value length mismatch. expected: %v, got: %v.", len(expected), len(got))
	}

	for i := 0; i < len(expected); i++ {
		if math.Abs(float64(expected[i]-got[i])) >= tol {
			t.Errorf("\nelement mismatch at %v: expected %v, got %v\n", i, expected[i], got[i])
		}
	}
}
