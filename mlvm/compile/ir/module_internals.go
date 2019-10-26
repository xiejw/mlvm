package ir

import (
	"fmt"
)

var (
	errForModuleFreeze    = fmt.Errorf("Outputs have been set for current Module already.")
	errForNameUsedAlready = "Name (\"%v\") has been used in Module already. Only allow once."
	errForNameUniqueness  = "In Module, tensor/instruction name must be unique. " +
		"\"%v\" has been used already."
)

// Register a `name` into Module with instance as `item`.
//
// If `registerOnce` is true, `name` must never be seen before. Otherwise, it is expected to be the
// same instance.
func (m *Module) registerName(name string, item interface{}, registerOnce bool) error {
	existedItem, existed := m.nameStore[name]
	if !existed {
		m.nameStore[name] = item
		return nil
	}

	if registerOnce {
		return fmt.Errorf(errForNameUsedAlready, name)
	}

	if existedItem != item {
		return fmt.Errorf(errForNameUniqueness, name)
	}
	return nil
}

func (m *Module) mustNotFreezed() error {
	if m.freezed {
		return errForModuleFreeze
	}
	return nil
}

func (m *Module) freeze() error {
	err := m.mustNotFreezed()
	if err != nil {
		return err
	}
	m.freezed = true
	return nil
}
