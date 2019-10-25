package ir

type OpKind int

const (
	OpKAdd OpKind = iota + 1 // 0 Is invalid.
)

var (
	opConstAdd = &Op{kind: OpKAdd}
)

// Oper
type Op struct {
	kind OpKind
}

func (op *Op) Kind() OpKind {
	return op.kind
}

func (op *Op) BaseName() string {
	switch op.kind {
	case OpKAdd:
		return "opAdd"
	default:
		panic("Op Kind is not expected.")
	}

}

func OpAdd() *Op {
	return opConstAdd
}
