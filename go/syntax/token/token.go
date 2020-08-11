package token

import "fmt"

type TokenType uint

const (
	Lparen TokenType = iota
	Rparen
	LSBRACKET
	RSBRACKET
	BACKSLASH
	IDENTIFIER
	INTEGER
	FLOAT
	STRING
	ILLEGAL
	EOF
)

type Location struct {
	Row      uint
	Column   uint
	Position uint
}

type Token struct {
	Type     TokenType
	Literal  string
	Location Location
}

func (tok *Token) String() string {
	return fmt.Sprintf("token{ type: %v, loc: (%3d, %3d), literal: \"%v\" }",
		tok.Type, tok.Location.Row, tok.Location.Column, tok.Literal)
}

func (t TokenType) String() string {
	switch t {
	case Lparen:
		return "Lparen    "
	case Rparen:
		return "Rparen    "
	case LSBRACKET:
		return "LSBRACKET "
	case RSBRACKET:
		return "RSBRACKET "
	case BACKSLASH:
		return "BACKSLASH "
	case IDENTIFIER:
		return "IDENTIFIER"
	case INTEGER:
		return "INTEGER   "
	case FLOAT:
		return "FLOAT     "
	case STRING:
		return "STRING    "
	case ILLEGAL:
		return "ILLEGAL   "
	case EOF:
		return "EOF       "
	default:
		return fmt.Sprintf("(unknown token type: %d)", uint(t))
	}
}
