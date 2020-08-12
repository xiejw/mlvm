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
		// For App, we do some heavy work here. 1. logs the entry point 2. logs the exit point.
		p.logP("App")
		defer func() {
			p.logP("App %v: %v",
				color.YellowString("source"), string(p.l.Bytes(startPos, endPos)))
			p.logP("App %v: %v",
				color.YellowString("result"), ast.Exprs([]ast.Expr{fc}))
		}()
	}

	// Consumes the Lparen.
	if err := p.parseSingleTokenWithType(token.Lparen); err != nil {
		return nil, err.EmitNote("matching starting Lparen for App")
	}

	// Consumes the App.Func.  TODO: Supports General Expr returning Fn.
	switch p.curToken.Type {
	case token.Id:
		id, err := p.parseId()
		if err != nil {
			return nil, err.EmitNote("parsing App.Func with Id")
		}
		fc.Func = id
	default:
		return nil, errors.New("unsupported App.Func. Only Id is supported. got: %v", p.curToken)
	}

	// Consumes App.Args.
	var args []ast.Expr
	for p.curToken.Type != token.Rparen {
		arg, err := p.parseExpr()
		args = append(args, arg)
		if err != nil {
			return nil, err.EmitNote("parsing the %v-th of App arg", len(args))
		}
	}
	fc.Args = args

	// Records the final position for tracing.
	endPos = p.curToken.Loc.Position + 1

	// Consumes the Rparen.
	if err := p.parseSingleTokenWithType(token.Rparen); err != nil {
		return nil, err.EmitNote("matching ending Rparen for App")
	}

	return fc, nil
}

// Parses Id.
func (p *Parser) parseId() (*ast.Id, *errors.DError) {
	if p.option.TraceParser {
		p.logP("Id")
	}

	if !p.isCurrentTokenType(token.Id) {
		return nil, errors.New("expected to match a token with Id type, got: %v", p.curToken)
	}

	id := &ast.Id{Value: p.curToken.Literal}
	p.advanceToken()
	return id, nil
}

// Parses IntLit.
func (p *Parser) parseIntLit() (*ast.IntLit, *errors.DError) {
	if p.option.TraceParser {
		p.logP("IntLit")
	}

	if !p.isCurrentTokenType(token.Int) {
		return nil, errors.New("expected to match a token with Int type, got: %v", p.curToken)
	}

	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		return nil, errors.From(err).EmitNote(
			"parsing Int expression for literal: %v", p.curToken.Literal)
	}

	i := &ast.IntLit{Value: v}
	p.advanceToken()
	return i, nil
}

// Parses IntLit.
func (p *Parser) parseStringLit() (*ast.StringLit, *errors.DError) {
	if p.option.TraceParser {
		p.logP("StringLit")
	}

	if !p.isCurrentTokenType(token.String) {
		return nil, errors.New("expected to match a token with String type, got: %v", p.curToken)
	}

	// Unwraps the `"abc"` and puts `abc` into StringLit.
	rawLiteral := p.curToken.Literal
	s := &ast.StringLit{Value: rawLiteral[1 : len(rawLiteral)-1]}
	p.advanceToken()
	return s, nil
}

// Parses ArrayLit.
func (p *Parser) parseArrayLit() (*ast.ArrayLit, *errors.DError) {
	if p.option.TraceParser {
		p.logP("ArrayLit")
	}

	if err := p.parseSingleTokenWithType(token.Lbrack); err != nil {
		return nil, err.EmitNote("matching starting Lbrack for ArrayLit")
	}

	if err := p.parseSingleTokenWithType(token.Rbrack); err != nil {
		return nil, err.EmitNote("matching ending Rbrack for ArrayLit")
	}

	panic("unimplemented.") // return nil, nil
}

func (p *Parser) parseSingleTokenWithType(t token.TokenType) *errors.DError {
	if p.option.TraceParser {
		p.logP("%v", t)
	}

	if !p.isCurrentTokenType(t) {
		return errors.New("expected to match a token with type: %v, got: %v", t, p.curToken)
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
