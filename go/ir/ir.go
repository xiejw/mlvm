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

type Value interface{}

type TensorValue struct {
	Shape *Shape
}

type TupleValue struct {
}

type Inst interface {
	GetResult() Value
	GetResults() []Value
}

// Conform Value
type IntLiteral struct {
	Value int64
}

func (lit *IntLiteral) GetResult() Value {
	return lit
}

func (lit *IntLiteral) GetResults() []Value {
	return []Value{lit}
}

type Fn struct {
	b         *Builder
	name      string
	finalized bool
}

func (f *Fn) DebugString(w io.Writer) {
	fmt.Fprintf(w, "fn %v()", f.name)
}

func (f *Fn) String() string {
	buf := bytes.Buffer{}
	f.DebugString(&buf)
	return buf.String()
}

type Module struct {
	fns []*Fn
}

func (m *Module) DebugString(w io.Writer) {
	fmt.Fprintf(w, "module {\n")
	for _, f := range m.fns {
		fmt.Fprintf(w, "\nfn %v() {\n", f.name)
		fmt.Fprintf(w, "}\n")
	}
	fmt.Fprintf(w, "\n}\n")
}

func (m *Module) String() string {
	buf := bytes.Buffer{}
	m.DebugString(&buf)
	return buf.String()
}

func (m *Module) Fns() []*Fn {
	return m.fns
}

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

func (f *Fn) IntLiteral(v int64) *IntLiteral {
	return &IntLiteral{v}
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
	f.finalize()
}

func (f *Fn) finalize() {
	if f.finalized {
		// internel bug.
		panic("fn has finalized already.")
	}

	f.finalized = true
	f.b.fns = append(f.b.fns, f)
	f.b.f_map[f.name] = f
}

func (b *Builder) Finalize() (*Module, *errors.DError) {
	return &Module{fns: b.fns}, nil
}
