package ir

type Func struct {
	// Internal fields to store instructions, outputs, updates.
	freezed      bool                   // If true, the module cannot be modified anymore.
	opNameIndex  int                    // The index to generate default name for Op.
	nameStore    map[string]interface{} // Name to object mapping.
	instructions []*Instruction         // Ordered Instructions
}

func NewFunc() *Func {
	f := &Func{
		nameStore: make(map[string]interface{}),
	}
	return f
}

func (f *Func) Instructions() []*Instruction {
	return f.instructions
}
