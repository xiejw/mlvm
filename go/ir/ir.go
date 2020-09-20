package ir

import (
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

type Fn struct{}

type Module struct{}

func (m *Module) Fns() []*Fn {
	return nil
}

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) NewFn(fn_name string) (*Fn, *errors.DError) {
	return &Fn{}, nil
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

func (f *Fn) SetOutput(v Value) {
}

func (b *Builder) Finalize() (*Module, *errors.DError) {
	return nil, nil
}
