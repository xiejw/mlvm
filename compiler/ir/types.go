package ir

type TypeKind int

type Type struct {
	Kind TypeKind
}

const (
	KInt TypeKind = iota
)
