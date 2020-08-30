package vm

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/object"
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

func (vm *VM) popTensor() (*object.Tensor, *errors.DError) {
	tensorObject, err := vm.stack.Pop()
	if err != nil {
		return nil, err.EmitNote("failed to pop tensor from stack.")
	}
	tensor, ok := tensorObject.(*object.Tensor)
	if !ok {
		return nil, errors.New("failed to pop tensor from stack: wrong type.")
	}
	return tensor, nil
}
