package mach

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/object"
)

const defaultStackSize = 256

var errPopOnEmptyStack = errors.New("pop an item from empty stack.")

type Stack struct {
	items []object.Object
	top   int // Current top.
}

func NewStack() *Stack {
	stack := &Stack{
		items: make([]object.Object, 0, defaultStackSize),
		top:   -1,
	}
	return stack
}

func (stack *Stack) Push(item object.Object) *errors.DError {
	stack.top++
	top := stack.top

	if top >= len(stack.items) {
		// Ask the golang slice to double the space.
		stack.items = append(stack.items, item)
		return nil
	}

	stack.items[top] = item
	return nil
}

func (stack *Stack) Pop() (object.Object, *errors.DError) {
	top := stack.top
	if top < 0 {
		return nil, errPopOnEmptyStack
	}

	item := stack.items[top]
	stack.items[top] = nil // Freeze memory.
	stack.top--
	return item, nil
}

func (stack *Stack) Top() object.Object {
	top := stack.top
	if top < 0 {
		return nil
	}

	return stack.items[top]
}
