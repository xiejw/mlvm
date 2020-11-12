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
	valueInterface()
	String() string
}

type Result struct {
	Name string
	Src  Instruction
}

// -- Conform Value
func (r *Result) valueInterface() {}
func (r *Result) String() string  { return r.Name }

// -----------------------------------------------------------------------------
// Instruction
// -----------------------------------------------------------------------------

type Instruction interface {
	GetOperand() Value // Could be nil
	GetResult() Value
	GetResults() []Value
	String() string
}

type IntLiteral struct {
	Value  int64
	Result *Result
}

type RngSeed struct {
	// TODO Expect a value with Int type
	Input  *IntLiteral
	Result *Result
}

type Return struct {
	Value Value
}

// -- Conform Instruction
func (lit *IntLiteral) GetOperand() Value   { return nil }
func (lit *IntLiteral) GetResult() Value    { return lit.Result }
func (lit *IntLiteral) GetResults() []Value { return []Value{lit.Result} }
func (lit *IntLiteral) String() string      { return fmt.Sprintf("%v = IntLit(%v)", lit.Result, lit.Value) }

func (rng *RngSeed) GetOperand() Value   { return rng.Input.Result }
func (rng *RngSeed) GetResult() Value    { return rng.Result }
func (rng *RngSeed) GetResults() []Value { return []Value{rng.Result} }
func (rng *RngSeed) String() string {
	return fmt.Sprintf("%v = RngSeed(%v)", rng.Result, rng.Input.Result)
}

func (r *Return) GetOperand() Value   { return r.Value }
func (r *Return) GetResult() Value    { return nil }
func (r *Return) GetResults() []Value { return nil }
func (r *Return) String() string      { return fmt.Sprintf("return %v", r.Value) }

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
	ins.Result = f.nextResult(ins)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngSeed(lit *IntLiteral) *RngSeed {
	ins := &RngSeed{
		Input:  lit,
		Result: nil,
	}
	ins.Result = f.nextResult(ins)
	f.insts = append(f.insts, ins)
	return ins
}

// -- Output

func (f *Fn) SetOutputAndDone(v Value) {
	f.insts = append(f.insts, &Return{v})
	f.Done()
}

// -- Helper Methods.

func (f *Fn) nextResult(src Instruction) *Result {
	i := f.res_index
	f.res_index++
	return &Result{
		Name: fmt.Sprintf("%%%v", i),
		Src:  src,
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
	return &Module{fns: b.fns}, nil
}

// -----------------------------------------------------------------------------
// All String Helpers
// -----------------------------------------------------------------------------

func (f *Fn) DebugString(w io.Writer) {
	fmt.Fprintf(w, "fn %v() {\n", f.name)
	for _, ins := range f.insts {
		fmt.Fprintf(w, "  %v\n", ins)
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
