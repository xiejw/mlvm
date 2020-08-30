package object

type ObjectType int

const (
	IntegerType ObjectType = iota // builtin.go
	StringType                    // builtin.go
	RngType                       // Rng.go
	ShapeType                     // shape.go
	ArrayType                     // array.go
	TensorType                    // tensor.go
)

type Object interface {
	Type() ObjectType
	String() string
}
