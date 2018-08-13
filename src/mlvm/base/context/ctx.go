package context

import (
	t "mlvm/base/tensor"
	w "mlvm/base/weight"
)

type ContextBuilder struct {
}

func (b *ContextBuilder) Build() *Context {
	c := Context{
		weights:     make(map[string]w.Weight),
		nameRecords: make(map[string]int),
	}
	return &c
}

type Context struct {
	weights     map[string]w.Weight
	nameRecords map[string]int
}

// Returns a new `Weight` with unique name.
func (ctx *Context) NewWeight(name string, shape t.Shape, dtype t.DType) w.Weight {
	unique_name := ctx.AssignUniqueName(name)
	weight := w.NewWeight(unique_name, shape, dtype)
	ctx.registerNewWeight(weight)
	return weight
}

// Assign a unique name in the context. It appends suffix at the end to make it
// unique.
func (ctx *Context) AssignUniqueName(name string) string {
	if _, existed := ctx.nameRecords[name]; existed {
		panic("not impl")
	}
	ctx.nameRecords[name] += 1
	return name
}

// Records a table for weight name to `Weight` mapping.
func (ctx *Context) registerNewWeight(weight w.Weight) {
	ctx.weights[weight.Name()] = weight
}
