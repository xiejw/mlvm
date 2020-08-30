package object

type ObjectType int

const (
	IntegerType ObjectType = iota // builtin.go
	StringType                    // builtin.go
	RngType                       // rng.go
	ShapeType                     // tensor.go
	ArrayType                     // tensor.go
	TensorType                    // tensor.go
)

type Object interface {
	Type() ObjectType
	String() string
}
