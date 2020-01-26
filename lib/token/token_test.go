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

func TestLookupKeywords(t *testing.T) {
	tokenType := LookupIdentifier("hello")
	expected := IDENTIFIER
	if string(tokenType) != expected {
		t.Errorf("Expect type %q, got %q", expected, tokenType)
	}

	tokenType = LookupIdentifier("func")
	expected = FUNC
	if string(tokenType) != expected {
		t.Errorf("Expect type %q, got %q", expected, tokenType)
	}
}
