package ir

import (
	"bytes"
	"fmt"

	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
)

type Flag int

const (
	F_DIST_TYPE_BEGIN Flag = iota
	F_DIST_TYPE_NORM
	F_DIST_TYPE_TRUNC_NORM
	F_DIST_TYPE_END // unused
)

type IntLiteral struct {
	Value  int64
	Result *Result
}

type ShapeLiteral struct {
	Dims   []int
	Result *Result
}

type ArrayLiteral struct {
	Value  []float32
	Result *Result
}

type NewTensor struct {
	Shape  Value
	Array  Value
	Result *Result
}

type RngSource struct {
	Seed   Value // Must KInt as Type
	Result *Result
}

type RngFill struct {
	Shape    Value
	Source   Value
	Result   *Result
	DistType Flag
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

func (f *Fn) ArrayLiteral(v []float32) *ArrayLiteral {
	ins := &ArrayLiteral{
		Value:  v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KArray, Dims: []int{len(v)}})
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) NewTensor(s Value, arr Value) *NewTensor {
	ins := &NewTensor{
		Shape:  s,
		Array:  arr,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KTensor, Dims: s.Type().Dims})
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngSource(v Value) *RngSource {
	ins := &RngSource{
		Seed:   v,
		Result: nil,
	}
	ins.Result = f.nextResult(ins, 0, RngType)
	f.insts = append(f.insts, ins)
	return ins
}

func (f *Fn) RngFill(s Value, src Value, dist_type Flag) *RngFill {
	ins := &RngFill{
		Shape:    s,
		Source:   src,
		Result:   nil,
		DistType: dist_type,
	}
	dims := s.Type().Dims // could be nil
	ins.Result = f.nextResult(ins, 0, &Type{Kind: KTensor, Dims: dims})
	f.insts = append(f.insts, ins)
	return ins
}

// -----------------------------------------------------------------------------
// Check (Propagate shape dims, validate types, etc )
// -----------------------------------------------------------------------------

func (lit *IntLiteral) Check() error { return nil }

func (lit *ShapeLiteral) Check() error {
	return lit.Result.Type().ValidateShape()
}

func (lit *ArrayLiteral) Check() error {
	if len(lit.Value) == 0 {
		return errors.New("ArrayLiteral cannot be empty")
	}
	return lit.Result.Type().ValidateArray()
}

func (t *NewTensor) Check() error {
	if !t.Shape.Type().IsShape() {
		return errors.New(
			"NewTensor expects Shape as the first operand, but got type: %v", t.Shape.Type())
	}
	if t.Array.Type().Kind != KArray {
		return errors.New(
			"NewTensor expects Array as the second operand, but got type: %v", t.Array.Type())
	}

	// Check the elements in Array matching Shape.
	dims := t.Shape.Type().Dims
	count := 1
	for _, d := range dims {
		count *= d
	}
	if count != t.Array.Type().Dims[0] {
		return errors.New(
			"NewTensor.Shape should match Array.size: shape elements: %v, array len: %v", count,
			t.Array.Type().Dims[0])
	}

	// forwards the shape
	t.Result.Type().Dims = dims
	return nil
}

func (rng *RngSource) Check() error {
	if !rng.Seed.Type().IsInt() {
		return errors.New(
			"RngSource expects int as seed input, but got type: %v", rng.Seed.Type())
	}
	return nil
}

func (rng *RngFill) Check() error {
	if rng.DistType <= F_DIST_TYPE_BEGIN || rng.DistType >= F_DIST_TYPE_END {
		return errors.New(
			"RngFill.DistType is not in the valid range: %v", rng.DistType)
	}
	if !rng.Shape.Type().IsShape() {
		return errors.New(
			"RngFill expects Shape as the second operand, but got type: %v", rng.Shape.Type())
	}
	if rng.Source.Type().Kind != KRng {
		return errors.New(
			"RngFill expects RngSource as the first operand, but got type: %v", rng.Source.Type())
	}
	// forwards the shape
	rng.Result.Type().Dims = rng.Shape.Type().Dims
	return nil
}

func (r *Return) Check() error { return nil }

// -----------------------------------------------------------------------------
// String
// -----------------------------------------------------------------------------

func (lit *IntLiteral) String() string {
	return fmt.Sprintf("%v = IntLit(%v)", lit.Result, lit.Value)
}

func (lit *ShapeLiteral) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v = ShapeLit(", lit.Result)
	object.NewShape(lit.Dims).DebugString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

func (lit *ArrayLiteral) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v = ArrayLit(", lit.Result)
	(&object.Array{lit.Value}).DebugString(&buf, 9)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

func (t *NewTensor) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%v = NewTensor(", t.Result)
	object.NewShape(t.Result.Type().Dims).DebugString(&buf)
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

func (rng *RngSource) String() string {
	return fmt.Sprintf("%v = RngSource(%v)", rng.Result, rng.Seed)
}

func (rng *RngFill) String() string {
	return fmt.Sprintf("%v = RngFill(%v, %v)", rng.Result, rng.Shape, rng.Source)
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

// -- ArrayLiteral
func (lit *ArrayLiteral) GetOperand() Value    { return nil }
func (lit *ArrayLiteral) GetOperands() []Value { return nil }
func (lit *ArrayLiteral) GetResult() Value     { return lit.Result }
func (lit *ArrayLiteral) GetResults() []Value  { return []Value{lit.Result} }

// -- NewTensor
func (t *NewTensor) GetOperand() Value    { panic("invalid with multiple operands.") }
func (t *NewTensor) GetOperands() []Value { return []Value{t.Shape, t.Array} }
func (t *NewTensor) GetResult() Value     { return t.Result }
func (t *NewTensor) GetResults() []Value  { return []Value{t.Result} }

// -- RngSource
func (rng *RngSource) GetOperand() Value    { return rng.Seed }
func (rng *RngSource) GetOperands() []Value { return []Value{rng.Seed} }
func (rng *RngSource) GetResult() Value     { return rng.Result }
func (rng *RngSource) GetResults() []Value  { return []Value{rng.Result} }

// -- RngFill
func (rng *RngFill) GetOperand() Value    { panic("invalid with multiple operands.") }
func (rng *RngFill) GetOperands() []Value { return []Value{rng.Shape, rng.Source} }
func (rng *RngFill) GetResult() Value     { return rng.Result }
func (rng *RngFill) GetResults() []Value  { return []Value{rng.Result} }

// -- Return
func (r *Return) GetOperand() Value    { return r.Value }
func (r *Return) GetOperands() []Value { return []Value{r.Value} }
func (r *Return) GetResult() Value     { return r.Value }
func (r *Return) GetResults() []Value  { return []Value{r.Value} }
