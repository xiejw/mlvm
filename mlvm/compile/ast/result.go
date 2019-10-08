package ast

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

type Result struct {
	name  string
	shape *array.Shape
	ins   *Instruction
	index int // Result index in Instruction
}

func (r *Result) Name() string {
	return r.name
}

func (r *Result) Shape() *array.Shape {
	return r.shape
}
