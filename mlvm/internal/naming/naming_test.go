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
