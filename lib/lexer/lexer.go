package lexer

import (
	"github.com/xiejw/mlvm/lib/token"
)

type Lexer struct {
	input        string
	loc          token.Loc // Current Location. This respects new line `\n`. 1-based index.
	position     int       // Points to current char. This and the companion below do not handle new line differently.
	readPosition int       // After position
	ch           byte      // Current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.loc.L, l.loc.C = 1, 0
	l.readChar()
	return l
}

// Returns the token, including the EOF.
//
// Invocation advances the lexer's position. Behavior is undefined after EOF is returned.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	tok.Loc = l.loc // Make a copy.
	ch := l.ch
	switch ch {
	case '=':
		fillToken(&tok, token.ASSIGN, ch)
	case '+':
		fillToken(&tok, token.PLUS, ch)
	case '-':
		fillToken(&tok, token.MINUS, ch)
	case '*':
		fillToken(&tok, token.ASTERISK, ch)
	case '/':
		fillToken(&tok, token.SLASH, ch)
	case '!':
		fillToken(&tok, token.BANG, ch)
	case '<':
		fillToken(&tok, token.LT, ch)
	case '>':
		fillToken(&tok, token.GT, ch)
	case ',':
		fillToken(&tok, token.COMMA, ch)
	case ';':
		fillToken(&tok, token.SEMICOLON, ch)
	case '(':
		fillToken(&tok, token.LPAREN, ch)
	case ')':
		fillToken(&tok, token.RPAREN, ch)
	case '{':
		fillToken(&tok, token.LBRACE, ch)
	case '}':
		fillToken(&tok, token.RBRACE, ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(ch) {
			tok.Literal, tok.Type = l.readIdentifier()
			return tok // Returns immediately to avoid readChar() again.
		} else if isDigit(ch) {
			tok.Literal, tok.Type = l.readNumber()
			return tok // Returns immediately to avoid readChar() again.
		} else {
			fillToken(&tok, token.ILLEGAL, ch)
		}
	}
	l.readChar()
	return tok
}

// Reads next character.
func (l *Lexer) readChar() {
	var ch byte // Zero value.
	if l.readPosition < len(l.input) {
		ch = l.input[l.readPosition]
	}

	// Updates the loc in lexer.
	switch ch {
	case '\r': // no-op
	case '\n':
		l.loc.L += 1
		l.loc.C = 0
	default:
		l.loc.C += 1
	}

	l.ch = ch
	l.position = l.readPosition
	l.readPosition += 1
}

// Skips all white space including newline.
func (l *Lexer) skipWhitespace() {
	ch := l.ch
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		l.readChar()
		ch = l.ch
	}
}

// Reads the identifier.
//
//     identifier = letter +
//     letter = [a-zA-Z_]
func (l *Lexer) readIdentifier() (string, token.TokenType) {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	literal := l.input[position:l.position]
	return literal, token.LookupIdentifier(literal)
}

// Reads the number.
//
//     number = (digit+ | digit+ `.` digit*)
//     digit = [0-9]
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
func fillToken(tok *token.Token, tokenType token.TokenType, ch byte) {
	tok.Type = tokenType
	tok.Literal = string(ch)
}

// Helper method to check whether a byte is considerd as letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Helper method to check whether a byte is considerd as a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
