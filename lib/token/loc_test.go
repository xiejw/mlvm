package token

import (
	"testing"
)

func TestLocEqualness(t *testing.T) {
	loc1 := Loc{L: 120, C: 23}
	loc2 := Loc{L: 120, C: 23}
	if loc1 != loc2 {
		t.Errorf("Expected equal loc.")
	}
}
