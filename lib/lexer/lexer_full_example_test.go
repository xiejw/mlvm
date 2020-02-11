package lexer

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/lib/token"
)

func TestLexerWithOneFullExample(t *testing.T) {
	input := `
let x = 3 / 1 - 10 ;
let y = 4 + 5 * 2;
let gt = func(x, y) {
    if (x > y) {
        return true;
    } else {
        return false;
    }
}
let result = !gt(3, 4);
let expected = 3 < 4;
`
	l := New(input)

	expectedNumTokensPerLine := []int{0, 0, 9, 9, 10, 7, 3, 3, 3, 1, 1, 11, 7} // Line is 1-based.
	gotNumTokensPerLine := make([]int, len(expectedNumTokensPerLine))
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		gotNumTokensPerLine[tok.Loc.L]++
	}
	if !reflect.DeepEqual(expectedNumTokensPerLine, gotNumTokensPerLine) {
		t.Errorf("Token count mismatch: expected %v got %v",
			expectedNumTokensPerLine, gotNumTokensPerLine)
	}
}
