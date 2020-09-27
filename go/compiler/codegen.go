package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/ir"
	"github.com/xiejw/mlvm/go/object"
)

func Compile(m *ir.Module) (*code.Program, *errors.DError) {
	fns := m.Fns()
	if len(fns) != 1 {
		return nil, errors.New("Compile only supports one fn in module now.")
	}

	fn := fns[0]
	if fn.Name() != "main" {
		return nil, errors.New("Compile only supports one fn called `main` in module now.")
	}

	return codeGen(fn)
}

func codeGen(fn *ir.Fn) (*code.Program, *errors.DError) {

	insts := make([]byte, 0)
	consts := make([]object.Object, 0)

	value_loader := make(map[ir.Result]func() ([]byte, *errors.DError))

	const_loader := func(c_index int) func() ([]byte, *errors.DError) {
		return func() ([]byte, *errors.DError) {
			ins, err := code.MakeOp(code.OpCONST, c_index)
			if err != nil {
				return nil, errors.From(err)
			}
			return ins, nil
		}
	}

	for _, ins := range fn.Insts() {
		switch v := ins.(type) {
		case *ir.IntLiteral:
			c := &object.Integer{Value: v.Value}
			index := len(consts)
			consts = append(consts, c)
			value_loader[*v.GetResult().(*ir.Result)] = const_loader(index)

		case *ir.Return:
			operand := v.GetOperand().(*ir.Result)
			ins, err := value_loader[*operand]()
			if err != nil {
				return nil, err
			}
			insts = append(insts, ins...)

		default:
			panic(fmt.Sprintf("unsupported ins type for codegen: %v", v))
		}
	}

	return &code.Program{insts, consts}, nil
}
