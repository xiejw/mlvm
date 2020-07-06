package lexer

import (
	"testing"

	"github.com/xiejw/mlvm/go/syntax/token"
)

func TestLexerTokens(t *testing.T) {
	l := New([]byte("(def a 123)\n(+ a a)\n"))

	expects := []struct {
		row       uint
		col       uint
		pos       uint
		tokenType token.TokenType
		literal   string
	}{
		{0, 0, 0, token.LPAREN, "("},
		{0, 1, 1, token.IDENTIFIER, "def"},
		{0, 5, 5, token.IDENTIFIER, "a"},
		{0, 7, 7, token.INTEGER, "123"},
		{0, 10, 10, token.RPAREN, ")"},
		{1, 0, 12, token.LPAREN, "("},
		{1, 1, 13, token.IDENTIFIER, "+"},
		{1, 3, 15, token.IDENTIFIER, "a"},
		{1, 5, 17, token.IDENTIFIER, "a"},
		{1, 6, 18, token.RPAREN, ")"},
		{2, 0, 20, token.EOF, ""},
	}

	for i, expected := range expects {
		token := l.NextToken()
		if token.Type != expected.tokenType {
			t.Errorf("type mismatch for index: %v", i)
		}
		if token.Literal != expected.literal {
			t.Errorf("literal mismatch for index: %v", i)
		}
		if token.Location.Row != expected.row {
			t.Errorf("row mismatch for index: %v", i)
		}
		if token.Location.Column != expected.col {
			t.Errorf("col mismatch for index: %v", i)
		}
		if token.Location.Position != expected.pos {
			t.Errorf("pos mismatch for index: %v", i)
		}
	}
}

func TestStringLiteral(t *testing.T) {
	l := New([]byte("(def \"123\")"))

	expects := []struct {
		row       uint
		col       uint
		pos       uint
		tokenType token.TokenType
		literal   string
	}{
		{0, 0, 0, token.LPAREN, "("},
		{0, 1, 1, token.IDENTIFIER, "def"},
		{0, 5, 5, token.STRING, `"123"`},
		{0, 10, 10, token.RPAREN, ")"},
		{0, 11, 11, token.EOF, ""},
	}

	for i, expected := range expects {
		token := l.NextToken()
		if token.Type != expected.tokenType {
			t.Errorf("type mismatch for index: %v", i)
		}
		if token.Literal != expected.literal {
			t.Errorf("literal mismatch for index: %v", i)
		}
		if token.Location.Row != expected.row {
			t.Errorf("row mismatch for index: %v", i)
		}
		if token.Location.Column != expected.col {
			t.Errorf("col mismatch for index: %v", i)
		}
		if token.Location.Position != expected.pos {
			t.Errorf("pos mismatch for index: %v", i)
		}
	}
}
