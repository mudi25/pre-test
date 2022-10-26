package errors_test

import (
	"fmt"
	"testing"

	"github.com/mudi25/pretest/errors"
)

func TestMain(m *testing.M) {
	m.Run()
}
func TestErrors(t *testing.T) {
	type Data struct {
		err        error
		code       int
		isInternal bool
	}
	data := []Data{
		{err: nil, code: 200, isInternal: false},
		{err: errors.BadRequest(), code: 400, isInternal: false},
		{err: errors.IdNotFound(), code: 404, isInternal: false},
		{err: errors.NotFound(), code: 404, isInternal: false},
		{err: errors.Internal(), code: 500, isInternal: true},
		{err: fmt.Errorf("unknown error"), code: 500, isInternal: true},
	}
	for i, d := range data {
		code, _ := errors.FromError(d.err)
		switch {
		case d.code != code:
			t.Errorf("%d expect error code equal to code", i+1)
		}
	}
}
func TestErrorString(t *testing.T) {
	e := errors.Internal()
	if e.Error() != fmt.Sprintf("%d: %s", e.Code, e.Message) {
		t.Error("invalid error string")
	}
}
