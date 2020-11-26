package ir

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/vm/base/errors"
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
	Check() error // Check should assume operands passing `Check`.
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

func (f *Fn) Name() string                { return f.name }
func (f *Fn) Instructions() []Instruction { return f.insts }

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

func (f *Fn) Done() { f.b.finishFn(f.name) }

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
	fns       []*Fn
	f_map     map[string]*Fn
	f_workset map[string]*Fn
}

func NewBuilder() *Builder {
	return &Builder{
		fns:       nil,
		f_map:     make(map[string]*Fn),
		f_workset: make(map[string]*Fn),
	}
}

func (b *Builder) NewFn(fn_name string) (*Fn, error) {
	if _, existed := b.f_map[fn_name]; existed {
		return nil, errors.New("fn name already existed in module: %v", fn_name)
	}
	if _, existed := b.f_workset[fn_name]; existed {
		return nil, errors.New("fn name already existed in module's workset: %v", fn_name)
	}
	fn := &Fn{b: b, name: fn_name}
	b.f_workset[fn_name] = fn
	b.fns = append(b.fns, fn)
	return fn, nil
}

func (b *Builder) finishFn(name string) {
	fn := b.f_workset[name]
	delete(b.f_workset, name)
	b.f_map[name] = fn
}

func (b *Builder) Done() (*Module, error) {
	// checks empty workset.
	if len(b.f_workset) != 0 {
		err := errors.New(
			"there is unfinished Fn(s) under construction (call `Done` or related methods).")
		for name, _ := range b.f_workset {
			err.EmitNote("found unfinished Fn: `%v`", name)
		}
		return nil, err
	}

	// validates all Fns.
	for _, fn := range b.fns {
		for _, ins := range fn.insts {
			err := ins.Check()
			if err != nil {
				return nil, errors.From(err).EmitNote(
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

func (f *Fn) DebugString(w io.Writer, printType bool) {
	fmt.Fprintf(w, "fn %v() {\n", f.name)
	for _, ins := range f.insts {
		if printType {
			fmt.Fprintf(w, "  %-40s:: %v\n", ins.String(), ins.GetResult().Type())
		} else {
			fmt.Fprintf(w, "  %v\n", ins)
		}
	}
	fmt.Fprintf(w, "}\n")
}

func (f *Fn) String() string {
	buf := bytes.Buffer{}
	f.DebugString(&buf, true /*printType*/)
	return buf.String()
}

func (m *Module) DebugString(w io.Writer, printType bool) {
	fmt.Fprintf(w, "module {\n")
	for _, f := range m.fns {
		fmt.Fprintf(w, "\n")
		f.DebugString(w, printType)
	}
	fmt.Fprintf(w, "\n}\n")
}

func (m *Module) String() string {
	buf := bytes.Buffer{}
	m.DebugString(&buf, true /*printType*/)
	return buf.String()
}
