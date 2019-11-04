package eager

import (
	"testing"
)

func TestNilFunc(t *testing.T) {
	RunFunc(nil)
}
