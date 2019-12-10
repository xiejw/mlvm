package array

import (
	"testing"
)

func TestNewShape(t *testing.T) {
	_, err := NewShape([]Dimension{4})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = NewShape([]Dimension{})

	if err == nil {
		t.Errorf("Expected error.")
	}
}

func TestShapeIsScalar(t *testing.T) {
	s1 := NewShapeOrDie([]Dimension{4})
	if s1.IsScalar() {
		t.Errorf("Should not be scalar.")
	}

	s2 := NewShapeOrDie([]Dimension{1})
	if !s2.IsScalar() {
		t.Errorf("Should be scalar.")
	}
}

func TestShapesEqual(t *testing.T) {
	s1 := NewShapeOrDie([]Dimension{4, 1})
	s2 := NewShapeOrDie([]Dimension{1, 4})
	s3 := NewShapeOrDie([]Dimension{4, 1})

	if ShapesEqual(s1, s2) {
		t.Errorf("Shapes should not be equal.")
	}
	if !ShapesEqual(s1, s3) {
		t.Errorf("Shapes should be equal.")
	}
}
