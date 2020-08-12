package errors

import (
	"bytes"
	"fmt"
)

// A special diagnosis error structure, which allows downstream call site to emit more notes after
// creation.
type DError struct {
	notes     []string // Reverse order
	rootCause error
}

// Creates a DError with root cause specified by the message.
func New(sfmt string, args ...interface{}) *DError {
	return &DError{
		rootCause: fmt.Errorf(sfmt, args...),
	}
}

// Creates a DError with root cause specified by `err`.
func From(err error) *DError {
	return &DError{
		rootCause: err,
	}
}

// Formats the error into string.
func (de *DError) String() string {
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

// Emit one more note to the DError.
func (de *DError) EmitNote(sfmt string, args ...interface{}) *DError {
	note := fmt.Sprintf(sfmt, args...)
	de.notes = append(de.notes, note)
	return de
}
