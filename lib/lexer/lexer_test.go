package lexer

import (
	"testing"

	"github.com/xiejw/mlvm/lib/token"
)

type ExpectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextTokenWithBasicChars(t *testing.T) {
	input := `=+(){},;`

	expectedTokens := []ExpectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	assertTokens(t, input, expectedTokens)
}

func TestNextTokenWithIdAndInts(t *testing.T) {
	input := `let five= 5;
            let ten =10;
            let add = func(x, y) {
                x + y;
            }
            let result = add(five, ten);`

	expectedTokens := []ExpectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		// {token.PLUS, "+"},
		// {token.LPAREN, "."},
		// {token.RPAREN, ")"},
		// {token.LBRACE, "{"},
		// {token.RBRACE, "}"},
		// {token.COMMA, ","},
		// {token.SEMICOLON, ";"},
		// {token.EOF, ""},
	}

	assertTokens(t, input, expectedTokens)
}

func assertTokens(t *testing.T, input string, expectedTokens []ExpectedToken) {
	t.Helper()

	l := New(input)

	for i, tt := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[index: %2d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[index: %2d] - token literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
