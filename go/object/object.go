package object

type ObjectType int

const (
	// basic_types.go
	IntegerType ObjectType = iota
	StringType

	PrngType   // prng.go
	ShapeType  // shape.go
	ArrayType  // array.go
	TensorType // tensor.go
)

type Object interface {
	Type() ObjectType
	String() string
}
