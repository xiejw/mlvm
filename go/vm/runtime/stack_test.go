package runtime

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
)

func checkNotErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}
}

func checkStringValue(t *testing.T, err error, got code.Object, expected string) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(*code.String).Value != expected {
		t.Fatalf("Mismatch value.")
	}
}

func checkIntValue(t *testing.T, err error, got code.Object, expected int64) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(*code.Integer).Value != expected {
		t.Fatalf("Mismatch value.")
	}
}

func TestPopAndPush(t *testing.T) {
	stack := NewStack()

	err := stack.Push(&code.Integer{1})
	checkNotErr(t, err)

	err = stack.Push(&code.String{"hello"})
	checkNotErr(t, err)

	v, err := stack.Pop()
	checkStringValue(t, err, v, "hello")

	v, err = stack.Pop()
	checkIntValue(t, err, v, 1)

	_, err = stack.Pop()
	if err == nil {
		t.Fatalf("Expected error for empty stack.")
	}
}
