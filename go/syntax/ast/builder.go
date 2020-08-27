package ast

type Builder struct {
	dims map[string]int
}

func NewBuilder() *Builder {
	return &Builder{
		dims: make(map[string]int),
	}
}

func (b *Builder) NewIntLit(value int64) Expr {
	e := IntLit{Value: value}
	e.baseExpr.etype = Int
	return &e
}
