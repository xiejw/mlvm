package algorithrms

import (
	"github.com/xiejw/mlvm/go/base/errors"
)

type Item interface{}

type Queue struct {
	front *qNode
	end   *qNode
}

func (q *Queue) Dequeue() (Item, *errors.DError) {
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

func (q *Queue) Enqueue(item Item) *errors.DError {
	if q.end == nil {
		n := &qNode{
			V: item,
			N: nil,
		}
		q.front = n
		q.end = n
		return nil
	}

	q.end.N = &qNode{
		V: item,
		N: nil,
	}
	return nil
}

type qNode struct {
	V Item
	N *qNode
}
