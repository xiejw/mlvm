package ast

// Instruction in Module.
type Instruction struct {
	name     string
	op       *Op
	operands []*Tensor
	results  []*Tensor
}

func (ins *Instruction) Name() string {
	return ins.name
}

func (ins *Instruction) Operands() []*Tensor {
	return ins.operands
}

func (ins *Instruction) OnlyResult() *Tensor {
	if len(ins.results) != 1 {
		panic("Should have only one results.")
	}
	return ins.results[0]
}

func (ins *Instruction) Results() []*Tensor {
	return ins.results
}
