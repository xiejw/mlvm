// Inference the Shapes of Instruction Results.
package shapes

import (
	"github.com/xiejw/mlvm/mlvm/array"
)

type BroadcastMode int

func InferResultShapesForElementWiseOp(operands []*array.Shape) []*array.Shape {
	return nil
}
