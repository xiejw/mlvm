package runtime

import (
	"errors"
)

const defaultStackSize = 256

var errPopOnEmptyStack = errors.New("Pop an item from empty stack.")

type Stack struct {
	items []interface{}
	top   int
}

func NewStack() *Stack {
	stack := &Stack{
		items: make([]interface{}, 0, defaultStackSize),
		top:   -1,
	}
	return stack
}

func (stack *Stack) Push(item interface{}) error {
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

func (stack *Stack) Pop() (interface{}, error) {
	top := stack.top
	if top < 0 {
		return nil, errPopOnEmptyStack
	}

	item := stack.items[top]
	stack.items[top] = nil // Freeze memory.
	stack.top--
	return item, nil
}
