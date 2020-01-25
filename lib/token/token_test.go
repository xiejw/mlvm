package token

import (
	"testing"
)

func TestTokenEqualness(t *testing.T) {
	tok1 := Token{Type: PLUS, Literal: "+"}
	tok2 := Token{Type: PLUS, Literal: "+"}
	tok3 := Token{Type: ASSIGN, Literal: "="}

	if tok1 != tok2 {
		t.Errorf("Expected equal tokens")
	}
	if tok1 == tok3 {
		t.Errorf("Not expected equal tokens")
	}
}
