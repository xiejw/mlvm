package tensor

type DType int

const (
	Float32 DType = 1
)

// Immutable
type Tensor interface {
	Shape() Shape
	Dtype() DType
}
