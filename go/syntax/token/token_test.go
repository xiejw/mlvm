package token

import (
	"testing"
)

func TestTokenString(t *testing.T) {
	expects := []struct {
		str       string
		tokenType TokenType
	}{
		{"LPAREN    ", LPAREN},
		{"RPAREN    ", RPAREN},
		{"LSBRACKET ", LSBRACKET},
		{"RSBRACKET ", RSBRACKET},
		{"BACKSLASH ", BACKSLASH},
		{"INTEGER   ", INTEGER},
		{"FLOAT     ", FLOAT},
		{"STRING    ", STRING},
		{"IDENTIFIER", IDENTIFIER},
		{"ILLEGAL   ", ILLEGAL},
		{"EOF       ", EOF},
	}

	for _, expected := range expects {
		got := expected.tokenType.String()
		if expected.str != got {
			t.Errorf("token type string mismatch. expected: `%v`, got: `%v`", expected.str, got)
		}
	}
}
