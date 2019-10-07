package ast

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
)

const (
	errorForModuleFreeze = "Outputs have been set for current Module already."
)

// TODO: Placeholder for paramter updates.
// type VariableUpdate struct {
// 	Variable *variable.Variable
// 	NewValue tensor.Tensor
// }

type Module struct {
	freezed     bool                   // If true, the module cannot be modified anymore.
	opNameIndex int                    // The index to generate default name for Op.
	nameStore   map[string]interface{} // Name to object mapping.

	// Internal fields to store instructions, outputs, updates.
	// instructions []*ast.Instruction
	// outputs      []*ast.Result
	// updates      []*VariableUpdate

	// Internal fields to store the name of objects.
	// instructionMap  map[string]*ast.Instruction
}

// Register a constant (array.Array) as Tensor in Module.
func (m *Module) NewConstant(arr *array.Array) *Tensor {
	m.registerName(arr.Name(), arr, true /* registerOnce */)
	return newConstantTensor(arr)
}

func NewModule() *Module {
	m := &Module{
		nameStore: make(map[string]interface{}),
		// instructionMap:  make(map[string]*ast.Instruction),
	}
	return m
}

// func (m *Module) SetOutputs(outputs []tensor.Tensor) (err error) {
// 	// if m.freezed {
// 	// 	return fmt.Errorf(errorForModuleFreeze)
// 	// }
// 	// m.freezed = true
//
// 	// if len(outputs) == 0 {
// 	// 	return fmt.Errorf("Cannot have empty outputs for Modules")
// 	// }
//
// 	// m.outputs, err = m.validateOutputsBelongToModule(outputs)
// 	// return
// }
//
// func (m *Module) Run() ([]tensor.Tensor, error) {
// 	// if m.outputs == nil {
// 	// 	return nil, fmt.Errorf("Outputs have NOT been set yet.")
// 	// }
//
// 	// inputs := mi.ModuleInputs{
// 	// 	Outputs:      m.outputs,
// 	// 	Instructions: m.instructions,
// 	// }
//
// 	// // Convert the udpates to correct type.
// 	// if m.updates != nil {
// 	// 	updates := make([]*mi.Update, 0, len(m.updates))
// 	// 	for _, u := range m.updates {
// 	// 		r, ok := u.NewValue.(*ast.Result)
// 	// 		if !ok {
// 	// 			panic("Expect a result.")
// 	// 		}
// 	// 		updates = append(updates, &mi.Update{
// 	// 			Variable: u.Variable,
// 	// 			Result:   r,
// 	// 		})
// 	// 	}
//
// 	// 	inputs.Updates = updates
// 	// }
//
// 	// engine := engine.SimpleEngine{
// 	// 	Inputs: &inputs,
// 	// }
// 	// return engine.Run()
// }
//

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
