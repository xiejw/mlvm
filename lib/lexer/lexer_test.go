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

func TestNextTokenWithIdentifiers(t *testing.T) {
	input := `abc e_f ghhh_ _hi`

	expectedTokens := []ExpectedToken{
		{token.IDENTIFIER, "abc"},
		{token.IDENTIFIER, "e_f"},
		{token.IDENTIFIER, "ghhh_"},
		{token.IDENTIFIER, "_hi"},
		{token.EOF, ""},
	}

	assertTokens(t, input, expectedTokens)
}

func TestNextTokenWithInvalidIdentifiers(t *testing.T) {
	input := `ab.c d123`

	expectedTokens := []ExpectedToken{
		{token.IDENTIFIER, "ab"},
		{token.ILLEGAL, "."},
		{token.IDENTIFIER, "c"},
		{token.IDENTIFIER, "d"},
		{token.INT, "123"},
		{token.EOF, ""},
	}

	assertTokens(t, input, expectedTokens)
}

func TestNextTokenWithNumbers(t *testing.T) {
	input := `20  20.  3.23 `

	expectedTokens := []ExpectedToken{
		{token.INT, "20"},
		{token.FLOAT, "20."},
		{token.FLOAT, "3.23"},
		{token.EOF, ""},
	}

	assertTokens(t, input, expectedTokens)
}

func TestNextTokenWithInvalidNumbers(t *testing.T) {
	input := `3.23. .45 `

	expectedTokens := []ExpectedToken{
		{token.FLOAT, "3.23"},
		{token.ILLEGAL, "."},
		{token.ILLEGAL, "."},
		{token.INT, "45"},
	}

	assertTokens(t, input, expectedTokens)
}

func TestNextTokenWithFuncIdAndInts(t *testing.T) {
	input := `let five= 5;
            let ten =10;
            let add = func(x, y) {
                x + y;
            }
            let result = add(five, ten);`

	expectedTokens := []ExpectedToken{
		/*  0 */ {token.LET, "let"},
		/*  1 */ {token.IDENTIFIER, "five"},
		/*  2 */ {token.ASSIGN, "="},
		/*  3 */ {token.INT, "5"},
		/*  4 */ {token.SEMICOLON, ";"},

		/*  5 */ {token.LET, "let"},
		/*  6 */ {token.IDENTIFIER, "ten"},
		/*  7 */ {token.ASSIGN, "="},
		/*  8 */ {token.INT, "10"},
		/*  9 */ {token.SEMICOLON, ";"},

		/* 10 */ {token.LET, "let"},
		/* 11 */ {token.IDENTIFIER, "add"},
		/* 12 */ {token.ASSIGN, "="},
		/* 13 */ {token.FUNC, "func"},
		/* 14 */ {token.LPAREN, "("},
		/* 15 */ {token.IDENTIFIER, "x"},
		/* 16 */ {token.COMMA, ","},
		/* 17 */ {token.IDENTIFIER, "y"},
		/* 18 */ {token.RPAREN, ")"},
		/*    */ // Func body
		/* 19 */ {token.LBRACE, "{"},
		/* 20 */ {token.IDENTIFIER, "x"},
		/* 21 */ {token.PLUS, "+"},
		/* 22 */ {token.IDENTIFIER, "y"},
		/* 23 */ {token.SEMICOLON, ";"},
		/* 24 */ {token.RBRACE, "}"},

		/* 26 */ {token.LET, "let"},
		/* 27 */ {token.IDENTIFIER, "result"},
		/* 28 */ {token.ASSIGN, "="},
		/* 29 */ {token.IDENTIFIER, "add"},
		/* 30 */ {token.LPAREN, "("},
		/* 31 */ {token.IDENTIFIER, "five"},
		/* 32 */ {token.COMMA, ","},
		/* 33 */ {token.IDENTIFIER, "ten"},
		/* 34 */ {token.RPAREN, ")"},
		/* 35 */ {token.SEMICOLON, ";"},
		/* 36 */ {token.EOF, ""},
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
