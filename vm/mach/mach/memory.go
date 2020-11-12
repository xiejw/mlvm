package mach

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
)

const defaultMemoryInitSize = 256

var (
	errRetrieveFromInvalidSlot = "retrieves an item from in valid slot at: %d " +
		"(current allocated slot count: %d)."
)

// An abstraction of memory to store and load item from slots. If allocated slot count is not enough, internally
// the memory will be enlarged automatically.
//
// Internally, we could use a slice or use list of pages.
type Memory struct {
	slots         []object.Object
	slot_count    int
	size_in_bytes int
}

func NewMemory() *Memory {
	memory := &Memory{
		slots:         make([]object.Object, defaultMemoryInitSize),
		slot_count:    defaultMemoryInitSize,
		size_in_bytes: 0,
	}

	return memory
}

func (m *Memory) ByteSize() int {
	return m.size_in_bytes
}

func (m *Memory) Get(index int) (object.Object, *errors.DError) {
	if index >= m.slot_count {
		return nil, errors.New(errRetrieveFromInvalidSlot, index, m.slot_count)
	}

	item := m.slots[index]
	if item == nil {
		return nil, errors.New("the memory slot at %v is empty.", index)
	}
	return item, nil
}

// Deletes the item in memory and returns it.
func (m *Memory) Drop(index int) (object.Object, *errors.DError) {
	if index >= m.slot_count {
		return nil, errors.New(errRetrieveFromInvalidSlot, index, m.slot_count)
	}
	item := m.slots[index]
	if item == nil {
		return nil, errors.New("the memory slot at %v is empty.", index)
	}
	m.size_in_bytes -= item.MemSize()
	m.slots[index] = nil
	return item, nil
}

func (m *Memory) Set(index int, item object.Object) *errors.DError {
	if index >= m.slot_count {
		panic("Index is too large. Enlarging is planning.")
	}

	old_item := m.slots[index]
	if old_item != nil {
		m.size_in_bytes -= old_item.MemSize()
	}

	m.slots[index] = item
	m.size_in_bytes += item.MemSize()
	return nil
}
