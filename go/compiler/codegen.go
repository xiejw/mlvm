package compiler

import (
	"fmt"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/ir"
	"github.com/xiejw/mlvm/go/object"
)

// Compiles the source module to program.
//
// The entry point is the `main` fn.
func Compile(m *ir.Module) (*code.Program, *errors.DError) {
	fns := m.Fns()
	if len(fns) != 1 {
		return nil, errors.New("Compile only supports one fn in module now.")
	}

	fn := fns[0]
	if fn.Name() != "main" {
		return nil, errors.New("Compile requires at least one fn called `main` as entry point.")
	}
	return codeGen(fn)
}

//-----------------------------------------------------------------------------
// Helper Methods
//-----------------------------------------------------------------------------

type LoaderFn func() ([]byte, *errors.DError)

func constLoaderFn(c_index int) LoaderFn {
	return func() ([]byte, *errors.DError) {
		ins, err := code.MakeOp(code.OpCONST, c_index)
		if err != nil {
			return nil, errors.From(err)
		}
		return ins, nil
	}
}

// Code gens the `fn`.
func codeGen(fn *ir.Fn) (*code.Program, *errors.DError) {
	insts := make([]byte, 0)
	consts := make([]object.Object, 0)

	value_loader := make(map[ir.Result]LoaderFn)

	for _, ins := range fn.Insts() {
		switch v := ins.(type) {

		case *ir.IntLiteral:
			c := &object.Integer{Value: v.Value}
			index := len(consts)
			consts = append(consts, c)
			value_loader[*v.GetResult().(*ir.Result)] = constLoaderFn(index)

		case *ir.RngSeed:
			int_lit := v.Input.Result
			ins, err := value_loader[*int_lit]()
			if err != nil {
				return nil, err
			}
			insts = append(insts, ins...)

			ins, err_1 := code.MakeOp(code.OpRNG)
			if err_1 != nil {
				return nil, errors.From(err_1)
			}
			insts = append(insts, ins...)

		case *ir.Return:
			operand := v.GetOperand().(*ir.Result)
			loader, existed := value_loader[*operand]
			if !existed {
				panic(fmt.Sprintf("value loader for result (%v) does not exist.", operand))
			}

			ins, err := loader()
			if err != nil {
				return nil, err
			}
			insts = append(insts, ins...)

		default:
			panic(fmt.Sprintf("unsupported ins type for codegen: %v", v)) // internal bug.
		}
	}

	return &code.Program{insts, consts}, nil
}
