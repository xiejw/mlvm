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
// - Invocation advances the lexer's position.
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
		} else if isDigit(ch) {
			tok.Literal, tok.Type = l.readNumber()
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

// Reads (and thereby advances lexer's position) the identifier.
//
// identifier = letter +
// letter = [a-zA-Z_]
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Reads (and thereby advances lexer's position) the number.
//
// number = (digit+ | digit+ `.` digit*)
// digit = [0-9]
func (l *Lexer) readNumber() (string, token.TokenType) {
	position := l.position

	// Stage 1: Read enough digits.
	for isDigit(l.ch) {
		l.readChar()
	}

	// If it is an int, fast return.
	if l.ch != '.' {
		literal := l.input[position:l.position]
		return literal, token.INT
	}

	// Eat the period `.`
	l.readChar()

	// Stage 2: Read enough digits.
	for isDigit(l.ch) {
		l.readChar()
	}

	literal := l.input[position:l.position]
	return literal, token.FLOAT
}

// Helper method to create a new Token.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// Helper method to check whether a byte is considerd as letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Helper method to check whether a byte is considerd as a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
