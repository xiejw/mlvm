package parser

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"

	"github.com/xiejw/mlvm/go/base/errors"
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

func (p *Parser) ParseAst() (*ast.Program, *errors.DiagnosisError) {
	program := &ast.Program{}
	expressions := make([]ast.Expr, 0)
	defer func() { p.level -= 1 }()

	for p.curToken.Type != token.Eof {
		expr, err := p.parseExpression()

		if err != nil {
			index := len(expressions)
			return nil, err.EmitDiagnosisNote(
				"during parsing ast, at %v-th expression", index+1)
		}

		expressions = append(expressions, expr)
	}

	program.Exprs = expressions
	return program, nil
}

func (p *Parser) parseExpression() (ast.Expr, *errors.DiagnosisError) {
	if p.option.TraceParser {
		p.logParserTracing("Expr")
	}

	switch p.curToken.Type {
	case token.Lparen:
		p.level += 1
		defer func() { p.level -= 1 }()
		return p.parseFunctionCallExpression()
	case token.Id:
		return p.parseIdentifider()
	case token.Int:
		return p.parseInteger()
	case token.String:
		return p.parseString()
	case token.Lbrack:
		return p.parseArray()
	default:
		err := errors.NewDiagnosisError(
			"unsupported starting token to be parsed for expression: %v", p.curToken)
		return nil, err.EmitDiagnosisNote(
			"supported starting token for expressions are: " +
				"function call, identifier, integer literal, string literal, array literal.")
	}
}

func (p *Parser) parseFunctionCallExpression() (
	ast.Expr, *errors.DiagnosisError,
) {
	fc := &ast.App{}

	startPos := p.curToken.Loc.Position
	var endPos uint

	if p.option.TraceParser {
		p.logParserTracing("FunctionCallExpression")
		defer func() {
			p.logParserTracing("FunctionCallExpression %v: %v",
				color.YellowString("source"), string(p.l.Bytes(startPos, endPos)))
			p.logParserTracing("FunctionCallExpression %v: %v",
				color.YellowString("result"), ast.Exprs([]ast.Expr{fc}))
		}()
	}

	err := p.parseSingleTokenWithType(token.Lparen)
	if err != nil {
		return nil, err.EmitDiagnosisNote(
			"matching starting Lparen for App expression")
	}

	// TODO: Supports `fn`
	switch p.curToken.Type {
	case token.Id:
		id, err := p.parseIdentifider()
		if err != nil {
			return nil, err.EmitDiagnosisNote("parsing function with identifier")
		}
		fc.Func = id
	default:
		return nil, errors.NewDiagnosisError(
			"unsupported function. currently only identifier is supported. got: %v",
			p.curToken)
	}

	var args []ast.Expr
	// Args
	for p.curToken.Type != token.Rparen {
		arg, err := p.parseExpression()
		if err != nil {
			return nil, err.EmitDiagnosisNote(
				"parsing the %v-th function argument", len(args)+1)
		}
		args = append(args, arg)
	}
	fc.Args = args

	endPos = p.curToken.Loc.Position + 1
	err = p.parseSingleTokenWithType(token.Rparen)
	if err != nil {
		return nil, err.EmitDiagnosisNote(
			"matching ending Rparen for App expression")
	}
	return fc, nil
}

func (p *Parser) parseIdentifider() (*ast.Id, *errors.DiagnosisError) {
	if p.option.TraceParser {
		p.logParserTracing("Id")
	}

	if !p.isCurrentTokenType(token.Id) {
		return nil, errors.NewDiagnosisError(
			"expected to match a token exactly with Id type, but got: %v",
			p.curToken)
	}

	id := &ast.Id{Value: p.curToken.Literal}
	p.advanceToken()
	return id, nil
}

func (p *Parser) parseInteger() (*ast.IntLit, *errors.DiagnosisError) {
	if p.option.TraceParser {
		p.logParserTracing("Integer")
	}

	if !p.isCurrentTokenType(token.Int) {
		return nil, errors.NewDiagnosisError(
			"expected to match a token exactly with Int type, but got: %v",
			p.curToken)
	}

	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		return nil, errors.EmitDiagnosisNote(
			err, "parsing integer expression for literal: %v",
			p.curToken.Literal)
	}

	i := &ast.IntLit{Value: v}
	p.advanceToken()
	return i, nil
}

func (p *Parser) parseString() (*ast.StringLit, *errors.DiagnosisError) {
	if p.option.TraceParser {
		p.logParserTracing("String")
	}

	if !p.isCurrentTokenType(token.String) {
		return nil, errors.NewDiagnosisError(
			"expected to match a token exactly with String type, but got: %v",
			p.curToken)
	}

	rawLiteral := p.curToken.Literal

	// Unwraps the `"abc"` and puts `abc` into StringLit.
	s := &ast.StringLit{Value: rawLiteral[1 : len(rawLiteral)-1]}
	p.advanceToken()
	return s, nil
}

func (p *Parser) parseArray() (*ast.ArrayLit, *errors.DiagnosisError) {
	var err *errors.DiagnosisError
	if p.option.TraceParser {
		p.logParserTracing("Array")

		// defer func() {
		// 	p.logParserTracing("ArrayLiteral %v: %v",
		// 		color.YellowString("source"), string(p.l.Bytes(startPos, endPos)))
		// 	p.logParserTracing("ArrayLiteral %v: %v",
		// 		color.YellowString("result"), ast.Exprs([]ast.Expr{fc}))
		// }()
	}

	err = p.parseSingleTokenWithType(token.Lbrack)
	if err != nil {
		return nil, err.EmitDiagnosisNote(
			"matching starting Lbrack for Array literal expression")
	}

	err = p.parseSingleTokenWithType(token.Rbrack)
	if err != nil {
		return nil, err.EmitDiagnosisNote(
			"matching ending Rparen for App expression")
	}

	return nil, nil
}

func (p *Parser) parseSingleTokenWithType(t token.TokenType) *errors.DiagnosisError {
	if p.option.TraceParser {
		p.logParserTracing("Token with type: %v", t)
	}

	if !p.isCurrentTokenType(t) {
		return errors.NewDiagnosisError(
			"expected to match a token exactly with specific type: %v, but got: %v",
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
		log.Printf(
			"%vTracer: Lexer: next token: %+v\n", levelToIndent(p.level), p.curToken)
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
	log.Printf("%vTracer: %v: %v",
		levelToIndent(p.level), color.GreenString("Parser"), line)
}
