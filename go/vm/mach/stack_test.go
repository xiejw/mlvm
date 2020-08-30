package mach

import (
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func checkNotErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}
}

func checkStringValue(t *testing.T, err error, got object.Object, expected string) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(*object.String).Value != expected {
		t.Fatalf("Mismatch value.")
	}
}

func checkIntValue(t *testing.T, err error, got object.Object, expected int64) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(*object.Integer).Value != expected {
		t.Fatalf("Mismatch value.")
	}
}

func TestPopAndPush(t *testing.T) {
	stack := NewStack()

	err := stack.Push(&object.Integer{1})
	checkNotErr(t, err)

	err = stack.Push(&object.String{"hello"})
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
