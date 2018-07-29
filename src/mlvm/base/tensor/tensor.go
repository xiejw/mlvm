package tensor

type DType int

const (
	Float32 DType = 1
)

// Immutable
type Tensor interface {
	Name() string
	Shape() Shape
	DType() DType
}
