package vm

import (
	"testing"

	"github.com/xiejw/mlvm/go/code"
)

func checkMemoryNilValue(t *testing.T, memory *Memory, index int) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v != nil {
		t.Errorf("Value mismatch. Expect nil.")
	}
}

func checkMemoryIntValue(t *testing.T, memory *Memory, index int, expected int64) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(*code.Integer).Value != expected {
		t.Errorf("Value mismatch.")
	}
}

func checkMemoryStringValue(t *testing.T, memory *Memory, index int, expected string) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(*code.String).Value != expected {
		t.Errorf("Value mismatch.")
	}
}

func TestGetAndSet(t *testing.T) {
	memory := NewMemory()

	memory.Set(10, &code.Integer{123})
	memory.Set(11, &code.String{"hello"})

	checkMemoryNilValue(t, memory, 0)
	checkMemoryIntValue(t, memory, 10, 123)
	checkMemoryStringValue(t, memory, 11, "hello")
}
