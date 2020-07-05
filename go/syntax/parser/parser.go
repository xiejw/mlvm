package parser

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"

	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/syntax/lexer"
	"github.com/xiejw/mlvm/go/syntax/token"
)

type Parser struct {
	option *Option

	l         *lexer.Lexer
	curToken  *token.Token
	peekToken *token.Token
	level     int
}

type Option struct {
	TraceLexer  bool
	TraceParser bool
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
	defer func() { p.level -= 1 }()

	for p.curToken.Type != token.EOF {
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		expressions = append(expressions, expr)
	}

	program.Expressions = expressions
	return program, nil
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	if p.option.TraceParser {
		p.logParserTracing("Expression")
	}

	switch p.curToken.Type {
	case token.LPAREN:
		p.level += 1
		defer func() { p.level -= 1 }()
		return p.parseFunctionCallExpression()
	case token.IDENTIFIER:
		return p.parseIdentifider()
	case token.INTEGER:
		return p.parseInteger()
	default:
		return nil, fmt.Errorf("unknown expression token: %v", p.curToken)
	}
}

func (p *Parser) parseFunctionCallExpression() (ast.Expression, error) {
	fc := &ast.FunctionCall{}

	if p.option.TraceParser {
		p.logParserTracing("FunctionCallExpression")
		defer func() {
			p.logParserTracing("FunctionCallExpression %v: %v",
				color.YellowString("result"), ast.Expressions([]ast.Expression{fc}))
		}()
	}

	err := p.parseSingleTokenWithType(token.LPAREN)
	if err != nil {
		return nil, err
	}

	// TODO: Supports `fn`
	switch p.curToken.Type {
	case token.IDENTIFIER:
		id, err := p.parseIdentifider()
		if err != nil {
			return nil, err
		}
		fc.Func = id
	default:
		return nil, fmt.Errorf("unsupport function: %v", p.curToken)
	}

	var args []ast.Expression
	// Args
	for p.curToken.Type != token.RPAREN {
		arg, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	fc.Args = args

	err = p.parseSingleTokenWithType(token.RPAREN)
	if err != nil {
		return nil, err
	}
	return fc, nil
}

func (p *Parser) parseIdentifider() (*ast.Identifier, error) {
	if p.option.TraceParser {
		p.logParserTracing("Identifier")
	}

	if !p.isCurrentTokenType(token.IDENTIFIER) {
		return nil, fmt.Errorf("expected to see token type: %v, got: %v",
			token.IDENTIFIER, p.curToken)
	}

	id := &ast.Identifier{Value: p.curToken.Literal}
	p.advanceToken()
	return id, nil
}

func (p *Parser) parseInteger() (*ast.IntegerLiteral, error) {
	if p.option.TraceParser {
		p.logParserTracing("Integer")
	}

	if !p.isCurrentTokenType(token.INTEGER) {
		return nil, fmt.Errorf("expected to see token type: %v, got: %v",
			token.INTEGER, p.curToken)
	}

	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		return nil, err
	}

	i := &ast.IntegerLiteral{Value: v}
	p.advanceToken()
	return i, nil
}

func (p *Parser) parseSingleTokenWithType(t token.TokenType) error {
	if p.option.TraceParser {
		p.logParserTracing("Token with type: %v", t)
	}

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

	if p.option.TraceLexer && p.curToken != nil {
		log.Printf("%vTracer: Lexer: next token: %+v\n", levelToIndent(p.level), p.curToken)
	}
}

func levelToIndent(level int) string {
	switch level {
	case 0:
		return ""
	case 1:
		return "  "
	case 2:
		return "    "
	case 3:
		return "      "
	default:
		return fmt.Sprintf("      (level %v)", level)
	}
}

func (p *Parser) logParserTracing(sfmt string, args ...interface{}) {
	line := fmt.Sprintf(sfmt, args...)
	log.Printf("%vTracer: %v: %v", levelToIndent(p.level), color.GreenString("Parser"), line)
}
