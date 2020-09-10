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
		{OpCONST, "OpCONST"},
		{OpPOP, "OpPOP"},
		{OpLOAD, "OpLOAD"},
		{OpSTORE, "OpSTORE"},
		{OpLOADS, "OpLOADS"},
		{OpSTORES, "OpSTORES"},
		{OpIOR, "OpIOR"},
		{OpRNG, "OpRNG"},
		{OpRNGT, "OpRNGT"},
		{OpRNGS, "OpRNGS"},
		{OpT, "OpT"},
		{OpTADD, "OpTADD"},
		{OpTMINUS, "OpTMINUS"},
		{OpTMUL, "OpTMUL"},
		{OpTBROAD, "OpTBROAD"},
		{OpTREDUCE, "OpTREDUCE"},
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
		{"\n000000 OpCONST    123\n", OpCONST, []int{123}},
		{"\n000000 OpPOP     \n", OpPOP, []int{}},
		{"\n000000 OpLOAD     456\n", OpLOAD, []int{456}},
		{"\n000000 OpSTORE    789\n", OpSTORE, []int{789}},
		{"\n000000 OpLOADS   \n", OpLOADS, []int{}},
		{"\n000000 OpSTORES  \n", OpSTORES, []int{}},
		{"\n000000 OpIOR      12\n", OpIOR, []int{12}},
		{"\n000000 OpRNG     \n", OpRNG, []int{}},
		{"\n000000 OpRNGT     111\n", OpRNGT, []int{111}},
		{"\n000000 OpRNGS    \n", OpRNGS, []int{}},
		{"\n000000 OpT       \n", OpT, []int{}},
		{"\n000000 OpTADD    \n", OpTADD, []int{}},
		{"\n000000 OpTMINUS  \n", OpTMINUS, []int{}},
		{"\n000000 OpTMUL    \n", OpTMUL, []int{}},
		{"\n000000 OpTBROAD  \n", OpTBROAD, []int{}},
		{"\n000000 OpTREDUCE  1\n", OpTREDUCE, []int{1}},
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

	ins1, err := MakeOp(OpCONST, 123)
	checkNoErr(t, err)
	ins2, err := MakeOp(OpSTORE, 456)
	checkNoErr(t, err)
	ins3, err := MakeOp(OpLOAD, 789)
	checkNoErr(t, err)

	ins = append(ins, ins1...)
	ins = append(ins, ins2...)
	ins = append(ins, ins3...)
	expected :=
		`
000000 OpCONST    123
000003 OpSTORE    456
000006 OpLOAD     789
`
	got := Instructions(ins).String()

	assertDisassemblyCode(t, expected, got)
}
