package tensor

import (
	"bytes"
	"strings"
)

type shapeImpl struct {
	dims []*Dimension
}

func (s *shapeImpl) Dim(index int) Dimension {
	return *s.dims[index]
}

func (s *shapeImpl) Rank() int {
	return len(s.dims)
}

func (s *shapeImpl) AsList() []Dimension {
	dimsCopy := make([]Dimension, 0, len(s.dims))
	for _, dim := range s.dims {
		dimsCopy = append(dimsCopy, *dim)
	}
	return dimsCopy
}

func (s *shapeImpl) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")

	dimStrs := make([]string, 0, s.Rank())
	for _, dim := range s.dims {
		dimStrs = append(dimStrs, dim.String())
	}
	buf.WriteString(strings.Join(dimStrs, ", "))

	buf.WriteString("]")
	return buf.String()
}
