package naming

import (
	"fmt"

	"regexp"
)

// Naming policy
// - All user provided tensor and instruction names should NOT have special
//   charactors. See `UserTensorNameRegexp` and `UserInstructionNameRegexp`.
//
// - Special characters.
//   - `%` is the leading charactor for result.
//   - `::` is the separator for result name and index.
//   - `#` is used as leading char for graph rewrites, like inserted op, added
//     contant tensor, etc.

var (
	userTensorNameRegexp      = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	userInstructionNameRegexp = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)
	// ContainerNameRegexp       = regexp.MustCompile(`^[a-zA-Z]+$`)
	// VariableNameRegexp        = regexp.MustCompile(`^[a-zA-Z]+$`)
)

const (
	resultTensorLeandingCharactor = "%"
	// GraphRewritePrefix            = "#"
)

const (
	DefaultContainerName = "*"
)

func ValidateUserTensor(name string) error {
	return nil
}

func CanonicalNameForResult(insName string, index int) string {
	return fmt.Sprintf("%s{%v,%v}", resultTensorLeandingCharactor, index, insName)
}

// func CanonicalNameForVariable(containerName, varName string) string {
// 	if containerName == "*" {
// 		return containerName + varName
// 	} else {
// 		return containerName + "/" + varName
// 	}
// }
