package ast

////////////////////////////////////////////////////////////////////////////////
/// Interfaces: Program, Expression
////////////////////////////////////////////////////////////////////////////////

type Program struct {
	Expressions []Expression
}

type Node interface{}

type Expression interface {
	Node
	expressionNode()
}

////////////////////////////////////////////////////////////////////////////////
/// Expressions
////////////////////////////////////////////////////////////////////////////////

type Identifier struct {
	Value string
}

type FunctionCall struct {
	Func *Identifier
	Args []Expression
}

type IntegerLiteral struct {
	Value int64
}

type StringLiteral struct {
	Value string
}

func (id *Identifier) expressionNode()          {}
func (fc *FunctionCall) expressionNode()        {}
func (literal *IntegerLiteral) expressionNode() {}
func (literal *StringLiteral) expressionNode()  {}
