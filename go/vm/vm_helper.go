package vm

import (
	"bytes"
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func (vm *VM) canonicalError(op code.Opcode, format string, args ...interface{}) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "program error: Opcode: %v: ", op)
	fmt.Fprintf(&buf, format, args...)
	return fmt.Errorf(buf.String())
}

func (vm *VM) pop(op code.Opcode) (object.Object, error) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, vm.canonicalError(op, "failed to pop object from stack: %v", err)
	}
	return o, nil
}

func (vm *VM) popArray(op code.Opcode) (*object.Array, error) {
	arrayObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: %v", op, err)
	}
	array, ok := arrayObject.(*object.Array)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop array from stack: wrong type", op)
	}
	return array, nil
}

func (vm *VM) popShape(op code.Opcode) (*object.Shape, error) {
	shapeObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: %v", op, err)
	}
	shape, ok := shapeObject.(*object.Shape)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop shape from stack: wrong type", op)
	}
	return shape, nil
}

func (vm *VM) popTensor(op code.Opcode) (*object.Tensor, error) {
	tensorObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop tensor from stack: %v", op, err)
	}
	tensor, ok := tensorObject.(*object.Tensor)
	if !ok {
		return nil, fmt.Errorf("program error: Opcode: %v: failed to pop tensor from stack: wrong type", op)
	}
	return tensor, nil
}
