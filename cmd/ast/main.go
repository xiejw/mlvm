package main

import (
	"fmt"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {
	seed := ast.Decl{
		Name: "seed",
	}

	fmt.Printf("Hello MLVM.\n")
	fmt.Printf("  %v.\n", seed)
}
