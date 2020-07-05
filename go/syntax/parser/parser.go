package parser

import (
	"log"

	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/syntax/lexer"
	"github.com/xiejw/mlvm/go/syntax/token"
)

type Parser struct {
	option *Option

	l         *lexer.Lexer
	curToken  *token.Token
	peekToken *token.Token
}

type Option struct {
	Trace bool
}

func New(input []byte) *Parser {
	return NewWithOption(input, &Option{})
}

func NewWithOption(input []byte, option *Option) *Parser {
	p := &Parser{
		option: option,
		l:      lexer.New(input),
	}
	// Fills curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseAst() *ast.Program {
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
	return nil
}
