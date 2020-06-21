package ast

import ()

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

type DeclStatement struct {
	Name  *Identifier
	Value Expression
	// Type Type
}

func (decl *DeclStatement) statementNode() {}

type Identifier struct {
	Value string
}

func (id *Identifier) expressionNode() {}

type FunctionCall struct {
	Name *Identifier
	Args []Expression
}

func (fc *FunctionCall) expressionNode() {}

type IntLiteral struct {
	Value int64
}

func (literal *IntLiteral) expressionNode() {}
