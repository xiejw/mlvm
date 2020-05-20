package object

type ObjectType int

const (
	// basic_types.go
	IntType ObjectType = iota
	StringType

	ShapeType  // shape.go
	ArrayType  // array.go
	TensorType // tensor.go
)

type Object interface {
	Type() ObjectType
	String() string
}
