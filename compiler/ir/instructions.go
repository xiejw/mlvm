package ir

import (
	"fmt"

	"github.com/xiejw/mlvm/compiler/base/errors"
)

type IntLiteral struct {
	Value  int64
	Result *Result
}

type ShapeLiteral struct {
	Dims   []int
	Result *Result
}

type RngSource struct {
	Input  Value // Must KInt as Type
	Result *Result
}

type RngTensor struct {
	Source Value
	Shape  Value
	Result *Result
	// Support more dist type
}

type Return struct {
	Value Value
}

// -- Conform Instruction
func (lit *IntLiteral) GetOperand() Value     { return nil }
func (lit *IntLiteral) GetOperands() []Value  { return nil }
func (lit *IntLiteral) GetResult() Value      { return lit.Result }
func (lit *IntLiteral) GetResults() []Value   { return []Value{lit.Result} }
func (lit *IntLiteral) String() string        { return fmt.Sprintf("%v = IntLit(%v)", lit.Result, lit.Value) }
func (lit *IntLiteral) Check() *errors.DError { return nil }

func (lit *ShapeLiteral) GetOperand() Value    { return nil }
func (lit *ShapeLiteral) GetOperands() []Value { return nil }
func (lit *ShapeLiteral) GetResult() Value     { return lit.Result }
func (lit *ShapeLiteral) GetResults() []Value  { return []Value{lit.Result} }
func (lit *ShapeLiteral) String() string {
	return fmt.Sprintf("%v = ShapeLit(%v)", lit.Result, lit.Dims)
}
func (lit *ShapeLiteral) Check() *errors.DError {
	for _, d := range lit.Dims {
		if d <= 0 {
			return errors.New("All Dims of ShapeLiteral must be positive, but got: %v", lit.Dims)
		}
	}
	return nil
}

func (rng *RngSource) GetOperand() Value    { return rng.Input }
func (rng *RngSource) GetOperands() []Value { return []Value{rng.Input} }
func (rng *RngSource) GetResult() Value     { return rng.Result }
func (rng *RngSource) GetResults() []Value  { return []Value{rng.Result} }
func (rng *RngSource) String() string {
	return fmt.Sprintf("%v = RngSource(%v)", rng.Result, rng.Input)
}
func (rng *RngSource) Check() *errors.DError {
	if !rng.Input.Type().IsInt() {
		return errors.New("RngSource expects int as seed input, but got: %v",
			rng.Input.Type()).EmitNote("RngSource: %v", rng)
	}
	return nil
}

func (rng *RngTensor) GetOperand() Value {
	panic("GetOperand should not be called with multiple operands.")
}
func (rng *RngTensor) GetOperands() []Value { return []Value{rng.Source, rng.Shape} }
func (rng *RngTensor) GetResult() Value     { return rng.Result }
func (rng *RngTensor) GetResults() []Value  { return []Value{rng.Result} }
func (rng *RngTensor) String() string {
	return fmt.Sprintf("%v = RngTensor(%v, %v)", rng.Result, rng.Source, rng.Shape)
}
func (rng *RngTensor) Check() *errors.DError {
	//if !rng.Input.Type().IsInt() {
	//	return errors.New("RngTensor expects int as seed input, but got: %v",
	//		rng.Input.Type()).EmitNote("RngTensor: %v", rng)
	//}
	return nil
}

func (r *Return) GetOperand() Value     { return r.Value }
func (r *Return) GetOperands() []Value  { return []Value{r.Value} }
func (r *Return) GetResult() Value      { return nil }
func (r *Return) GetResults() []Value   { return nil }
func (r *Return) String() string        { return fmt.Sprintf("return %v", r.Value) }
func (r *Return) Check() *errors.DError { return nil }
