package object

import "fmt"

type Integer struct {
	Value int64
}

func (o *Integer) Type() ObjectType {
	return IntegerType
}

func (o *Integer) String() string {
	return fmt.Sprintf("Integer(%v)", o.Value)
}

type String struct {
	Value string
}

func (o *String) Type() ObjectType {
	return StringType
}

func (o *String) String() string {
	return fmt.Sprintf("String(\"%v\")", o.Value)
}
