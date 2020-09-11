package mach

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/object"
)

const defaultMemoryInitSize = 256

var (
	errRetrieveFromInvalidSlot = "retrieves an item from in valid slot at: %d (current memory size: %d)."
)

// An abstraction of memory to store and load item from slots. If the size is not enough, internally
// the memory size will be enlarged automatically.
//
// Internally, we could use a slice or use list of pages.
type Memory struct {
	slots []object.Object
	size  int
}

func NewMemory() *Memory {
	memory := &Memory{
		slots: make([]object.Object, defaultMemoryInitSize),
		size:  defaultMemoryInitSize,
	}

	return memory
}

func (m *Memory) Get(index int) (object.Object, *errors.DError) {
	if index >= m.size {
		return nil, errors.New(errRetrieveFromInvalidSlot, index, m.size)
	}

	item := m.slots[index]
	if item == nil {
		return nil, errors.New("the memory slot at %v is empty.", index)
	}
	return item, nil
}

// Deletes the item in memory and returns it.
func (m *Memory) Drop(index int) (object.Object, *errors.DError) {
	if index >= m.size {
		return nil, errors.New(errRetrieveFromInvalidSlot, index, m.size)
	}
	return m.slots[index], nil
}

func (m *Memory) Set(index int, item object.Object) *errors.DError {
	if index >= m.size {
		panic("Index is too large. Enlarging is planning.")
	}
	m.slots[index] = item
	return nil
}
