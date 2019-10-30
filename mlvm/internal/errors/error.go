package errors

import (
	"bytes"
	"fmt"
)

// Wrapper of fmt.Errorf.
func Errorf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

// Wrapper of fmt.Errorf with a root cause `original`.
func ErrorfW(original error, msg string, args ...interface{}) error {
	ierr := &internalError{
		Message:  fmt.Sprintf(msg, args...),
		Original: original,
	}
	return ierr
}

type internalError struct {
	Message  string
	Original error
}

// Format the nested error.
//
// Example output:
//
//     Unexpected error during creating instruction for opAdd_001
//       \-> Fail to infer result shapes for Op kind `opAdd`
//         \-> Expected two operands, got: 1
func (ierr *internalError) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString(ierr.Message)

	var err error
	err = ierr.Original
	indent := 0
	for {
		if err == nil {
			break
		}

		// Formats the indent.
		buf.WriteString("\n  ")
		for i := 0; i < indent; i++ {
			buf.WriteString("  ")
		}
		buf.WriteString("\\-> ")
		indent++

		if original, ok := err.(*internalError); ok {
			buf.WriteString(original.Message)
			err = original.Original
			continue
		} else {
			buf.WriteString(err.Error())
			break
		}
	}

	return buf.String()
}
