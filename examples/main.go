package main

import (
	"fmt"

	"github.com/xiejw/mlvm/mlvm/array"
)

func main() {
	arr := array.NewArrayOrDie("a", []array.Dimension{2, 1}, []array.Float{1.0, 2.0})
	fmt.Printf("%v", arr)
}
