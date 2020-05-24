package object

import "testing"

func TestObjectType(t *testing.T) {
	if int(IntType) != 0 {
		t.Errorf("unexpected value for IntType.")
	}
	if int(ShapeType) == 0 {
		t.Errorf("unexpected value for ShapeType.")
	}
}
