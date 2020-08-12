package ast

import (
	"bytes"
	"fmt"
	"io"
)

////////////////////////////////////////////////////////////////////////////////
/// Interfaces: Program, Expression
////////////////////////////////////////////////////////////////////////////////

type Exprs []Expr

type Program struct {
	Exprs Exprs
}

type Expr interface {
	ToHumanReadableString(w io.Writer)
}

func String(e Expr) string {
	var buf bytes.Buffer
	e.ToHumanReadableString(&buf)
	return buf.String()
}

func (exprs Exprs) String() string {
	var buf bytes.Buffer
	for _, expr := range exprs {
		expr.ToHumanReadableString(&buf)
		fmt.Fprint(&buf, "\n")
	}
	return buf.String()
}

////////////////////////////////////////////////////////////////////////////////
/// Base Expr
////////////////////////////////////////////////////////////////////////////////

type baseExpr struct {
}

////////////////////////////////////////////////////////////////////////////////
/// Expressions
////////////////////////////////////////////////////////////////////////////////

type Id struct {
	Value string

	baseExpr
}

type App struct {
	Func *Id
	Args []Expr

	baseExpr
}

type IntLit struct {
	Value int64

	baseExpr
}

type FloatLit struct {
	Value float32

	baseExpr
}

type ShapeLit struct {
	Dims []*Id

	baseExpr
}

type ArrayLit struct {
	Values []*FloatLit

	baseExpr
}

type StringLit struct {
	Value string

	baseExpr
}

func (id *Id) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Id(%v)", id.Value)
}
func (fc *App) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "App(")
	fc.Func.ToHumanReadableString(w)
	for _, arg := range fc.Args {
		fmt.Fprintf(w, ", ")
		arg.ToHumanReadableString(w)
	}
	fmt.Fprintf(w, ")")
}
func (literal *IntLit) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Int(%v)", literal.Value)
}
func (literal *FloatLit) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Float(%v)", literal.Value)
}
func (literal *ShapeLit) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Shape(")
	for i, dim := range literal.Dims {
		dim.ToHumanReadableString(w)
		if i != len(literal.Dims)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ")")
}
func (literal *ArrayLit) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "Array(")
	for i, f := range literal.Values {
		f.ToHumanReadableString(w)
		if i != len(literal.Values)-1 {
			fmt.Fprintf(w, ", ")
		}
	}
	fmt.Fprintf(w, ")")
}
func (literal *StringLit) ToHumanReadableString(w io.Writer) {
	fmt.Fprintf(w, "String(\"%v\")", literal.Value)
}
