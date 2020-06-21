package main

import (
	"fmt"
	"log"

	"encoding/json"
	_ "github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {

	statements := make([]ast.Statement, 0)
	statements = append(statements, &ast.DeclStatement{
		ID: "seed",
	})

	// seed := ast.Decl{
	// 	ID:   "seed",
	// 	Type: makeType(ast.TpKdPrng),
	// 	Value: &ast.Expression{
	// 		Type: makeType(ast.TpKdPrng),
	// 		Kind: ast.EpKdCall,
	// 		Left: makeStringIdExpr("prng_create"),
	// 		Right: &ast.Expression{
	// 			Type: &ast.Type{
	// 				Kind: ast.TpKdNA,
	// 			},
	// 			Kind: ast.EpKdArg,
	// 			Left: &ast.Expression{
	// 				Type: makeType(ast.TpKdInt),
	// 				Kind: ast.EpKdIntLiteral,
	// 				Value: &object.Integer{
	// 					Value: 123,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	//
	p := &ast.Program{
		Statements: statements,
	}

	prettyJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
