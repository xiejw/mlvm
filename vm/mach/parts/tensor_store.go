package parts

import (
	"github.com/xiejw/mlvm/vm/base/errors"
	"github.com/xiejw/mlvm/vm/object"
)

var (
	ErrTSTensorNotFound = errors.New("object not found in key-value store ")
)

// TensorStore is a key-Tensor store.
type TensorStore interface {
	Load(key string) (*object.Tensor, error)
	Store(key string, value *object.Tensor) error
}

type storeImpl struct {
	db map[string]*object.Tensor
}

func NewTensorStore() TensorStore {
	return &storeImpl{
		db: make(map[string]*object.Tensor),
	}
}

func (st *storeImpl) Load(key string) (*object.Tensor, error) {
	tensor, ok := st.db[key]
	if !ok {
		return nil, ErrTSTensorNotFound
	}
	return tensor, nil
}

func (st *storeImpl) Store(key string, value *object.Tensor) error {
	st.db[key] = value
	return nil
}
