package code

import (
	"encoding/binary"
	"fmt"
)

type Opcode byte

const (
	// Opcode string name should be at most 10 chars. Enforced in test.
	//
	// For meaning, see Definition.
	OpCONST Opcode = iota
	OpPOP
	OpLOAD
	OpMOVE
	OpSTORE
	OpLOADS
	OpSTORES
	OpIOR
	OpRNG
	OpRNGT
	OpRNGS
	OpT
	OpTSHAPE
	OpTADD
	OpTMINUS
	OpTMUL
	OpTBROAD
	OpTREDUCE
)

// Defines the string name and operand requirements.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	// ---------------------------------------------------------------------------
	// Consts, Memory, and Store.
	// ---------------------------------------------------------------------------

	// Loads constant object from Program into stack.
	//
	// Operand: (uint16) object index.
	// Stack  : push the object to the top.
	OpCONST: {"OpCONST", []int{2}},

	// Pops out top item on stack.
	//
	// Operand: no.
	// Stack  : pop the object from the top.
	OpPOP: {"OpPOP", []int{}},

	// Loads the object from global memory.
	//
	// Operand: (uint16) object index.
	// Stack  : push the object to the top.
	OpLOAD: {"OpLOAD", []int{2}},

	// Loads the object from global memory (and deletes it from memory)
	//
	// Operand: (uint16) object index.
	// Stack  : push the object to the top.
	OpMOVE: {"OpMOVE", []int{2}},

	// Stores object into global memory.
	//
	// Operand: (uint16) object index.
	// Stack  : pop the object from the top.
	OpSTORE: {"OpSTORE", []int{2}},

	// Loads object from key-value store.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as string key.
	OpLOADS: {"OpLOADS", []int{}},

	// Stores the object into key-value store.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as the object.
	//          pops the second item and uses it as (String) key.
	OpSTORES: {"OpSTORES", []int{}},

	// ---------------------------------------------------------------------------
	// I/O.
	// ---------------------------------------------------------------------------

	// Reads objects from infeed channel.
	//
	// Operand: (uint16) num of objects to read.
	// Stack  : reads N objects from infeed channel and pushes to stack one by one in seq.
	OpIOR: {"OpIOR", []int{2}},

	// ---------------------------------------------------------------------------
	// RNG.
	// ---------------------------------------------------------------------------

	// Creates a new rng source.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as (int type) seed.
	//          stores the RNG source into the stack.
	OpRNG: {"OpRNG", []int{}},

	// Creates a Tensor with selectd distribution.
	//
	// Operand: (uint16) distribution type index.
	// Stack  : pops the top item and uses it as RNG source.
	//          pops the second item and uses it as (Shape).
	//          stores the Tensor into the stack.
	OpRNGT: {"OpRNGT", []int{2}},

	// Splits the rng source.
	//
	// Operand: no
	// Stack  : pops the top item and uses it as RNG source.
	//          stores the first item of the split result into the stack.
	//          then stores the second item of the split result into the stack.
	OpRNGS: {"OpRNGS", []int{}},

	// ---------------------------------------------------------------------------
	// Tensor.
	// ---------------------------------------------------------------------------

	// Creates a new Tensor.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as (Array).
	//          pops the second item and uses it as (Shape).
	//          stores the Tensor into the stack.
	OpT: {"OpT", []int{}},

	// Gets the shape from the Tensor.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as Tensor operand.
	//          stores the shape of it into the stack.
	OpTSHAPE: {"OpTSHAPE", []int{}},

	// Does addition two tensors.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as the second operand.
	//          then pops the top item and uses it as the first operand.
	//          stores the result into the stack.
	// Shape  : both operands must have same shapes.
	OpTADD: {"OpTADD", []int{}},

	// Does minus for two tensors.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as the second operand.
	//          then pops the top item and uses it as the first operand.
	//          stores the result into the stack.
	// Shape  : both operands must have same shapes.
	OpTMINUS: {"OpTMINUS", []int{}},

	// Does (element-wise) multiplication for two tensors.
	//
	// Operand: no.
	// Stack  : pops the top item and uses it as the second operand.
	//          then pops the top item and uses it as the first operand.
	//          stores the result into the stack.
	// Shape  : both operands must have same shapes.
	OpTMUL: {"OpTMUL", []int{}},

	// Broadcasts a tensor of from its original shape to a tensor with new shape.
	//
	// OpTBROAD supports one simple form of broadcasting, on top of the normal Tensor representation.
	// The basic idea is: The underlying data can be repeated to represent the broadcasted shape.
	//
	// To be specific, for any shape [a1, a2, a3] with a1 * a2 * a3 elements for a tensor, it is
	// called non-compressed tensor array, as all elements have been filled in value buffer.
	//
	// - If a1 is not 1, the only valid broadcasting case is [b1, b2, ..., bm, a1, a2, a3], i.e.,
	//   major dimension extension, where m >= 1.  For example, all of the following cases are
	//   supported
	//
	//     - yes [2, 3] -> [3, 2, 3]
	//     - yes [2, 3] -> [5, 4, 3, 2, 3]
	//     - yes [2, 1] -> [4, 3, 2, 1]
	//
	//     - no [2, 3] -> [4, 3, 2, 1]   the final dim can only be 3, but got 1.
	//     - no [2, 3] -> [3, 4, 3]      the second from last dim can only be 2, but got 4.
	//     - no [2, 1] -> [3, 2, 3]      the final dim can only be 1, but got 3. numpy allows this.
	//
	// - If a1, a2, ..., ak is 1, where k<=3 as in the example, then valid broadcasting case is
	//
	//     [b1, b2, ..., bm, a'1, a'2, ..., a'k, a{k+1}, ..., a3]
	//
	//   i.e,
	//
	//     - yes [1, 3] -> [3, 2, 3]
	//     - yes [1, 3] -> [5, 3, 1, 3]
	//     - yes [1, 1] -> [5, 3, 2, 3]
	//
	// - As a special but super useful use case, shape [1] is allowed to be broadcasted to all other
	//   shapes, i.e.,
	//
	//   - yes [1] -> [3, 2, 1]
	//   - yes [1] -> [4, 3, 2, 3]
	//
	//   where is a natual extention of the rule 2.
	//
	// Operand: (uint8) flag.
	//          - 0: Lazy broadcast.
	//          - 1: Materialize broadcast, i.e., by creating the non-compressions tensors
	//
	// Stack  : pops the top item and uses it as tensor operand.
	//          then pops the top item and uses it as new Shape.
	//          stores the new result (Tensor) into the stack.
	// Shape  : the new shape is the left extension of the old shape.
	OpTBROAD: {"OpTBROAD", []int{1}},

	// ---------------------------------------------------------------------------
	// Return.
	// ---------------------------------------------------------------------------

	// Reduce the tensor.
	//
	// Operand: (uint8) reduce merge Op index (see MergeType)
	// Stack  : pops the top item and uses it as tensor operand.
	//          stores the new result into the stack.
	OpTREDUCE: {"OpTREDUCE", []int{1}},
}

