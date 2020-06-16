package main

import (
	"fmt"

	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {
	seed := ast.Decl{
		Name: "seed",
		Type: &ast.Type{
			Kind: ast.TpKdPrng,
		},
		Value: &ast.Expression{
			Type: *ast.Type{
				Kind: ast.TpKdFunc,
				SubType: ast.TpKdPrng,
			},
		},
	}

	fmt.Printf("Hello MLVM.\n")
	fmt.Printf("  %v.\n", seed)
}
