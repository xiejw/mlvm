package shapes

import (
	"testing"

	"github.com/xiejw/mlvm/mlvm/array"
)

func TestCompatibleElementWiseShapes(t *testing.T) {
	s1 := array.NewShapeOrDie([]array.Dimension{4, 1})
	s2 := array.NewShapeOrDie([]array.Dimension{4, 1})
	_, err := InferResultShapesForElementWiseOp([]*array.Shape{s1, s2})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestInvalidElementWiseInference(t *testing.T) {
	_, err := InferResultShapesForElementWiseOp(nil)
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestIncompatibleElementWiseShapes(t *testing.T) {
	s1 := array.NewShapeOrDie([]array.Dimension{4, 1})
	s2 := array.NewShapeOrDie([]array.Dimension{1, 4})
	_, err := InferResultShapesForElementWiseOp([]*array.Shape{s1, s2})
	if err == nil {
		t.Errorf("Expected error")
	}
}
