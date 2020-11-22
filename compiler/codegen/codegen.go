package codegen

import (
	"fmt"

	"github.com/xiejw/mlvm/compiler/ir"
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/code"
	"github.com/xiejw/mlvm/vm/object"
)

// Compiles the source module to program.
//
// The entry point is the `main` fn.
func Compile(m *ir.Module) (*code.Program, error) {
	fns := m.Fns()
	if len(fns) != 1 {
		return nil, errors.New("Compile only supports one fn in module now.")
	}

	fn := fns[0]
	if fn.Name() != "main" {
		return nil, errors.New("Compile requires at least one fn called `main` as entry point.")
	}
	p, err := codeGen(fn)
	if err != nil {
		return nil, errors.From(err).EmitNote("fn:\n\n%v", fn).EmitNote("compile program error")
	}
	return p, nil
}

//-----------------------------------------------------------------------------
// Loader.
//-----------------------------------------------------------------------------

type LoaderFn func() ([]byte, error)

func constLoaderFn(c_index int) LoaderFn {
	return func() ([]byte, error) {
		ins, err := code.MakeOp(code.OpCONST, c_index)
		if err != nil {
			return nil, err
		}
		return ins, nil
	}
}

func memLoaderFn(m_index int) LoaderFn {
	return func() ([]byte, error) {
		ins, err := code.MakeOp(code.OpLOAD, m_index)
		if err != nil {
			return nil, err
		}
		return ins, nil
	}
}

//-----------------------------------------------------------------------------
// Codegen.
//-----------------------------------------------------------------------------

// Generates the code for `fn`.
func codeGen(fn *ir.Fn) (*code.Program, error) {
	mem_slot_i := 0                             // points to next available index.
	value_loader := make(map[ir.Value]LoaderFn) // ir.Value to loader mapping.

	insts := make([]byte, 0)
	consts := make([]object.Object, 0)

	for _, ins := range fn.Instructions() {

		switch v := ins.(type) {
		case *ir.IntLiteral:
			c := &object.Integer{Value: v.Value}
			index := len(consts)
			consts = append(consts, c)
			value_loader[v.GetResult()] = constLoaderFn(index)

		case *ir.ShapeLiteral:
			s := object.NewShape(v.Dims)
			index := len(consts)
			consts = append(consts, s)
			value_loader[v.GetResult()] = constLoaderFn(index)

		case *ir.ArrayLiteral:
			s := &object.Array{v.Value}
			index := len(consts)
			consts = append(consts, s)
			value_loader[v.GetResult()] = constLoaderFn(index)

		case *ir.NewTensor:
			loadValueToInsts(&insts, v.Shape, value_loader)
			loadValueToInsts(&insts, v.Array, value_loader)
			appendOpcode(&insts, code.OpT)
			value_loader[v.Result] = memLoaderFn(mem_slot_i)
			pushToMemoryAndIncrIndex(&insts, &mem_slot_i)

		case *ir.RngSource:
			loadValueToInsts(&insts, v.Seed, value_loader)
			appendOpcode(&insts, code.OpRNG)
			value_loader[v.Result] = memLoaderFn(mem_slot_i)
			pushToMemoryAndIncrIndex(&insts, &mem_slot_i)

		case *ir.RngFill:
			loadValueToInsts(&insts, v.Shape, value_loader)
			loadValueToInsts(&insts, v.Source, value_loader)
			appendOpcode(&insts, code.OpRNGT, 0)
			value_loader[v.Result] = memLoaderFn(mem_slot_i)
			pushToMemoryAndIncrIndex(&insts, &mem_slot_i)

		case *ir.Return:
			loadValueToInsts(&insts, v.GetOperand(), value_loader)

		default:
			panic(fmt.Sprintf("codegen: unsupported instruction type: %v", v)) // internal bug.
		}
	}

	return &code.Program{insts, consts}, nil
}

// -----------------------------------------------------------------------------
// Helper methods
// -----------------------------------------------------------------------------

func loadValueToInsts(insts *[]byte, v ir.Value, value_loader map[ir.Value]LoaderFn) {
	loader, existed := value_loader[v]
	if !existed {
		panic(fmt.Sprintf("value loader for result (%v) does not exist.", v))
	}

	ins, err := loader()
	if err != nil {
		panic(err)
	}
	*insts = append(*insts, ins...)
}

func appendOpcode(insts *[]byte, c code.Opcode, args ...int) {
	ins, err := code.MakeOp(c, args...)
	if err != nil {
		panic(err)
	}
	*insts = append(*insts, ins...)
}

func pushToMemoryAndIncrIndex(insts *[]byte, m_index *int) {
	ins, err := code.MakeOp(code.OpSTORE, *m_index)
	if err != nil {
		panic(err)
	}
	*m_index++
	*insts = append(*insts, ins...)
}
