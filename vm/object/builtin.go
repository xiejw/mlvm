package object

import (
	"fmt"
	"unsafe"
)

var (
	sizeInt64  int = int(unsafe.Sizeof(int64(1)))
	sizeUint64 int = int(unsafe.Sizeof(uint64(1)))
	sizeByte   int = int(unsafe.Sizeof(byte(1)))
)

type Integer struct {
	Value int64
}

type String struct {
	Value string
}

func (o *Integer) Type() ObjectType { return IntegerType }
func (o *Integer) String() string   { return fmt.Sprintf("Integer(%v)", o.Value) }
func (o *Integer) MemSize() int     { return sizeInt64 }
func (o *String) Type() ObjectType  { return StringType }
func (o *String) String() string    { return fmt.Sprintf("String(\"%v\")", o.Value) }
func (o *String) MemSize() int      { return len(o.Value) * sizeByte }
