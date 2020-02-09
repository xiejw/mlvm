package token

import (
	"testing"
)

func TestTokenEqualness(t *testing.T) {
	loc := Loc{1, 2}
	tok1 := Token{Type: PLUS, Literal: "+", Loc: loc}
	tok2 := Token{Type: PLUS, Literal: "+", Loc: loc}
	tok3 := Token{Type: ASSIGN, Literal: "=", Loc: loc}

	if tok1 != tok2 {
		t.Errorf("Expected equal tokens")
	}
	if tok1 == tok3 {
		t.Errorf("Not expected equal tokens")
	}
}

func TestLookupNonKeywords(t *testing.T) {
	tokenType := LookupIdentifier("hello")
	expected := IDENTIFIER
	if string(tokenType) != expected {
		t.Errorf("Expect type %q, got %q", expected, tokenType)
	}
}

func TestLookupKeywords(t *testing.T) {
	expectedKeywords := []struct {
		literal   string
		tokenType TokenType
	}{
		{"func", FUNC},
		{"let", LET},
		{"return", RETURN},
		{"if", IF},
		{"else", ELSE},
		{"true", TRUE},
		{"false", FALSE},
	}

	for _, expected := range expectedKeywords {
		got := LookupIdentifier(expected.literal)
		if got != expected.tokenType {
			t.Errorf("Expect type %q, got %q", expected, got)
		}
	}
}
