package errors

import (
	"bytes"
	"fmt"
)

type internalError struct {
	Message  string
	Original error
}

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

// Wrapper of fmt.Errorf.
func Errorf(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

// Wrapper of fmt.Errorf with a root cause.
func ErrorfW(original error, msg string, args ...interface{}) error {
	ierr := &internalError{
		Message:  fmt.Sprintf(msg, args...),
		Original: original,
	}
	return ierr
}
