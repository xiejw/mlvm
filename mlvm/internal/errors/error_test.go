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
