package parser

import (
	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/syntax/ast"
	"github.com/xiejw/mlvm/go/syntax/lexer"
	"github.com/xiejw/mlvm/go/syntax/token"
)

// Parser, which consumes Tokens from input bytes and produces Expr as ast.Program.
type Parser struct {
	option *Option

	l         *lexer.Lexer
	curToken  *token.Token
	peekToken *token.Token
	level     int // records the current depth of the ast level.
}

type Option struct {
	TraceLexer  bool // If true, print tracing info related to lexer.
	TraceParser bool // If true, print tracing info related to parser.
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

// Parse Ast from input bytes.
func (p *Parser) ParseAst() (*ast.Program, *errors.DError) {
	program := &ast.Program{}
	exprs := make([]ast.Expr, 0)

	for p.curToken.Type != token.Eof {
		expr, err := p.parseExpr()
		exprs = append(exprs, expr)

		if err != nil {
			return nil, err.EmitNote(
				"during parsing ast, at %v-th top level expression", len(exprs)+1)
		}
	}

	program.Exprs = exprs
	return program, nil
}
