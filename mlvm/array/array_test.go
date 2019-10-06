package array

import (
	"testing"
)

func TestRankOneTensor(t *testing.T) {
	a := NewArrayOrDie("a", []Dimension{4}, []Float{1.0, 2.0, 3.0, 4.0})
	got := a.String()
	expected := "[ 1.000 2.000 3.000 4.000 ]"

	if got != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, got)
	}
}

func TestRankTwoVerticalTensor(t *testing.T) {
	a := NewArrayOrDie("a", []Dimension{4, 1}, []Float{1.0, 2.0, 3.0, 4.0})
	got := a.String()

	expected := `
[ [ 1.000 ]
  [ 2.000 ]
  [ 3.000 ]
  [ 4.000 ]
]
`
	if got != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, got)
	}
}

func TestRankTwoHorizontalTensor(t *testing.T) {
	a := NewArrayOrDie("a", []Dimension{1, 4}, []Float{1.0, 2.0, 3.0, 4.0})
	got := a.String()

	expected := `
[ [ 1.000 2.000 3.000 4.000 ]
]
`
	if got != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, got)
	}
}

func TestRankTwoMatrixTensor(t *testing.T) {
	a := NewArrayOrDie("a", []Dimension{2, 2}, []Float{1.0, 2.0, 3.0, 4.0})
	got := a.String()

	expected := `
[ [ 1.000 2.000 ]
  [ 3.000 4.000 ]
]
`
	if got != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, got)
	}
}

func TestRankTreeMatrixTensor(t *testing.T) {
	a := NewArrayOrDie("a", []Dimension{2, 2, 1}, []Float{1.0, 2.0, 3.0, 4.0})
	got := a.String()

	expected := `
[ [ [ 1.000 ]
    [ 2.000 ]
  ]
  [ [ 3.000 ]
    [ 4.000 ]
  ]
]
`
	if got != expected {
		t.Errorf("Expected: %v\nGot: %v\n", expected, got)
	}
}
