package vm

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
	"github.com/xiejw/mlvm/vm/vm/tensorarray"
)

func (vm *VM) pop() (object.Object, *errors.DError) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop object from stack.")
	}
	return o, nil
}

func (vm *VM) popInteger() (*object.Integer, *errors.DError) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop integer from stack.")
	}
	v, ok := o.(*object.Integer)
	if !ok {
		return nil, errors.New("failed to pop integer from stack: wrong type.")
	}
	return v, nil
}

func (vm *VM) popString() (*object.String, *errors.DError) {
	o, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop string from stack.")
	}
	v, ok := o.(*object.String)
	if !ok {
		return nil, errors.New("failed to pop string from stack: wrong type.")
	}
	return v, nil
}

func (vm *VM) popArray() (*object.Array, *errors.DError) {
	arrayObject, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop array from stack.")
	}
	array, ok := arrayObject.(*object.Array)
	if !ok {
		return nil, errors.New("failed to pop array from stack: wrong type.")
	}
	return array, nil
}

func (vm *VM) popShape() (*object.Shape, *errors.DError) {
	shapeObject, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop shape from stack.")
	}
	shape, ok := shapeObject.(*object.Shape)
	if !ok {
		return nil, errors.New("failed to pop shape from stack: wrong type.")
	}
	return shape, nil
}

func (vm *VM) popTensor() (*tensorarray.TensorArray, *errors.DError) {
	tensorObject, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop tensor from stack.")
	}
	ta, ok := tensorObject.(*tensorarray.TensorArray)
	if !ok {
		return nil, errors.New("failed to pop tensor (array) from stack: wrong type.")
	}
	return ta, nil
}
