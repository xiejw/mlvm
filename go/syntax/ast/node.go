package ast

import (
	"encoding/json"
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
	TpKdPrng
)

func (kind TypeKind) MarshalJSON() ([]byte, error) {
	switch kind {
	case TpKdInt:
		return json.Marshal("int")
	case TpKdString:
		return json.Marshal("string")
	case TpKdArray:
		return json.Marshal("array")
	case TpKdNamedDim:
		return json.Marshal("named_dim")
	case TpKdTensor:
		return json.Marshal("tensor")
	case TpKdFunc:
		return json.Marshal("func")
	case TpKdPrng:
		return json.Marshal("prng")
	default:
		return json.Marshal("unknown_type")
	}
}

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

const (
	EpKdUnspecified ExpressionKind = iota
	EpKdAdd
	EpKdMul
	EpKdCall
	EpKdID
	EpKdIntLiteral
)

func (kind ExpressionKind) MarshalJSON() ([]byte, error) {
	switch kind {
	case EpKdCall:
		return json.Marshal("func_call")
	case EpKdID:
		return json.Marshal("id")
	default:
		return json.Marshal("unknown_expr_kind")
	}
}

type Expression struct {
	Type  *Type
	Kind  ExpressionKind
	Left  *Expression   `json:",omitempty"`
	Right *Expression   `json:",omitempty"`
	Value object.Object `json:",omitempty"`
}
