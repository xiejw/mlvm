package mach

import (
	"strings"
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func TestGetAndSet(t *testing.T) {
	memory := NewMemory()

	memory.Set(10, &object.Integer{123})
	memory.Set(11, &object.String{"hello"})

	checkMemoryNilValue(t, memory, 0)
	checkMemoryIntValue(t, memory, 10, 123)
	checkMemoryStringValue(t, memory, 11, "hello")
}

func TestDrop(t *testing.T) {
	memory := NewMemory()

	memory.Set(10, &object.Integer{123})

	checkMemoryNilValueViaDrop(t, memory, 0)
	checkMemoryIntValue(t, memory, 10, 123)

	v, err := memory.Drop(10)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(*object.Integer).Value != 123 {
		t.Errorf("Value mismatch.")
	}

	checkMemoryNilValueViaDrop(t, memory, 10)
}

func checkMemoryNilValue(t *testing.T, memory *Memory, index int) {
	t.Helper()
	_, err := memory.Get(index)
	if err == nil {
		t.Errorf("Should fail due to empty slot.")
	}
	if !strings.Contains(err.String(), "empty") {
		t.Errorf("should see empty error.")
	}
}

func checkMemoryNilValueViaDrop(t *testing.T, memory *Memory, index int) {
	t.Helper()
	_, err := memory.Drop(index)
	if err == nil {
		t.Errorf("Should fail due to empty slot.")
	}
	if !strings.Contains(err.String(), "empty") {
		t.Errorf("should see empty error.")
	}
}

func checkMemoryIntValue(t *testing.T, memory *Memory, index int, expected int64) {
	t.Helper()
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(*object.Integer).Value != expected {
		t.Errorf("Value mismatch.")
	}
}

func checkMemoryStringValue(t *testing.T, memory *Memory, index int, expected string) {
	t.Helper()
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(*object.String).Value != expected {
		t.Errorf("Value mismatch.")
	}
}

func assertIntEq(t *testing.T, expected, got int) {
	t.Helper()
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}
