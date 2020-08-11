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

type App struct {
	Func *Id
	Args []Expr
}

type IntLit struct {
	Value int64
}

type FloatLit struct {
	Value float32
}

type ShapeLit struct {
	Dimensions []*Id
}

type ArrayLit struct {
	Values []*FloatLit
}

type StringLit struct {
	Value string
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
	for i, dim := range literal.Dimensions {
		dim.ToHumanReadableString(w)
		if i != len(literal.Dimensions)-1 {
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
