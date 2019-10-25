package ir

type Module struct {
	// Internal fields to store instructions, outputs, updates.
	freezed      bool                   // If true, the module cannot be modified anymore.
	opNameIndex  int                    // The index to generate default name for Op.
	nameStore    map[string]interface{} // Name to object mapping.
	instructions []*Instruction         // Ordered Instructions
}

func NewModule() *Module {
	m := &Module{
		nameStore: make(map[string]interface{}),
	}
	return m
}

func (m *Module) Instructions() []*Instruction {
	return m.instructions
}
