package token

type TokenType string

type Token struct {
	Type    TokenType // Token type.
	Literal string    // The Literal in the source file.
	Loc     Loc       // Location of the Token. Length can be deduced by Literal
}

// TokenTypes
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	FLOAT      = "FLOAT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNC   = "FUNC"
	LET    = "LET"
	RETURN = "RETURN"
	IF     = "IF"
	ELSE   = "ELSE"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
)

// Keywords.
var keywords = map[string]TokenType{
	"func":   FUNC,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

// Returns the TokenType for the identifier, including keywords.
//
// Returns the corresponding keyword token type. Otherwise, returns IDENTIFIER as type.
func LookupIdentifier(id string) TokenType {
	if tokenType, ok := keywords[id]; ok {
		return tokenType
	}
	return IDENTIFIER
}
