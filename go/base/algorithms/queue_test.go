package algorithrms

import (
	"testing"

	"github.com/xiejw/mlvm/go/base/errors"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(2)
	v, err := q.Dequeue()
	assertNoErr(t, err)
	if v != 2 {
		t.Errorf("value mismatch.")
	}
}

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
