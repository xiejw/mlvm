package ir

import (
	"bytes"
	"fmt"

	"github.com/xiejw/mlvm/compiler/base/errors"
	"github.com/xiejw/mlvm/vm/object"
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

// -----------------------------------------------------------------------------
// Constructors (from Fn)
// -----------------------------------------------------------------------------

func (f *Fn) IntLiteral(v int64) *IntLiteral {
	ins := &IntLiteral{
		Value:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, IntType)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) ShapeLiteral(dims []int) *ShapeLiteral {
	ins := &ShapeLiteral{
		Dims:   dims, // `Check` checks validness.
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KShape, Dims: dims})
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngSource(v Value) *RngSource {
	ins := &RngSource{
		Input:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, RngType)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngTensor(src Value, s Value) *RngTensor {
	ins := &RngTensor{
		Source: src,
		Shape:  s,
		Result: nil,
	}
	dims := s.Type().Dims // could be nil
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KTensor, Dims: dims})
	f.insts = append(f.insts, ins)
	return ins
}

// -----------------------------------------------------------------------------
// Check (Propagate shape dims, validate types, etc )
// -----------------------------------------------------------------------------

func (lit *IntLiteral) Check() *errors.DError { return nil }

func (lit *ShapeLiteral) Check() *errors.DError {
	for _, d := range lit.Dims {
		if d <= 0 {
			return errors.New("All Dims of ShapeLiteral must be positive, but got: %v", lit.Dims)
		}
	}
	return nil
}

func (rng *RngSource) Check() *errors.DError {
	if !rng.Input.Type().IsInt() {
		return errors.New(
			"RngSource expects int as seed input, but got type: %v", rng.Input.Type())
	}
	return nil
}

func (rng *RngTensor) Check() *errors.DError {
	if rng.Source.Type().Kind != KRng {
		return errors.New(
			"RngTensor expects RngSource as the first operand, but got type: %v", rng.Source.Type())
	}
	if rng.Shape.Type().Kind != KShape {
		return errors.New(
			"RngTensor expects Shape as the second operand, but got type: %v", rng.Shape.Type())
	}
	// forwards the shape
	rng.Result.Type().Dims = rng.Shape.Type().Dims
	return nil
}

func (r *Return) Check() *errors.DError { return nil }

// -----------------------------------------------------------------------------
// String
// -----------------------------------------------------------------------------

func (lit *IntLiteral) String() string {
	return fmt.Sprintf("%v = IntLit(%v)", lit.Result, lit.Value)
}

func (lit *ShapeLiteral) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v = ShapeLit(", lit.Result)
	object.NewShape(lit.Dims).ToHumanReadableString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

func (rng *RngSource) String() string {
	return fmt.Sprintf("%v = RngSource(%v)", rng.Result, rng.Input)
}

func (rng *RngTensor) String() string {
	return fmt.Sprintf("%v = RngTensor(%v, %v)", rng.Result, rng.Source, rng.Shape)
}

func (r *Return) String() string { return fmt.Sprintf("return %v", r.Value) }

// -----------------------------------------------------------------------------
// Conform Instruction
// -----------------------------------------------------------------------------

// -- IntLiteral
func (lit *IntLiteral) GetOperand() Value    { return nil }
func (lit *IntLiteral) GetOperands() []Value { return nil }
func (lit *IntLiteral) GetResult() Value     { return lit.Result }
func (lit *IntLiteral) GetResults() []Value  { return []Value{lit.Result} }

// -- ShapeLiteral
func (lit *ShapeLiteral) GetOperand() Value    { return nil }
func (lit *ShapeLiteral) GetOperands() []Value { return nil }
func (lit *ShapeLiteral) GetResult() Value     { return lit.Result }
func (lit *ShapeLiteral) GetResults() []Value  { return []Value{lit.Result} }

// -- RngSource
func (rng *RngSource) GetOperand() Value    { return rng.Input }
func (rng *RngSource) GetOperands() []Value { return []Value{rng.Input} }
func (rng *RngSource) GetResult() Value     { return rng.Result }
func (rng *RngSource) GetResults() []Value  { return []Value{rng.Result} }

// -- RngTensor
func (rng *RngTensor) GetOperand() Value    { panic("invalid with multiple operands.") }
func (rng *RngTensor) GetOperands() []Value { return []Value{rng.Source, rng.Shape} }
func (rng *RngTensor) GetResult() Value     { return rng.Result }
func (rng *RngTensor) GetResults() []Value  { return []Value{rng.Result} }

// -- Return
func (r *Return) GetOperand() Value    { return r.Value }
func (r *Return) GetOperands() []Value { return []Value{r.Value} }
func (r *Return) GetResult() Value     { return r.Value }
func (r *Return) GetResults() []Value  { return []Value{r.Value} }
