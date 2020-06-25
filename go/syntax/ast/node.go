package ast

////////////////////////////////////////////////////////////////////////////////
/// Interfaces: Program, Statement, Expression
////////////////////////////////////////////////////////////////////////////////

type Program struct {
	Statements []Statement
}

type Node interface{}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

////////////////////////////////////////////////////////////////////////////////
/// Statements
////////////////////////////////////////////////////////////////////////////////

type DeclStatement struct {
	Name  *Identifier
	Value Expression
	// Type Type
}

type ExprStatement struct {
	Value Expression
}

func (s *DeclStatement) statementNode() {}
func (s *ExprStatement) statementNode() {}

////////////////////////////////////////////////////////////////////////////////
/// Expressions
////////////////////////////////////////////////////////////////////////////////

type Identifier struct {
	Value string
}

type FunctionCall struct {
	Name *Identifier
	Args []Expression
}

type IntegerLiteral struct {
	Value int64
}

func (id *Identifier) expressionNode()          {}
func (fc *FunctionCall) expressionNode()        {}
func (literal *IntegerLiteral) expressionNode() {}
