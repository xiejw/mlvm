package rng

import (
	"math"
	"testing"
)

func assertAllClose(t *testing.T, expected, got []float32, tol float64) {
	t.Helper()
	if len(expected) != len(got) {
		t.Fatalf("length mismatch. expected: %v, got: %v.", len(expected), len(got))
	}

	for i := 0; i < len(expected); i++ {
		if math.Abs(float64(expected[i]-got[i])) >= tol {
			t.Errorf("\nelement mismatch at %v: expected %v, got %v\n", i, expected[i], got[i])
		}
	}
}

func TestNormEven(t *testing.T) {
	size := 6
	value := make([]float32, size)

	prng := NewPrng64(456)
	prng.Norm(value)

	expected := []float32{1.3481823, -1.6701441, 1.4310317, 0.6320735, 0.28827125, 1.6303506}
	assertAllClose(t, expected, value, 1e-6)
}

func TestNormOdd(t *testing.T) {
	size := 5
	value := make([]float32, size)

	prng := NewPrng64(456)
	prng.Norm(value)

	expected := []float32{1.3481823, -1.6701441, 1.4310317, 0.6320735, 0.28827125}
	assertAllClose(t, expected, value, 1e-6)
}
