package ast

import (
	"bytes"
	"fmt"
	"io"
)

////////////////////////////////////////////////////////////////////////////////
/// Interfaces: Program, Expression
////////////////////////////////////////////////////////////////////////////////

type Expressions []Expression

type Program struct {
	Expressions Expressions
}

type Node interface {
	ToHumanReadableString(w io.Writer)
}

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

func (id *Identifier) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "ID(%v)", id.Value)
}
func (fc *FunctionCall) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Func(")
	fc.Func.ToHumanReadableString(w)
	for _, arg := range fc.Args {
		fmt.Fprintf(w, ", ")
		arg.ToHumanReadableString(w)
	}
	fmt.Fprintf(w, ")")
}
func (literal *IntegerLiteral) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Int(%v)", literal.Value)
}
func (literal *StringLiteral) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Str(\"%v\")", literal.Value)
}

func (exprs Expressions) String() string {
	var buf bytes.Buffer
	for _, expr := range exprs {
		expr.ToHumanReadableString(&buf)
		fmt.Fprint(&buf, "\n")
	}
	return buf.String()
}
