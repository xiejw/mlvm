package linalg

import (
	"github.com/xiejw/mlvm/vm/base/errors"
)

type Context struct {
}

type opType int

const (
	addOp opType = iota
	mulOp
)

func Add(ctx *Context, lhs, rhs, o []float32) error {
	return elementWiseOp(ctx, addOp, lhs, rhs, o)
}

func Mul(ctx *Context, lhs, rhs, o []float32) error {
	return elementWiseOp(ctx, mulOp, lhs, rhs, o)
}

func elementWiseOp(ctx *Context, op opType, lhs, rhs, o []float32) error {
	if len(lhs) != len(rhs) {
		return errors.New("lhs and rhs must be same length; but got %v vs %v.", len(lhs), len(rhs))
	}

	l := len(rhs)

	if len(o) != l {
		return errors.New("operand and output must be same length; but got %v vs %v.", len(o), l)
	}

	// consider to use parallel for.
	switch op {
	case addOp:
		for i := 0; i < l; i++ {
			o[i] = lhs[i] + rhs[i]
		}
	case mulOp:
		for i := 0; i < l; i++ {
			o[i] = lhs[i] * rhs[i]
		}
	default:
		return errors.New("unknown element-wise op (%v).", op)
	}

	return nil
}
