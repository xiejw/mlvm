package object

type NamedDimension struct {
	Name string
	Size uint // Noozero
}

type Shape struct {
	Dimensions []NamedDimension // Cannot have dup names.
	Rank       uint
}

func NewShape(dims []NamedDimension) *Shape {
	return &Shape{
		Dimensions: dims,
		Rank:       uint(len(dims)),
	}
}

func (shape *Shape) Type() ObjectType {
	return ShapeType
}

func (shape *Shape) String() string {
	return "A shape"
}
