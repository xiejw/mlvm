package ast

import "github.com/xiejw/mlvm/go/base/errors"

type Builder struct {
	dims map[string]int
}

func NewBuilder() *Builder {
	return &Builder{
		dims: make(map[string]int),
	}
}

func (b *Builder) AddNamedDim(name string, size int) *errors.DError {
	_, existed := b.dims[name]
	if existed {
		return errors.New("Named dim %s already existed", name)
	}
	b.dims[name] = size
	return nil
}

func (b *Builder) NewIntLit(value int64) Expr {
	e := IntLit{Value: value}
	e.baseExpr.etype = Int
	return &e
}
