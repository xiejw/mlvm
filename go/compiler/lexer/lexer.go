package lexer

import (
	"fmt"

	"github.com/xiejw/mlvm/go/compiler/token"
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
	case '+':
		tok.Type = token.PLUS
		tok.Literal = "+"
	case 0:
		tok.Type = token.EOF
	default:
		panic(fmt.Sprintf("unkknown character: %v. location: %+v", l.ch, l.loc))
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
}
