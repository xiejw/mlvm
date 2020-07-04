package token

type TokenType uint

const (
	LPAREN TokenType = iota
	RPAREN
	IDENTIFIER
	INTEGER
	PLUS
	DEF
	ILLEGAL
	EOF
)

type Location struct {
	Row    uint
	Column uint
}

type Token struct {
	Type     TokenType
	Literal  string
	Location Location
}
