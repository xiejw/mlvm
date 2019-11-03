package eager

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/compile/ir"
)

func RunFunc(fn *ir.Func) ([]*array.Array, error) {
	fmt.Printf("Running: %v\n", fn)
	return nil, nil
}
