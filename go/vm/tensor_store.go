package vm

import (
	"github.com/xiejw/mlvm/go/object"
)

// TensorStore is a key-tensor store.
type TensorStore interface {
	Load(key string) (*object.Tensor, error)
	Store(key string, value *object.Tensor) error
}
