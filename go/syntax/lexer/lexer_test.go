package lexer

import (
	"testing"

	"github.com/xiejw/mlvm/go/syntax/token"
)

type expectedToken struct {
	row       uint
	col       uint
	pos       uint
	tokenType token.TokenType
	literal   string
}

func assertToken(t *testing.T, expected expectedToken, token *token.Token, i int) {
	t.Helper()
	if token.Type != expected.tokenType {
		t.Errorf("type mismatch for index: %v", i)
	}
	if token.Literal != expected.literal {
		t.Errorf("literal mismatch for index: %v. expected: %v, got: %v",
			i, expected.literal, token.Literal)
	}
	if token.Loc.Row != expected.row {
		t.Errorf("row mismatch for index: %v", i)
	}
	if token.Loc.Column != expected.col {
		t.Errorf("col mismatch for index: %v", i)
	}
	if token.Loc.Position != expected.pos {
		t.Errorf("pos mismatch for index: %v", i)
	}
}

func TestLexerTokens(t *testing.T) {
	l := New([]byte("(def a 123)\n(+ a a)\n"))

	expects := []expectedToken{
		{0, 0, 0, token.Lparen, "("},
		{0, 1, 1, token.Id, "def"},
		{0, 5, 5, token.Id, "a"},
		{0, 7, 7, token.Int, "123"},
		{0, 10, 10, token.Rparen, ")"},
		{1, 0, 12, token.Lparen, "("},
		{1, 1, 13, token.Id, "+"},
		{1, 3, 15, token.Id, "a"},
		{1, 5, 17, token.Id, "a"},
		{1, 6, 18, token.Rparen, ")"},
		{2, 0, 20, token.Eof, ""},
	}

	for i, expected := range expects {
		got := l.NextToken()
		assertToken(t, expected, got, i)
	}
}

func TestStringLiteral(t *testing.T) {
	l := New([]byte("(def \"123\")"))

	expects := []expectedToken{
		{0, 0, 0, token.Lparen, "("},
		{0, 1, 1, token.Id, "def"},
		{0, 5, 5, token.String, `"123"`},
		{0, 10, 10, token.Rparen, ")"},
		{0, 11, 11, token.Eof, ""},
	}

	for i, expected := range expects {
		got := l.NextToken()
		assertToken(t, expected, got, i)
	}
}

func TestPunctuation(t *testing.T) {
	l := New([]byte("[]\\[]"))

	expects := []expectedToken{
		{0, 0, 0, token.Lbrack, "["},
		{0, 1, 1, token.Rbrack, "]"},
		{0, 2, 2, token.Bslash, `\`},
		{0, 3, 3, token.Lbrack, "["},
		{0, 4, 4, token.Rbrack, "]"},
		{0, 5, 5, token.Eof, ""},
	}

	for i, expected := range expects {
		got := l.NextToken()
		assertToken(t, expected, got, i)
	}
}

func TestFloatPoint(t *testing.T) {
	l := New([]byte("[1.23 2.34]"))

	expects := []expectedToken{
		{0, 0, 0, token.Lbrack, "["},
		{0, 1, 1, token.Float, "1.23"},
		{0, 6, 6, token.Float, "2.34"},
		{0, 10, 10, token.Rbrack, "]"},
		{0, 11, 11, token.Eof, ""},
	}

	for i, expected := range expects {
		got := l.NextToken()
		assertToken(t, expected, got, i)
	}
}
