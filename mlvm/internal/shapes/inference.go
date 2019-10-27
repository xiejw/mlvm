// Inference the Shapes of Instruction Results.
package shapes

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
)

func InferResultShapesForElementWiseOp(operands []*array.Shape) ([]*array.Shape, error) {
	if len(operands) != 2 {
		return nil, fmt.Errorf("Expected two operands, got: %v", len(operands))
	}

	// TODO: Check shape compatible
	return []*array.Shape{operands[0]}, nil
}
