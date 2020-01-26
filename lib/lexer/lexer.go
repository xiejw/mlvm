package lexer

import (
	"github.com/xiejw/mlvm/lib/token"
)

type Lexer struct {
	input        string
	position     int  // Points to current char
	readPosition int  // After position
	ch           byte // Current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Returns the token, including the EOF.
//
// - Invocation advances the input reading.
// - Behavior is undefined after EOF is returned.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	ch := l.ch
	switch ch {
	case '=':
		tok = newToken(token.ASSIGN, ch)
	case '+':
		tok = newToken(token.PLUS, ch)
	case ',':
		tok = newToken(token.COMMA, ch)
	case ';':
		tok = newToken(token.SEMICOLON, ch)
	case '(':
		tok = newToken(token.LPAREN, ch)
	case ')':
		tok = newToken(token.RPAREN, ch)
	case '{':
		tok = newToken(token.LBRACE, ch)
	case '}':
		tok = newToken(token.RBRACE, ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok // Returns immediately to avoid readChar() again.
		} else {
			tok = newToken(token.ILLEGAL, ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	ch := l.ch
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		l.readChar()
		ch = l.ch
	}
}

// Reads (and thereby advances internal reader) the identifier.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Helper method to create a new Token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper method to check whether a byte is considerd as letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
