package token

import "fmt"

type TokenType uint

const (
	Lparen TokenType = iota // (
	Rparen                  // )
	Lbrack                  // [
	Rbrack                  // ]
	Bslash                  // \
	Id
	Int
	Float
	String
	Illegal
	Eof
)

type Loc struct {
	Row      uint
	Column   uint
	Position uint
}

type Token struct {
	Type    TokenType
	Literal string
	Loc     Loc
}

func (tok *Token) String() string {
	return fmt.Sprintf("token{ type: %v, loc: (%3d, %3d), literal: \"%v\" }",
		tok.Type, tok.Loc.Row, tok.Loc.Column, tok.Literal)
}

func (t TokenType) String() string {
	switch t {
	case Lparen:
		return "Lparen    "
	case Rparen:
		return "Rparen    "
	case Lbrack:
		return "Lbrack    "
	case Rbrack:
		return "Rbrack    "
	case Bslash:
		return "Bslash    "
	case Id:
		return "Id        "
	case Int:
		return "Int       "
	case Float:
		return "Float     "
	case String:
		return "String    "
	case Illegal:
		return "Illegal   "
	case Eof:
		return "Eof       "
	default:
		return fmt.Sprintf("(unknown token type: %d)", uint(t))
	}
}
