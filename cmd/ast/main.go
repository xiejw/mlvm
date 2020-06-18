package main

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func makeType(kind ast.TypeKind) *ast.Type {
	return &ast.Type{Kind: kind}
}

func makeStringIdExpr(id string) *ast.Expression {
	return &ast.Expression{
		Type: makeType(ast.TpKdNA), // This is not correct ast. Each ID should have type from table.
		Kind: ast.EpKdID,
		Value: &object.String{
			Value: id,
		},
	}
}

func main() {
	seed := ast.Decl{
		ID:   "seed",
		Type: makeType(ast.TpKdPrng),
		Value: &ast.Expression{
			Type: makeType(ast.TpKdPrng),
			Kind: ast.EpKdCall,
			Left: makeStringIdExpr("prng_create"),
			Right: &ast.Expression{
				Type: &ast.Type{
					Kind: ast.TpKdNA,
				},
				Kind: ast.EpKdArg,
				Left: &ast.Expression{
					Type: makeType(ast.TpKdInt),
					Kind: ast.EpKdIntLiteral,
					Value: &object.Integer{
						Value: 123,
					},
				},
			},
		},
	}

	fmt.Printf("Hello MLVM.\n")
	prettyJSON, err := json.MarshalIndent(seed, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
