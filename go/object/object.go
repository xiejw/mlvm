package object

import "fmt"

type ObjectType int

const (
	IntType ObjectType = iota
	StringType
	ShapeType
)

type Object interface {
	Type() ObjectType
	String() string
}

type Integer struct {
	Value int64
}

func (o *Integer) Type() ObjectType {
	return IntType
}

func (o *Integer) String() string {
	return fmt.Sprintf("Int(%v)", o.Value)
}

type String struct {
	Value string
}

func (o *String) Type() ObjectType {
	return StringType
}

func (o *String) String() string {
	return fmt.Sprintf("String(`%v`)", o.Value)
}
