package object

import "fmt"

type Integer struct {
	Value int64
}

type String struct {
	Value string
}

func (o *Integer) Type() ObjectType { return IntegerType }
func (o *Integer) String() string   { return fmt.Sprintf("Integer(%v)", o.Value) }
func (o *String) Type() ObjectType  { return StringType }
func (o *String) String() string    { return fmt.Sprintf("String(\"%v\")", o.Value) }
