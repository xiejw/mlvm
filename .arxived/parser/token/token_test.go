package token

import (
	"testing"
)

func TestTokenString(t *testing.T) {
	expects := []struct {
		str       string
		tokenType TokenType
	}{
		{"Lparen    ", Lparen},
		{"Rparen    ", Rparen},
		{"Lbrack    ", Lbrack},
		{"Rbrack    ", Rbrack},
		{"Bslash    ", Bslash},
		{"Int       ", Int},
		{"Float     ", Float},
		{"String    ", String},
		{"Id        ", Id},
		{"Illegal   ", Illegal},
		{"Eof       ", Eof},
	}

	for _, expected := range expects {
		got := expected.tokenType.String()
		if expected.str != got {
			t.Errorf("token type string mismatch. expected: `%v`, got: `%v`", expected.str, got)
		}
	}
}
