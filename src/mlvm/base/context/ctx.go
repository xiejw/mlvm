package context

import (
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

type ContextBuilder struct {
	IsTraining bool
}

func (b *ContextBuilder) Build() *Context {
	c := Context{
		isTraining:  b.IsTraining,
		weights:     make(map[string]w.Weight),
		nameRecords: make(map[string]int),
	}
	return &c
}

type Context struct {
	isTraining  bool
	weights     map[string]w.Weight
	nameRecords map[string]int
}

func (ctx *Context) IsTraining() bool {
	return ctx.isTraining
}

func (ctx *Context) NewWeight(name string, shape t.Shape, dtype t.DType) w.Weight {
	unique_name := ctx.AssignUniqueName(name)
	weight := w.NewWeight(unique_name, shape, dtype)
	ctx.registerNewWeight(weight)
	return weight
}

func (ctx *Context) AssignUniqueName(name string) string {
	if _, existed := ctx.nameRecords[name]; existed {
		panic("not impl")
	}
	ctx.nameRecords[name] += 1
	return name
}

func (ctx *Context) registerNewWeight(weight w.Weight) {
	ctx.weights[weight.Name()] = weight
}
