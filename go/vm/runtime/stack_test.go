package runtime

import (
	"testing"
)

func checkNotErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}
}

func checkStringValue(t *testing.T, err error, got interface{}, expected string) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(string) != expected {
		t.Fatalf("Mismatch value.")
	}
}

func checkIntValue(t *testing.T, err error, got interface{}, expected int) {
	if err != nil {
		t.Fatalf("Unexpected err: %v", err)
	}

	if got.(int) != expected {
		t.Fatalf("Mismatch value.")
	}
}

func TestPopAndPush(t *testing.T) {
	stack := NewStack()

	err := stack.Push(1)
	checkNotErr(t, err)

	err = stack.Push("hello")
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
