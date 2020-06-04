package ast

import (
	"github.com/xiejw/mlvm/go/object"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
}

// Represents the Kind (enum) of Type of Expression.
type TypeKind uint

const (
	TpKdVoid TypeKind = iota
	TpKdInt
	TpKdString
	TpKdArray
	TpKdNamedDim
	TpKdTensor
	TpKdFunc
)

// Represents a parameter in a func signature.
type Param struct {
	Name string
	Type *Type
}

// Represents Type of Expression.
//
// For Array, the base type is SubType.
type Type struct {
	Kind    TypeKind
	SubType *Type
	Params  []Param
}

type Decl struct {
	Name  string
	Type  *Type
	Value *Expression
	Code  Statement
	Next  *Decl
}

type ExpressionType uint

const (
	ExTpValue ExpressionType = iota
	ExTpAdd
	ExTpMul
)

type Expression struct {
	Type  ExpressionType
	Left  *Expression
	Right *Expression
	Value object.Object
}
