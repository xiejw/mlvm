package ir

import (
	"bytes"
	"fmt"
	"io"

	"github.com/xiejw/mlvm/go/base/errors"
)

type Shape struct {
	Dims []int
}

///////////////////////////////////////////////////////////////////////////////
// Value
///////////////////////////////////////////////////////////////////////////////

type Value interface {
	valueInterface()
	String() string
}

type TensorValue struct {
	Shape *Shape
}

type TupleValue struct {
}

type Result struct {
	Name string
	Src  Inst
}

// Conform Value
func (r *Result) valueInterface() {}
func (r *Result) String() string  { return r.Name }

///////////////////////////////////////////////////////////////////////////////
// Inst
///////////////////////////////////////////////////////////////////////////////

type Inst interface {
	GetResult() Value
	GetResults() []Value
	String() string
}

type IntLiteral struct {
	Value  int64
	Result *Result
}

// Conform Inst
func (lit *IntLiteral) GetResult() Value    { return lit.Result }
func (lit *IntLiteral) GetResults() []Value { return []Value{lit.Result} }
func (lit *IntLiteral) String() string {
	return fmt.Sprintf("%v = IntLit(%v)", lit.Result, lit.Value)
}

type Return struct {
	Value Value
}

// Conform Inst
func (r *Return) GetResult() Value    { return nil }
func (r *Return) GetResults() []Value { return nil }
func (r *Return) String() string {
	return fmt.Sprintf("return %v", r.Value)
}

///////////////////////////////////////////////////////////////////////////////
// Fn
///////////////////////////////////////////////////////////////////////////////

type Fn struct {
	b         *Builder
	name      string
	finalized bool
	result_i  int
	inss      []Inst
}

func (f *Fn) Name() string {
	return f.name
}

func (f *Fn) Insts() []Inst {
	return f.inss
}

func (f *Fn) nextResult(src Inst) *Result {
	i := f.result_i
	f.result_i++
	return &Result{
		Name: fmt.Sprintf("%%%v", i),
		Src:  src,
	}
}

func (f *Fn) IntLiteral(v int64) *IntLiteral {
	ins := &IntLiteral{
		Value:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins)
	f.inss = append(f.inss, ins)
	return ins
}

func (f *Fn) ReadKVStore(key string, s *Shape) Inst {
	return nil
}

func (f *Fn) TAdd(lhs, rhs *TensorValue) Inst {
	return nil
}

func (f *Fn) MakeTuple(args ...*TensorValue) Inst {
	return nil
}

func (f *Fn) SetInput(v Value) Inst {
	return nil
}

func (f *Fn) NoOutput() {
	f.finalize()
}

func (f *Fn) SetOutput(v Value) {
	f.inss = append(f.inss, &Return{v})
	f.finalize()
}

func (f *Fn) finalize() {
	if f.finalized {
		panic("fn has finalized already.") // internel bug.
	}

	f.finalized = true
	f.b.fns = append(f.b.fns, f)
	f.b.f_map[f.name] = f
}

///////////////////////////////////////////////////////////////////////////////
// Module
///////////////////////////////////////////////////////////////////////////////

type Module struct {
	fns []*Fn
}

func (m *Module) Fns() []*Fn {
	return m.fns
}

///////////////////////////////////////////////////////////////////////////////
// Builder
///////////////////////////////////////////////////////////////////////////////

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

func (b *Builder) Finalize() (*Module, *errors.DError) {
	return &Module{fns: b.fns}, nil
}

///////////////////////////////////////////////////////////////////////////////
// All String Helpers
///////////////////////////////////////////////////////////////////////////////

func (f *Fn) DebugString(w io.Writer) {
	fmt.Fprintf(w, "fn %v() {\n", f.name)
	for _, ins := range f.inss {
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
