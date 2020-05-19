package code

import "fmt"

type ObjectType int

const (
	IntType ObjectType = iota
)

type Object interface {
	Type() ObjectType
	String() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return IntType
}

func (i *Integer) String() string {
	return fmt.Sprintf("Int(%v)", i.Value)
}
