package compiler

import (
	"fmt"
	"log"

	"github.com/xiejw/mlvm/go/base/errors"
	"github.com/xiejw/mlvm/go/code"
	"github.com/xiejw/mlvm/go/ir"
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

	for _, ins := range fn.Insts() {
		switch t := ins.(type) {
		case *ir.IntLiteral:
			log.Printf("found intliteral")
		case *ir.Return:
			log.Printf("found return")

		default:
			panic(fmt.Sprintf("unsupported ins type for codegen: %v", t))
		}
	}

	return code.NewProgram(), nil
}
