package object

type ObjectType int

const (
	// basic_types.go
	IntType ObjectType = iota
	StringType

	PrngType // prng64 package

	ShapeType  // shape.go
	ArrayType  // array.go
	TensorType // tensor.go
)

type Object interface {
	Type() ObjectType
	String() string
}
