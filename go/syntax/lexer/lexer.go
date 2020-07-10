package lexer

import (
	"github.com/xiejw/mlvm/go/syntax/token"
)

type Lexer struct {
	input        []byte
	position     uint
	readPosition uint // after position
	ch           byte
	size         uint
	loc          token.Location
}

func New(input []byte) *Lexer {
	l := &Lexer{
		input: input,
		size:  uint(len(input)),
	}
	l.readChar()
	return l
}

func (l *Lexer) Bytes(startPos, endPos uint) []byte {
	return l.input[startPos:endPos]
}

func (l *Lexer) NextToken() *token.Token {
	var tok token.Token
	l.skipWhiteSpaces()

	tok.Location = l.loc // a copy

	switch l.ch {
	case '(':
		tok.Type = token.LPAREN
		tok.Literal = "("
	case ')':
		tok.Type = token.RPAREN
		tok.Literal = ")"
	case '[':
		tok.Type = token.LSBRACKET
		tok.Literal = "["
	case ']':
		tok.Type = token.RSBRACKET
		tok.Literal = "]"
	case '\\':
		tok.Type = token.BACKSLASH
		tok.Literal = "\\"
	case 0:
		tok.Type = token.EOF
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	default:
		if isIdentifierChar(l.ch) {
			tok.Type = token.IDENTIFIER
			tok.Literal = l.readIdentifider()
		} else if isDigit(l.ch) {
			tok.Type = token.INTEGER
			tok.Literal = l.readInteger()
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.ch)
		}
		return &tok // skip the next readChar.
	}

	l.readChar()
	return &tok
}

func (l *Lexer) skipWhiteSpaces() {
	for {
		c := l.ch
		if c == ' ' || c == '\n' || c == '\t' || c == '\a' {
			l.readChar()
			continue
		}
		break
	}
}

func (l *Lexer) readChar() {
	// handles the location. two special cases.
	// case 1. just started
	if l.readPosition > 0 {
		l.loc.Column += 1
	}
	// case 2. just read a newline
	if l.ch == '\n' {
		l.loc.Row += 1
		l.loc.Column = 0
	}

	// sets the new ch and advances the positions
	if l.readPosition >= l.size {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
	l.loc.Position = l.position
}

func (l *Lexer) readIdentifider() string {
	pos := l.position
	for isIdentifierChar(l.ch) {
		l.readChar()
	}
	return string(l.input[pos:l.position])
}

func (l *Lexer) readInteger() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[pos:l.position])
}

func (l *Lexer) readString() string {
	pos := l.position
	l.readChar()
	// TODO: handle EOF and newline case.
	for l.ch != '"' {
		l.readChar()
	}
	return string(l.input[pos : l.position+1])
}

func isIdentifierChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '+' == ch || '_' == ch
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
