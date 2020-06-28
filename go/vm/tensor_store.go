package vm

import (
	"github.com/xiejw/mlvm/go/object"
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
