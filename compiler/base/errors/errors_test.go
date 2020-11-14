package errors

import (
	"testing"
)

func TestDiagnosisErrorOutput(t *testing.T) {
	err := New("root 1")
	err.EmitNote("during stack 1")
	err.EmitNote("during stack 2")

	expected := `Diagnosis Error:
+-+ during stack 2
  +-+ during stack 1
    +-> root 1
`

	got := err.String()

	if got != expected {
		t.Errorf("Error() mismatch. expected: %v, got: %v", expected, got)
	}

	if err == nil {
		t.Errorf("error should be nil.")
	}
}
