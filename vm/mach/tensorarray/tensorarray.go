// Tensor implemention in vm.
package tensorarray

import (
	"fmt"
	"unsafe"

	"github.com/xiejw/mlvm/vm/object"
)

const (
	sizeInt     int = int(unsafe.Sizeof(int(1)))
	sizeFloat32 int = int(unsafe.Sizeof(float32(1.0)))
)

// TensorArray supports one simple form of broadcasting, on top of the normal Tensor representation.
// See OpTBROAD for detials.
type TensorArray struct {
	Dims     []int
	Rank     int
	Size     int
	RealSize int
	Value    []float32
}

// Creates TensorArray from raw components.
func FromRaw(dims []int, value []float32) *TensorArray {
	rank := len(dims)

	var size = 1
	for _, dim := range dims {
		size *= dim
	}

	realSize := len(value)

	if size%realSize != 0 {
		panic(fmt.Sprintf(
			"Tensor.FromRaw: size (%v) is not multiple of real size (%v).", size, realSize))
	}

	return &TensorArray{
		Dims:     dims,
		Rank:     rank,
		Size:     size,
		RealSize: realSize,
		Value:    value,
	}
}

// ----------------------------------------------------------------------------
// Helper Method to create TensorArray from Tensor.
// ----------------------------------------------------------------------------

func (ta *TensorArray) IsCompressed() bool {
	return ta.Size != ta.RealSize
}

func FromTensor(t *object.Tensor) *TensorArray {
	return FromRaw(t.Shape.Dims, t.Array.Value)
}

func (ta *TensorArray) ToTensor() *object.Tensor {
	if ta.IsCompressed() {
		panic("Converting compressed TensorArray to Tensor is not impl'ed.")
	}

	return object.NewTensor(ta.Dims, ta.Value)
}

// Converts the compressed tensor array to full array, i.e., `!IsCompressed()`.
func (ta *TensorArray) ToFullArray() *TensorArray {
	if !ta.IsCompressed() {
		return ta
	}

	var dst []float32

	if ta.RealSize > 1 {
		src := ta.Value
		dst = make([]float32, 0, ta.Size)
		repeated_times := ta.Size / ta.RealSize
		// Apprently, we can append quicker by append with a larger blocker.
		for i := 0; i < repeated_times; i++ {
			dst = append(dst, src...)
		}
	} else {
		// Special optimization for single element
		v := ta.Value[0]
		dst = make([]float32, ta.Size)
		size := ta.Size
		for i := 0; i < size; i++ {
			dst[i] = v
		}
	}

	return &TensorArray{
		Dims:     ta.Dims,
		Rank:     ta.Rank,
		Size:     ta.Size,
		RealSize: ta.Size,
		Value:    dst,
	}
}

// ----------------------------------------------------------------------------
// Conform object.Object
// ----------------------------------------------------------------------------

func (ta *TensorArray) String() string          { return "TensorArray" }
func (ta *TensorArray) MemSize() int            { return (ta.Rank+3)*sizeInt + ta.RealSize*sizeFloat32 }
func (ta *TensorArray) Type() object.ObjectType { return object.TensorType }
