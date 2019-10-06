package array

type Dimension int

type Float float32

type Shape struct {
	Dims []Dimension
}

type Data struct {
	Value []Float
}

type Array struct {
	name  string // Name of the Array
	shape *Shape // The real shape.
	data  *Data  // The backing data.
}

func (arr *Array) Name() string {
	return arr.name
}

func (arr *Array) Shape() *Shape {
	return arr.shape
}

func (arr *Array) Data() *Data {
	return arr.data
}

// func IsScalar(shape *Shape) bool {
// 	return len(shape) == 1 && shape[0] == 1
// }
//
// func ElementCount(shape []Dimension) int {
// 	count := 1
// 	for _, dim := range shape {
// 		count *= int(dim)
// 	}
// 	return count
// }
