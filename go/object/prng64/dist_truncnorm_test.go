package prng64

import (
	"testing"
)

func TestTrucNorm(t *testing.T) {
	size := 5
	value := make([]float32, size)

	// First assert norm generates the number outside the bounds, for this seed.
	prng := NewPrng64(123456789)
	prng.Norm(value)
	expected := []float32{1.1747507, 2.3442311, -1.0833434, -0.26150864, -0.08424469}
	assertAllClose(t, expected, value, 1e-6)

	// Then assert that number is skipped.
	size = 4
	value = make([]float32, size)
	prng = NewPrng64(123456789)
	prng.TruncNorm(value)
	expected = []float32{1.1747507, -1.0833434, -0.26150864, -0.08424469}
	assertAllClose(t, expected, value, 1e-6)
}
