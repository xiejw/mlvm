package code

import "testing"

const size = 100000

func BenchmarkDirectInt(b *testing.B) {
	slice := make([]int, size)
	for k := 0; k < size; k++ {
		slice[k] = k
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sum := 0

		for k := 0; k < size; k++ {
			sum += slice[k]
		}
	}
}

func BenchmarkInterfaceCastInt(b *testing.B) {
	slice := make([]interface{}, size)
	for k := 0; k < size; k++ {
		slice[k] = k
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sum := 0

		for k := 0; k < size; k++ {
			sum += slice[k].(int)
		}
	}
}

type Object interface {
	Type() int
}

type IntBox struct {
	Value int
}

func (ib *IntBox) Type() int {
	return 0
}

func BenchmarkInterfaceCastIntboxAsValue(b *testing.B) {
	slice := make([]interface{}, size)
	for k := 0; k < size; k++ {
		slice[k] = IntBox{Value: k}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sum := 0

		for k := 0; k < size; k++ {
			sum += slice[k].(IntBox).Value
		}
	}
}

func BenchmarkInterfaceCastIntPtr(b *testing.B) {
	slice := make([]interface{}, size)

	// Allocate memory might take times.
	for k := 0; k < size; k++ {
		slice[k] = &IntBox{Value: k}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sum := 0

		for k := 0; k < size; k++ {
			sum += slice[k].(*IntBox).Value
		}
	}
}

func BenchmarkIntBoxCastIntPtr(b *testing.B) {
	slice := make([]Object, size)

	for k := 0; k < size; k++ {
		slice[k] = &IntBox{Value: k}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		sum := 0

		for k := 0; k < size; k++ {
			sum += slice[k].(*IntBox).Value
		}
	}
}
