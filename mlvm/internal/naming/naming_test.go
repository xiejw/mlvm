package naming

import (
	"testing"
)

func TestValidArrayName(t *testing.T) {
	validNames := []string{"abc", "a_1", "A2_"}
	for _, validName := range validNames {
		if err := ValidateArrayName(validName); err != nil {
			t.Errorf("Did not expect error; got %v", err)
		}
	}
}

func TestInvalidArrayName(t *testing.T) {
	invalidNames := []string{"_1", "$1", "2a"}
	for _, invalidName := range invalidNames {
		if err := ValidateArrayName(invalidName); err == nil {
			t.Errorf("Expect error")
		}
	}
}

func TestValidInstructionName(t *testing.T) {
	validNames := []string{"abc", "a_1", "A2_"}
	for _, validName := range validNames {
		if err := ValidateInstructionName(validName); err != nil {
			t.Errorf("Did not expect error; got %v", err)
		}
	}
}

func TestInvalidInstructionName(t *testing.T) {
	invalidNames := []string{"_1", "$1", "2a"}
	for _, invalidName := range invalidNames {
		if err := ValidateInstructionName(invalidName); err == nil {
			t.Errorf("Expect error")
		}
	}
}

func TestCanonicalResultName(t *testing.T) {
	name := CanonicalResultName("add", 1)
	if name != "%o{1,add}" {
		t.Errorf("Unexpected result name.")
	}
}

func TestDefaultInstructionName(t *testing.T) {
	name := DefaultInstructionName("add", 1)
	if name != "add_001" {
		t.Errorf("Unexpected result name.")
	}
}
