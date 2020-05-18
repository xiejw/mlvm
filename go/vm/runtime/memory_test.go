package runtime

import (
	"testing"
)

func checkNilValue(t *testing.T, memory *Memory, index int) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v != nil {
		t.Errorf("Value mismatch. Expect nil.")
	}
}

func checkIntValue(t *testing.T, memory *Memory, index int, expected int) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(int) != expected {
		t.Errorf("Value mismatch.")
	}
}

func checkStringValue(t *testing.T, memory *Memory, index int, expected string) {
	v, err := memory.Get(index)
	if err != nil {
		t.Errorf("Should not fail.")
	}
	if v.(string) != expected {
		t.Errorf("Value mismatch.")
	}
}

func TestGetAndSet(t *testing.T) {
	memory := NewMemory()

	memory.Set(10, 123)
	memory.Set(11, "hello")

	checkNilValue(t, memory, 0)
	checkIntValue(t, memory, 10, 123)
	checkStringValue(t, memory, 11, "hello")
}
