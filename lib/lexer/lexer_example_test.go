package lexer

import (
	"testing"

	"github.com/xiejw/mlvm/lib/token"
)

func TestLexerWithOneFullExample(t *testing.T) {
	input := `
let x = 3;
let y = 4;
let gt = func(x, y) {
    if (x > y) {
        return true;
    } else {
        return false;
    }
}
let result = gt(3, 4);
`
	l := New(input)

	i := 0
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		i += 1
	}
	expected := 5 + 5 + 10 + 7 + 3 + 3 + 3 + 1 + 1 + 10
	if i != expected {
		t.Errorf("Token count mismatch: expected %v got %v", expected, i)
	}
}
