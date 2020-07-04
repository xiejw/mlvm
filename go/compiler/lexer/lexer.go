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

func NewLexer(input []byte) *Lexer {
	l := &Lexer{
		input: input,
		size:  uint(len(input)),
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() *token.Token {
	var tok token.Token
	tok.Location = l.loc // a copy

	switch l.ch {
	case 0:
		tok.Type = token.EOF
	default:
		panic(fmt.Sprintf("unkknown character: %v. location: %+v", l.ch, l.loc))
	}

	l.readChar()
	return &tok
}

func (l *Lexer) readChar() {
	if l.readPosition > 0 {
		l.loc.Column += 1
	}

	if l.readPosition >= l.size {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
