package eager

import (
	"log"

	"github.com/xiejw/mlvm/mlvm/array"
	"github.com/xiejw/mlvm/mlvm/compile/ir"
)

func RunFunc(fn *ir.Func) ([]*array.Array, error) {
	log.Printf("Running: %v\n", fn)
	if fn == nil {
		return nil, nil
	}

	for _, ins := range fn.Instructions() {
		log.Printf("Run: %v", ins)
	}
	return nil, nil
}
