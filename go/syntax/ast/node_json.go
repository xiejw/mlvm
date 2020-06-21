package ast

// import (
// 	"encoding/json"
// 	"fmt"
// )
//
// func (kind TypeKind) MarshalJSON() ([]byte, error) {
// 	switch kind {
// 	case TpKdNA:
// 		return json.Marshal("n/a")
// 	case TpKdInt:
// 		return json.Marshal("int")
// 	case TpKdString:
// 		return json.Marshal("string")
// 	case TpKdArray:
// 		return json.Marshal("array")
// 	case TpKdNamedDim:
// 		return json.Marshal("named_dim")
// 	case TpKdTensor:
// 		return json.Marshal("tensor")
// 	case TpKdFunc:
// 		return json.Marshal("func")
// 	case TpKdPrng:
// 		return json.Marshal("prng")
// 	default:
// 		panic(fmt.Errorf("unknown type kind: %v", kind))
// 	}
// }
//
// func (kind ExpressionKind) MarshalJSON() ([]byte, error) {
// 	switch kind {
// 	case EpKdCall:
// 		return json.Marshal("func_call")
// 	case EpKdID:
// 		return json.Marshal("id")
// 	case EpKdArg:
// 		return json.Marshal("func_arg")
// 	case EpKdIntLiteral:
// 		return json.Marshal("int_literal")
// 	default:
// 		panic(fmt.Errorf("unknown expr kind: %v", kind))
// 	}
// }
