package ast

import (
	"bytes"
	"fmt"
	"io"
)

////////////////////////////////////////////////////////////////////////////////
/// Interfaces: Program, Expression
////////////////////////////////////////////////////////////////////////////////

type Expressions []Expr

type Program struct {
	Expressions Expressions
}

type Expr interface {
	ToHumanReadableString(w io.Writer)
}

////////////////////////////////////////////////////////////////////////////////
/// Expressions
////////////////////////////////////////////////////////////////////////////////

type Id struct {
	Value string
}

type FunctionCall struct {
	Func *Id
	Args []Expr
}

type IntegerLiteral struct {
	Value int64
}

type FloatLiteral struct {
	Value float32
}

type ShapeLiteral struct {
	Dimensions []*Id
}

type ArrayLiteral struct {
	Values []*FloatLiteral
}

type StringLiteral struct {
	Value string
}

func (id *Id) ToHumanReadableString(w io.Writer) {
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
func (literal *FloatLiteral) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Float(%v)", literal.Value)
}
func (literal *ShapeLiteral) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Shape(")
	for i, dim := range literal.Dimensions {
		dim.ToHumanReadableString(w)
		if i != len(literal.Dimensions)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ")")
}
func (literal *ArrayLiteral) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Array(")
	for i, f := range literal.Values {
		f.ToHumanReadableString(w)
		if i != len(literal.Values)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ")")
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
