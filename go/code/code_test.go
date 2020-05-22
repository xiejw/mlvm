package code

import (
	"testing"
)

func checkNoErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Did not expect error: %v.", err)
	}
}

func TestOpcodes(t *testing.T) {
	ops := []struct {
		op   Opcode
		name string
	}{
		{OpData, "OpData"},
		{OpLoad, "OpLoad"},
		{OpStore, "OpStore"},
		{OpTensor, "OpTensor"},
		{OpAdd, "OpAdd"},
	}

	for _, testOp := range ops {
		op, err := Lookup(testOp.op)
		checkNoErr(t, err)
		if op.Name != testOp.name {
			t.Fatalf("Op name mistmatch")
		}
	}
}

func TestInstructionDisassembly(t *testing.T) {
	ops := []struct {
		expected string
		op       Opcode
		args     []int
	}{
		{"000000 OpData 123\n", OpData, []int{123}},
		{"000000 OpLoad 123\n", OpLoad, []int{123}},
		{"000000 OpStore 123\n", OpStore, []int{123}},
		{"000000 OpTensor\n", OpTensor, []int{}},
		{"000000 OpAdd\n", OpAdd, []int{}},
	}

	for _, testOp := range ops {
		expected := testOp.expected
		ins, err := MakeOp(testOp.op, testOp.args...)
		checkNoErr(t, err)

		got := Instructions(ins).String()

		if expected != got {
			t.Errorf("Unexpected Instructions String(): expected:\n%v\ngot:\n%v\n", expected, got)
		}
	}
}
