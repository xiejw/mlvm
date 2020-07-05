package token

import "fmt"

type TokenType uint

const (
	LPAREN TokenType = iota
	RPAREN
	IDENTIFIER
	INTEGER
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
	case LPAREN:
		return "LPAREN    "
	case RPAREN:
		return "RPAREN    "
	case IDENTIFIER:
		return "IDENTIFIER"
	case INTEGER:
		return "INTEGER   "
	case ILLEGAL:
		return "ILLEGAL   "
	case EOF:
		return "EOF       "
	default:
		return fmt.Sprintf("(unknown token type: %d)", uint(t))
	}
}
