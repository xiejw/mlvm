package ast

import (
	"fmt"
)

const (
	errorForModuleFreeze = "Outputs have been set for current Module already."
)

// Register a `name` into Module with instance as `item`.
//
// If `registerOnce` is true, `name` must never be seen before. Otherwise, it is expected to be the
// same instance.
func (m *Module) registerName(name string, item interface{}, registerOnce bool) {
	existedItem, existed := m.nameStore[name]
	if !existed {
		m.nameStore[name] = item
		return
	}

	if registerOnce {
		panic(fmt.Sprintf("Name (\"%v\") has been used in Module already. Only allow once.", name))
	}

	if existedItem != item {
		panic(fmt.Sprintf("In Module, tensor/instruction name must be unique. "+
			"\"%v\" has been used already.", name))
	}
}

func (m *Module) mustNotFreezed() {
	if m.freezed {
		panic(errorForModuleFreeze)
	}
}
