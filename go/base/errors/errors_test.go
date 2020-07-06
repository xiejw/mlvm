package errors

import (
	"testing"
)

func TestDiagnosisErrorOutput(t *testing.T) {
	var err error
	err = NewDiagnosisError("root 1")
	err = EmitDiagnosisNote(err, "during stack 2")
	err = EmitDiagnosisNote(err, "during stack 1")

	expected := `Diagnosis Error
  +-+ during stack 1
    +-+ during stack 2
      +-> root 1
`

	got := err.Error()

	if got != expected {
		t.Errorf("Error() mismatch. expected: %v, got: %v", expected, got)
	}
}
