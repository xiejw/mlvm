package errors

import (
	"bytes"
	"fmt"
)

// A special diagnosis error structure, which allows downstream call site to emit more notes after
// creation.
type DiagnosisError struct {
	notes     []string // Reverse order
	rootCause error
}

// Creates a DiagnosisError with root cause specified by the message.
func NewDiagnosisError(sfmt string, args ...interface{}) *DiagnosisError {
	return &DiagnosisError{
		rootCause: fmt.Errorf(sfmt, args...),
	}
}

// Creates a DiagnosisError with root cause specified by `err` and emit a note immediately..
func EmitDiagnosisNote(err error, sfmt string, args ...interface{}) *DiagnosisError {
	de := &DiagnosisError{
		rootCause: err,
	}

	note := fmt.Sprintf(sfmt, args...)
	de.notes = append(de.notes, note)
	return de
}

// Formats the error into string.
func (de *DiagnosisError) String() string {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "\nDiagnosis Error\n")

	indentLevel := "  "
	for index := len(de.notes) - 1; index >= 0; index-- {
		fmt.Fprintf(&buf, "%v+-+ %v\n", indentLevel, de.notes[index])
		indentLevel += "  "
	}

	fmt.Fprintf(&buf, "%v+-> %v\n", indentLevel, de.rootCause)
	return buf.String()
}

// Emit one more note to the DiagnosisError.
func (de *DiagnosisError) EmitDiagnosisNote(sfmt string, args ...interface{}) *DiagnosisError {
	note := fmt.Sprintf(sfmt, args...)
	de.notes = append(de.notes, note)
	return de
}
