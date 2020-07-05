package parser

import (
	"log"
	"fmt"

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

	if p.option.Trace && p.curToken != nil {
		log.Printf("trace parser: current token: %+v\n", p.curToken)
	}
}

func (p *Parser) ParseAst() (*ast.Program, error) {
	program := &ast.Program{}
	expressions := make([]ast.Expression, 0)

	for p.curToken.Type != token.EOF {
		expr, err := p.parseFunctionCallExpression()
		if err != nil {
			return nil, err
		}

		expressions = append(expressions, expr)
		p.nextToken()
	}

	program.Expressions = expressions
	return program, nil
}

func (p *Parser) parseFunctionCallExpression() (ast.Expression, error) {
	err := p.consumeTokenType(token.LPAREN)
	if err != nil {
		return nil, err
	}

	fc := &ast.FunctionCall{
	}

	// Supports `fn`
	id, err := p.parseIdentifider()
	fc.Name = id

	err = p.consumeTokenType(token.RPAREN)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Parser) parseIdentifider() (*ast.Identifier, error) {
	err := p.expectTokenType(token.IDENTIFIER)
	if err != nil {
		return nil, err
	}

	id := &ast.Identifier{Value: p.curToken.Literal}
	p.nextToken()
	return id, nil
}

func (p *Parser) consumeTokenType(t token.TokenType) error {
	err := p.expectTokenType(t)
	if err != nil {
		return err
	}
	p.nextToken()
	return nil
}

func (p *Parser) expectTokenType(t token.TokenType) error {
	if p.isCurrentTokenType(t) {
		return fmt.Errorf("expected to see token type: %v, got",
		t, p.curToken)
	}
	return nil
}

func (p *Parser) isCurrentTokenType(t token.TokenType) bool {
	return p.curToken.Type == t
}
