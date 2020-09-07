package tensorarray

import (
	"testing"

	"github.com/xiejw/mlvm/go/object"
)

func TestObjectInterface(t *testing.T) {
	var o object.Object

	ta := &TensorArray{}
	o = ta

	if o.String() != "TensorArray" {
		t.Fatalf("String() mistmatch.")
	}
}
