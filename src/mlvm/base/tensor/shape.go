package tensor

// Immutable
type Shape interface {
	Dim(index int) Dimension
	Rank() int
	AsList() []Dimension
}

func NewShape(dimArgs ...int) Shape {
	dims := make([]*Dimension, 0, len(dimArgs))
	dims = fillDims(dims, dimArgs...)
	return &shapeImpl{dims: dims}
}

func NewShapeFromDims(dimArgs ...Dimension) Shape {
	dims := make([]*Dimension, 0, len(dimArgs))
	for _, dim := range dimArgs {
		var dimCopy Dimension
		dimCopy = dim
		dims = append(dims, &dimCopy)
	}
	return &shapeImpl{dims: dims}
}

func NewShapeWithBatchSize(dimArgs ...int) Shape {
	dims := make([]*Dimension, 0, len(dimArgs)+1)
	dims = append(dims, &Dimension{
		Type: BatchSizeType,
	})
	dims = fillDims(dims, dimArgs...)
	return &shapeImpl{dims: dims}
}

// Fills `dims` with int dimension specified by `dimArgs`.
func fillDims(dims []*Dimension, dimArgs ...int) []*Dimension {
	for _, dim := range dimArgs {
		dims = append(dims, &Dimension{
			Type:  IntType,
			Value: dim,
		})
	}
	return dims
}
