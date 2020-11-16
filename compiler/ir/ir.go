package ir

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/compiler/base/errors"
)

// -----------------------------------------------------------------------------
// Value
// -----------------------------------------------------------------------------

type Value interface {
	Type() *Type
	valueInterface()
	String() string
}

type Result struct {
	Name     string
	T        *Type
	Src      Instruction
	SrcIndex int
}

// -- Conform Value
func (r *Result) valueInterface() {}
func (r *Result) Type() *Type     { return r.T }
func (r *Result) String() string  { return r.Name }

// -----------------------------------------------------------------------------
// Instruction
// -----------------------------------------------------------------------------

type Instruction interface {
	GetOperand() Value    // Could be nil
	GetOperands() []Value // Could be nil
	GetResult() Value
	GetResults() []Value
	String() string
	Check() *errors.DError
}

// See instructions.go for all implementations.

// -----------------------------------------------------------------------------
// Fn
// -----------------------------------------------------------------------------

type Fn struct {
	b         *Builder      // Parent builder.
	name      string        // Func name.
	res_index int           // Next result index.
	insts     []Instruction // Instructions
}

func (f *Fn) Name() string {
	return f.name
}

func (f *Fn) Instructions() []Instruction {
	return f.insts
}

func (f *Fn) IntLiteral(v int64) *IntLiteral {
	ins := &IntLiteral{
		Value:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, IntType)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) ShapeLiteral(dims []int) *ShapeLiteral {
	ins := &ShapeLiteral{
		Dims:   dims,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KShape, Dims: dims})
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngSource(v Value) *RngSource {
	ins := &RngSource{
		Input:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, RngType)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngTensor(src Value, s Value) *RngTensor {
	ins := &RngTensor{
		Source: src,
		Shape:  s,
		Result: nil,
	}
	dims := s.Type().Dims // could be nil
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KTensor, Dims: dims})
	f.insts = append(f.insts, ins)
	return ins
}

// -- Output

func (f *Fn) SetOutputAndDone(v Value) {
	f.insts = append(f.insts, &Return{v})
	f.Done()
}

// -- Helper Methods.

func (f *Fn) nextResult(src Instruction, outputIndex int, t *Type) *Result {
	i := f.res_index
	f.res_index++
	return &Result{
		Name:     fmt.Sprintf("%%%v", i),
		T:        t,
		Src:      src,
		SrcIndex: outputIndex,
	}
}

func (f *Fn) Done() {
	f.b.fns = append(f.b.fns, f)
	f.b.f_map[f.name] = f
}

// -----------------------------------------------------------------------------
// Module
// -----------------------------------------------------------------------------

type Module struct {
	fns []*Fn
}

func (m *Module) Fns() []*Fn {
	return m.fns
}

// -----------------------------------------------------------------------------
// Builder (Producing Module)
// -----------------------------------------------------------------------------

type Builder struct {
	fns   []*Fn
	f_map map[string]*Fn
}

func NewBuilder() *Builder {
	return &Builder{
		fns:   nil,
		f_map: make(map[string]*Fn),
	}
}

func (b *Builder) NewFn(fn_name string) (*Fn, *errors.DError) {
	if _, existed := b.f_map[fn_name]; existed {
		return nil, errors.New("fn name already existed in module: %v", fn_name)
	}
	return &Fn{b: b, name: fn_name}, nil
}

func (b *Builder) Done() (*Module, *errors.DError) {
	for _, fn := range b.fns {
		for _, ins := range fn.insts {
			err := ins.Check()
			if err != nil {
				return nil, err.EmitNote(
					"Failed to check Instrunction: %v", ins,
				).EmitNote(
					"Failed to check the fn (`%v`):\n\n%v", fn.Name(), fn,
				)
			}
		}
	}
	return &Module{fns: b.fns}, nil
}

// -----------------------------------------------------------------------------
// All String Helpers
// -----------------------------------------------------------------------------

func (f *Fn) DebugString(w io.Writer) {
	fmt.Fprintf(w, "fn %v() {\n", f.name)
	for _, ins := range f.insts {
		// prints ins with type.
		fmt.Fprintf(w, "  %-40s:: %v\n", ins.String(), ins.GetResult().Type())
	}
	fmt.Fprintf(w, "}\n")
}

func (f *Fn) String() string {
	buf := bytes.Buffer{}
	f.DebugString(&buf)
	return buf.String()
}

func (m *Module) DebugString(w io.Writer) {
	fmt.Fprintf(w, "module {\n")
	for _, f := range m.fns {
		fmt.Fprintf(w, "\n%v", f)
	}
	fmt.Fprintf(w, "\n}\n")
}

func (m *Module) String() string {
	buf := bytes.Buffer{}
	m.DebugString(&buf)
	return buf.String()
}
