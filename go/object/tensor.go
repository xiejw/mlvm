package object

import (
	"fmt"
)

type Tensor struct {
	Shape *Shape
	Value *Array
}

func (t *Tensor) Type() ObjectType {
	return TensorType
}

func (t *Tensor) String() string {
	return fmt.Sprintf("%v %v", t.Shape, t.Value)
}
