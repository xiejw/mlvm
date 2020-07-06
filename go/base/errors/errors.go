package errors

import (
	"bytes"
	"fmt"
)

type DiagnosisError struct {
	notes     []string // Reverse order
	rootCause error
}

func NewDiagnosisError(sfmt string, args ...interface{}) *DiagnosisError {
	return &DiagnosisError{
		rootCause: fmt.Errorf(sfmt, args...),
	}
}

func EmitDiagnosisNote(err error, sfmt string, args ...interface{}) *DiagnosisError {
	de, ok := err.(*DiagnosisError)

	if !ok {
		de = &DiagnosisError{
			rootCause: err,
		}
	}

	note := fmt.Sprintf(sfmt, args...)
	de.notes = append(de.notes, note)
	return de
}

// Conform
func (de *DiagnosisError) Error() string {
	var buf bytes.Buffer

	fmt.Fprint(&buf, "Diagnosis Error\n")

	indentLevel := "  "
	for index := len(de.notes) - 1; index >= 0; index-- {
		fmt.Fprintf(&buf, "%v+-+ %v\n", indentLevel, de.notes[index])
		indentLevel+= "  "
	}

	fmt.Fprintf(&buf, "%v+-> %v\n", indentLevel, de.rootCause)
	return buf.String()
}
