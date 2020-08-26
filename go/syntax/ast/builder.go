package ast

func NewIntLit(value int64) Expr {
	e := IntLit{Value: value}
	e.baseExpr.etype = Int
	return &e
}
