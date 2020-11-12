package algorithrms

import (
	"testing"

	"github.com/xiejw/mlvm/vm/base/errors"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Enqueue(2)
	v, err := q.Dequeue()
	assertNoErr(t, err)
	if v != 2 {
		t.Errorf("value mismatch.")
	}
	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("expect error.")
	}

	q.Enqueue(22)
	q.Enqueue(33)
	v, err = q.Dequeue()
	assertNoErr(t, err)
	if v != 22 {
		t.Errorf("value mismatch.")
	}
	v, err = q.Dequeue()
	assertNoErr(t, err)
	if v != 33 {
		t.Errorf("value mismatch.")
	}
	_, err = q.Dequeue()
	if err == nil {
		t.Errorf("expect error.")
	}
}

func assertNoErr(t *testing.T, err *errors.DError) {
	t.Helper()
	if err != nil {
		t.Fatalf("did not expect error. got: %v", err)
	}
}
