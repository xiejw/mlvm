package tests

import (
	"reflect"
	"testing"

	"github.com/xiejw/mlvm/go/compiler"
	"github.com/xiejw/mlvm/go/object"
	"github.com/xiejw/mlvm/go/vm"
	"github.com/xiejw/mlvm/go/syntax/parser"
)

func createSimpleTensor() *object.Tensor {
	shape := object.NewShape([]object.NamedDim{{"x", 2}})
	array := &object.Array{[]float32{1.0, 2.0}}
	return &object.Tensor{shape, array}
}

func TestExprTensorStoreLoad(t *testing.T) {
	p, err := parser.New([]byte(`(store_load "a")`)).ParseAst()
	assertNoErr(t, err)

	o, err := compiler.Compile(p)
	assertNoErr(t, err)

	ts := vm.NewTensorStore()
	err = ts.Store("a", createSimpleTensor())
	assertNoErr(t, err)

	m := vm.NewVMWithTensorStore(o, ts)

	outputs, err := m.Run()
	got := assertSingleOutput(t, outputs, err)

	expected := createSimpleTensor()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("unexpected output.")
	}
}
