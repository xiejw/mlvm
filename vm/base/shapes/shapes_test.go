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

func assertTrue(t *testing.T, r bool) {
	t.Helper()
	if !r {
		t.Errorf("unexpected result.")
	}
}
