package ast

type OpKind int

const (
	opAdd OpKind = iota + 1 // 0 Is invalid.
)

var (
	opConstAdd = &Op{kind: opAdd}
)

// Oper
type Op struct {
	kind OpKind
}

func (op *Op) Kind() OpKind {
	return op.kind
}

func OpAdd() *Op {
	return opConstAdd
}
