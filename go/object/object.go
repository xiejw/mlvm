// Package `object` provides the Object representation in the compiled code and vm.
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

// All object implementations must be immutable and serializable.
type Object interface {
	Type() ObjectType
	String() string
}
