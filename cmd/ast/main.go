package main

import (
	"fmt"
	"log"

	"encoding/json"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/syntax/ast"
)

func main() {
	seed := ast.Decl{
		ID: "seed",
		Type: &ast.Type{
			Kind: ast.TpKdPrng,
		},
		Value: &ast.Expression{
			Type: &ast.Type{
				Kind: ast.TpKdPrng,
			},
			Kind: ast.EpKdCall,
			Left: &ast.Expression{
				Kind: ast.EpKdID,
				Value: &object.String{
					Value: "prng_create",
				},
			},
			Right: &ast.Expression{},
		},
	}

	fmt.Printf("Hello MLVM.\n")
	prettyJSON, err := json.MarshalIndent(seed, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
