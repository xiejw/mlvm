package ast

import (
	_ "bytes"
	"fmt"
	"io"
)

// Decl{
//   Token: "name"
//   Value: value.String()
// }\n
func (decl *Decl) WriteDebugString(w io.Writer, indent string) {
	newIndent := indent + "  "

	fmt.Fprintf(w, "Decl{\n%sToken: \"%s\"\n%sValue: ",
		newIndent, decl.Name, newIndent)
	decl.Value.WriteDebugString(w, newIndent)
	fmt.Fprintf(w, "%s}\n", indent)
}

// Expr{
//   Value: value.String()
// }\n
func (expr *Expression) WriteDebugString(w io.Writer, indent string) {
	if expr.Type != ExTpValue {
		panic("unsupported expression type.")
	}
	newIndent := indent + "  "
	fmt.Fprintf(w, "Expr{\n%sValue: %s\n%s}\n",
		newIndent, expr.Value, indent)
}
