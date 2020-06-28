package vm

import (
	"bytes"
	"fmt"

	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
)

func (vm *VM) canonicalError(op code.Opcode, format string, args ...interface{}) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "program error: current opcode: %v: ", op)
	fmt.Fprintf(&buf, format, args...)
	return fmt.Errorf(buf.String())
}

func (vm *VM) pop() (object.Object, error) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop object from stack: %v.", err)
	}
	return o, nil
}

func (vm *VM) popInteger() (*object.Integer, error) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop integer from stack: %v.", err)
	}
	v, ok := o.(*object.Integer)
	if !ok {
		return nil, fmt.Errorf("failed to pop integer from stack: wrong type.")
	}
	return v, nil
}

func (vm *VM) popString() (*object.String, error) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop string from stack: %v.", err)
	}
	v, ok := o.(*object.String)
	if !ok {
		return nil, fmt.Errorf("failed to pop string from stack: wrong type.")
	}
	return v, nil
}

func (vm *VM) popArray() (*object.Array, error) {
	arrayObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop array from stack: %v.", err)
	}
	array, ok := arrayObject.(*object.Array)
	if !ok {
		return nil, fmt.Errorf("failed to pop array from stack: wrong type.")
	}
	return array, nil
}

func (vm *VM) popShape() (*object.Shape, error) {
	shapeObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop shape from stack: %v.", err)
	}
	shape, ok := shapeObject.(*object.Shape)
	if !ok {
		return nil, fmt.Errorf("failed to pop shape from stack: wrong type.")
	}
	return shape, nil
}

func (vm *VM) popTensor() (*object.Tensor, error) {
	tensorObject, err := vm.stack.Pop()
	if err != nil {
		return nil, fmt.Errorf("failed to pop tensor from stack: %v.", err)
	}
	tensor, ok := tensorObject.(*object.Tensor)
	if !ok {
		return nil, fmt.Errorf("failed to pop tensor from stack: wrong type.")
	}
	return tensor, nil
}
