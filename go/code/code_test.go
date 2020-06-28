package code

import (
	"testing"
)

const maxOpcodeNameLen = 10

func checkNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Did not expect error: %v.", err)
	}
}

// expected is allowed to have one newline.
func assertDisassemblyCode(t *testing.T, expected, gotWithoutNewLine string) {
	t.Helper()
	got := "\n" + gotWithoutNewLine

	if expected != got {
		t.Errorf("unexpected Instructions String(): expected:\n`%v`\ngot:\n`%v`\n", expected, got)
	}
}

func TestOpcodes(t *testing.T) {
	ops := []struct {
		op   Opcode
		name string
	}{
		{OpConstant, "OpConstant"},
		{OpLoadG, "OpLoadG"},
		{OpStoreG, "OpStoreG"},
		{OpLoadT, "OpLoadT"},
		{OpStoreT, "OpStoreT"},
		{OpPrngNew, "OpPrngNew"},
		{OpPrngDist, "OpPrngDist"},
		{OpTensor, "OpTensor"},
		{OpAdd, "OpAdd"},
	}

	for _, testOp := range ops {
		op, err := Lookup(testOp.op)
		checkNoErr(t, err)
		if op.Name != testOp.name {
			t.Fatalf("Op name mistmatch")
		}

		if len(op.Name) > maxOpcodeNameLen {
			t.Errorf("Opcode name is too long: %v", op.Name)
		}
	}
}

func TestInstructionDisassembly(t *testing.T) {
	ops := []struct {
		expected string
		op       Opcode
		args     []int
	}{
		{"\n000000 OpConstant 123\n", OpConstant, []int{123}},
		{"\n000000 OpLoadG    456\n", OpLoadG, []int{456}},
		{"\n000000 OpStoreG   789\n", OpStoreG, []int{789}},
		{"\n000000 OpLoadT   \n", OpLoadT, []int{}},
		{"\n000000 OpStoreT  \n", OpStoreT, []int{}},
		{"\n000000 OpPrngNew \n", OpPrngNew, []int{}},
		{"\n000000 OpPrngDist 111\n", OpPrngDist, []int{111}},
		{"\n000000 OpTensor  \n", OpTensor, []int{}},
		{"\n000000 OpAdd     \n", OpAdd, []int{}},
	}

	for _, testOp := range ops {
		expected := testOp.expected
		ins, err := MakeOp(testOp.op, testOp.args...)
		checkNoErr(t, err)

		got := Instructions(ins).String()
		assertDisassemblyCode(t, expected, got)
	}
}

func TestMultiInstructionsDisassembly(t *testing.T) {
	var ins []byte

	ins1, err := MakeOp(OpConstant, 123)
	checkNoErr(t, err)
	ins2, err := MakeOp(OpStoreG, 456)
	checkNoErr(t, err)
	ins3, err := MakeOp(OpLoadG, 789)
	checkNoErr(t, err)

	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)
	expected :=
		`
000000 OpConstant 123
000003 OpStoreG   456
000006 OpLoadG    789
`
	got := Instructions(ins).String()

	assertDisassemblyCode(t, expected, got)
}
