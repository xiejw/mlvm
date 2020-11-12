package errors

import (
	"github.com/xiejw/mlvm/vm/base/errors"
)

type DError = errors.DError

var (
	New  = errors.New
	From = errors.From
)
