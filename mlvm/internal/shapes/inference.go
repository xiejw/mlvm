// Inference the Shapes of Instruction Results.
package shapes

import (
	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/internal/errors"
)

var ()

func InferResultShapesForElementWiseOp(operands []*array.Shape) ([]*array.Shape, error) {
	var err error
	switch {
	case len(operands) != 2:
		err = errors.Errorf("Expected two operands, got: %v", len(operands))
	case !array.ShapesEqual(operands[0], operands[1]):
		err = errors.Errorf(
			"The shapes for element-wise op are not compatible: LHS %v vs RHS %v",
			operands[0], operands[1],
		)
	}
	if err != nil {
		return nil, err
	}

	return []*array.Shape{operands[0]}, nil
}
