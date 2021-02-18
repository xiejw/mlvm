package ops

type OpCode int

const (
	OP_RNG OpCode = iota
	OP_ADD
	OP_REDUCE
)

type Option interface {
	option()
}