// Looks up the Definition of the op code.
func Lookup(op Opcode) (*Definition, error) {
	def, ok := definitions[op]
	if !ok {
		return nil, fmt.Errorf("Opcode %d undefined", op)
	}

	return def, nil
}

// Encodes the op code with its operands, if any.
//
// This is Opcode encoder. The decoder is in vm.
func MakeOp(op Opcode, operands ...int) ([]byte, error) {
	def, err := Lookup(op)
	if err != nil {
		return nil, err
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	if len(def.OperandWidths) != len(operands) {
		return nil, fmt.Errorf("MakeOp: operand counts for `%v` mismatch: expected %v, got %v",
			op, len(def.OperandWidths), len(operands))
	}

	offset := 1
	for i, o := range operands {
		w := def.OperandWidths[i]
		switch w {
		case 1:
			instruction[offset] = byte(o)
			offset += 1
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
			offset += 2
		default:
			panic(fmt.Sprintf("unsupported width (%v) for op: %v", w, op))
		}
	}

	return instruction, nil
}

////////////////////////////////////////////////////////////////////////////////
// Public Helper Methods
////////////////////////////////////////////////////////////////////////////////

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

////////////////////////////////////////////////////////////////////////////////
// String Related.
////////////////////////////////////////////////////////////////////////////////

func (opcode Opcode) String() string {
	def, ok := definitions[opcode]
	if !ok {
		panic(fmt.Errorf("Opcode %d undefined", opcode))
	}
	return def.Name
}
