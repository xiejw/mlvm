package ir

import (
	"fmt"
)

var (
	errForFuncFreeze      = fmt.Errorf("Outputs have been set for current Func already.")
	errForNameUsedAlready = "Name (\"%v\") has been used in Func already. Only allow once."
	errForNameUniqueness  = "In Func, tensor/instruction name must be unique. " +
		"\"%v\" has been used already."
)

// Register a `name` into Func with instance as `item`.
//
// If `registerOnce` is true, `name` must never be seen before. Otherwise, it is expected to be the
// same instance.
func (f *Func) registerName(name string, item interface{}, registerOnce bool) error {
	existedItem, existed := f.nameStore[name]
	if !existed {
		f.nameStore[name] = item
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

func (f *Func) mustNotFreezed() error {
	if f.freezed {
		return errForFuncFreeze
	}
	return nil
}

func (f *Func) freeze() error {
	err := f.mustNotFreezed()
	if err != nil {
		return err
	}
	f.freezed = true
	return nil
}
