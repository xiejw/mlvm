package code

import "testing"

func TestProgram(t *testing.T) {
	p := NewProgram()

	if p.Instructions == nil {
		t.Fatalf("Instructions should be initialized.")
	}
}
