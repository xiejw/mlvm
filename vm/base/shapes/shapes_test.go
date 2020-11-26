package shapes

import (
	"testing"
)

func TestBroadcastableForIdentitcalShapes(t *testing.T) {
	assertTrue(t, IsBroadcastable([]int{2, 3}, []int{2, 3}))
	assertTrue(t, IsBroadcastable([]int{2, 1}, []int{2, 1}))
	assertTrue(t, IsBroadcastable([]int{1, 3}, []int{1, 3}))
	assertTrue(t, IsBroadcastable([]int{1, 1}, []int{1, 1}))
	assertTrue(t, IsBroadcastable([]int{1}, []int{1}))
}

func TestBroadcastableForShorterDestRank(t *testing.T) {
	assertFalse(t, IsBroadcastable([]int{2, 3}, []int{2}))
	assertFalse(t, IsBroadcastable([]int{2, 1}, []int{2}))
	assertFalse(t, IsBroadcastable([]int{1, 3}, []int{1}))
	assertFalse(t, IsBroadcastable([]int{1, 1}, []int{1}))
}

func TestBroadcastableIncompatibleDest(t *testing.T) {
	assertFalse(t, IsBroadcastable([]int{2, 3}, []int{2, 4}))
	assertFalse(t, IsBroadcastable([]int{2, 1}, []int{2, 2}))
	assertFalse(t, IsBroadcastable([]int{1, 3}, []int{2, 1, 1}))
	assertFalse(t, IsBroadcastable([]int{1, 2}, []int{1, 1}))
}

func TestBroadcastableCompatibleDest(t *testing.T) {
	assertTrue(t, IsBroadcastable([]int{2, 3}, []int{2, 3}))
	assertTrue(t, IsBroadcastable([]int{2, 3}, []int{2, 2, 3}))
	assertTrue(t, IsBroadcastable([]int{2, 1}, []int{2, 1}))
	assertTrue(t, IsBroadcastable([]int{2, 1}, []int{2, 2, 1}))
	assertTrue(t, IsBroadcastable([]int{2, 1}, []int{1, 2, 1}))
	assertTrue(t, IsBroadcastable([]int{1, 3}, []int{2, 1, 3}))
	assertTrue(t, IsBroadcastable([]int{1, 3}, []int{2, 2, 3}))
	assertTrue(t, IsBroadcastable([]int{1}, []int{1, 1}))
	assertTrue(t, IsBroadcastable([]int{1}, []int{2, 2}))
	assertTrue(t, IsBroadcastable([]int{1}, []int{2, 1}))
}

// -----------------------------------------------------------------------------
// Helper methods.
// -----------------------------------------------------------------------------

func assertTrue(t *testing.T, r bool) {
	t.Helper()
	if !r {
		t.Errorf("unexpected result.")
	}
}

func assertFalse(t *testing.T, r bool) {
	t.Helper()
	if r {
		t.Errorf("unexpected result.")
	}
}
