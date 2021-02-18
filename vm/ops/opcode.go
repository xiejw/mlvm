package ops

type OpCode int

const (
	OP_FILL OpCode = iota
)

type Option interface {
	option()
}
