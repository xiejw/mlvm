package compiler

import (
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func (b *Builder) emitIntegerConstant(literal *ast.IntegerLiteral) int {
	var o object.Object
	o = &object.Integer{literal.Value}
	index := len(b.output.Constants)
	b.output.Constants = append(b.output.Constants, o)
	return index
}

func (b *Builder) emitStringConstant(literal *ast.StringLiteral) int {
	var o object.Object
	o = &object.String{literal.Value}
	index := len(b.output.Constants)
	b.output.Constants = append(b.output.Constants, o)
	return index
}

func (b *Builder) emitLoadConstant(constIndex int) {
	ins, err := code.MakeOp(code.OpConstant, constIndex)
	if err != nil {
		panic(err)
	}
	b.output.Instructions = append(b.output.Instructions, ins...)
}

func (b *Builder) emitLoadTensor() {
	ins, err := code.MakeOp(code.OpLoadT)
	if err != nil {
		panic(err)
	}
	b.output.Instructions = append(b.output.Instructions, ins...)
}

func (b *Builder) emitPop() {
	ins, err := code.MakeOp(code.OpPop)
	if err != nil {
		panic(err)
	}
	b.output.Instructions = append(b.output.Instructions, ins...)
}
