package tensor

type DimensionType int

const (
	IntType       DimensionType = 1
	BatchSizeType DimensionType = 2
)

type Dimension struct {
	Type  DimensionType
	Value int
}

type Shape struct {
	Dimensions []Dimension
}

type DType int

const (
	Float32 DType = 1
)

type Tensor interface {
}
