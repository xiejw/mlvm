package ir

import (
	"testing"
)

func TestTypeStrings(t *testing.T) {
	assertStringEquals(t, (&Type{Kind: KInt}).String(), "Int")
	assertStringEquals(t, (&Type{Kind: KRng}).String(), "Rng")
	assertStringEquals(t, (&Type{Kind: KShape, Dims: []int{1, 2}}).String(), "Shape(<1, 2>)")
	assertStringEquals(t, (&Type{Kind: KArray, Dims: []int{2}}).String(), "Array(<2>)")
	assertStringEquals(t, (&Type{Kind: KTensor, Dims: []int{1, 2}}).String(), "Tensor(<1, 2>)")
}

func assertStringEquals(t *testing.T, expected, got string) {
	t.Helper()
	if expected != got {
		t.Errorf("expected: %v got: %v", expected, got)
	}
}
