package runtime

const defaultMemoryInitSize = 256

type Memory struct {
	items []interface{}
	size  int
}

func NewMemory() *Memory {
	memory := &Memory{
		items: make([]interface{}, defaultMemoryInitSize),
		size:  defaultMemoryInitSize,
	}

	return memory
}

func (m *Memory) Get(index int) interface{} {
	return m.items[index]
}
