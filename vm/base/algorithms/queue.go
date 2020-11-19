package algorithrms

import (
	"github.com/xiejw/mlvm/vm/base/errors"
)

type Item interface{}

type Queue struct {
	front *qNode
	end   *qNode
}

func (q *Queue) Dequeue() (Item, error) {
	n := q.front
	if n == nil {
		return nil, errors.New("queue is empty; cannot dequeue.")
	}

	q.front = n.N

	if q.front == nil {
		q.end = nil
	}

	return n.V, nil
}

func (q *Queue) Enqueue(item Item) error {
	if q.end == nil {
		n := &qNode{
			V: item,
			N: nil,
		}
		q.front = n
		q.end = n
		return nil
	}

	n := &qNode{
		V: item,
		N: nil,
	}
	q.end.N = n
	q.end = n
	return nil
}

type qNode struct {
	V Item
	N *qNode
}
