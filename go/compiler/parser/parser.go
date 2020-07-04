package parser

import (
	"fmt"

	"github.com/xiejw/mlvm/go/compiler/lexer"
	"github.com/xiejw/mlvm/go/compiler/token"
)

type Parser struct {
	option Option
	l      *lexer.Lexer
}

type Option struct {
	Trace bool
}

func New(input []byte, option Option) *Parser {
	return &Parser{
		option: option,
		l:      lexer.New(input),
	}
}

func (p *Parser) Loop() {

	for {
		tok := p.l.NextToken()
		if p.option.Trace {
			fmt.Printf("token: %+v\n", tok)
		}
		if tok.Type == token.EOF {
			break
		}
	}
}
