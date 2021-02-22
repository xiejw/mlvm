package linalg

import (
	"github.com/xiejw/mlvm/vm/base/errors"
)

func Sum(ctx *Context, x []float32, dims []int, reduce_dims []int, o []float32) error {
	if len(o) != 1 {
		return errors.New("linalg.Sum only supports reducing all dims.")
	}
	var y float32
	for i := 0; i < len(x); i++ {
		y += x[i]
	}
	o[0] = y
	return nil
}
