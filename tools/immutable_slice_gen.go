package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const code = `//
// DO NOT EDIT. Generated by 'make generate'
package {{.Package}}

{{.Comment}}
type {{.Name}} interface {
	Iterator() <-chan {{.Type}}
}

type {{.ImplName}} struct {
	data      []{{.Type}}
	finalized bool
}

func (impl *{{.ImplName}}) Iterator() <-chan {{.Type}} {
	if !impl.finalized {
		panic("The Build should be called first.")
	}
	c := make(chan {{.Type}})
	go func() {
		for _, item := range impl.data {
			c <- item
		}
	}()
	return c
}

func (impl *{{.ImplName}}) Append(item {{.Type}}) {
	impl.data = append(impl.data, item)
}

func (impl *{{.ImplName}}) Build() {
	if impl.finalized {
		panic("The Build should not be called twice.")
	}
	impl.finalized = true
}
`

	type Context struct {
		Package, Name, Type, Comment string
		ImplName                     string
	}

	c := Context{
		Package:  "layers",
		Comment:  "// Inputs of layers. Typically a list of `Layer`s.",
		Name:     "Inputs",
		Type:     "Layer",
		ImplName: "InputsBuilder",
	}

	t := template.Must(template.New("code").Parse(code))

	// Execute the template.
	err := t.Execute(os.Stdout, c)
	if err != nil {
		log.Println("executing template:", err)
	}
}
