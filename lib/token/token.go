package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	FLOAT      = "FLOAT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNC = "FUNC"
	LET  = "LET"
)

var keywords = map[string]TokenType{
	"func": FUNC,
	"let":  LET,
}

// Returns the TokenType for the identifier.
//
// - Returns the corresponding keyword token type.
// - Otherwise, returns IDENTIFIER as type.
func LookupIdentifier(id string) TokenType {
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return IDENTIFIER
}
