package parser

import (
	"log"

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
			log.Printf("trace parser: got next token: %+v\n", tok)
		}
		if tok.Type == token.EOF {
			break
		}
		if tok.Type == token.ILLEGAL {
			log.Fatalf("illegal: %+v", tok)
			break
		}
	}
}
