package naming

import (
	"fmt"

	"regexp"
)

// Naming policy
// - All user provided tensor and instruction names should NOT have special
//   charactors. See `arrayNameRegexp` and `userInstructionNameRegexp`.
//
// - Special characters.
//   - `%` is the leading charactor for result.
//   - `#` is used as leading char for graph rewrites, like inserted op, added
//     contant tensor, etc.

var (
	// Regexp
	arrayNameRegexp           = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	userInstructionNameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)

	// Errors
	errInvalidArrayName       = "Array name `%v` is invalid. Must be legal identifier name."
	errInvalidInstructionName = "Instruction name `%v` is invalid. Must be legal identifier name."
)

const (
	resultTensorLeandingCharactor = "%o"
	// GraphRewritePrefix            = "#"
)

// Validates whether array name is valid.
func ValidateArrayName(name string) error {
	if arrayNameRegexp.MatchString(name) {
		return nil
	}
	return fmt.Errorf(errInvalidArrayName, name)
}

// Validates whether user provided instruction name is valid.
func ValidateInstructionName(name string) error {
	if userInstructionNameRegexp.MatchString(name) {
		return nil
	}
	return fmt.Errorf(errInvalidInstructionName, name)
}

// Returns canonical name for result.
//
// %{i,insName}
func CanonicalResultName(insName string, index int) string {
	return fmt.Sprintf("%s{%v,%v}", resultTensorLeandingCharactor, index, insName)
}

// Returns a default name for instruction based on baseName and globalIndex.
func DefaultInstructionName(baseName string, globalIndex int) string {
	return fmt.Sprintf("%v_%03v", baseName, globalIndex)
}