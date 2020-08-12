package parser

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/syntax/token"
)

func (p *Parser) parseExpr() (ast.Expr, *errors.DError) {
	if p.option.TraceParser {
		p.logP("Expr")
	}

	switch p.curToken.Type {
	case token.Lparen:
		p.level += 1
		defer func() { p.level -= 1 }()
		return p.parseApp()

	case token.Id:
		return p.parseId()
	case token.Int:
		return p.parseIntLit()
	case token.String:
		return p.parseStringLit()
	case token.Lbrack:
		return p.parseArrayLit()

	default:
		return nil, errors.New(
			"unsupported starting token for Expr: %v", p.curToken).EmitNote(
			"supported starting token for Expr are: " +
				"App, Id, IntLiti, StringLit, ArrayLit.")
	}
}

// Parses applicaiton (function call).
func (p *Parser) parseApp() (ast.Expr, *errors.DError) {
	fc := &ast.App{}

	startPos := p.curToken.Loc.Position
	var endPos uint

	if p.option.TraceParser {
		p.logP("App")
		defer func() {
			p.logP("App %v: %v",
				color.YellowString("source"), string(p.l.Bytes(startPos, endPos)))
			p.logP("App %v: %v",
				color.YellowString("result"), ast.Exprs([]ast.Expr{fc}))
		}()
	}

	if err := p.parseSingleTokenWithType(token.Lparen); err != nil {
		return nil, err.EmitNote("matching starting Lparen for App")
	}

	// TODO: Supports `fn`
	switch p.curToken.Type {
	case token.Id:
		id, err := p.parseId()
		if err != nil {
			return nil, err.EmitNote("parsing function with identifier")
		}
		fc.Func = id
	default:
		return nil, errors.New(
			"unsupported function. currently only identifier is supported. got: %v",
			p.curToken)
	}

	var args []ast.Expr
	// Args
	for p.curToken.Type != token.Rparen {
		arg, err := p.parseExpr()
		if err != nil {
			return nil, err.EmitNote(
				"parsing the %v-th function argument", len(args)+1)
		}
		args = append(args, arg)
	}
	fc.Args = args

	endPos = p.curToken.Loc.Position + 1
	err := p.parseSingleTokenWithType(token.Rparen)
	if err != nil {
		return nil, err.EmitNote(
			"matching ending Rparen for App expression")
	}
	return fc, nil
}

func (p *Parser) parseId() (*ast.Id, *errors.DError) {
	if p.option.TraceParser {
		p.logP("Id")
	}

	if !p.isCurrentTokenType(token.Id) {
		return nil, errors.New(
			"expected to match a token exactly with Id type, but got: %v",
			p.curToken)
	}

	id := &ast.Id{Value: p.curToken.Literal}
	p.advanceToken()
	return id, nil
}

func (p *Parser) parseIntLit() (*ast.IntLit, *errors.DError) {
	if p.option.TraceParser {
		p.logP("Integer")
	}

	if !p.isCurrentTokenType(token.Int) {
		return nil, errors.New(
			"expected to match a token exactly with Int type, but got: %v",
			p.curToken)
	}

	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		return nil, errors.From(err).EmitNote(
			"parsing integer expression for literal: %v",
			p.curToken.Literal)
	}

	i := &ast.IntLit{Value: v}
	p.advanceToken()
	return i, nil
}

func (p *Parser) parseStringLit() (*ast.StringLit, *errors.DError) {
	if p.option.TraceParser {
		p.logP("String")
	}

	if !p.isCurrentTokenType(token.String) {
		return nil, errors.New(
			"expected to match a token exactly with String type, but got: %v",
			p.curToken)
	}

	rawLiteral := p.curToken.Literal

	// Unwraps the `"abc"` and puts `abc` into StringLit.
	s := &ast.StringLit{Value: rawLiteral[1 : len(rawLiteral)-1]}
	p.advanceToken()
	return s, nil
}

func (p *Parser) parseArrayLit() (*ast.ArrayLit, *errors.DError) {
	var err *errors.DError
	if p.option.TraceParser {
		p.logP("Array")

		// defer func() {
		// 	p.logP("ArrayLiteral %v: %v",
		// 		color.YellowString("source"), string(p.l.Bytes(startPos, endPos)))
		// 	p.logP("ArrayLiteral %v: %v",
		// 		color.YellowString("result"), ast.Exprs([]ast.Expr{fc}))
		// }()
	}

	err = p.parseSingleTokenWithType(token.Lbrack)
	if err != nil {
		return nil, err.EmitNote(
			"matching starting Lbrack for Array literal expression")
	}

	err = p.parseSingleTokenWithType(token.Rbrack)
	if err != nil {
		return nil, err.EmitNote(
			"matching ending Rparen for App expression")
	}

	return nil, nil
}

func (p *Parser) parseSingleTokenWithType(t token.TokenType) *errors.DError {
	if p.option.TraceParser {
		p.logP("Token with type: %v", t)
	}

	if !p.isCurrentTokenType(t) {
		return errors.New(
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

func (p *Parser) logP(sfmt string, args ...interface{}) {
	line := fmt.Sprintf(sfmt, args...)
	log.Printf("%vTracer: %v: %v",
		levelToIndent(p.level), color.GreenString("Parser"), line)
}
