package lexer

import (
	"github.com/xiejw/mlvm/go/syntax/token"
)

// Lexer, which emits Tokens from input bytes.
type Lexer struct {
	input        []byte
	position     uint
	readPosition uint // after position
	ch           byte
	size         uint
	loc          token.Loc
}

// Creates a Lexer, reay for use.
//
// Input is not expected to be changed later.
func New(input []byte) *Lexer {
	l := &Lexer{
		input: input,
		size:  uint(len(input)),
	}
	l.readChar()
	return l
}

// Given a valid range startPos..endPos, returns the bytes.
func (l *Lexer) Bytes(startPos, endPos uint) []byte {
	return l.input[startPos:endPos]
}

// Returns the next token.
func (l *Lexer) NextToken() *token.Token {
	var tok token.Token
	l.skipWhiteSpaces()

	tok.Loc = l.loc // a "deep" copy

	switch l.ch {
	case '(':
		tok.Type = token.Lparen
		tok.Literal = "("
	case ')':
		tok.Type = token.Rparen
		tok.Literal = ")"
	case '[':
		tok.Type = token.Lbrack
		tok.Literal = "["
	case ']':
		tok.Type = token.Rbrack
		tok.Literal = "]"
	case '\\':
		tok.Type = token.Bslash
		tok.Literal = "\\"
	case 0:
		tok.Type = token.Eof
	case '"':
		tok.Type = token.String
		tok.Literal = l.readString()
	default:
		if isIdentifierChar(l.ch) {
			tok.Type = token.Id
			tok.Literal = l.readIdentifider()
		} else if isDigit(l.ch) {
			tok.Literal, tok.Type = l.readNumber()
		} else {
			tok.Type = token.Illegal
			tok.Literal = string(l.ch)
		}
		return &tok // skip the next readChar.
	}

	l.readChar()
	return &tok
}

// Skips all white spaces.
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

// Reads the next char. Updates Loc info.
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

func (l *Lexer) readNumber() (string, token.TokenType) {
	hitDecimalPoint := false
	tokenType := token.Int

	pos := l.position
	for {
		ch := l.ch
		if isDigit(ch) {
			l.readChar()
			continue
		}

		// decimal point should be hit at most once.
		if l.ch == '.' && !hitDecimalPoint {
			hitDecimalPoint = true
			tokenType = token.Float
			l.readChar()
			continue
		}
		break
	}
	return string(l.input[pos:l.position]), tokenType
}

func (l *Lexer) readString() string {
	pos := l.position
	l.readChar()
	// TODO: handle Eof and newline case.
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
