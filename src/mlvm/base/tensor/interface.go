package tensor

type DimensionType int

const (
	IntValue DimensionType = 1
	BatchSize DimensionType = 2
)

type Dimension struct {
	Type DimensionType
	Value int
}

type Shape struct {
	Dimensions []Dimension
}

type Tensor interface {
}
