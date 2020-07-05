package parser

import (
	"fmt"
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
	p.advanceToken()
	p.advanceToken()
	return p
}

func (p *Parser) ParseAst() (*ast.Program, error) {
	program := &ast.Program{}
	expressions := make([]ast.Expression, 0)

	for p.curToken.Type != token.EOF {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		expressions = append(expressions, expr)
		p.advanceToken()
	}

	program.Expressions = expressions
	return program, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	switch p.curToken.Type {
	case token.LPAREN:
		return p.parseFunctionCallExpression()
	case token.IDENTIFIER:
		return p.parseIdentifider()
	default:
		return nil, fmt.Errorf("unknown expression token: %v", p.curToken)
	}
}

func (p *Parser) parseFunctionCallExpression() (ast.Expression, error) {
	err := p.parseSingleTokenWithType(token.LPAREN)
	if err != nil {
		return nil, err
	}

	fc := &ast.FunctionCall{}

	// Supports `fn`
	id, err := p.parseIdentifider()
	fc.Func = id

	err = p.parseSingleTokenWithType(token.RPAREN)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p *Parser) parseIdentifider() (*ast.Identifier, error) {
	if !p.isCurrentTokenType(token.IDENTIFIER) {
		return nil, fmt.Errorf("expected to see token type: %v, got: %v",
			token.IDENTIFIER, p.curToken)
	}

	id := &ast.Identifier{Value: p.curToken.Literal}
	p.advanceToken()
	return id, nil
}

func (p *Parser) parseSingleTokenWithType(t token.TokenType) error {
	if !p.isCurrentTokenType(t) {
		return fmt.Errorf("expected to see token type: %v, got: %v",
			t, p.curToken)
	}

	p.advanceToken()
	return nil
}

func (p *Parser) isCurrentTokenType(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) isPeekTokenType(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) advanceToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()

	if p.option.Trace && p.curToken != nil {
		log.Printf("trace parser: current token: %+v\n", p.curToken)
	}
}
