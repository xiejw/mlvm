package errors

import (
	"testing"
)

func TestErrorf(t *testing.T) {
	err := Errorf("a %v", "b")
	if err.Error() != "a b" {
		t.Errorf("Unexpected error, got %v", err)
	}
}

func TestErrorfW(t *testing.T) {
	err1 := Errorf("a")
	err2 := ErrorfW(err1, "b")
	err3 := ErrorfW(err2, "c")
	if err3.Error() != `c
  \-> b
    \-> a` {
		t.Errorf("Unexpected error, got %v", err3)
	}
}
