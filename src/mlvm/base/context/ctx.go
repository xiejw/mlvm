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
		isTraining:        b.IsTraining,
		weights:           make(map[string]w.Weight),
		weightNameRecords: make(map[string]int),
		tensorNameRecords: make(map[string]int),
	}
	return &c
}

type Context struct {
	isTraining        bool
	weights           map[string]w.Weight
	weightNameRecords map[string]int
	tensorNameRecords map[string]int
}

func (ctx *Context) IsTraining() bool {
	return ctx.isTraining
}

func (ctx *Context) NewWeight(name string, shape t.Shape, dtype t.DType) w.Weight {
	unique_name := ctx.getUniqueNameForWeight(name)
	weight := w.NewWeight(unique_name, shape, dtype)
	ctx.registerNewWeight(weight)
	return weight
}

func (ctx *Context) GetUniqueNameForTensor(name string) string {
	if _, existed := ctx.tensorNameRecords[name]; existed {
		panic("not impl")
	}
	ctx.tensorNameRecords[name] += 1
	return name
}

func (ctx *Context) getUniqueNameForWeight(name string) string {
	if _, existed := ctx.weights[name]; existed {
		panic("not impl")
	}
	return name
}

func (ctx *Context) registerNewWeight(weight w.Weight) {
	ctx.weights[weight.Name()] = weight
}
