package ast

import (
	_ "fmt"

	_ "github.com/xiejw/mlvm/mlvm/array"
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
	freezed bool // If true, the module cannot be modified anymore.

	// Internal fields to store instructions, outputs, updates.
	// instructions []*ast.Instruction
	// outputs      []*ast.Result
	// updates      []*VariableUpdate

	// Internal fields to store the name of objects.
	nameStore map[string]interface{}
	// instructionMap  map[string]*ast.Instruction
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

func (m *Module) mustNotFreezed() {
	if m.freezed {
		panic(errorForModuleFreeze)
	}
}
