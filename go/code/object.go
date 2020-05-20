package code

import "fmt"

type ObjectType int

const (
	IntType ObjectType = iota
	StringType
	PRNGType
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

type PRNG struct {
	// Consider to use original seed, and state hash.
}

func (o *PRNG) Type() ObjectType {
	return PRNGType
}

func (o *PRNG) String() string {
	return fmt.Sprintf("PRNG(seed: `??` state: `??`)")
}
