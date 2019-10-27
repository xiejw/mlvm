package shapes

import (
	"testing"
)

func TestInvalidElementWiseInference(t *testing.T) {
	_, err := InferResultShapesForElementWiseOp(nil)
	if err == nil {
		t.Errorf("Expected error")
	}
}
