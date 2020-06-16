package ast

import (
	_ "bytes"
	_ "fmt"
	_ "io"
)

// // Decl{
// //   Token: "name"
// //   Value: value.String()
// // }\n
// func (decl *Decl) WriteDebugString(w io.Writer, indent string) {
// 	newIndent := indent + "  "
//
// 	fmt.Fprintf(w, "Decl{\n%sName: \"%s\"\n%sType: %v\n%sValue: ",
// 		newIndent, decl.ID,
// 		newIndent, decl.Type,
// 		newIndent)
// 	decl.Value.WriteDebugString(w, newIndent)
// 	fmt.Fprintf(w, "%s}\n", indent)
// }
//
// // Expr{
// //   Value: value.String()
// // }\n
// func (expr *Expression) WriteDebugString(w io.Writer, indent string) {
// 	if expr.Kind != EpKdIntLiteral {
// 		panic("unsupported expression type.")
// 	}
// 	newIndent := indent + "  "
// 	fmt.Fprintf(w, "Expr{\n%sValue: %s\n%s}\n",
// 		newIndent, expr.Value, indent)
// }
//
// func (tp *Type) String() string {
// 	if tp.Kind != TpKdNamedDim {
// 		panic("unsupported type kind.")
// 	}
//
// 	return "Type(NamedDim)"
// }
