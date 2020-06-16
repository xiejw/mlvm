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

// MarshalJSON in node_json.go
const (
	TpKdUnspecified TypeKind = iota
	TpKdNA
	TpKdVoid
	TpKdInt
	TpKdString
	TpKdArray
	TpKdNamedDim
	TpKdTensor
	TpKdFunc
	TpKdPrng
)

// Represents a parameter in a func signature.
type Param struct {
	ID   string
	Type *Type
}

// Represents Type of Expression.
//
// For Array, the base type is SubType.
type Type struct {
	Kind    TypeKind
	SubType *Type   `json:",omitempty"`
	Params  []Param `json:",omitempty"`
}

type Decl struct {
	ID    string
	Type  *Type
	Value *Expression `json:",omitempty"`
	Code  Statement   `json:",omitempty"`
	Next  *Decl       `json:",omitempty"`
}

type ExpressionKind uint

// MarshalJSON in node_json.go
const (
	EpKdUnspecified ExpressionKind = iota
	EpKdAdd
	EpKdMul
	EpKdCall
	EpKdID
	EpKdArg
	EpKdIntLiteral
)

type Expression struct {
	Type  *Type
	Kind  ExpressionKind
	Left  *Expression   `json:",omitempty"`
	Right *Expression   `json:",omitempty"`
	Value object.Object `json:",omitempty"`
}
